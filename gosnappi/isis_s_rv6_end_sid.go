package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRv6EndSid *****
type isisSRv6EndSid struct {
	validation
	obj          *otg.IsisSRv6EndSid
	marshaller   marshalIsisSRv6EndSid
	unMarshaller unMarshalIsisSRv6EndSid
}

func NewIsisSRv6EndSid() IsisSRv6EndSid {
	obj := isisSRv6EndSid{obj: &otg.IsisSRv6EndSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRv6EndSid) msg() *otg.IsisSRv6EndSid {
	return obj.obj
}

func (obj *isisSRv6EndSid) setMsg(msg *otg.IsisSRv6EndSid) IsisSRv6EndSid {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRv6EndSid struct {
	obj *isisSRv6EndSid
}

type marshalIsisSRv6EndSid interface {
	// ToProto marshals IsisSRv6EndSid to protobuf object *otg.IsisSRv6EndSid
	ToProto() (*otg.IsisSRv6EndSid, error)
	// ToPbText marshals IsisSRv6EndSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRv6EndSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRv6EndSid to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRv6EndSid struct {
	obj *isisSRv6EndSid
}

type unMarshalIsisSRv6EndSid interface {
	// FromProto unmarshals IsisSRv6EndSid from protobuf object *otg.IsisSRv6EndSid
	FromProto(msg *otg.IsisSRv6EndSid) (IsisSRv6EndSid, error)
	// FromPbText unmarshals IsisSRv6EndSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRv6EndSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRv6EndSid from JSON text
	FromJson(value string) error
}

func (obj *isisSRv6EndSid) Marshal() marshalIsisSRv6EndSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRv6EndSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRv6EndSid) Unmarshal() unMarshalIsisSRv6EndSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRv6EndSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRv6EndSid) ToProto() (*otg.IsisSRv6EndSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRv6EndSid) FromProto(msg *otg.IsisSRv6EndSid) (IsisSRv6EndSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRv6EndSid) ToPbText() (string, error) {
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

func (m *unMarshalisisSRv6EndSid) FromPbText(value string) error {
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

func (m *marshalisisSRv6EndSid) ToYaml() (string, error) {
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

func (m *unMarshalisisSRv6EndSid) FromYaml(value string) error {
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

func (m *marshalisisSRv6EndSid) ToJson() (string, error) {
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

func (m *unMarshalisisSRv6EndSid) FromJson(value string) error {
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

func (obj *isisSRv6EndSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRv6EndSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRv6EndSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRv6EndSid) Clone() (IsisSRv6EndSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRv6EndSid()
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

// IsisSRv6EndSid is sRv6 End SID Sub-TLV (sub-TLV type 5) nested within the enclosing SRv6 Locator TLV (TLV type 27, RFC 9352 Section 7.2). Advertises a locally instantiated, node-level SRv6 SID and its associated endpoint behavior. The locator prefix is taken from the parent IsisSRv6.Locator; only the Function and Argument fields need to be specified here. The full 128-bit SID is assembled as:
// <parent locator prefix> | Function | Argument
// Example - given parent locator fc00:0:1:: /48 with sid_structure lb=32, ln=16, fn=16, arg=0:
// function "0001", argument "0000" => SID fc00:0:1:1::
// function "0064", argument "0000" => SID fc00:0:1:64::
// Valid behaviors are End variants (with PSP/USP/USD flavors) and decapsulation behaviors (End.DT4, End.DT6, End.DT46). Reference: RFC 9352 Section 7.2, RFC 8986.
type IsisSRv6EndSid interface {
	Validation
	// msg marshals IsisSRv6EndSid to protobuf object *otg.IsisSRv6EndSid
	// and doesn't set defaults
	msg() *otg.IsisSRv6EndSid
	// setMsg unmarshals IsisSRv6EndSid from protobuf object *otg.IsisSRv6EndSid
	// and doesn't set defaults
	setMsg(*otg.IsisSRv6EndSid) IsisSRv6EndSid
	// provides marshal interface
	Marshal() marshalIsisSRv6EndSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRv6EndSid
	// validate validates IsisSRv6EndSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRv6EndSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Function returns string, set in IsisSRv6EndSid.
	Function() string
	// SetFunction assigns string provided by user to IsisSRv6EndSid
	SetFunction(value string) IsisSRv6EndSid
	// HasFunction checks if Function has been set in IsisSRv6EndSid
	HasFunction() bool
	// Argument returns string, set in IsisSRv6EndSid.
	Argument() string
	// SetArgument assigns string provided by user to IsisSRv6EndSid
	SetArgument(value string) IsisSRv6EndSid
	// HasArgument checks if Argument has been set in IsisSRv6EndSid
	HasArgument() bool
	// EndpointBehavior returns IsisSRv6EndSidEndpointBehaviorEnum, set in IsisSRv6EndSid
	EndpointBehavior() IsisSRv6EndSidEndpointBehaviorEnum
	// SetEndpointBehavior assigns IsisSRv6EndSidEndpointBehaviorEnum provided by user to IsisSRv6EndSid
	SetEndpointBehavior(value IsisSRv6EndSidEndpointBehaviorEnum) IsisSRv6EndSid
	// HasEndpointBehavior checks if EndpointBehavior has been set in IsisSRv6EndSid
	HasEndpointBehavior() bool
	// CFlag returns bool, set in IsisSRv6EndSid.
	CFlag() bool
	// SetCFlag assigns bool provided by user to IsisSRv6EndSid
	SetCFlag(value bool) IsisSRv6EndSid
	// HasCFlag checks if CFlag has been set in IsisSRv6EndSid
	HasCFlag() bool
}

// The Function part of this End SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal function_length from the parent IsisSRv6.Locator.sid_structure divided by 4 - for example, a function_length of 16 requires a 4-nibble string. Example: "0001" with parent locator fc00:0:1:: /48 gives SID fc00:0:1:1::.
// Function returns a string
func (obj *isisSRv6EndSid) Function() string {

	return *obj.obj.Function

}

// The Function part of this End SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal function_length from the parent IsisSRv6.Locator.sid_structure divided by 4 - for example, a function_length of 16 requires a 4-nibble string. Example: "0001" with parent locator fc00:0:1:: /48 gives SID fc00:0:1:1::.
// Function returns a string
func (obj *isisSRv6EndSid) HasFunction() bool {
	return obj.obj.Function != nil
}

// The Function part of this End SID expressed as a hex string, occupying the function bits immediately after the locator prefix in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal function_length from the parent IsisSRv6.Locator.sid_structure divided by 4 - for example, a function_length of 16 requires a 4-nibble string. Example: "0001" with parent locator fc00:0:1:: /48 gives SID fc00:0:1:1::.
// SetFunction sets the string value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetFunction(value string) IsisSRv6EndSid {

	obj.obj.Function = &value
	return obj
}

// The Argument part of this End SID expressed as a hex string, occupying the argument bits immediately after the function in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal argument_length from the parent IsisSRv6.Locator.sid_structure divided by 4. Use the default "0000" when no argument is carried (argument_length is 0).
// Argument returns a string
func (obj *isisSRv6EndSid) Argument() string {

	return *obj.obj.Argument

}

// The Argument part of this End SID expressed as a hex string, occupying the argument bits immediately after the function in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal argument_length from the parent IsisSRv6.Locator.sid_structure divided by 4. Use the default "0000" when no argument is carried (argument_length is 0).
// Argument returns a string
func (obj *isisSRv6EndSid) HasArgument() bool {
	return obj.obj.Argument != nil
}

// The Argument part of this End SID expressed as a hex string, occupying the argument bits immediately after the function in the 128-bit SID (RFC 8986 Section 3.1). The number of hex digits must equal argument_length from the parent IsisSRv6.Locator.sid_structure divided by 4. Use the default "0000" when no argument is carried (argument_length is 0).
// SetArgument sets the string value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetArgument(value string) IsisSRv6EndSid {

	obj.obj.Argument = &value
	return obj
}

type IsisSRv6EndSidEndpointBehaviorEnum string

// Enum of EndpointBehavior on IsisSRv6EndSid
var IsisSRv6EndSidEndpointBehavior = struct {
	END                  IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP         IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_USP         IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USP     IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_USD         IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USD     IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_USP_USD     IsisSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USP_USD IsisSRv6EndSidEndpointBehaviorEnum
	END_DT4              IsisSRv6EndSidEndpointBehaviorEnum
	END_DT6              IsisSRv6EndSidEndpointBehaviorEnum
	END_DT46             IsisSRv6EndSidEndpointBehaviorEnum
}{
	END:                  IsisSRv6EndSidEndpointBehaviorEnum("end"),
	END_WITH_PSP:         IsisSRv6EndSidEndpointBehaviorEnum("end_with_psp"),
	END_WITH_USP:         IsisSRv6EndSidEndpointBehaviorEnum("end_with_usp"),
	END_WITH_PSP_USP:     IsisSRv6EndSidEndpointBehaviorEnum("end_with_psp_usp"),
	END_WITH_USD:         IsisSRv6EndSidEndpointBehaviorEnum("end_with_usd"),
	END_WITH_PSP_USD:     IsisSRv6EndSidEndpointBehaviorEnum("end_with_psp_usd"),
	END_WITH_USP_USD:     IsisSRv6EndSidEndpointBehaviorEnum("end_with_usp_usd"),
	END_WITH_PSP_USP_USD: IsisSRv6EndSidEndpointBehaviorEnum("end_with_psp_usp_usd"),
	END_DT4:              IsisSRv6EndSidEndpointBehaviorEnum("end_dt4"),
	END_DT6:              IsisSRv6EndSidEndpointBehaviorEnum("end_dt6"),
	END_DT46:             IsisSRv6EndSidEndpointBehaviorEnum("end_dt46"),
}

func (obj *isisSRv6EndSid) EndpointBehavior() IsisSRv6EndSidEndpointBehaviorEnum {
	return IsisSRv6EndSidEndpointBehaviorEnum(obj.obj.EndpointBehavior.Enum().String())
}

// The endpoint behavior for this SRv6 End SID, encoded as a 2-octet behavior codepoint (RFC 8986 Section 7.4). Defines how the router processes packets arriving with this SID as the Active Segment. Valid node-level behaviors for End SID Sub-TLV type 5 are End variants (with PSP/USP/USD flavors) and decapsulation behaviors.
// EndpointBehavior returns a string
func (obj *isisSRv6EndSid) HasEndpointBehavior() bool {
	return obj.obj.EndpointBehavior != nil
}

func (obj *isisSRv6EndSid) SetEndpointBehavior(value IsisSRv6EndSidEndpointBehaviorEnum) IsisSRv6EndSid {
	intValue, ok := otg.IsisSRv6EndSid_EndpointBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisSRv6EndSidEndpointBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisSRv6EndSid_EndpointBehavior_Enum(intValue)
	obj.obj.EndpointBehavior = &enumValue

	return obj
}

// Compression (uSID) flag. When set, this End SID is a Micro-SID (uSID) per RFC 9800. A headend router can pack it into a shared 128-bit uSID container alongside uSIDs from other nodes in the same /32 uSID block, steering a multi-hop path without a Segment Routing Header. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the parent IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "0001" => uSID fc00:0:1:1::  (End)
// c_flag=true, function "0002" => uSID fc00:0:1:2::  (End.DT46)
// For a 2-hop path via node 1 (End "0001") then node 2 (End "0001"), the headend builds container DA = fc00:0:1:1:2:1:: - two 32-bit uSID slots [LN=1,FN=1] and [LN=2,FN=1] packed after the /32 block prefix. Reference: RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6EndSid) CFlag() bool {

	return *obj.obj.CFlag

}

// Compression (uSID) flag. When set, this End SID is a Micro-SID (uSID) per RFC 9800. A headend router can pack it into a shared 128-bit uSID container alongside uSIDs from other nodes in the same /32 uSID block, steering a multi-hop path without a Segment Routing Header. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the parent IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "0001" => uSID fc00:0:1:1::  (End)
// c_flag=true, function "0002" => uSID fc00:0:1:2::  (End.DT46)
// For a 2-hop path via node 1 (End "0001") then node 2 (End "0001"), the headend builds container DA = fc00:0:1:1:2:1:: - two 32-bit uSID slots [LN=1,FN=1] and [LN=2,FN=1] packed after the /32 block prefix. Reference: RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6EndSid) HasCFlag() bool {
	return obj.obj.CFlag != nil
}

// Compression (uSID) flag. When set, this End SID is a Micro-SID (uSID) per RFC 9800. A headend router can pack it into a shared 128-bit uSID container alongside uSIDs from other nodes in the same /32 uSID block, steering a multi-hop path without a Segment Routing Header. The node-level IsisSRv6.NodeCapability.c_flag must also be set. The SID bit layout is governed by the parent IsisSRv6.Locator.sid_structure. Example using F3216 (lb=32, ln=16, fn=16, arg=0), uSID block fc00::/32, locator fc00:0:1:: /48:
// c_flag=true, function "0001" => uSID fc00:0:1:1::  (End)
// c_flag=true, function "0002" => uSID fc00:0:1:2::  (End.DT46)
// For a 2-hop path via node 1 (End "0001") then node 2 (End "0001"), the headend builds container DA = fc00:0:1:1:2:1:: - two 32-bit uSID slots [LN=1,FN=1] and [LN=2,FN=1] packed after the /32 block prefix. Reference: RFC 9800.
// SetCFlag sets the bool value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetCFlag(value bool) IsisSRv6EndSid {

	obj.obj.CFlag = &value
	return obj
}

func (obj *isisSRv6EndSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Function != nil {

		err := obj.validateHex(obj.Function())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRv6EndSid.Function"))
		}

	}

	if obj.obj.Argument != nil {

		err := obj.validateHex(obj.Argument())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRv6EndSid.Argument"))
		}

	}

}

func (obj *isisSRv6EndSid) setDefault() {
	if obj.obj.Argument == nil {
		obj.SetArgument("0000")
	}
	if obj.obj.EndpointBehavior == nil {
		obj.SetEndpointBehavior(IsisSRv6EndSidEndpointBehavior.END)

	}
	if obj.obj.CFlag == nil {
		obj.SetCFlag(false)
	}

}
