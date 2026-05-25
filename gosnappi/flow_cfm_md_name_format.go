package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmMDNameFormat *****
type flowCfmMDNameFormat struct {
	validation
	obj           *otg.FlowCfmMDNameFormat
	marshaller    marshalFlowCfmMDNameFormat
	unMarshaller  unMarshalFlowCfmMDNameFormat
	lengthHolder  FlowCfmLength
	macAddrHolder FlowCfmMDNameMacAddr
}

func NewFlowCfmMDNameFormat() FlowCfmMDNameFormat {
	obj := flowCfmMDNameFormat{obj: &otg.FlowCfmMDNameFormat{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmMDNameFormat) msg() *otg.FlowCfmMDNameFormat {
	return obj.obj
}

func (obj *flowCfmMDNameFormat) setMsg(msg *otg.FlowCfmMDNameFormat) FlowCfmMDNameFormat {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmMDNameFormat struct {
	obj *flowCfmMDNameFormat
}

type marshalFlowCfmMDNameFormat interface {
	// ToProto marshals FlowCfmMDNameFormat to protobuf object *otg.FlowCfmMDNameFormat
	ToProto() (*otg.FlowCfmMDNameFormat, error)
	// ToPbText marshals FlowCfmMDNameFormat to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmMDNameFormat to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmMDNameFormat to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmMDNameFormat struct {
	obj *flowCfmMDNameFormat
}

type unMarshalFlowCfmMDNameFormat interface {
	// FromProto unmarshals FlowCfmMDNameFormat from protobuf object *otg.FlowCfmMDNameFormat
	FromProto(msg *otg.FlowCfmMDNameFormat) (FlowCfmMDNameFormat, error)
	// FromPbText unmarshals FlowCfmMDNameFormat from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmMDNameFormat from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmMDNameFormat from JSON text
	FromJson(value string) error
}

func (obj *flowCfmMDNameFormat) Marshal() marshalFlowCfmMDNameFormat {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmMDNameFormat{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmMDNameFormat) Unmarshal() unMarshalFlowCfmMDNameFormat {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmMDNameFormat{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmMDNameFormat) ToProto() (*otg.FlowCfmMDNameFormat, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmMDNameFormat) FromProto(msg *otg.FlowCfmMDNameFormat) (FlowCfmMDNameFormat, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmMDNameFormat) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmMDNameFormat) FromPbText(value string) error {
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

func (m *marshalflowCfmMDNameFormat) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmMDNameFormat) FromYaml(value string) error {
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

func (m *marshalflowCfmMDNameFormat) ToJson() (string, error) {
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

func (m *unMarshalflowCfmMDNameFormat) FromJson(value string) error {
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

func (obj *flowCfmMDNameFormat) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmMDNameFormat) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmMDNameFormat) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmMDNameFormat) Clone() (FlowCfmMDNameFormat, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmMDNameFormat()
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

func (obj *flowCfmMDNameFormat) setNil() {
	obj.lengthHolder = nil
	obj.macAddrHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmMDNameFormat is maintenance domain name encoding. The choice selects the MD Name Format as defined in IEEE 802.1Q-2018 Table 21-19 and provides a typed input for each format. Length is auto-computed from the value unless overridden.
type FlowCfmMDNameFormat interface {
	Validation
	// msg marshals FlowCfmMDNameFormat to protobuf object *otg.FlowCfmMDNameFormat
	// and doesn't set defaults
	msg() *otg.FlowCfmMDNameFormat
	// setMsg unmarshals FlowCfmMDNameFormat from protobuf object *otg.FlowCfmMDNameFormat
	// and doesn't set defaults
	setMsg(*otg.FlowCfmMDNameFormat) FlowCfmMDNameFormat
	// provides marshal interface
	Marshal() marshalFlowCfmMDNameFormat
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmMDNameFormat
	// validate validates FlowCfmMDNameFormat
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmMDNameFormat, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmMDNameFormatChoiceEnum, set in FlowCfmMDNameFormat
	Choice() FlowCfmMDNameFormatChoiceEnum
	// setChoice assigns FlowCfmMDNameFormatChoiceEnum provided by user to FlowCfmMDNameFormat
	setChoice(value FlowCfmMDNameFormatChoiceEnum) FlowCfmMDNameFormat
	// HasChoice checks if Choice has been set in FlowCfmMDNameFormat
	HasChoice() bool
	// Length returns FlowCfmLength, set in FlowCfmMDNameFormat.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	Length() FlowCfmLength
	// SetLength assigns FlowCfmLength provided by user to FlowCfmMDNameFormat.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	SetLength(value FlowCfmLength) FlowCfmMDNameFormat
	// HasLength checks if Length has been set in FlowCfmMDNameFormat
	HasLength() bool
	// DomainNameStr returns string, set in FlowCfmMDNameFormat.
	DomainNameStr() string
	// SetDomainNameStr assigns string provided by user to FlowCfmMDNameFormat
	SetDomainNameStr(value string) FlowCfmMDNameFormat
	// HasDomainNameStr checks if DomainNameStr has been set in FlowCfmMDNameFormat
	HasDomainNameStr() bool
	// MacAddr returns FlowCfmMDNameMacAddr, set in FlowCfmMDNameFormat.
	// FlowCfmMDNameMacAddr is mD Name Format 3 encoding as defined in IEEE 802.1Q-2018 Table 21-19. Consists of a 6-octet MAC address followed by a 2-octet unsigned integer in network byte order. Total wire length is always 8 octets.
	MacAddr() FlowCfmMDNameMacAddr
	// SetMacAddr assigns FlowCfmMDNameMacAddr provided by user to FlowCfmMDNameFormat.
	// FlowCfmMDNameMacAddr is mD Name Format 3 encoding as defined in IEEE 802.1Q-2018 Table 21-19. Consists of a 6-octet MAC address followed by a 2-octet unsigned integer in network byte order. Total wire length is always 8 octets.
	SetMacAddr(value FlowCfmMDNameMacAddr) FlowCfmMDNameFormat
	// HasMacAddr checks if MacAddr has been set in FlowCfmMDNameFormat
	HasMacAddr() bool
	// CharStr returns string, set in FlowCfmMDNameFormat.
	CharStr() string
	// SetCharStr assigns string provided by user to FlowCfmMDNameFormat
	SetCharStr(value string) FlowCfmMDNameFormat
	// HasCharStr checks if CharStr has been set in FlowCfmMDNameFormat
	HasCharStr() bool
	setNil()
}

type FlowCfmMDNameFormatChoiceEnum string

// Enum of Choice on FlowCfmMDNameFormat
var FlowCfmMDNameFormatChoice = struct {
	DOMAIN_NAME_STR FlowCfmMDNameFormatChoiceEnum
	MAC_ADDR        FlowCfmMDNameFormatChoiceEnum
	CHAR_STR        FlowCfmMDNameFormatChoiceEnum
}{
	DOMAIN_NAME_STR: FlowCfmMDNameFormatChoiceEnum("domain_name_str"),
	MAC_ADDR:        FlowCfmMDNameFormatChoiceEnum("mac_addr"),
	CHAR_STR:        FlowCfmMDNameFormatChoiceEnum("char_str"),
}

func (obj *flowCfmMDNameFormat) Choice() FlowCfmMDNameFormatChoiceEnum {
	return FlowCfmMDNameFormatChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowCfmMDNameFormat) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmMDNameFormat) setChoice(value FlowCfmMDNameFormatChoiceEnum) FlowCfmMDNameFormat {
	intValue, ok := otg.FlowCfmMDNameFormat_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmMDNameFormatChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmMDNameFormat_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CharStr = nil
	obj.obj.MacAddr = nil
	obj.macAddrHolder = nil
	obj.obj.DomainNameStr = nil

	if value == FlowCfmMDNameFormatChoice.DOMAIN_NAME_STR {
		defaultValue := ""
		obj.obj.DomainNameStr = &defaultValue
	}

	if value == FlowCfmMDNameFormatChoice.MAC_ADDR {
		obj.obj.MacAddr = NewFlowCfmMDNameMacAddr().msg()
	}

	if value == FlowCfmMDNameFormatChoice.CHAR_STR {
		defaultValue := ""
		obj.obj.CharStr = &defaultValue
	}

	return obj
}

// Length of the MD Name value in octets. Auto-computed by default.
// Length returns a FlowCfmLength
func (obj *flowCfmMDNameFormat) Length() FlowCfmLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowCfmLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowCfmLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// Length of the MD Name value in octets. Auto-computed by default.
// Length returns a FlowCfmLength
func (obj *flowCfmMDNameFormat) HasLength() bool {
	return obj.obj.Length != nil
}

// Length of the MD Name value in octets. Auto-computed by default.
// SetLength sets the FlowCfmLength value in the FlowCfmMDNameFormat object
func (obj *flowCfmMDNameFormat) SetLength(value FlowCfmLength) FlowCfmMDNameFormat {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// DNS-like domain name, e.g. "example.com" (RFC 1034 encoding). Maximum 43 octets (IEEE 802.1Q-2018 Table 21-19).
// DomainNameStr returns a string
func (obj *flowCfmMDNameFormat) DomainNameStr() string {

	if obj.obj.DomainNameStr == nil {
		obj.setChoice(FlowCfmMDNameFormatChoice.DOMAIN_NAME_STR)
	}

	return *obj.obj.DomainNameStr

}

// DNS-like domain name, e.g. "example.com" (RFC 1034 encoding). Maximum 43 octets (IEEE 802.1Q-2018 Table 21-19).
// DomainNameStr returns a string
func (obj *flowCfmMDNameFormat) HasDomainNameStr() bool {
	return obj.obj.DomainNameStr != nil
}

// DNS-like domain name, e.g. "example.com" (RFC 1034 encoding). Maximum 43 octets (IEEE 802.1Q-2018 Table 21-19).
// SetDomainNameStr sets the string value in the FlowCfmMDNameFormat object
func (obj *flowCfmMDNameFormat) SetDomainNameStr(value string) FlowCfmMDNameFormat {
	obj.setChoice(FlowCfmMDNameFormatChoice.DOMAIN_NAME_STR)
	obj.obj.DomainNameStr = &value
	return obj
}

// MAC address + 2-octet integer (IEEE Format 3). See Flow.Cfm.MDNameMacAddr.
// MacAddr returns a FlowCfmMDNameMacAddr
func (obj *flowCfmMDNameFormat) MacAddr() FlowCfmMDNameMacAddr {
	if obj.obj.MacAddr == nil {
		obj.setChoice(FlowCfmMDNameFormatChoice.MAC_ADDR)
	}
	if obj.macAddrHolder == nil {
		obj.macAddrHolder = &flowCfmMDNameMacAddr{obj: obj.obj.MacAddr}
	}
	return obj.macAddrHolder
}

// MAC address + 2-octet integer (IEEE Format 3). See Flow.Cfm.MDNameMacAddr.
// MacAddr returns a FlowCfmMDNameMacAddr
func (obj *flowCfmMDNameFormat) HasMacAddr() bool {
	return obj.obj.MacAddr != nil
}

// MAC address + 2-octet integer (IEEE Format 3). See Flow.Cfm.MDNameMacAddr.
// SetMacAddr sets the FlowCfmMDNameMacAddr value in the FlowCfmMDNameFormat object
func (obj *flowCfmMDNameFormat) SetMacAddr(value FlowCfmMDNameMacAddr) FlowCfmMDNameFormat {
	obj.setChoice(FlowCfmMDNameFormatChoice.MAC_ADDR)
	obj.macAddrHolder = nil
	obj.obj.MacAddr = value.msg()

	return obj
}

// Character string (ASCII), up to 43 octets (IEEE 802.1Q-2018 Table 21-19).
// CharStr returns a string
func (obj *flowCfmMDNameFormat) CharStr() string {

	if obj.obj.CharStr == nil {
		obj.setChoice(FlowCfmMDNameFormatChoice.CHAR_STR)
	}

	return *obj.obj.CharStr

}

// Character string (ASCII), up to 43 octets (IEEE 802.1Q-2018 Table 21-19).
// CharStr returns a string
func (obj *flowCfmMDNameFormat) HasCharStr() bool {
	return obj.obj.CharStr != nil
}

// Character string (ASCII), up to 43 octets (IEEE 802.1Q-2018 Table 21-19).
// SetCharStr sets the string value in the FlowCfmMDNameFormat object
func (obj *flowCfmMDNameFormat) SetCharStr(value string) FlowCfmMDNameFormat {
	obj.setChoice(FlowCfmMDNameFormatChoice.CHAR_STR)
	obj.obj.CharStr = &value
	return obj
}

func (obj *flowCfmMDNameFormat) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.MacAddr != nil {

		obj.MacAddr().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmMDNameFormat) setDefault() {
	var choices_set int = 0
	var choice FlowCfmMDNameFormatChoiceEnum

	if obj.obj.DomainNameStr != nil {
		choices_set += 1
		choice = FlowCfmMDNameFormatChoice.DOMAIN_NAME_STR
	}

	if obj.obj.MacAddr != nil {
		choices_set += 1
		choice = FlowCfmMDNameFormatChoice.MAC_ADDR
	}

	if obj.obj.CharStr != nil {
		choices_set += 1
		choice = FlowCfmMDNameFormatChoice.CHAR_STR
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmMDNameFormatChoice.CHAR_STR)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmMDNameFormat")
			}
		} else {
			intVal := otg.FlowCfmMDNameFormat_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmMDNameFormat_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
