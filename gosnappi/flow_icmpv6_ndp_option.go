package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIcmpv6NdpOption *****
type flowIcmpv6NdpOption struct {
	validation
	obj                          *otg.FlowIcmpv6NdpOption
	marshaller                   marshalFlowIcmpv6NdpOption
	unMarshaller                 unMarshalFlowIcmpv6NdpOption
	sourceLinkLayerAddressHolder FlowIcmpv6NdpLinkLayerAddressOption
	targetLinkLayerAddressHolder FlowIcmpv6NdpLinkLayerAddressOption
	prefixInformationHolder      FlowIcmpv6NdpPrefixInformationOption
	redirectedHeaderHolder       FlowIcmpv6NdpRedirectedHeaderOption
	mtuHolder                    FlowIcmpv6NdpMtuOption
}

func NewFlowIcmpv6NdpOption() FlowIcmpv6NdpOption {
	obj := flowIcmpv6NdpOption{obj: &otg.FlowIcmpv6NdpOption{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIcmpv6NdpOption) msg() *otg.FlowIcmpv6NdpOption {
	return obj.obj
}

func (obj *flowIcmpv6NdpOption) setMsg(msg *otg.FlowIcmpv6NdpOption) FlowIcmpv6NdpOption {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIcmpv6NdpOption struct {
	obj *flowIcmpv6NdpOption
}

type marshalFlowIcmpv6NdpOption interface {
	// ToProto marshals FlowIcmpv6NdpOption to protobuf object *otg.FlowIcmpv6NdpOption
	ToProto() (*otg.FlowIcmpv6NdpOption, error)
	// ToPbText marshals FlowIcmpv6NdpOption to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIcmpv6NdpOption to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIcmpv6NdpOption to JSON text
	ToJson() (string, error)
}

type unMarshalflowIcmpv6NdpOption struct {
	obj *flowIcmpv6NdpOption
}

type unMarshalFlowIcmpv6NdpOption interface {
	// FromProto unmarshals FlowIcmpv6NdpOption from protobuf object *otg.FlowIcmpv6NdpOption
	FromProto(msg *otg.FlowIcmpv6NdpOption) (FlowIcmpv6NdpOption, error)
	// FromPbText unmarshals FlowIcmpv6NdpOption from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIcmpv6NdpOption from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIcmpv6NdpOption from JSON text
	FromJson(value string) error
}

func (obj *flowIcmpv6NdpOption) Marshal() marshalFlowIcmpv6NdpOption {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIcmpv6NdpOption{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIcmpv6NdpOption) Unmarshal() unMarshalFlowIcmpv6NdpOption {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIcmpv6NdpOption{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIcmpv6NdpOption) ToProto() (*otg.FlowIcmpv6NdpOption, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIcmpv6NdpOption) FromProto(msg *otg.FlowIcmpv6NdpOption) (FlowIcmpv6NdpOption, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIcmpv6NdpOption) ToPbText() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpOption) FromPbText(value string) error {
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

func (m *marshalflowIcmpv6NdpOption) ToYaml() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpOption) FromYaml(value string) error {
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

func (m *marshalflowIcmpv6NdpOption) ToJson() (string, error) {
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

func (m *unMarshalflowIcmpv6NdpOption) FromJson(value string) error {
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

func (obj *flowIcmpv6NdpOption) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpOption) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIcmpv6NdpOption) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIcmpv6NdpOption) Clone() (FlowIcmpv6NdpOption, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIcmpv6NdpOption()
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

func (obj *flowIcmpv6NdpOption) setNil() {
	obj.sourceLinkLayerAddressHolder = nil
	obj.targetLinkLayerAddressHolder = nil
	obj.prefixInformationHolder = nil
	obj.redirectedHeaderHolder = nil
	obj.mtuHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIcmpv6NdpOption is nDP option (RFC 4861 Section 4.6) carried in NDP messages. The option Length is in units of 8 octets including the type and length bytes themselves.
type FlowIcmpv6NdpOption interface {
	Validation
	// msg marshals FlowIcmpv6NdpOption to protobuf object *otg.FlowIcmpv6NdpOption
	// and doesn't set defaults
	msg() *otg.FlowIcmpv6NdpOption
	// setMsg unmarshals FlowIcmpv6NdpOption from protobuf object *otg.FlowIcmpv6NdpOption
	// and doesn't set defaults
	setMsg(*otg.FlowIcmpv6NdpOption) FlowIcmpv6NdpOption
	// provides marshal interface
	Marshal() marshalFlowIcmpv6NdpOption
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIcmpv6NdpOption
	// validate validates FlowIcmpv6NdpOption
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIcmpv6NdpOption, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowIcmpv6NdpOptionChoiceEnum, set in FlowIcmpv6NdpOption
	Choice() FlowIcmpv6NdpOptionChoiceEnum
	// setChoice assigns FlowIcmpv6NdpOptionChoiceEnum provided by user to FlowIcmpv6NdpOption
	setChoice(value FlowIcmpv6NdpOptionChoiceEnum) FlowIcmpv6NdpOption
	// HasChoice checks if Choice has been set in FlowIcmpv6NdpOption
	HasChoice() bool
	// SourceLinkLayerAddress returns FlowIcmpv6NdpLinkLayerAddressOption, set in FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpLinkLayerAddressOption is nDP Source Link-Layer Address (Type 1) or Target Link-Layer Address (Type 2) option (RFC 4861 Section 4.6.1). The option Type byte is set from the parent choice.
	SourceLinkLayerAddress() FlowIcmpv6NdpLinkLayerAddressOption
	// SetSourceLinkLayerAddress assigns FlowIcmpv6NdpLinkLayerAddressOption provided by user to FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpLinkLayerAddressOption is nDP Source Link-Layer Address (Type 1) or Target Link-Layer Address (Type 2) option (RFC 4861 Section 4.6.1). The option Type byte is set from the parent choice.
	SetSourceLinkLayerAddress(value FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpOption
	// HasSourceLinkLayerAddress checks if SourceLinkLayerAddress has been set in FlowIcmpv6NdpOption
	HasSourceLinkLayerAddress() bool
	// TargetLinkLayerAddress returns FlowIcmpv6NdpLinkLayerAddressOption, set in FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpLinkLayerAddressOption is nDP Source Link-Layer Address (Type 1) or Target Link-Layer Address (Type 2) option (RFC 4861 Section 4.6.1). The option Type byte is set from the parent choice.
	TargetLinkLayerAddress() FlowIcmpv6NdpLinkLayerAddressOption
	// SetTargetLinkLayerAddress assigns FlowIcmpv6NdpLinkLayerAddressOption provided by user to FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpLinkLayerAddressOption is nDP Source Link-Layer Address (Type 1) or Target Link-Layer Address (Type 2) option (RFC 4861 Section 4.6.1). The option Type byte is set from the parent choice.
	SetTargetLinkLayerAddress(value FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpOption
	// HasTargetLinkLayerAddress checks if TargetLinkLayerAddress has been set in FlowIcmpv6NdpOption
	HasTargetLinkLayerAddress() bool
	// PrefixInformation returns FlowIcmpv6NdpPrefixInformationOption, set in FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpPrefixInformationOption is nDP Prefix Information option (Type 3). Defined in RFC 4861 Section 4.6.2 and extended for Mobile IPv6 Home Agent by RFC 6275 Section 7.2 (R flag). Provides hosts with on-link prefixes and prefixes for address autoconfiguration.
	PrefixInformation() FlowIcmpv6NdpPrefixInformationOption
	// SetPrefixInformation assigns FlowIcmpv6NdpPrefixInformationOption provided by user to FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpPrefixInformationOption is nDP Prefix Information option (Type 3). Defined in RFC 4861 Section 4.6.2 and extended for Mobile IPv6 Home Agent by RFC 6275 Section 7.2 (R flag). Provides hosts with on-link prefixes and prefixes for address autoconfiguration.
	SetPrefixInformation(value FlowIcmpv6NdpPrefixInformationOption) FlowIcmpv6NdpOption
	// HasPrefixInformation checks if PrefixInformation has been set in FlowIcmpv6NdpOption
	HasPrefixInformation() bool
	// RedirectedHeader returns FlowIcmpv6NdpRedirectedHeaderOption, set in FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpRedirectedHeaderOption is nDP Redirected Header option (Type 4). Defined in RFC 4861 Section 4.6.3. Used in Redirect messages to include as much as possible of the IP packet that triggered the Redirect, without exceeding the minimum IPv6 MTU.
	RedirectedHeader() FlowIcmpv6NdpRedirectedHeaderOption
	// SetRedirectedHeader assigns FlowIcmpv6NdpRedirectedHeaderOption provided by user to FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpRedirectedHeaderOption is nDP Redirected Header option (Type 4). Defined in RFC 4861 Section 4.6.3. Used in Redirect messages to include as much as possible of the IP packet that triggered the Redirect, without exceeding the minimum IPv6 MTU.
	SetRedirectedHeader(value FlowIcmpv6NdpRedirectedHeaderOption) FlowIcmpv6NdpOption
	// HasRedirectedHeader checks if RedirectedHeader has been set in FlowIcmpv6NdpOption
	HasRedirectedHeader() bool
	// Mtu returns FlowIcmpv6NdpMtuOption, set in FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpMtuOption is nDP MTU option (Type 5). Defined in RFC 4861 Section 4.6.4. Used in Router Advertisement messages to ensure that all nodes on a link use the same MTU value in those cases where the link MTU is not well known.
	Mtu() FlowIcmpv6NdpMtuOption
	// SetMtu assigns FlowIcmpv6NdpMtuOption provided by user to FlowIcmpv6NdpOption.
	// FlowIcmpv6NdpMtuOption is nDP MTU option (Type 5). Defined in RFC 4861 Section 4.6.4. Used in Router Advertisement messages to ensure that all nodes on a link use the same MTU value in those cases where the link MTU is not well known.
	SetMtu(value FlowIcmpv6NdpMtuOption) FlowIcmpv6NdpOption
	// HasMtu checks if Mtu has been set in FlowIcmpv6NdpOption
	HasMtu() bool
	setNil()
}

type FlowIcmpv6NdpOptionChoiceEnum string

// Enum of Choice on FlowIcmpv6NdpOption
var FlowIcmpv6NdpOptionChoice = struct {
	SOURCE_LINK_LAYER_ADDRESS FlowIcmpv6NdpOptionChoiceEnum
	TARGET_LINK_LAYER_ADDRESS FlowIcmpv6NdpOptionChoiceEnum
	PREFIX_INFORMATION        FlowIcmpv6NdpOptionChoiceEnum
	REDIRECTED_HEADER         FlowIcmpv6NdpOptionChoiceEnum
	MTU                       FlowIcmpv6NdpOptionChoiceEnum
}{
	SOURCE_LINK_LAYER_ADDRESS: FlowIcmpv6NdpOptionChoiceEnum("source_link_layer_address"),
	TARGET_LINK_LAYER_ADDRESS: FlowIcmpv6NdpOptionChoiceEnum("target_link_layer_address"),
	PREFIX_INFORMATION:        FlowIcmpv6NdpOptionChoiceEnum("prefix_information"),
	REDIRECTED_HEADER:         FlowIcmpv6NdpOptionChoiceEnum("redirected_header"),
	MTU:                       FlowIcmpv6NdpOptionChoiceEnum("mtu"),
}

func (obj *flowIcmpv6NdpOption) Choice() FlowIcmpv6NdpOptionChoiceEnum {
	return FlowIcmpv6NdpOptionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowIcmpv6NdpOption) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowIcmpv6NdpOption) setChoice(value FlowIcmpv6NdpOptionChoiceEnum) FlowIcmpv6NdpOption {
	intValue, ok := otg.FlowIcmpv6NdpOption_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowIcmpv6NdpOptionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowIcmpv6NdpOption_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Mtu = nil
	obj.mtuHolder = nil
	obj.obj.RedirectedHeader = nil
	obj.redirectedHeaderHolder = nil
	obj.obj.PrefixInformation = nil
	obj.prefixInformationHolder = nil
	obj.obj.TargetLinkLayerAddress = nil
	obj.targetLinkLayerAddressHolder = nil
	obj.obj.SourceLinkLayerAddress = nil
	obj.sourceLinkLayerAddressHolder = nil

	if value == FlowIcmpv6NdpOptionChoice.SOURCE_LINK_LAYER_ADDRESS {
		obj.obj.SourceLinkLayerAddress = NewFlowIcmpv6NdpLinkLayerAddressOption().msg()
	}

	if value == FlowIcmpv6NdpOptionChoice.TARGET_LINK_LAYER_ADDRESS {
		obj.obj.TargetLinkLayerAddress = NewFlowIcmpv6NdpLinkLayerAddressOption().msg()
	}

	if value == FlowIcmpv6NdpOptionChoice.PREFIX_INFORMATION {
		obj.obj.PrefixInformation = NewFlowIcmpv6NdpPrefixInformationOption().msg()
	}

	if value == FlowIcmpv6NdpOptionChoice.REDIRECTED_HEADER {
		obj.obj.RedirectedHeader = NewFlowIcmpv6NdpRedirectedHeaderOption().msg()
	}

	if value == FlowIcmpv6NdpOptionChoice.MTU {
		obj.obj.Mtu = NewFlowIcmpv6NdpMtuOption().msg()
	}

	return obj
}

// description is TBD
// SourceLinkLayerAddress returns a FlowIcmpv6NdpLinkLayerAddressOption
func (obj *flowIcmpv6NdpOption) SourceLinkLayerAddress() FlowIcmpv6NdpLinkLayerAddressOption {
	if obj.obj.SourceLinkLayerAddress == nil {
		obj.setChoice(FlowIcmpv6NdpOptionChoice.SOURCE_LINK_LAYER_ADDRESS)
	}
	if obj.sourceLinkLayerAddressHolder == nil {
		obj.sourceLinkLayerAddressHolder = &flowIcmpv6NdpLinkLayerAddressOption{obj: obj.obj.SourceLinkLayerAddress}
	}
	return obj.sourceLinkLayerAddressHolder
}

// description is TBD
// SourceLinkLayerAddress returns a FlowIcmpv6NdpLinkLayerAddressOption
func (obj *flowIcmpv6NdpOption) HasSourceLinkLayerAddress() bool {
	return obj.obj.SourceLinkLayerAddress != nil
}

// description is TBD
// SetSourceLinkLayerAddress sets the FlowIcmpv6NdpLinkLayerAddressOption value in the FlowIcmpv6NdpOption object
func (obj *flowIcmpv6NdpOption) SetSourceLinkLayerAddress(value FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpOption {
	obj.setChoice(FlowIcmpv6NdpOptionChoice.SOURCE_LINK_LAYER_ADDRESS)
	obj.sourceLinkLayerAddressHolder = nil
	obj.obj.SourceLinkLayerAddress = value.msg()

	return obj
}

// description is TBD
// TargetLinkLayerAddress returns a FlowIcmpv6NdpLinkLayerAddressOption
func (obj *flowIcmpv6NdpOption) TargetLinkLayerAddress() FlowIcmpv6NdpLinkLayerAddressOption {
	if obj.obj.TargetLinkLayerAddress == nil {
		obj.setChoice(FlowIcmpv6NdpOptionChoice.TARGET_LINK_LAYER_ADDRESS)
	}
	if obj.targetLinkLayerAddressHolder == nil {
		obj.targetLinkLayerAddressHolder = &flowIcmpv6NdpLinkLayerAddressOption{obj: obj.obj.TargetLinkLayerAddress}
	}
	return obj.targetLinkLayerAddressHolder
}

// description is TBD
// TargetLinkLayerAddress returns a FlowIcmpv6NdpLinkLayerAddressOption
func (obj *flowIcmpv6NdpOption) HasTargetLinkLayerAddress() bool {
	return obj.obj.TargetLinkLayerAddress != nil
}

// description is TBD
// SetTargetLinkLayerAddress sets the FlowIcmpv6NdpLinkLayerAddressOption value in the FlowIcmpv6NdpOption object
func (obj *flowIcmpv6NdpOption) SetTargetLinkLayerAddress(value FlowIcmpv6NdpLinkLayerAddressOption) FlowIcmpv6NdpOption {
	obj.setChoice(FlowIcmpv6NdpOptionChoice.TARGET_LINK_LAYER_ADDRESS)
	obj.targetLinkLayerAddressHolder = nil
	obj.obj.TargetLinkLayerAddress = value.msg()

	return obj
}

// description is TBD
// PrefixInformation returns a FlowIcmpv6NdpPrefixInformationOption
func (obj *flowIcmpv6NdpOption) PrefixInformation() FlowIcmpv6NdpPrefixInformationOption {
	if obj.obj.PrefixInformation == nil {
		obj.setChoice(FlowIcmpv6NdpOptionChoice.PREFIX_INFORMATION)
	}
	if obj.prefixInformationHolder == nil {
		obj.prefixInformationHolder = &flowIcmpv6NdpPrefixInformationOption{obj: obj.obj.PrefixInformation}
	}
	return obj.prefixInformationHolder
}

// description is TBD
// PrefixInformation returns a FlowIcmpv6NdpPrefixInformationOption
func (obj *flowIcmpv6NdpOption) HasPrefixInformation() bool {
	return obj.obj.PrefixInformation != nil
}

// description is TBD
// SetPrefixInformation sets the FlowIcmpv6NdpPrefixInformationOption value in the FlowIcmpv6NdpOption object
func (obj *flowIcmpv6NdpOption) SetPrefixInformation(value FlowIcmpv6NdpPrefixInformationOption) FlowIcmpv6NdpOption {
	obj.setChoice(FlowIcmpv6NdpOptionChoice.PREFIX_INFORMATION)
	obj.prefixInformationHolder = nil
	obj.obj.PrefixInformation = value.msg()

	return obj
}

// description is TBD
// RedirectedHeader returns a FlowIcmpv6NdpRedirectedHeaderOption
func (obj *flowIcmpv6NdpOption) RedirectedHeader() FlowIcmpv6NdpRedirectedHeaderOption {
	if obj.obj.RedirectedHeader == nil {
		obj.setChoice(FlowIcmpv6NdpOptionChoice.REDIRECTED_HEADER)
	}
	if obj.redirectedHeaderHolder == nil {
		obj.redirectedHeaderHolder = &flowIcmpv6NdpRedirectedHeaderOption{obj: obj.obj.RedirectedHeader}
	}
	return obj.redirectedHeaderHolder
}

// description is TBD
// RedirectedHeader returns a FlowIcmpv6NdpRedirectedHeaderOption
func (obj *flowIcmpv6NdpOption) HasRedirectedHeader() bool {
	return obj.obj.RedirectedHeader != nil
}

// description is TBD
// SetRedirectedHeader sets the FlowIcmpv6NdpRedirectedHeaderOption value in the FlowIcmpv6NdpOption object
func (obj *flowIcmpv6NdpOption) SetRedirectedHeader(value FlowIcmpv6NdpRedirectedHeaderOption) FlowIcmpv6NdpOption {
	obj.setChoice(FlowIcmpv6NdpOptionChoice.REDIRECTED_HEADER)
	obj.redirectedHeaderHolder = nil
	obj.obj.RedirectedHeader = value.msg()

	return obj
}

// description is TBD
// Mtu returns a FlowIcmpv6NdpMtuOption
func (obj *flowIcmpv6NdpOption) Mtu() FlowIcmpv6NdpMtuOption {
	if obj.obj.Mtu == nil {
		obj.setChoice(FlowIcmpv6NdpOptionChoice.MTU)
	}
	if obj.mtuHolder == nil {
		obj.mtuHolder = &flowIcmpv6NdpMtuOption{obj: obj.obj.Mtu}
	}
	return obj.mtuHolder
}

// description is TBD
// Mtu returns a FlowIcmpv6NdpMtuOption
func (obj *flowIcmpv6NdpOption) HasMtu() bool {
	return obj.obj.Mtu != nil
}

// description is TBD
// SetMtu sets the FlowIcmpv6NdpMtuOption value in the FlowIcmpv6NdpOption object
func (obj *flowIcmpv6NdpOption) SetMtu(value FlowIcmpv6NdpMtuOption) FlowIcmpv6NdpOption {
	obj.setChoice(FlowIcmpv6NdpOptionChoice.MTU)
	obj.mtuHolder = nil
	obj.obj.Mtu = value.msg()

	return obj
}

func (obj *flowIcmpv6NdpOption) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SourceLinkLayerAddress != nil {

		obj.SourceLinkLayerAddress().validateObj(vObj, set_default)
	}

	if obj.obj.TargetLinkLayerAddress != nil {

		obj.TargetLinkLayerAddress().validateObj(vObj, set_default)
	}

	if obj.obj.PrefixInformation != nil {

		obj.PrefixInformation().validateObj(vObj, set_default)
	}

	if obj.obj.RedirectedHeader != nil {

		obj.RedirectedHeader().validateObj(vObj, set_default)
	}

	if obj.obj.Mtu != nil {

		obj.Mtu().validateObj(vObj, set_default)
	}

}

func (obj *flowIcmpv6NdpOption) setDefault() {
	var choices_set int = 0
	var choice FlowIcmpv6NdpOptionChoiceEnum

	if obj.obj.SourceLinkLayerAddress != nil {
		choices_set += 1
		choice = FlowIcmpv6NdpOptionChoice.SOURCE_LINK_LAYER_ADDRESS
	}

	if obj.obj.TargetLinkLayerAddress != nil {
		choices_set += 1
		choice = FlowIcmpv6NdpOptionChoice.TARGET_LINK_LAYER_ADDRESS
	}

	if obj.obj.PrefixInformation != nil {
		choices_set += 1
		choice = FlowIcmpv6NdpOptionChoice.PREFIX_INFORMATION
	}

	if obj.obj.RedirectedHeader != nil {
		choices_set += 1
		choice = FlowIcmpv6NdpOptionChoice.REDIRECTED_HEADER
	}

	if obj.obj.Mtu != nil {
		choices_set += 1
		choice = FlowIcmpv6NdpOptionChoice.MTU
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowIcmpv6NdpOptionChoice.SOURCE_LINK_LAYER_ADDRESS)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowIcmpv6NdpOption")
			}
		} else {
			intVal := otg.FlowIcmpv6NdpOption_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowIcmpv6NdpOption_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
