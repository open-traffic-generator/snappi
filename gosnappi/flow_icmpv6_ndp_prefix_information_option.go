package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NdpPrefixInformationOption *****
type flowIcmpv6NdpPrefixInformationOption struct {
	validation
	obj                     *otg.FlowIcmpv6NdpPrefixInformationOption
	marshaller              marshalFlowIcmpv6NdpPrefixInformationOption
	unMarshaller            unMarshalFlowIcmpv6NdpPrefixInformationOption
	lengthHolder            PatternFlowIcmpv6NdpPrefixInformationOptionLength
	prefixLengthHolder      PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	lFlagHolder             PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	aFlagHolder             PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	rFlagHolder             PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	reserved1Holder         PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	validLifetimeHolder     PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	preferredLifetimeHolder PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	reserved2Holder         PatternFlowIcmpv6NdpPrefixInformationOptionReserved2
	prefixHolder            PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
}

func NewFlowIcmpv6NdpPrefixInformationOption() FlowIcmpv6NdpPrefixInformationOption {
	obj := flowIcmpv6NdpPrefixInformationOption{obj: &otg.FlowIcmpv6NdpPrefixInformationOption{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NdpPrefixInformationOption) msg() *otg.FlowIcmpv6NdpPrefixInformationOption {
	return obj.obj
}

func (obj *flowIcmpv6NdpPrefixInformationOption) setMsg(msg *otg.FlowIcmpv6NdpPrefixInformationOption) FlowIcmpv6NdpPrefixInformationOption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NdpPrefixInformationOption struct {
	obj *flowIcmpv6NdpPrefixInformationOption
}

type marshalFlowIcmpv6NdpPrefixInformationOption interface {
	// ToProto marshals FlowIcmpv6NdpPrefixInformationOption to protobuf object *otg.FlowIcmpv6NdpPrefixInformationOption
	ToProto() (*otg.FlowIcmpv6NdpPrefixInformationOption, error)
	// ToPbText marshals FlowIcmpv6NdpPrefixInformationOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NdpPrefixInformationOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NdpPrefixInformationOption to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NdpPrefixInformationOption struct {
	obj *flowIcmpv6NdpPrefixInformationOption
}

type unMarshalFlowIcmpv6NdpPrefixInformationOption interface {
	// FromProto unmarshals FlowIcmpv6NdpPrefixInformationOption from protobuf object *otg.FlowIcmpv6NdpPrefixInformationOption
	FromProto(msg *otg.FlowIcmpv6NdpPrefixInformationOption) (FlowIcmpv6NdpPrefixInformationOption, error)
	// FromPbText unmarshals FlowIcmpv6NdpPrefixInformationOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NdpPrefixInformationOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NdpPrefixInformationOption from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NdpPrefixInformationOption) Marshal() marshalFlowIcmpv6NdpPrefixInformationOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NdpPrefixInformationOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NdpPrefixInformationOption) Unmarshal() unMarshalFlowIcmpv6NdpPrefixInformationOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NdpPrefixInformationOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NdpPrefixInformationOption) ToProto() (*otg.FlowIcmpv6NdpPrefixInformationOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NdpPrefixInformationOption) FromProto(msg *otg.FlowIcmpv6NdpPrefixInformationOption) (FlowIcmpv6NdpPrefixInformationOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NdpPrefixInformationOption) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpPrefixInformationOption) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6NdpPrefixInformationOption) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpPrefixInformationOption) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6NdpPrefixInformationOption) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpPrefixInformationOption) FromJson(value string) error {
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

func (obj *flowIcmpv6NdpPrefixInformationOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpPrefixInformationOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpPrefixInformationOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NdpPrefixInformationOption) Clone() (FlowIcmpv6NdpPrefixInformationOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NdpPrefixInformationOption()
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

func (obj *flowIcmpv6NdpPrefixInformationOption) setNil() {
	obj.lengthHolder = nil
	obj.prefixLengthHolder = nil
	obj.lFlagHolder = nil
	obj.aFlagHolder = nil
	obj.rFlagHolder = nil
	obj.reserved1Holder = nil
	obj.validLifetimeHolder = nil
	obj.preferredLifetimeHolder = nil
	obj.reserved2Holder = nil
	obj.prefixHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NdpPrefixInformationOption is nDP Prefix Information option (Type 3). Defined in RFC 4861 Section 4.6.2 and extended for Mobile IPv6 Home Agent by RFC 6275 Section 7.2 (R flag). Provides hosts with on-link prefixes and prefixes for address autoconfiguration.
type FlowIcmpv6NdpPrefixInformationOption interface {
	Validation
	// msg marshals FlowIcmpv6NdpPrefixInformationOption to protobuf object *otg.FlowIcmpv6NdpPrefixInformationOption
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NdpPrefixInformationOption
	// setMsg unmarshals FlowIcmpv6NdpPrefixInformationOption from protobuf object *otg.FlowIcmpv6NdpPrefixInformationOption
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NdpPrefixInformationOption) FlowIcmpv6NdpPrefixInformationOption
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NdpPrefixInformationOption
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NdpPrefixInformationOption
	// validate validates FlowIcmpv6NdpPrefixInformationOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NdpPrefixInformationOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Length returns PatternFlowIcmpv6NdpPrefixInformationOptionLength, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLength is option length in units of 8 octets. Fixed at 4 for the Prefix Information option.
	Length() PatternFlowIcmpv6NdpPrefixInformationOptionLength
	// SetLength assigns PatternFlowIcmpv6NdpPrefixInformationOptionLength provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLength is option length in units of 8 octets. Fixed at 4 for the Prefix Information option.
	SetLength(value PatternFlowIcmpv6NdpPrefixInformationOptionLength) FlowIcmpv6NdpPrefixInformationOption
	// HasLength checks if Length has been set in FlowIcmpv6NdpPrefixInformationOption
	HasLength() bool
	// PrefixLength returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength is number of leading bits of the Prefix field that are valid. Range 0..128.
	PrefixLength() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
	// SetPrefixLength assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength is number of leading bits of the Prefix field that are valid. Range 0..128.
	SetPrefixLength(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FlowIcmpv6NdpPrefixInformationOption
	// HasPrefixLength checks if PrefixLength has been set in FlowIcmpv6NdpPrefixInformationOption
	HasPrefixLength() bool
	// LFlag returns PatternFlowIcmpv6NdpPrefixInformationOptionLFlag, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlag is on-link (L) flag. When set, the prefix can be used for on-link determination. RFC 4861 Section 4.6.2.
	LFlag() PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
	// SetLFlag assigns PatternFlowIcmpv6NdpPrefixInformationOptionLFlag provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionLFlag is on-link (L) flag. When set, the prefix can be used for on-link determination. RFC 4861 Section 4.6.2.
	SetLFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FlowIcmpv6NdpPrefixInformationOption
	// HasLFlag checks if LFlag has been set in FlowIcmpv6NdpPrefixInformationOption
	HasLFlag() bool
	// AFlag returns PatternFlowIcmpv6NdpPrefixInformationOptionAFlag, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlag is autonomous address-configuration (A) flag. When set, the prefix can be used for stateless address autoconfiguration. RFC 4861 Section 4.6.2.
	AFlag() PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
	// SetAFlag assigns PatternFlowIcmpv6NdpPrefixInformationOptionAFlag provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionAFlag is autonomous address-configuration (A) flag. When set, the prefix can be used for stateless address autoconfiguration. RFC 4861 Section 4.6.2.
	SetAFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FlowIcmpv6NdpPrefixInformationOption
	// HasAFlag checks if AFlag has been set in FlowIcmpv6NdpPrefixInformationOption
	HasAFlag() bool
	// RFlag returns PatternFlowIcmpv6NdpPrefixInformationOptionRFlag, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlag is router address (R) flag. When set, the Prefix field contains a complete IP address assigned to the sending router. Used by Mobile IPv6 (RFC 6275 Section 7.2).
	RFlag() PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
	// SetRFlag assigns PatternFlowIcmpv6NdpPrefixInformationOptionRFlag provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionRFlag is router address (R) flag. When set, the Prefix field contains a complete IP address assigned to the sending router. Used by Mobile IPv6 (RFC 6275 Section 7.2).
	SetRFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FlowIcmpv6NdpPrefixInformationOption
	// HasRFlag checks if RFlag has been set in FlowIcmpv6NdpPrefixInformationOption
	HasRFlag() bool
	// Reserved1 returns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 is reserved. Must be initialized to zero by the sender.
	Reserved1() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
	// SetReserved1 assigns PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 is reserved. Must be initialized to zero by the sender.
	SetReserved1(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FlowIcmpv6NdpPrefixInformationOption
	// HasReserved1 checks if Reserved1 has been set in FlowIcmpv6NdpPrefixInformationOption
	HasReserved1() bool
	// ValidLifetime returns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime is length of time in seconds that the prefix is valid for on-link determination. A value of 0xffffffff represents infinity.
	ValidLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
	// SetValidLifetime assigns PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime is length of time in seconds that the prefix is valid for on-link determination. A value of 0xffffffff represents infinity.
	SetValidLifetime(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FlowIcmpv6NdpPrefixInformationOption
	// HasValidLifetime checks if ValidLifetime has been set in FlowIcmpv6NdpPrefixInformationOption
	HasValidLifetime() bool
	// PreferredLifetime returns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime is length of time in seconds that addresses generated from the prefix via stateless address autoconfiguration remain preferred. Must be less than or equal to valid_lifetime.
	PreferredLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
	// SetPreferredLifetime assigns PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime is length of time in seconds that addresses generated from the prefix via stateless address autoconfiguration remain preferred. Must be less than or equal to valid_lifetime.
	SetPreferredLifetime(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FlowIcmpv6NdpPrefixInformationOption
	// HasPreferredLifetime checks if PreferredLifetime has been set in FlowIcmpv6NdpPrefixInformationOption
	HasPreferredLifetime() bool
	// Reserved2 returns PatternFlowIcmpv6NdpPrefixInformationOptionReserved2, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved2 is reserved. Must be initialized to zero by the sender.
	Reserved2() PatternFlowIcmpv6NdpPrefixInformationOptionReserved2
	// SetReserved2 assigns PatternFlowIcmpv6NdpPrefixInformationOptionReserved2 provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionReserved2 is reserved. Must be initialized to zero by the sender.
	SetReserved2(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved2) FlowIcmpv6NdpPrefixInformationOption
	// HasReserved2 checks if Reserved2 has been set in FlowIcmpv6NdpPrefixInformationOption
	HasReserved2() bool
	// Prefix returns PatternFlowIcmpv6NdpPrefixInformationOptionPrefix, set in FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefix is iP address or prefix of an IP address. The Prefix Length field contains the number of valid leading bits.
	Prefix() PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
	// SetPrefix assigns PatternFlowIcmpv6NdpPrefixInformationOptionPrefix provided by user to FlowIcmpv6NdpPrefixInformationOption.
	// PatternFlowIcmpv6NdpPrefixInformationOptionPrefix is iP address or prefix of an IP address. The Prefix Length field contains the number of valid leading bits.
	SetPrefix(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FlowIcmpv6NdpPrefixInformationOption
	// HasPrefix checks if Prefix has been set in FlowIcmpv6NdpPrefixInformationOption
	HasPrefix() bool
	setNil()
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpPrefixInformationOptionLength
func (obj *flowIcmpv6NdpPrefixInformationOption) Length() PatternFlowIcmpv6NdpPrefixInformationOptionLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowIcmpv6NdpPrefixInformationOptionLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowIcmpv6NdpPrefixInformationOptionLength
func (obj *flowIcmpv6NdpPrefixInformationOption) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowIcmpv6NdpPrefixInformationOptionLength value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetLength(value PatternFlowIcmpv6NdpPrefixInformationOptionLength) FlowIcmpv6NdpPrefixInformationOption {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// PrefixLength returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
func (obj *flowIcmpv6NdpPrefixInformationOption) PrefixLength() PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength {
	if obj.obj.PrefixLength == nil {
		obj.obj.PrefixLength = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength().msg()
	}
	if obj.prefixLengthHolder == nil {
		obj.prefixLengthHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefixLength{obj: obj.obj.PrefixLength}
	}
	return obj.prefixLengthHolder
}

// description is TBD
// PrefixLength returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength
func (obj *flowIcmpv6NdpPrefixInformationOption) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// description is TBD
// SetPrefixLength sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetPrefixLength(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefixLength) FlowIcmpv6NdpPrefixInformationOption {

	obj.prefixLengthHolder = nil
	obj.obj.PrefixLength = value.msg()

	return obj
}

// description is TBD
// LFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) LFlag() PatternFlowIcmpv6NdpPrefixInformationOptionLFlag {
	if obj.obj.LFlag == nil {
		obj.obj.LFlag = NewPatternFlowIcmpv6NdpPrefixInformationOptionLFlag().msg()
	}
	if obj.lFlagHolder == nil {
		obj.lFlagHolder = &patternFlowIcmpv6NdpPrefixInformationOptionLFlag{obj: obj.obj.LFlag}
	}
	return obj.lFlagHolder
}

// description is TBD
// LFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionLFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) HasLFlag() bool {
	return obj.obj.LFlag != nil
}

// description is TBD
// SetLFlag sets the PatternFlowIcmpv6NdpPrefixInformationOptionLFlag value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetLFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionLFlag) FlowIcmpv6NdpPrefixInformationOption {

	obj.lFlagHolder = nil
	obj.obj.LFlag = value.msg()

	return obj
}

// description is TBD
// AFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) AFlag() PatternFlowIcmpv6NdpPrefixInformationOptionAFlag {
	if obj.obj.AFlag == nil {
		obj.obj.AFlag = NewPatternFlowIcmpv6NdpPrefixInformationOptionAFlag().msg()
	}
	if obj.aFlagHolder == nil {
		obj.aFlagHolder = &patternFlowIcmpv6NdpPrefixInformationOptionAFlag{obj: obj.obj.AFlag}
	}
	return obj.aFlagHolder
}

// description is TBD
// AFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionAFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) HasAFlag() bool {
	return obj.obj.AFlag != nil
}

// description is TBD
// SetAFlag sets the PatternFlowIcmpv6NdpPrefixInformationOptionAFlag value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetAFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionAFlag) FlowIcmpv6NdpPrefixInformationOption {

	obj.aFlagHolder = nil
	obj.obj.AFlag = value.msg()

	return obj
}

// description is TBD
// RFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) RFlag() PatternFlowIcmpv6NdpPrefixInformationOptionRFlag {
	if obj.obj.RFlag == nil {
		obj.obj.RFlag = NewPatternFlowIcmpv6NdpPrefixInformationOptionRFlag().msg()
	}
	if obj.rFlagHolder == nil {
		obj.rFlagHolder = &patternFlowIcmpv6NdpPrefixInformationOptionRFlag{obj: obj.obj.RFlag}
	}
	return obj.rFlagHolder
}

// description is TBD
// RFlag returns a PatternFlowIcmpv6NdpPrefixInformationOptionRFlag
func (obj *flowIcmpv6NdpPrefixInformationOption) HasRFlag() bool {
	return obj.obj.RFlag != nil
}

// description is TBD
// SetRFlag sets the PatternFlowIcmpv6NdpPrefixInformationOptionRFlag value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetRFlag(value PatternFlowIcmpv6NdpPrefixInformationOptionRFlag) FlowIcmpv6NdpPrefixInformationOption {

	obj.rFlagHolder = nil
	obj.obj.RFlag = value.msg()

	return obj
}

// description is TBD
// Reserved1 returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
func (obj *flowIcmpv6NdpPrefixInformationOption) Reserved1() PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 {
	if obj.obj.Reserved1 == nil {
		obj.obj.Reserved1 = NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved1().msg()
	}
	if obj.reserved1Holder == nil {
		obj.reserved1Holder = &patternFlowIcmpv6NdpPrefixInformationOptionReserved1{obj: obj.obj.Reserved1}
	}
	return obj.reserved1Holder
}

// description is TBD
// Reserved1 returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved1
func (obj *flowIcmpv6NdpPrefixInformationOption) HasReserved1() bool {
	return obj.obj.Reserved1 != nil
}

// description is TBD
// SetReserved1 sets the PatternFlowIcmpv6NdpPrefixInformationOptionReserved1 value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetReserved1(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved1) FlowIcmpv6NdpPrefixInformationOption {

	obj.reserved1Holder = nil
	obj.obj.Reserved1 = value.msg()

	return obj
}

// description is TBD
// ValidLifetime returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
func (obj *flowIcmpv6NdpPrefixInformationOption) ValidLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime {
	if obj.obj.ValidLifetime == nil {
		obj.obj.ValidLifetime = NewPatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime().msg()
	}
	if obj.validLifetimeHolder == nil {
		obj.validLifetimeHolder = &patternFlowIcmpv6NdpPrefixInformationOptionValidLifetime{obj: obj.obj.ValidLifetime}
	}
	return obj.validLifetimeHolder
}

// description is TBD
// ValidLifetime returns a PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime
func (obj *flowIcmpv6NdpPrefixInformationOption) HasValidLifetime() bool {
	return obj.obj.ValidLifetime != nil
}

// description is TBD
// SetValidLifetime sets the PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetValidLifetime(value PatternFlowIcmpv6NdpPrefixInformationOptionValidLifetime) FlowIcmpv6NdpPrefixInformationOption {

	obj.validLifetimeHolder = nil
	obj.obj.ValidLifetime = value.msg()

	return obj
}

