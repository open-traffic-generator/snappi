package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6AdjSid *****
type isisSRv6AdjSid struct {
	validation
	obj                *otg.IsisSRv6AdjSid
	marshaller         marshalIsisSRv6AdjSid
	unMarshaller       unMarshalIsisSRv6AdjSid
	sidStructureHolder IsisSRv6SidStructure
}

func NewIsisSRv6AdjSid() IsisSRv6AdjSid {
	obj := isisSRv6AdjSid{obj: &otg.IsisSRv6AdjSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6AdjSid) msg() *otg.IsisSRv6AdjSid {
	return obj.obj
}

func (obj *isisSRv6AdjSid) setMsg(msg *otg.IsisSRv6AdjSid) IsisSRv6AdjSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6AdjSid struct {
	obj *isisSRv6AdjSid
}

type marshalIsisSRv6AdjSid interface {
	// ToProto marshals IsisSRv6AdjSid to protobuf object *otg.IsisSRv6AdjSid
	ToProto() (*otg.IsisSRv6AdjSid, error)
	// ToPbText marshals IsisSRv6AdjSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6AdjSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6AdjSid to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6AdjSid struct {
	obj *isisSRv6AdjSid
}

type unMarshalIsisSRv6AdjSid interface {
	// FromProto unmarshals IsisSRv6AdjSid from protobuf object *otg.IsisSRv6AdjSid
	FromProto(msg *otg.IsisSRv6AdjSid) (IsisSRv6AdjSid, error)
	// FromPbText unmarshals IsisSRv6AdjSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6AdjSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6AdjSid from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6AdjSid) Marshal() marshalIsisSRv6AdjSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6AdjSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6AdjSid) Unmarshal() unMarshalIsisSRv6AdjSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6AdjSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6AdjSid) ToProto() (*otg.IsisSRv6AdjSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6AdjSid) FromProto(msg *otg.IsisSRv6AdjSid) (IsisSRv6AdjSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6AdjSid) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6AdjSid) FromPbText(value string) error {
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

func (m *marshalisisSRv6AdjSid) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6AdjSid) FromYaml(value string) error {
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

func (m *marshalisisSRv6AdjSid) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6AdjSid) FromJson(value string) error {
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

func (obj *isisSRv6AdjSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6AdjSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6AdjSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6AdjSid) Clone() (IsisSRv6AdjSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6AdjSid()
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

func (obj *isisSRv6AdjSid) setNil() {
	obj.sidStructureHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6AdjSid is sRv6 Adjacency SID Sub-TLV for IS-IS interfaces. Point-to-point adjacencies use End.X SID Sub-TLV (sub-TLV type 43); LAN adjacencies use LAN End.X SID Sub-TLV (sub-TLV type 44). The End.X SID binds a 128-bit SRv6 SID to a specific outgoing interface and next-hop, enabling traffic steering across a specific L3 adjacency. The full 128-bit SID is assembled as:
// Locator (selected via locator/custom_locator_reference) | Function | Argument
// Example  - given locator fc00:0:1:: /48 with sid_structure lb=32, ln=16, fn=16, arg=0:
// locator auto (=> fc00:0:1::), function "00c8"
// => final adjacency SID: fc00:0:1:c8::
// locator custom_locator_reference "loc2" (fc00:0:2:: /48), function "00c9"
// => final adjacency SID: fc00:0:2:c9::
// Valid behaviors include End.X variants (with PSP/USP/USD flavors) and cross-connect decapsulation behaviors (End.DX4, End.DX6). Reference: RFC 9352 Sections 8.1-8.2, RFC 8986 Section 4.3.
type IsisSRv6AdjSid interface {
	Validation
	// msg marshals IsisSRv6AdjSid to protobuf object *otg.IsisSRv6AdjSid
	// and doesn't set defaults
	msg() *otg.IsisSRv6AdjSid
	// setMsg unmarshals IsisSRv6AdjSid from protobuf object *otg.IsisSRv6AdjSid
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6AdjSid) IsisSRv6AdjSid
	// provides marshal interface
	Marshal() marshalIsisSRv6AdjSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6AdjSid
	// validate validates IsisSRv6AdjSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6AdjSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns IsisSRv6AdjSidLocatorEnum, set in IsisSRv6AdjSid
	Locator() IsisSRv6AdjSidLocatorEnum
	// SetLocator assigns IsisSRv6AdjSidLocatorEnum provided by user to IsisSRv6AdjSid
	SetLocator(value IsisSRv6AdjSidLocatorEnum) IsisSRv6AdjSid
	// HasLocator checks if Locator has been set in IsisSRv6AdjSid
	HasLocator() bool
	// CustomLocatorReference returns string, set in IsisSRv6AdjSid.
	CustomLocatorReference() string
	// SetCustomLocatorReference assigns string provided by user to IsisSRv6AdjSid
	SetCustomLocatorReference(value string) IsisSRv6AdjSid
	// HasCustomLocatorReference checks if CustomLocatorReference has been set in IsisSRv6AdjSid
	HasCustomLocatorReference() bool
	// Function returns string, set in IsisSRv6AdjSid.
	Function() string
	// SetFunction assigns string provided by user to IsisSRv6AdjSid
	SetFunction(value string) IsisSRv6AdjSid
	// HasFunction checks if Function has been set in IsisSRv6AdjSid
	HasFunction() bool
	// EndpointBehavior returns IsisSRv6AdjSidEndpointBehaviorEnum, set in IsisSRv6AdjSid
	EndpointBehavior() IsisSRv6AdjSidEndpointBehaviorEnum
	// SetEndpointBehavior assigns IsisSRv6AdjSidEndpointBehaviorEnum provided by user to IsisSRv6AdjSid
	SetEndpointBehavior(value IsisSRv6AdjSidEndpointBehaviorEnum) IsisSRv6AdjSid
	// HasEndpointBehavior checks if EndpointBehavior has been set in IsisSRv6AdjSid
	HasEndpointBehavior() bool
	// BFlag returns bool, set in IsisSRv6AdjSid.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisSRv6AdjSid
	SetBFlag(value bool) IsisSRv6AdjSid
	// HasBFlag checks if BFlag has been set in IsisSRv6AdjSid
	HasBFlag() bool
	// SFlag returns bool, set in IsisSRv6AdjSid.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisSRv6AdjSid
	SetSFlag(value bool) IsisSRv6AdjSid
	// HasSFlag checks if SFlag has been set in IsisSRv6AdjSid
	HasSFlag() bool
	// PFlag returns bool, set in IsisSRv6AdjSid.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisSRv6AdjSid
	SetPFlag(value bool) IsisSRv6AdjSid
	// HasPFlag checks if PFlag has been set in IsisSRv6AdjSid
	HasPFlag() bool
	// Weight returns uint32, set in IsisSRv6AdjSid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisSRv6AdjSid
	SetWeight(value uint32) IsisSRv6AdjSid
	// HasWeight checks if Weight has been set in IsisSRv6AdjSid
	HasWeight() bool
	// Algorithm returns uint32, set in IsisSRv6AdjSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisSRv6AdjSid
	SetAlgorithm(value uint32) IsisSRv6AdjSid
	// HasAlgorithm checks if Algorithm has been set in IsisSRv6AdjSid
	HasAlgorithm() bool
	// CFlag returns bool, set in IsisSRv6AdjSid.
	CFlag() bool
	// SetCFlag assigns bool provided by user to IsisSRv6AdjSid
	SetCFlag(value bool) IsisSRv6AdjSid
	// HasCFlag checks if CFlag has been set in IsisSRv6AdjSid
	HasCFlag() bool
	// SidStructure returns IsisSRv6SidStructure, set in IsisSRv6AdjSid.
	// IsisSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1), carried within SRv6 SID Sub-TLVs (End SID type 5, End.X SID type 43/44). Describes the internal bit-field decomposition of the SRv6 SID value so that receiving routers can interpret each field independently. The four length fields (lb_length + ln_length + function_length + argument_length) MUST NOT exceed 128 bits. Required when advertising Micro-SID (uSID) SIDs to describe the compressed encoding. Example for common uSID F3216 format:
	// lb_length=32, ln_length=16, function_length=16, argument_length=0
	// Reference: RFC 9352 Section 9, RFC 9800.
	SidStructure() IsisSRv6SidStructure
	// SetSidStructure assigns IsisSRv6SidStructure provided by user to IsisSRv6AdjSid.
	// IsisSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1), carried within SRv6 SID Sub-TLVs (End SID type 5, End.X SID type 43/44). Describes the internal bit-field decomposition of the SRv6 SID value so that receiving routers can interpret each field independently. The four length fields (lb_length + ln_length + function_length + argument_length) MUST NOT exceed 128 bits. Required when advertising Micro-SID (uSID) SIDs to describe the compressed encoding. Example for common uSID F3216 format:
	// lb_length=32, ln_length=16, function_length=16, argument_length=0
	// Reference: RFC 9352 Section 9, RFC 9800.
	SetSidStructure(value IsisSRv6SidStructure) IsisSRv6AdjSid
	// HasSidStructure checks if SidStructure has been set in IsisSRv6AdjSid
	HasSidStructure() bool
	setNil()
}

type IsisSRv6AdjSidLocatorEnum string

// Enum of Locator on IsisSRv6AdjSid
var IsisSRv6AdjSidLocator = struct {
	AUTO                     IsisSRv6AdjSidLocatorEnum
	CUSTOM_LOCATOR_REFERENCE IsisSRv6AdjSidLocatorEnum
}{
	AUTO:                     IsisSRv6AdjSidLocatorEnum("auto"),
	CUSTOM_LOCATOR_REFERENCE: IsisSRv6AdjSidLocatorEnum("custom_locator_reference"),
}

func (obj *isisSRv6AdjSid) Locator() IsisSRv6AdjSidLocatorEnum {
	return IsisSRv6AdjSidLocatorEnum(obj.obj.Locator.Enum().String())
}

// Selects which locator from the IsisSRv6.Locator list to use as the locator part of this adjacency SID. 'auto' (default) uses the first locator defined in isis.segment_routing.srv6_locators  - suitable when only one locator is configured. 'custom_locator_reference' uses the specific locator identified by custom_locator_reference  - use this when multiple locators are configured and a particular one must be selected, e.g. for a Flex-Algo binding.
// Locator returns a string
func (obj *isisSRv6AdjSid) HasLocator() bool {
	return obj.obj.Locator != nil
}

func (obj *isisSRv6AdjSid) SetLocator(value IsisSRv6AdjSidLocatorEnum) IsisSRv6AdjSid {
	intValue, ok := otg.IsisSRv6AdjSid_Locator_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6AdjSidLocatorEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6AdjSid_Locator_Enum(intValue)
	obj.obj.Locator = &enumValue

	return obj
}

// Name of the IsisSRv6.Locator to use when locator is set to 'custom_locator_reference'. Must match the locator_name of a locator configured in isis.segment_routing.srv6_locators. Example: "loc2" selects the locator whose locator_name is "loc2".
//
// x-constraint:
// - /components/schemas/IsisSRv6.Locator/properties/locator_name
//
// CustomLocatorReference returns a string
func (obj *isisSRv6AdjSid) CustomLocatorReference() string {

	return *obj.obj.CustomLocatorReference

}

// Name of the IsisSRv6.Locator to use when locator is set to 'custom_locator_reference'. Must match the locator_name of a locator configured in isis.segment_routing.srv6_locators. Example: "loc2" selects the locator whose locator_name is "loc2".
//
// x-constraint:
// - /components/schemas/IsisSRv6.Locator/properties/locator_name
//
// CustomLocatorReference returns a string
func (obj *isisSRv6AdjSid) HasCustomLocatorReference() bool {
	return obj.obj.CustomLocatorReference != nil
}

// Name of the IsisSRv6.Locator to use when locator is set to 'custom_locator_reference'. Must match the locator_name of a locator configured in isis.segment_routing.srv6_locators. Example: "loc2" selects the locator whose locator_name is "loc2".
//
// x-constraint:
// - /components/schemas/IsisSRv6.Locator/properties/locator_name
//
// SetCustomLocatorReference sets the string value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetCustomLocatorReference(value string) IsisSRv6AdjSid {

	obj.obj.CustomLocatorReference = &value
	return obj
}

// The Function part of this adjacency SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must match function_length in the selected IsisSRv6.Locator.sid_structure divided by 4  - e.g. function_length 16 requires a 4-nibble string. Example: "00c8" places the value 200 (0xc8) in the function field; with locator fc00:0:1:: /48 (selected via auto or custom_locator_reference) the resulting SID is fc00:0:1:c8::.
// Function returns a string
func (obj *isisSRv6AdjSid) Function() string {

	return *obj.obj.Function

}

// The Function part of this adjacency SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must match function_length in the selected IsisSRv6.Locator.sid_structure divided by 4  - e.g. function_length 16 requires a 4-nibble string. Example: "00c8" places the value 200 (0xc8) in the function field; with locator fc00:0:1:: /48 (selected via auto or custom_locator_reference) the resulting SID is fc00:0:1:c8::.
// Function returns a string
func (obj *isisSRv6AdjSid) HasFunction() bool {
	return obj.obj.Function != nil
}

// The Function part of this adjacency SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must match function_length in the selected IsisSRv6.Locator.sid_structure divided by 4  - e.g. function_length 16 requires a 4-nibble string. Example: "00c8" places the value 200 (0xc8) in the function field; with locator fc00:0:1:: /48 (selected via auto or custom_locator_reference) the resulting SID is fc00:0:1:c8::.
// SetFunction sets the string value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetFunction(value string) IsisSRv6AdjSid {

	obj.obj.Function = &value
	return obj
}

type IsisSRv6AdjSidEndpointBehaviorEnum string

// Enum of EndpointBehavior on IsisSRv6AdjSid
var IsisSRv6AdjSidEndpointBehavior = struct {
	END_X                  IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP         IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USP         IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USP     IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USD         IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USD     IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USP_USD     IsisSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USP_USD IsisSRv6AdjSidEndpointBehaviorEnum
	END_DX4                IsisSRv6AdjSidEndpointBehaviorEnum
	END_DX6                IsisSRv6AdjSidEndpointBehaviorEnum
}{
	END_X:                  IsisSRv6AdjSidEndpointBehaviorEnum("end_x"),
	END_X_WITH_PSP:         IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp"),
	END_X_WITH_USP:         IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_usp"),
	END_X_WITH_PSP_USP:     IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usp"),
	END_X_WITH_USD:         IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_usd"),
	END_X_WITH_PSP_USD:     IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usd"),
	END_X_WITH_USP_USD:     IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_usp_usd"),
	END_X_WITH_PSP_USP_USD: IsisSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usp_usd"),
	END_DX4:                IsisSRv6AdjSidEndpointBehaviorEnum("end_dx4"),
	END_DX6:                IsisSRv6AdjSidEndpointBehaviorEnum("end_dx6"),
}

func (obj *isisSRv6AdjSid) EndpointBehavior() IsisSRv6AdjSidEndpointBehaviorEnum {
	return IsisSRv6AdjSidEndpointBehaviorEnum(obj.obj.EndpointBehavior.Enum().String())
}

// The endpoint behavior for this adjacency SID, encoded as a 2-octet behavior codepoint (RFC 8986 Section 7.4). Valid adjacency-level behaviors for End.X SID Sub-TLV types 43/44 are End.X variants (with PSP/USP/USD flavors) and per-CE cross-connect decapsulation behaviors (End.DX4, End.DX6).
// EndpointBehavior returns a string
func (obj *isisSRv6AdjSid) HasEndpointBehavior() bool {
	return obj.obj.EndpointBehavior != nil
}

func (obj *isisSRv6AdjSid) SetEndpointBehavior(value IsisSRv6AdjSidEndpointBehaviorEnum) IsisSRv6AdjSid {
	intValue, ok := otg.IsisSRv6AdjSid_EndpointBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6AdjSidEndpointBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6AdjSid_EndpointBehavior_Enum(intValue)
	obj.obj.EndpointBehavior = &enumValue

	return obj
}

// Backup flag (B-flag, bit 0). When set, the End.X SID is eligible for IP Fast-ReRoute (IP-FRR) protection, indicating a backup adjacency that can protect traffic when the primary adjacency fails (RFC 9352 Section 8.1).
// BFlag returns a bool
func (obj *isisSRv6AdjSid) BFlag() bool {

	return *obj.obj.BFlag

}

// Backup flag (B-flag, bit 0). When set, the End.X SID is eligible for IP Fast-ReRoute (IP-FRR) protection, indicating a backup adjacency that can protect traffic when the primary adjacency fails (RFC 9352 Section 8.1).
// BFlag returns a bool
func (obj *isisSRv6AdjSid) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// Backup flag (B-flag, bit 0). When set, the End.X SID is eligible for IP Fast-ReRoute (IP-FRR) protection, indicating a backup adjacency that can protect traffic when the primary adjacency fails (RFC 9352 Section 8.1).
// SetBFlag sets the bool value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetBFlag(value bool) IsisSRv6AdjSid {

	obj.obj.BFlag = &value
	return obj
}

