package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSidEntry *****
type actionProtocolIsisSrv6MyLocalSidEntry struct {
	validation
	obj          *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	marshaller   marshalActionProtocolIsisSrv6MyLocalSidEntry
	unMarshaller unMarshalActionProtocolIsisSrv6MyLocalSidEntry
}

func NewActionProtocolIsisSrv6MyLocalSidEntry() ActionProtocolIsisSrv6MyLocalSidEntry {
	obj := actionProtocolIsisSrv6MyLocalSidEntry{obj: &otg.ActionProtocolIsisSrv6MyLocalSidEntry{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) msg() *otg.ActionProtocolIsisSrv6MyLocalSidEntry {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidEntry {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSidEntry struct {
	obj *actionProtocolIsisSrv6MyLocalSidEntry
}

type marshalActionProtocolIsisSrv6MyLocalSidEntry interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSidEntry to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidEntry, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSidEntry to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSidEntry to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSidEntry to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSidEntry struct {
	obj *actionProtocolIsisSrv6MyLocalSidEntry
}

type unMarshalActionProtocolIsisSrv6MyLocalSidEntry interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSidEntry from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidEntry) (ActionProtocolIsisSrv6MyLocalSidEntry, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSidEntry from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSidEntry from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSidEntry from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) Marshal() marshalActionProtocolIsisSrv6MyLocalSidEntry {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSidEntry{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidEntry {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSidEntry{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidEntry) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSidEntry, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidEntry) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSidEntry) (ActionProtocolIsisSrv6MyLocalSidEntry, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSidEntry) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidEntry) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidEntry) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidEntry) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSidEntry) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSidEntry) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) Clone() (ActionProtocolIsisSrv6MyLocalSidEntry, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSidEntry()
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

// ActionProtocolIsisSrv6MyLocalSidEntry is one SRv6 My Local SID entry. Corresponds to a single row in the MY_LOCAL_SID table (RFC 8986 Section 3.3), indexed by (SID prefix, behavior).
type ActionProtocolIsisSrv6MyLocalSidEntry interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSidEntry to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSidEntry from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSidEntry
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSidEntry) ActionProtocolIsisSrv6MyLocalSidEntry
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSidEntry
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSidEntry
	// validate validates ActionProtocolIsisSrv6MyLocalSidEntry
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSidEntry, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// SidPrefix returns string, set in ActionProtocolIsisSrv6MyLocalSidEntry.
	SidPrefix() string
	// SetSidPrefix assigns string provided by user to ActionProtocolIsisSrv6MyLocalSidEntry
	SetSidPrefix(value string) ActionProtocolIsisSrv6MyLocalSidEntry
	// PrefixLength returns uint32, set in ActionProtocolIsisSrv6MyLocalSidEntry.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to ActionProtocolIsisSrv6MyLocalSidEntry
	SetPrefixLength(value uint32) ActionProtocolIsisSrv6MyLocalSidEntry
	// HasPrefixLength checks if PrefixLength has been set in ActionProtocolIsisSrv6MyLocalSidEntry
	HasPrefixLength() bool
	// Behavior returns ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum, set in ActionProtocolIsisSrv6MyLocalSidEntry
	Behavior() ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	// SetBehavior assigns ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum provided by user to ActionProtocolIsisSrv6MyLocalSidEntry
	SetBehavior(value ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum) ActionProtocolIsisSrv6MyLocalSidEntry
	// HasBehavior checks if Behavior has been set in ActionProtocolIsisSrv6MyLocalSidEntry
	HasBehavior() bool
	// VrfName returns string, set in ActionProtocolIsisSrv6MyLocalSidEntry.
	VrfName() string
	// SetVrfName assigns string provided by user to ActionProtocolIsisSrv6MyLocalSidEntry
	SetVrfName(value string) ActionProtocolIsisSrv6MyLocalSidEntry
	// HasVrfName checks if VrfName has been set in ActionProtocolIsisSrv6MyLocalSidEntry
	HasVrfName() bool
	// NextHop returns string, set in ActionProtocolIsisSrv6MyLocalSidEntry.
	NextHop() string
	// SetNextHop assigns string provided by user to ActionProtocolIsisSrv6MyLocalSidEntry
	SetNextHop(value string) ActionProtocolIsisSrv6MyLocalSidEntry
	// HasNextHop checks if NextHop has been set in ActionProtocolIsisSrv6MyLocalSidEntry
	HasNextHop() bool
}

// IPv6 address of this My Local SID (e.g., "fcbb:bbbb:1::" for a SONiC uN node using the fcbb:bbbb::/32 locator block).
// SidPrefix returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SidPrefix() string {

	return *obj.obj.SidPrefix

}

// IPv6 address of this My Local SID (e.g., "fcbb:bbbb:1::" for a SONiC uN node using the fcbb:bbbb::/32 locator block).
// SetSidPrefix sets the string value in the ActionProtocolIsisSrv6MyLocalSidEntry object
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SetSidPrefix(value string) ActionProtocolIsisSrv6MyLocalSidEntry {

	obj.obj.SidPrefix = &value
	return obj
}

// Prefix length of the My Local SID in bits. Common values: 48 for u_n with F3216 locator; 64 for end and DT behaviors.
// PrefixLength returns a uint32
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// Prefix length of the My Local SID in bits. Common values: 48 for u_n with F3216 locator; 64 for end and DT behaviors.
// PrefixLength returns a uint32
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// Prefix length of the My Local SID in bits. Common values: 48 for u_n with F3216 locator; 64 for end and DT behaviors.
// SetPrefixLength sets the uint32 value in the ActionProtocolIsisSrv6MyLocalSidEntry object
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SetPrefixLength(value uint32) ActionProtocolIsisSrv6MyLocalSidEntry {

	obj.obj.PrefixLength = &value
	return obj
}

type ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum string

// Enum of Behavior on ActionProtocolIsisSrv6MyLocalSidEntry
var ActionProtocolIsisSrv6MyLocalSidEntryBehavior = struct {
	U_N      ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	U_DT4    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	U_DT6    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	U_DT46   ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	U_DX4    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	U_DX6    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	END      ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	END_DT4  ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	END_DT6  ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
	END_DT46 ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum
}{
	U_N:      ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_n"),
	U_DT4:    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_dt4"),
	U_DT6:    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_dt6"),
	U_DT46:   ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_dt46"),
	U_DX4:    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_dx4"),
	U_DX6:    ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("u_dx6"),
	END:      ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("end"),
	END_DT4:  ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("end_dt4"),
	END_DT6:  ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("end_dt6"),
	END_DT46: ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum("end_dt46"),
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) Behavior() ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum {
	return ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum(obj.obj.Behavior.Enum().String())
}

// SRv6 endpoint behavior applied when a packet's DA matches this My Local SID (RFC 8986 Section 7.4 codepoints, RFC 9800). 'u_n' - Micro-SID node: shift active uSID out of DA and forward
// (RFC 9800 Section 4).
// 'u_dt4' / 'u_dt6' / 'u_dt46' - Micro-SID decap + IPv4 / IPv6 /
// dual-stack table lookup in the VRF given by vrf_name (RFC 9800).
// 'u_dx4' / 'u_dx6' - Micro-SID decap + L3 cross-connect to the
// IPv4 / IPv6 next-hop given by next_hop (RFC 9800).
// 'end' / 'end_dt4' / 'end_dt6' / 'end_dt46' - Full-SID
// equivalents (RFC 8986 Sections 4.1, 4.6, 4.7, 4.8).
// Behavior returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) HasBehavior() bool {
	return obj.obj.Behavior != nil
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SetBehavior(value ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum) ActionProtocolIsisSrv6MyLocalSidEntry {
	intValue, ok := otg.ActionProtocolIsisSrv6MyLocalSidEntry_Behavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisSrv6MyLocalSidEntryBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsisSrv6MyLocalSidEntry_Behavior_Enum(intValue)
	obj.obj.Behavior = &enumValue

	return obj
}

// VRF name for decapsulation behaviors (u_dt4, u_dt6, u_dt46, u_dx4, u_dx6, end_dt4, end_dt6, end_dt46). Omit or leave empty for the default VRF.
// VrfName returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) VrfName() string {

	return *obj.obj.VrfName

}

// VRF name for decapsulation behaviors (u_dt4, u_dt6, u_dt46, u_dx4, u_dx6, end_dt4, end_dt6, end_dt46). Omit or leave empty for the default VRF.
// VrfName returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) HasVrfName() bool {
	return obj.obj.VrfName != nil
}

// VRF name for decapsulation behaviors (u_dt4, u_dt6, u_dt46, u_dx4, u_dx6, end_dt4, end_dt6, end_dt46). Omit or leave empty for the default VRF.
// SetVrfName sets the string value in the ActionProtocolIsisSrv6MyLocalSidEntry object
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SetVrfName(value string) ActionProtocolIsisSrv6MyLocalSidEntry {

	obj.obj.VrfName = &value
	return obj
}

// Next-hop address for cross-connect behaviors: an IPv4 address for u_dx4, an IPv6 address for u_dx6. Not applicable for DT or u_n behaviors.
// NextHop returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) NextHop() string {

	return *obj.obj.NextHop

}

// Next-hop address for cross-connect behaviors: an IPv4 address for u_dx4, an IPv6 address for u_dx6. Not applicable for DT or u_n behaviors.
// NextHop returns a string
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) HasNextHop() bool {
	return obj.obj.NextHop != nil
}

// Next-hop address for cross-connect behaviors: an IPv4 address for u_dx4, an IPv6 address for u_dx6. Not applicable for DT or u_n behaviors.
// SetNextHop sets the string value in the ActionProtocolIsisSrv6MyLocalSidEntry object
func (obj *actionProtocolIsisSrv6MyLocalSidEntry) SetNextHop(value string) ActionProtocolIsisSrv6MyLocalSidEntry {

	obj.obj.NextHop = &value
	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// SidPrefix is required
	if obj.obj.SidPrefix == nil {
		vObj.validationErrors = append(vObj.validationErrors, "SidPrefix is required field on interface ActionProtocolIsisSrv6MyLocalSidEntry")
	}
	if obj.obj.SidPrefix != nil {

		err := obj.validateIpv6(obj.SidPrefix())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on ActionProtocolIsisSrv6MyLocalSidEntry.SidPrefix"))
		}

	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength < 1 || *obj.obj.PrefixLength > 128 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= ActionProtocolIsisSrv6MyLocalSidEntry.PrefixLength <= 128 but Got %d", *obj.obj.PrefixLength))
		}

	}

}

func (obj *actionProtocolIsisSrv6MyLocalSidEntry) setDefault() {
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(48)
	}
	if obj.obj.Behavior == nil {
		obj.SetBehavior(ActionProtocolIsisSrv6MyLocalSidEntryBehavior.U_N)

	}

}