// description is TBD
// PreferredLifetime returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
func (obj *flowIcmpv6NdpPrefixInformationOption) PreferredLifetime() PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime {
	if obj.obj.PreferredLifetime == nil {
		obj.obj.PreferredLifetime = NewPatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime().msg()
	}
	if obj.preferredLifetimeHolder == nil {
		obj.preferredLifetimeHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime{obj: obj.obj.PreferredLifetime}
	}
	return obj.preferredLifetimeHolder
}

// description is TBD
// PreferredLifetime returns a PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime
func (obj *flowIcmpv6NdpPrefixInformationOption) HasPreferredLifetime() bool {
	return obj.obj.PreferredLifetime != nil
}

// description is TBD
// SetPreferredLifetime sets the PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetPreferredLifetime(value PatternFlowIcmpv6NdpPrefixInformationOptionPreferredLifetime) FlowIcmpv6NdpPrefixInformationOption {

	obj.preferredLifetimeHolder = nil
	obj.obj.PreferredLifetime = value.msg()

	return obj
}

// description is TBD
// Reserved2 returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved2
func (obj *flowIcmpv6NdpPrefixInformationOption) Reserved2() PatternFlowIcmpv6NdpPrefixInformationOptionReserved2 {
	if obj.obj.Reserved2 == nil {
		obj.obj.Reserved2 = NewPatternFlowIcmpv6NdpPrefixInformationOptionReserved2().msg()
	}
	if obj.reserved2Holder == nil {
		obj.reserved2Holder = &patternFlowIcmpv6NdpPrefixInformationOptionReserved2{obj: obj.obj.Reserved2}
	}
	return obj.reserved2Holder
}