// Set flag (S-flag, bit 1). When set, the End.X SID refers to a set of adjacencies (e.g., parallel links, ECMP group). The same SID value may be assigned to multiple adjacencies forming the set (RFC 9352 Section 8.1).
// SFlag returns a bool
func (obj *isisSRv6AdjSid) SFlag() bool {

	return *obj.obj.SFlag

}

// Set flag (S-flag, bit 1). When set, the End.X SID refers to a set of adjacencies (e.g., parallel links, ECMP group). The same SID value may be assigned to multiple adjacencies forming the set (RFC 9352 Section 8.1).
// SFlag returns a bool
func (obj *isisSRv6AdjSid) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// Set flag (S-flag, bit 1). When set, the End.X SID refers to a set of adjacencies (e.g., parallel links, ECMP group). The same SID value may be assigned to multiple adjacencies forming the set (RFC 9352 Section 8.1).
// SetSFlag sets the bool value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetSFlag(value bool) IsisSRv6AdjSid {

	obj.obj.SFlag = &value
	return obj
}

// Persistent flag (P-flag, bit 2). When set, the End.X SID value is persistently allocated and remains consistent across router restarts and interface flap events (RFC 9352 Section 8.1).
// PFlag returns a bool
func (obj *isisSRv6AdjSid) PFlag() bool {

	return *obj.obj.PFlag

}

