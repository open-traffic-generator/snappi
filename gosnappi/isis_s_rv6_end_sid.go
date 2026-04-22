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
	obj                *otg.IsisSRv6EndSid
	marshaller         marshalIsisSRv6EndSid
	unMarshaller       unMarshalIsisSRv6EndSid
	sidStructureHolder IsisSRv6SidStructure
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
	obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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
	m.obj.setNil()
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

func (obj *isisSRv6EndSid) setNil() {
	obj.sidStructureHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRv6EndSid is sRv6 End SID Sub-TLV (sub-TLV type 5) carried within the SRv6 Locator TLV (TLV type 27). Advertises a locally instantiated node-local SRv6 SID and its associated endpoint behavior. The SID value must be allocated from within the parent locator's prefix. Valid behaviors for node-local End SIDs are End (with PSP/USP/USD flavor variants) and decapsulation behaviors (End.DT4, End.DT6, End.DT46). Reference: RFC 9352 Section 7.2, RFC 8986.
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
	// Sid returns string, set in IsisSRv6EndSid.
	Sid() string
	// SetSid assigns string provided by user to IsisSRv6EndSid
	SetSid(value string) IsisSRv6EndSid
	// HasSid checks if Sid has been set in IsisSRv6EndSid
	HasSid() bool
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
	// SidStructure returns IsisSRv6SidStructure, set in IsisSRv6EndSid.
	// IsisSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1), carried within SRv6 SID Sub-TLVs (End SID type 5, End.X SID type 43/44). Describes the internal bit-field decomposition of the SRv6 SID value so that receiving routers can interpret each field independently. The four length fields (lb_length + ln_length + function_length + argument_length) MUST NOT exceed 128 bits. Required when advertising Micro-SID (uSID) SIDs to describe the compressed encoding. Example for common uSID F3216 format:
	// lb_length=32, ln_length=16, function_length=16, argument_length=0
	// Reference: RFC 9352 Section 9, RFC 9800.
	SidStructure() IsisSRv6SidStructure
	// SetSidStructure assigns IsisSRv6SidStructure provided by user to IsisSRv6EndSid.
	// IsisSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1), carried within SRv6 SID Sub-TLVs (End SID type 5, End.X SID type 43/44). Describes the internal bit-field decomposition of the SRv6 SID value so that receiving routers can interpret each field independently. The four length fields (lb_length + ln_length + function_length + argument_length) MUST NOT exceed 128 bits. Required when advertising Micro-SID (uSID) SIDs to describe the compressed encoding. Example for common uSID F3216 format:
	// lb_length=32, ln_length=16, function_length=16, argument_length=0
	// Reference: RFC 9352 Section 9, RFC 9800.
	SetSidStructure(value IsisSRv6SidStructure) IsisSRv6EndSid
	// HasSidStructure checks if SidStructure has been set in IsisSRv6EndSid
	HasSidStructure() bool
	setNil()
}

// The 128-bit SRv6 SID value in IPv6 address format. Must be allocated from within the parent locator's prefix space, following the structure: Locator | Function | Argument (RFC 8986 Section 3.1).
// Sid returns a string
func (obj *isisSRv6EndSid) Sid() string {

	return *obj.obj.Sid

}

// The 128-bit SRv6 SID value in IPv6 address format. Must be allocated from within the parent locator's prefix space, following the structure: Locator | Function | Argument (RFC 8986 Section 3.1).
// Sid returns a string
func (obj *isisSRv6EndSid) HasSid() bool {
	return obj.obj.Sid != nil
}

// The 128-bit SRv6 SID value in IPv6 address format. Must be allocated from within the parent locator's prefix space, following the structure: Locator | Function | Argument (RFC 8986 Section 3.1).
// SetSid sets the string value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetSid(value string) IsisSRv6EndSid {

	obj.obj.Sid = &value
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

// Compression (uSID) flag. When set, indicates that this SID supports Micro-SID (uSID) compressed encoding, enabling multiple node segments to be packed within a single 128-bit SID carrier using the Argument field. Requires the SRv6 SID Structure Sub-Sub-TLV to be included to describe the bit allocation. Reference: draft-filsfils-spring-net-pgm-extension-srv6-usid, RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6EndSid) CFlag() bool {

	return *obj.obj.CFlag

}

// Compression (uSID) flag. When set, indicates that this SID supports Micro-SID (uSID) compressed encoding, enabling multiple node segments to be packed within a single 128-bit SID carrier using the Argument field. Requires the SRv6 SID Structure Sub-Sub-TLV to be included to describe the bit allocation. Reference: draft-filsfils-spring-net-pgm-extension-srv6-usid, RFC 9800.
// CFlag returns a bool
func (obj *isisSRv6EndSid) HasCFlag() bool {
	return obj.obj.CFlag != nil
}

// Compression (uSID) flag. When set, indicates that this SID supports Micro-SID (uSID) compressed encoding, enabling multiple node segments to be packed within a single 128-bit SID carrier using the Argument field. Requires the SRv6 SID Structure Sub-Sub-TLV to be included to describe the bit allocation. Reference: draft-filsfils-spring-net-pgm-extension-srv6-usid, RFC 9800.
// SetCFlag sets the bool value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetCFlag(value bool) IsisSRv6EndSid {

	obj.obj.CFlag = &value
	return obj
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this SID value. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. The sum of all four lengths must not exceed 128. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SidStructure returns a IsisSRv6SidStructure
func (obj *isisSRv6EndSid) SidStructure() IsisSRv6SidStructure {
	if obj.obj.SidStructure == nil {
		obj.obj.SidStructure = NewIsisSRv6SidStructure().msg()
	}
	if obj.sidStructureHolder == nil {
		obj.sidStructureHolder = &isisSRv6SidStructure{obj: obj.obj.SidStructure}
	}
	return obj.sidStructureHolder
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this SID value. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. The sum of all four lengths must not exceed 128. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SidStructure returns a IsisSRv6SidStructure
func (obj *isisSRv6EndSid) HasSidStructure() bool {
	return obj.obj.SidStructure != nil
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) describing the internal bit-field decomposition of this SID value. Specifies the Locator Block, Locator Node, Function, and Argument lengths in bits. The sum of all four lengths must not exceed 128. When present, the sub-sub-TLV is included in the advertisement. Typically set when c_flag is true (RFC 9352 Section 9).
// SetSidStructure sets the IsisSRv6SidStructure value in the IsisSRv6EndSid object
func (obj *isisSRv6EndSid) SetSidStructure(value IsisSRv6SidStructure) IsisSRv6EndSid {

	obj.sidStructureHolder = nil
	obj.obj.SidStructure = value.msg()

	return obj
}

func (obj *isisSRv6EndSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Sid != nil {

		err := obj.validateIpv6(obj.Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisSRv6EndSid.Sid"))
		}

	}

	if obj.obj.SidStructure != nil {

		obj.SidStructure().validateObj(vObj, set_default)
	}

}

func (obj *isisSRv6EndSid) setDefault() {
	if obj.obj.EndpointBehavior == nil {
		obj.SetEndpointBehavior(IsisSRv6EndSidEndpointBehavior.END)

	}
	if obj.obj.CFlag == nil {
		obj.SetCFlag(false)
	}

}