// description is TBD
// Reserved2 returns a PatternFlowIcmpv6NdpPrefixInformationOptionReserved2
func (obj *flowIcmpv6NdpPrefixInformationOption) HasReserved2() bool {
	return obj.obj.Reserved2 != nil
}

// description is TBD
// SetReserved2 sets the PatternFlowIcmpv6NdpPrefixInformationOptionReserved2 value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetReserved2(value PatternFlowIcmpv6NdpPrefixInformationOptionReserved2) FlowIcmpv6NdpPrefixInformationOption {

	obj.reserved2Holder = nil
	obj.obj.Reserved2 = value.msg()

	return obj
}

// description is TBD
// Prefix returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
func (obj *flowIcmpv6NdpPrefixInformationOption) Prefix() PatternFlowIcmpv6NdpPrefixInformationOptionPrefix {
	if obj.obj.Prefix == nil {
		obj.obj.Prefix = NewPatternFlowIcmpv6NdpPrefixInformationOptionPrefix().msg()
	}
	if obj.prefixHolder == nil {
		obj.prefixHolder = &patternFlowIcmpv6NdpPrefixInformationOptionPrefix{obj: obj.obj.Prefix}
	}
	return obj.prefixHolder
}

// description is TBD
// Prefix returns a PatternFlowIcmpv6NdpPrefixInformationOptionPrefix
func (obj *flowIcmpv6NdpPrefixInformationOption) HasPrefix() bool {
	return obj.obj.Prefix != nil
}