// Persistent flag (P-flag, bit 2). When set, the End.X SID value is persistently allocated and remains consistent across router restarts and interface flap events (RFC 9352 Section 8.1).
// PFlag returns a bool
func (obj *isisSRv6AdjSid) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// Persistent flag (P-flag, bit 2). When set, the End.X SID value is persistently allocated and remains consistent across router restarts and interface flap events (RFC 9352 Section 8.1).
// SetPFlag sets the bool value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetPFlag(value bool) IsisSRv6AdjSid {

	obj.obj.PFlag = &value
	return obj
}

// Weight for load balancing across adjacencies sharing the same End.X SID (applicable when S-flag is set). Traffic is distributed proportionally across the adjacency set according to the weight values (RFC 9352 Section 8.1).
// Weight returns a uint32
func (obj *isisSRv6AdjSid) Weight() uint32 {

	return *obj.obj.Weight

}

// Weight for load balancing across adjacencies sharing the same End.X SID (applicable when S-flag is set). Traffic is distributed proportionally across the adjacency set according to the weight values (RFC 9352 Section 8.1).
// Weight returns a uint32
func (obj *isisSRv6AdjSid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// Weight for load balancing across adjacencies sharing the same End.X SID (applicable when S-flag is set). Traffic is distributed proportionally across the adjacency set according to the weight values (RFC 9352 Section 8.1).
// SetWeight sets the uint32 value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetWeight(value uint32) IsisSRv6AdjSid {

	obj.obj.Weight = &value
	return obj
}

