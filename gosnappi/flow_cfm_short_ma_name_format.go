package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmShortMANameFormat *****
type flowCfmShortMANameFormat struct {
	validation
	obj          *otg.FlowCfmShortMANameFormat
	marshaller   marshalFlowCfmShortMANameFormat
	unMarshaller unMarshalFlowCfmShortMANameFormat
	lengthHolder FlowCfmLength
}

func NewFlowCfmShortMANameFormat() FlowCfmShortMANameFormat {
	obj := flowCfmShortMANameFormat{obj: &otg.FlowCfmShortMANameFormat{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmShortMANameFormat) msg() *otg.FlowCfmShortMANameFormat {
	return obj.obj
}

func (obj *flowCfmShortMANameFormat) setMsg(msg *otg.FlowCfmShortMANameFormat) FlowCfmShortMANameFormat {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmShortMANameFormat struct {
	obj *flowCfmShortMANameFormat
}

type marshalFlowCfmShortMANameFormat interface {
	// ToProto marshals FlowCfmShortMANameFormat to protobuf object *otg.FlowCfmShortMANameFormat
	ToProto() (*otg.FlowCfmShortMANameFormat, error)
	// ToPbText marshals FlowCfmShortMANameFormat to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmShortMANameFormat to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmShortMANameFormat to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmShortMANameFormat struct {
	obj *flowCfmShortMANameFormat
}

type unMarshalFlowCfmShortMANameFormat interface {
	// FromProto unmarshals FlowCfmShortMANameFormat from protobuf object *otg.FlowCfmShortMANameFormat
	FromProto(msg *otg.FlowCfmShortMANameFormat) (FlowCfmShortMANameFormat, error)
	// FromPbText unmarshals FlowCfmShortMANameFormat from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmShortMANameFormat from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmShortMANameFormat from JSON text
	FromJson(value string) error
}

func (obj *flowCfmShortMANameFormat) Marshal() marshalFlowCfmShortMANameFormat {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmShortMANameFormat{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmShortMANameFormat) Unmarshal() unMarshalFlowCfmShortMANameFormat {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmShortMANameFormat{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmShortMANameFormat) ToProto() (*otg.FlowCfmShortMANameFormat, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmShortMANameFormat) FromProto(msg *otg.FlowCfmShortMANameFormat) (FlowCfmShortMANameFormat, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmShortMANameFormat) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmShortMANameFormat) FromPbText(value string) error {
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

func (m *marshalflowCfmShortMANameFormat) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmShortMANameFormat) FromYaml(value string) error {
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

func (m *marshalflowCfmShortMANameFormat) ToJson() (string, error) {
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

func (m *unMarshalflowCfmShortMANameFormat) FromJson(value string) error {
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

func (obj *flowCfmShortMANameFormat) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmShortMANameFormat) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmShortMANameFormat) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmShortMANameFormat) Clone() (FlowCfmShortMANameFormat, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmShortMANameFormat()
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

func (obj *flowCfmShortMANameFormat) setNil() {
	obj.lengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmShortMANameFormat is short MA Name encoding (IEEE 802.1Q-2018 Table 21-20). Choose one of four formats and provide its typed value. Length is auto-computed from the value unless overridden.
type FlowCfmShortMANameFormat interface {
	Validation
	// msg marshals FlowCfmShortMANameFormat to protobuf object *otg.FlowCfmShortMANameFormat
	// and doesn't set defaults
	msg() *otg.FlowCfmShortMANameFormat
	// setMsg unmarshals FlowCfmShortMANameFormat from protobuf object *otg.FlowCfmShortMANameFormat
	// and doesn't set defaults
	setMsg(*otg.FlowCfmShortMANameFormat) FlowCfmShortMANameFormat
	// provides marshal interface
	Marshal() marshalFlowCfmShortMANameFormat
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmShortMANameFormat
	// validate validates FlowCfmShortMANameFormat
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmShortMANameFormat, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmShortMANameFormatChoiceEnum, set in FlowCfmShortMANameFormat
	Choice() FlowCfmShortMANameFormatChoiceEnum
	// setChoice assigns FlowCfmShortMANameFormatChoiceEnum provided by user to FlowCfmShortMANameFormat
	setChoice(value FlowCfmShortMANameFormatChoiceEnum) FlowCfmShortMANameFormat
	// HasChoice checks if Choice has been set in FlowCfmShortMANameFormat
	HasChoice() bool
	// Length returns FlowCfmLength, set in FlowCfmShortMANameFormat.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	Length() FlowCfmLength
	// SetLength assigns FlowCfmLength provided by user to FlowCfmShortMANameFormat.
	// FlowCfmLength is length field with auto-compute or explicit-value choice. When auto, the implementation derives the value from the associated data field.
	SetLength(value FlowCfmLength) FlowCfmShortMANameFormat
	// HasLength checks if Length has been set in FlowCfmShortMANameFormat
	HasLength() bool
	// PrimaryVid returns uint32, set in FlowCfmShortMANameFormat.
	PrimaryVid() uint32
	// SetPrimaryVid assigns uint32 provided by user to FlowCfmShortMANameFormat
	SetPrimaryVid(value uint32) FlowCfmShortMANameFormat
	// HasPrimaryVid checks if PrimaryVid has been set in FlowCfmShortMANameFormat
	HasPrimaryVid() bool
	// CharStr returns string, set in FlowCfmShortMANameFormat.
	CharStr() string
	// SetCharStr assigns string provided by user to FlowCfmShortMANameFormat
	SetCharStr(value string) FlowCfmShortMANameFormat
	// HasCharStr checks if CharStr has been set in FlowCfmShortMANameFormat
	HasCharStr() bool
	// TwoOctInt returns uint32, set in FlowCfmShortMANameFormat.
	TwoOctInt() uint32
	// SetTwoOctInt assigns uint32 provided by user to FlowCfmShortMANameFormat
	SetTwoOctInt(value uint32) FlowCfmShortMANameFormat
	// HasTwoOctInt checks if TwoOctInt has been set in FlowCfmShortMANameFormat
	HasTwoOctInt() bool
	// Rfc2685VpnId returns string, set in FlowCfmShortMANameFormat.
	Rfc2685VpnId() string
	// SetRfc2685VpnId assigns string provided by user to FlowCfmShortMANameFormat
	SetRfc2685VpnId(value string) FlowCfmShortMANameFormat
	// HasRfc2685VpnId checks if Rfc2685VpnId has been set in FlowCfmShortMANameFormat
	HasRfc2685VpnId() bool
	setNil()
}

type FlowCfmShortMANameFormatChoiceEnum string

// Enum of Choice on FlowCfmShortMANameFormat
var FlowCfmShortMANameFormatChoice = struct {
	PRIMARY_VID     FlowCfmShortMANameFormatChoiceEnum
	CHAR_STR        FlowCfmShortMANameFormatChoiceEnum
	TWO_OCT_INT     FlowCfmShortMANameFormatChoiceEnum
	RFC_2685_VPN_ID FlowCfmShortMANameFormatChoiceEnum
}{
	PRIMARY_VID:     FlowCfmShortMANameFormatChoiceEnum("primary_vid"),
	CHAR_STR:        FlowCfmShortMANameFormatChoiceEnum("char_str"),
	TWO_OCT_INT:     FlowCfmShortMANameFormatChoiceEnum("two_oct_int"),
	RFC_2685_VPN_ID: FlowCfmShortMANameFormatChoiceEnum("rfc_2685_vpn_id"),
}

func (obj *flowCfmShortMANameFormat) Choice() FlowCfmShortMANameFormatChoiceEnum {
	return FlowCfmShortMANameFormatChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *flowCfmShortMANameFormat) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmShortMANameFormat) setChoice(value FlowCfmShortMANameFormatChoiceEnum) FlowCfmShortMANameFormat {
	intValue, ok := otg.FlowCfmShortMANameFormat_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmShortMANameFormatChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmShortMANameFormat_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Rfc_2685VpnId = nil
	obj.obj.TwoOctInt = nil
	obj.obj.CharStr = nil
	obj.obj.PrimaryVid = nil

	if value == FlowCfmShortMANameFormatChoice.PRIMARY_VID {
		defaultValue := uint32(1)
		obj.obj.PrimaryVid = &defaultValue
	}

	if value == FlowCfmShortMANameFormatChoice.CHAR_STR {
		defaultValue := ""
		obj.obj.CharStr = &defaultValue
	}

	if value == FlowCfmShortMANameFormatChoice.TWO_OCT_INT {
		defaultValue := uint32(0)
		obj.obj.TwoOctInt = &defaultValue
	}

	if value == FlowCfmShortMANameFormatChoice.RFC_2685_VPN_ID {
		defaultValue := "00000000000000"
		obj.obj.Rfc_2685VpnId = &defaultValue
	}

	return obj
}

// Length of the Short MA Name value in octets. Auto-computed by default.
// Length returns a FlowCfmLength
func (obj *flowCfmShortMANameFormat) Length() FlowCfmLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewFlowCfmLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &flowCfmLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// Length of the Short MA Name value in octets. Auto-computed by default.
// Length returns a FlowCfmLength
func (obj *flowCfmShortMANameFormat) HasLength() bool {
	return obj.obj.Length != nil
}

// Length of the Short MA Name value in octets. Auto-computed by default.
// SetLength sets the FlowCfmLength value in the FlowCfmShortMANameFormat object
func (obj *flowCfmShortMANameFormat) SetLength(value FlowCfmLength) FlowCfmShortMANameFormat {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// Primary VLAN ID (1-4094). Encoded as a 2-octet unsigned integer.
// PrimaryVid returns a uint32
func (obj *flowCfmShortMANameFormat) PrimaryVid() uint32 {

	if obj.obj.PrimaryVid == nil {
		obj.setChoice(FlowCfmShortMANameFormatChoice.PRIMARY_VID)
	}

	return *obj.obj.PrimaryVid

}

// Primary VLAN ID (1-4094). Encoded as a 2-octet unsigned integer.
// PrimaryVid returns a uint32
func (obj *flowCfmShortMANameFormat) HasPrimaryVid() bool {
	return obj.obj.PrimaryVid != nil
}

// Primary VLAN ID (1-4094). Encoded as a 2-octet unsigned integer.
// SetPrimaryVid sets the uint32 value in the FlowCfmShortMANameFormat object
func (obj *flowCfmShortMANameFormat) SetPrimaryVid(value uint32) FlowCfmShortMANameFormat {
	obj.setChoice(FlowCfmShortMANameFormatChoice.PRIMARY_VID)
	obj.obj.PrimaryVid = &value
	return obj
}

// Character string (ASCII), up to 45 octets (RFC 2579 DisplayString).
// CharStr returns a string
func (obj *flowCfmShortMANameFormat) CharStr() string {

	if obj.obj.CharStr == nil {
		obj.setChoice(FlowCfmShortMANameFormatChoice.CHAR_STR)
	}

	return *obj.obj.CharStr

}

// Character string (ASCII), up to 45 octets (RFC 2579 DisplayString).
// CharStr returns a string
func (obj *flowCfmShortMANameFormat) HasCharStr() bool {
	return obj.obj.CharStr != nil
}

// Character string (ASCII), up to 45 octets (RFC 2579 DisplayString).
// SetCharStr sets the string value in the FlowCfmShortMANameFormat object
func (obj *flowCfmShortMANameFormat) SetCharStr(value string) FlowCfmShortMANameFormat {
	obj.setChoice(FlowCfmShortMANameFormatChoice.CHAR_STR)
	obj.obj.CharStr = &value
	return obj
}

// 2-octet unsigned integer, e.g. a locally assigned VLAN tag.
// TwoOctInt returns a uint32
func (obj *flowCfmShortMANameFormat) TwoOctInt() uint32 {

	if obj.obj.TwoOctInt == nil {
		obj.setChoice(FlowCfmShortMANameFormatChoice.TWO_OCT_INT)
	}

	return *obj.obj.TwoOctInt

}

// 2-octet unsigned integer, e.g. a locally assigned VLAN tag.
// TwoOctInt returns a uint32
func (obj *flowCfmShortMANameFormat) HasTwoOctInt() bool {
	return obj.obj.TwoOctInt != nil
}

// 2-octet unsigned integer, e.g. a locally assigned VLAN tag.
// SetTwoOctInt sets the uint32 value in the FlowCfmShortMANameFormat object
func (obj *flowCfmShortMANameFormat) SetTwoOctInt(value uint32) FlowCfmShortMANameFormat {
	obj.setChoice(FlowCfmShortMANameFormatChoice.TWO_OCT_INT)
	obj.obj.TwoOctInt = &value
	return obj
}

// RFC 2685 VPN ID, exactly 7 octets (3-octet OUI + 4-octet index). Encode as 14 hex characters, e.g. "00000000000000".
// Rfc2685VpnId returns a string
func (obj *flowCfmShortMANameFormat) Rfc2685VpnId() string {

	if obj.obj.Rfc_2685VpnId == nil {
		obj.setChoice(FlowCfmShortMANameFormatChoice.RFC_2685_VPN_ID)
	}

	return *obj.obj.Rfc_2685VpnId

}

// RFC 2685 VPN ID, exactly 7 octets (3-octet OUI + 4-octet index). Encode as 14 hex characters, e.g. "00000000000000".
// Rfc2685VpnId returns a string
func (obj *flowCfmShortMANameFormat) HasRfc2685VpnId() bool {
	return obj.obj.Rfc_2685VpnId != nil
}

// RFC 2685 VPN ID, exactly 7 octets (3-octet OUI + 4-octet index). Encode as 14 hex characters, e.g. "00000000000000".
// SetRfc2685VpnId sets the string value in the FlowCfmShortMANameFormat object
func (obj *flowCfmShortMANameFormat) SetRfc2685VpnId(value string) FlowCfmShortMANameFormat {
	obj.setChoice(FlowCfmShortMANameFormatChoice.RFC_2685_VPN_ID)
	obj.obj.Rfc_2685VpnId = &value
	return obj
}

func (obj *flowCfmShortMANameFormat) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.PrimaryVid != nil {

		if *obj.obj.PrimaryVid < 1 || *obj.obj.PrimaryVid > 4094 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= FlowCfmShortMANameFormat.PrimaryVid <= 4094 but Got %d", *obj.obj.PrimaryVid))
		}

	}

	if obj.obj.TwoOctInt != nil {

		if *obj.obj.TwoOctInt > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= FlowCfmShortMANameFormat.TwoOctInt <= 65535 but Got %d", *obj.obj.TwoOctInt))
		}

	}

	if obj.obj.Rfc_2685VpnId != nil {

		err := obj.validateHex(obj.Rfc2685VpnId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmShortMANameFormat.Rfc2685VpnId"))
		}

	}

}

func (obj *flowCfmShortMANameFormat) setDefault() {
	var choices_set int = 0
	var choice FlowCfmShortMANameFormatChoiceEnum

	if obj.obj.PrimaryVid != nil {
		choices_set += 1
		choice = FlowCfmShortMANameFormatChoice.PRIMARY_VID
	}

	if obj.obj.CharStr != nil {
		choices_set += 1
		choice = FlowCfmShortMANameFormatChoice.CHAR_STR
	}

	if obj.obj.TwoOctInt != nil {
		choices_set += 1
		choice = FlowCfmShortMANameFormatChoice.TWO_OCT_INT
	}

	if obj.obj.Rfc_2685VpnId != nil {
		choices_set += 1
		choice = FlowCfmShortMANameFormatChoice.RFC_2685_VPN_ID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmShortMANameFormatChoice.CHAR_STR)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmShortMANameFormat")
			}
		} else {
			intVal := otg.FlowCfmShortMANameFormat_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmShortMANameFormat_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Rfc_2685VpnId == nil && choice == FlowCfmShortMANameFormatChoice.RFC_2685_VPN_ID {
		obj.SetRfc2685VpnId("00000000000000")
	}

}