// description is TBD
// SetPrefix sets the PatternFlowIcmpv6NdpPrefixInformationOptionPrefix value in the FlowIcmpv6NdpPrefixInformationOption object
func (obj *flowIcmpv6NdpPrefixInformationOption) SetPrefix(value PatternFlowIcmpv6NdpPrefixInformationOptionPrefix) FlowIcmpv6NdpPrefixInformationOption {

	obj.prefixHolder = nil
	obj.obj.Prefix = value.msg()

	return obj
}

func (obj *flowIcmpv6NdpPrefixInformationOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixLength != nil {

		obj.PrefixLength().validateObj(vObj, set_default)
	}

	if obj.obj.LFlag != nil {

		obj.LFlag().validateObj(vObj, set_default)
	}

	if obj.obj.AFlag != nil {

		obj.AFlag().validateObj(vObj, set_default)
	}

	if obj.obj.RFlag != nil {

		obj.RFlag().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved1 != nil {

		obj.Reserved1().validateObj(vObj, set_default)
	}

	if obj.obj.ValidLifetime != nil {

		obj.ValidLifetime().validateObj(vObj, set_default)
	}

	if obj.obj.PreferredLifetime != nil {

		obj.PreferredLifetime().validateObj(vObj, set_default)
	}

	if obj.obj.Reserved2 != nil {

		obj.Reserved2().validateObj(vObj, set_default)
	}

	if obj.obj.Prefix != nil {

		obj.Prefix().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6NdpPrefixInformationOption) setDefault() {

}