// IGP algorithm associated with this adjacency SID (RFC 8665). 0 = standard SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350). Binds the adjacency SID to the corresponding topology or path computation algorithm.
// Algorithm returns a uint32
func (obj *isisSRv6AdjSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// IGP algorithm associated with this adjacency SID (RFC 8665). 0 = standard SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350). Binds the adjacency SID to the corresponding topology or path computation algorithm.
// Algorithm returns a uint32
func (obj *isisSRv6AdjSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// IGP algorithm associated with this adjacency SID (RFC 8665). 0 = standard SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350). Binds the adjacency SID to the corresponding topology or path computation algorithm.
// SetAlgorithm sets the uint32 value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetAlgorithm(value uint32) IsisSRv6AdjSid {

	obj.obj.Algorithm = &value
	return obj
}

// Compression (uSID) flag. When set, this adjacency SID is a Micro-SID (uSID) End.X per RFC 9800. A headend can pack it into a shared 128-bit uSID container alongside node uSIDs from the same /32 block, directing traffic over a specific L3 adjacency within a compressed segment list. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the selected IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "00c8" => uSID fc00:0:1:c8::  (End.X)
// c_flag=true, function "00c9" => uSID fc00:0:1:c9::  (End.X alternate link)
// For a path routed via node 1's adjacency "00c8" then node 2 End "0001", the headend builds DA = fc00:0:1:c8:2:1:: encoding both hops in a single 128-bit address. Reference: RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6AdjSid) CFlag() bool {

	return *obj.obj.CFlag

}

