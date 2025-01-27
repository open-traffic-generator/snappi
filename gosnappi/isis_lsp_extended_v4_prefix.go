package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspExtendedV4Prefix *****
type isisLspExtendedV4Prefix struct {
	validation
	obj                    *otg.IsisLspExtendedV4Prefix
	marshaller             marshalIsisLspExtendedV4Prefix
	unMarshaller           unMarshalIsisLspExtendedV4Prefix
	prefixAttributesHolder IsisLspPrefixAttributes
}

func NewIsisLspExtendedV4Prefix() IsisLspExtendedV4Prefix {
	obj := isisLspExtendedV4Prefix{obj: &otg.IsisLspExtendedV4Prefix{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspExtendedV4Prefix) msg() *otg.IsisLspExtendedV4Prefix {
	return obj.obj
}

func (obj *isisLspExtendedV4Prefix) setMsg(msg *otg.IsisLspExtendedV4Prefix) IsisLspExtendedV4Prefix {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspExtendedV4Prefix struct {
	obj *isisLspExtendedV4Prefix
}

type marshalIsisLspExtendedV4Prefix interface {
	// ToProto marshals IsisLspExtendedV4Prefix to protobuf object *otg.IsisLspExtendedV4Prefix
	ToProto() (*otg.IsisLspExtendedV4Prefix, error)
	// ToPbText marshals IsisLspExtendedV4Prefix to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspExtendedV4Prefix to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspExtendedV4Prefix to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspExtendedV4Prefix to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspExtendedV4Prefix struct {
	obj *isisLspExtendedV4Prefix
}

type unMarshalIsisLspExtendedV4Prefix interface {
	// FromProto unmarshals IsisLspExtendedV4Prefix from protobuf object *otg.IsisLspExtendedV4Prefix
	FromProto(msg *otg.IsisLspExtendedV4Prefix) (IsisLspExtendedV4Prefix, error)
	// FromPbText unmarshals IsisLspExtendedV4Prefix from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspExtendedV4Prefix from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspExtendedV4Prefix from JSON text
	FromJson(value string) error
}

func (obj *isisLspExtendedV4Prefix) Marshal() marshalIsisLspExtendedV4Prefix {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspExtendedV4Prefix{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspExtendedV4Prefix) Unmarshal() unMarshalIsisLspExtendedV4Prefix {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspExtendedV4Prefix{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspExtendedV4Prefix) ToProto() (*otg.IsisLspExtendedV4Prefix, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspExtendedV4Prefix) FromProto(msg *otg.IsisLspExtendedV4Prefix) (IsisLspExtendedV4Prefix, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspExtendedV4Prefix) ToPbText() (string, error) {
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

func (m *unMarshalisisLspExtendedV4Prefix) FromPbText(value string) error {
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

func (m *marshalisisLspExtendedV4Prefix) ToYaml() (string, error) {
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

func (m *unMarshalisisLspExtendedV4Prefix) FromYaml(value string) error {
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

func (m *marshalisisLspExtendedV4Prefix) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspExtendedV4Prefix) ToJson() (string, error) {
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

func (m *unMarshalisisLspExtendedV4Prefix) FromJson(value string) error {
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

func (obj *isisLspExtendedV4Prefix) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspExtendedV4Prefix) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspExtendedV4Prefix) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspExtendedV4Prefix) Clone() (IsisLspExtendedV4Prefix, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspExtendedV4Prefix()
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

func (obj *isisLspExtendedV4Prefix) setNil() {
	obj.prefixAttributesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspExtendedV4Prefix is this group defines attributes of an IPv4 standard prefix.
type IsisLspExtendedV4Prefix interface {
	Validation
	// msg marshals IsisLspExtendedV4Prefix to protobuf object *otg.IsisLspExtendedV4Prefix
	// and doesn't set defaults
	msg() *otg.IsisLspExtendedV4Prefix
	// setMsg unmarshals IsisLspExtendedV4Prefix from protobuf object *otg.IsisLspExtendedV4Prefix
	// and doesn't set defaults
	setMsg(*otg.IsisLspExtendedV4Prefix) IsisLspExtendedV4Prefix
	// provides marshal interface
	Marshal() marshalIsisLspExtendedV4Prefix
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspExtendedV4Prefix
	// validate validates IsisLspExtendedV4Prefix
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspExtendedV4Prefix, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Ipv4Address returns string, set in IsisLspExtendedV4Prefix.
	Ipv4Address() string
	// SetIpv4Address assigns string provided by user to IsisLspExtendedV4Prefix
	SetIpv4Address(value string) IsisLspExtendedV4Prefix
	// HasIpv4Address checks if Ipv4Address has been set in IsisLspExtendedV4Prefix
	HasIpv4Address() bool
	// PrefixLength returns uint32, set in IsisLspExtendedV4Prefix.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to IsisLspExtendedV4Prefix
	SetPrefixLength(value uint32) IsisLspExtendedV4Prefix
	// HasPrefixLength checks if PrefixLength has been set in IsisLspExtendedV4Prefix
	HasPrefixLength() bool
	// Metric returns uint32, set in IsisLspExtendedV4Prefix.
	Metric() uint32
	// SetMetric assigns uint32 provided by user to IsisLspExtendedV4Prefix
	SetMetric(value uint32) IsisLspExtendedV4Prefix
	// HasMetric checks if Metric has been set in IsisLspExtendedV4Prefix
	HasMetric() bool
	// RedistributionType returns IsisLspExtendedV4PrefixRedistributionTypeEnum, set in IsisLspExtendedV4Prefix
	RedistributionType() IsisLspExtendedV4PrefixRedistributionTypeEnum
	// SetRedistributionType assigns IsisLspExtendedV4PrefixRedistributionTypeEnum provided by user to IsisLspExtendedV4Prefix
	SetRedistributionType(value IsisLspExtendedV4PrefixRedistributionTypeEnum) IsisLspExtendedV4Prefix
	// HasRedistributionType checks if RedistributionType has been set in IsisLspExtendedV4Prefix
	HasRedistributionType() bool
	// PrefixAttributes returns IsisLspPrefixAttributes, set in IsisLspExtendedV4Prefix.
	// IsisLspPrefixAttributes is this contains the properties of ISIS Prefix attributes for  the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html
	PrefixAttributes() IsisLspPrefixAttributes
	// SetPrefixAttributes assigns IsisLspPrefixAttributes provided by user to IsisLspExtendedV4Prefix.
	// IsisLspPrefixAttributes is this contains the properties of ISIS Prefix attributes for  the extended IPv4 and IPv6 reachability. https://www.rfc-editor.org/rfc/rfc7794.html
	SetPrefixAttributes(value IsisLspPrefixAttributes) IsisLspExtendedV4Prefix
	// HasPrefixAttributes checks if PrefixAttributes has been set in IsisLspExtendedV4Prefix
	HasPrefixAttributes() bool
	setNil()
}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// Ipv4Address returns a string
func (obj *isisLspExtendedV4Prefix) Ipv4Address() string {

	return *obj.obj.Ipv4Address

}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// Ipv4Address returns a string
func (obj *isisLspExtendedV4Prefix) HasIpv4Address() bool {
	return obj.obj.Ipv4Address != nil
}

// An IPv4 unicast prefix reachable via the originator of this LSP.
// SetIpv4Address sets the string value in the IsisLspExtendedV4Prefix object
func (obj *isisLspExtendedV4Prefix) SetIpv4Address(value string) IsisLspExtendedV4Prefix {

	obj.obj.Ipv4Address = &value
	return obj
}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *isisLspExtendedV4Prefix) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// The length of the IPv4 prefix.
// PrefixLength returns a uint32
func (obj *isisLspExtendedV4Prefix) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// The length of the IPv4 prefix.
// SetPrefixLength sets the uint32 value in the IsisLspExtendedV4Prefix object
func (obj *isisLspExtendedV4Prefix) SetPrefixLength(value uint32) IsisLspExtendedV4Prefix {

	obj.obj.PrefixLength = &value
	return obj
}

// ISIS wide metric.
// Metric returns a uint32
func (obj *isisLspExtendedV4Prefix) Metric() uint32 {

	return *obj.obj.Metric

}

// ISIS wide metric.
// Metric returns a uint32
func (obj *isisLspExtendedV4Prefix) HasMetric() bool {
	return obj.obj.Metric != nil
}

// ISIS wide metric.
// SetMetric sets the uint32 value in the IsisLspExtendedV4Prefix object
func (obj *isisLspExtendedV4Prefix) SetMetric(value uint32) IsisLspExtendedV4Prefix {

	obj.obj.Metric = &value
	return obj
}

type IsisLspExtendedV4PrefixRedistributionTypeEnum string

// Enum of RedistributionType on IsisLspExtendedV4Prefix
var IsisLspExtendedV4PrefixRedistributionType = struct {
	UP   IsisLspExtendedV4PrefixRedistributionTypeEnum
	DOWN IsisLspExtendedV4PrefixRedistributionTypeEnum
}{
	UP:   IsisLspExtendedV4PrefixRedistributionTypeEnum("up"),
	DOWN: IsisLspExtendedV4PrefixRedistributionTypeEnum("down"),
}

func (obj *isisLspExtendedV4Prefix) RedistributionType() IsisLspExtendedV4PrefixRedistributionTypeEnum {
	return IsisLspExtendedV4PrefixRedistributionTypeEnum(obj.obj.RedistributionType.Enum().String())
}

// Up (0)-used when a prefix is initially advertised within the ISIS L3 hierarchy,
// and for all other prefixes in L1 and L2 LSPs. (default)
// Down (1)-used when an L1/L2 router advertises L2 prefixes in L1 LSPs.
// The prefixes are being advertised from a higher level (L2) down to a lower level (L1).
// RedistributionType returns a string
func (obj *isisLspExtendedV4Prefix) HasRedistributionType() bool {
	return obj.obj.RedistributionType != nil
}

func (obj *isisLspExtendedV4Prefix) SetRedistributionType(value IsisLspExtendedV4PrefixRedistributionTypeEnum) IsisLspExtendedV4Prefix {
	intValue, ok := otg.IsisLspExtendedV4Prefix_RedistributionType_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspExtendedV4PrefixRedistributionTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspExtendedV4Prefix_RedistributionType_Enum(intValue)
	obj.obj.RedistributionType = &enumValue

	return obj
}

// description is TBD
// PrefixAttributes returns a IsisLspPrefixAttributes
func (obj *isisLspExtendedV4Prefix) PrefixAttributes() IsisLspPrefixAttributes {
	if obj.obj.PrefixAttributes == nil {
		obj.obj.PrefixAttributes = NewIsisLspPrefixAttributes().msg()
	}
	if obj.prefixAttributesHolder == nil {
		obj.prefixAttributesHolder = &isisLspPrefixAttributes{obj: obj.obj.PrefixAttributes}
	}
	return obj.prefixAttributesHolder
}

// description is TBD
// PrefixAttributes returns a IsisLspPrefixAttributes
func (obj *isisLspExtendedV4Prefix) HasPrefixAttributes() bool {
	return obj.obj.PrefixAttributes != nil
}

// description is TBD
// SetPrefixAttributes sets the IsisLspPrefixAttributes value in the IsisLspExtendedV4Prefix object
func (obj *isisLspExtendedV4Prefix) SetPrefixAttributes(value IsisLspPrefixAttributes) IsisLspExtendedV4Prefix {

	obj.prefixAttributesHolder = nil
	obj.obj.PrefixAttributes = value.msg()

	return obj
}

func (obj *isisLspExtendedV4Prefix) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Ipv4Address != nil {

		err := obj.validateIpv4(obj.Ipv4Address())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspExtendedV4Prefix.Ipv4Address"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspExtendedV4Prefix.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if obj.obj.PrefixAttributes != nil {

		obj.PrefixAttributes().validateObj(vObj, set_default)
	}

}

func (obj *isisLspExtendedV4Prefix) setDefault() {

}