// Compression (uSID) flag. When set, this adjacency SID is a Micro-SID (uSID) End.X per RFC 9800. A headend can pack it into a shared 128-bit uSID container alongside node uSIDs from the same /32 block, directing traffic over a specific L3 adjacency within a compressed segment list. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the selected IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "00c8" => uSID fc00:0:1:c8::  (End.X)
// c_flag=true, function "00c9" => uSID fc00:0:1:c9::  (End.X alternate link)
// For a path routed via node 1's adjacency "00c8" then node 2 End "0001", the headend builds DA = fc00:0:1:c8:2:1:: encoding both hops in a single 128-bit address. Reference: RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6AdjSid) HasCFlag() bool {
	return obj.obj.CFlag != nil
}

// Compression (uSID) flag. When set, this adjacency SID is a Micro-SID (uSID) End.X per RFC 9800. A headend can pack it into a shared 128-bit uSID container alongside node uSIDs from the same /32 block, directing traffic over a specific L3 adjacency within a compressed segment list. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the selected IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "00c8" => uSID fc00:0:1:c8::  (End.X)
// c_flag=true, function "00c9" => uSID fc00:0:1:c9::  (End.X alternate link)
// For a path routed via node 1's adjacency "00c8" then node 2 End "0001", the headend builds DA = fc00:0:1:c8:2:1:: encoding both hops in a single 128-bit address. Reference: RFC 9800.
// SetCFlag sets the bool value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetCFlag(value bool) IsisSRv6AdjSid {

	obj.obj.CFlag = &value
	return obj
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this adjacency SID. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SidStructure returns a IsisSRv6SidStructure
func (obj *isisSRv6AdjSid) SidStructure() IsisSRv6SidStructure {
	if obj.obj.SidStructure == nil {
		obj.obj.SidStructure = NewIsisSRv6SidStructure().msg()
	}
	if obj.sidStructureHolder == nil {
		obj.sidStructureHolder = &isisSRv6SidStructure{obj: obj.obj.SidStructure}
	}
	return obj.sidStructureHolder
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this adjacency SID. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SidStructure returns a IsisSRv6SidStructure
func (obj *isisSRv6AdjSid) HasSidStructure() bool {
	return obj.obj.SidStructure != nil
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this adjacency SID. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SetSidStructure sets the IsisSRv6SidStructure value in the IsisSRv6AdjSid object
func (obj *isisSRv6AdjSid) SetSidStructure(value IsisSRv6SidStructure) IsisSRv6AdjSid {

	obj.sidStructureHolder = nil
	obj.obj.SidStructure = value.msg()

	return obj
}

func (obj *isisSRv6AdjSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Function != nil {

		err := obj.validateHex(obj.Function())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRv6AdjSid.Function"))
		}

	}

	if obj.obj.Weight != nil {

		if *obj.obj.Weight > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6AdjSid.Weight <= 255 but Got %d", *obj.obj.Weight))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisSRv6AdjSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

	if obj.obj.SidStructure != nil {

		obj.SidStructure().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6AdjSid) setDefault() {
	if obj.obj.Locator == nil {
		obj.SetLocator(IsisSRv6AdjSidLocator.AUTO)

	}
	if obj.obj.EndpointBehavior == nil {
		obj.SetEndpointBehavior(IsisSRv6AdjSidEndpointBehavior.END_X)

	}
	if obj.obj.BFlag == nil {
		obj.SetBFlag(false)
	}
	if obj.obj.SFlag == nil {
		obj.SetSFlag(false)
	}
	if obj.obj.PFlag == nil {
		obj.SetPFlag(false)
	}
	if obj.obj.Weight == nil {
		obj.SetWeight(0)
	}
	if obj.obj.Algorithm == nil {
		obj.SetAlgorithm(0)
	}
	if obj.obj.CFlag == nil {
		obj.SetCFlag(false)
	}

}
