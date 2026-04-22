package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6EndSid *****
type isisLspSRv6EndSid struct {
	validation
	obj                *otg.IsisLspSRv6EndSid
	marshaller         marshalIsisLspSRv6EndSid
	unMarshaller       unMarshalIsisLspSRv6EndSid
	sidStructureHolder IsisLspSRv6SidStructure
}

func NewIsisLspSRv6EndSid() IsisLspSRv6EndSid {
	obj := isisLspSRv6EndSid{obj: &otg.IsisLspSRv6EndSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6EndSid) msg() *otg.IsisLspSRv6EndSid {
	return obj.obj
}

func (obj *isisLspSRv6EndSid) setMsg(msg *otg.IsisLspSRv6EndSid) IsisLspSRv6EndSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6EndSid struct {
	obj *isisLspSRv6EndSid
}

type marshalIsisLspSRv6EndSid interface {
	// ToProto marshals IsisLspSRv6EndSid to protobuf object *otg.IsisLspSRv6EndSid
	ToProto() (*otg.IsisLspSRv6EndSid, error)
	// ToPbText marshals IsisLspSRv6EndSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6EndSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6EndSid to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6EndSid struct {
	obj *isisLspSRv6EndSid
}

type unMarshalIsisLspSRv6EndSid interface {
	// FromProto unmarshals IsisLspSRv6EndSid from protobuf object *otg.IsisLspSRv6EndSid
	FromProto(msg *otg.IsisLspSRv6EndSid) (IsisLspSRv6EndSid, error)
	// FromPbText unmarshals IsisLspSRv6EndSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6EndSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6EndSid from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6EndSid) Marshal() marshalIsisLspSRv6EndSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6EndSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6EndSid) Unmarshal() unMarshalIsisLspSRv6EndSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6EndSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6EndSid) ToProto() (*otg.IsisLspSRv6EndSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6EndSid) FromProto(msg *otg.IsisLspSRv6EndSid) (IsisLspSRv6EndSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6EndSid) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6EndSid) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6EndSid) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6EndSid) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6EndSid) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6EndSid) FromJson(value string) error {
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

func (obj *isisLspSRv6EndSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6EndSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6EndSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6EndSid) Clone() (IsisLspSRv6EndSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6EndSid()
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

func (obj *isisLspSRv6EndSid) setNil() {
	obj.sidStructureHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRv6EndSid is sRv6 End SID Sub-TLV (sub-TLV type 5) learned from an SRv6 Locator TLV (type 27) in a received LSP. Describes a node-local SRv6 SID, its endpoint behavior, and optional SID structure information. Reference: RFC 9352 Section 7.2, RFC 8986.
type IsisLspSRv6EndSid interface {
	Validation
	// msg marshals IsisLspSRv6EndSid to protobuf object *otg.IsisLspSRv6EndSid
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6EndSid
	// setMsg unmarshals IsisLspSRv6EndSid from protobuf object *otg.IsisLspSRv6EndSid
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6EndSid) IsisLspSRv6EndSid
	// provides marshal interface
	Marshal() marshalIsisLspSRv6EndSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6EndSid
	// validate validates IsisLspSRv6EndSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6EndSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Sid returns string, set in IsisLspSRv6EndSid.
	Sid() string
	// SetSid assigns string provided by user to IsisLspSRv6EndSid
	SetSid(value string) IsisLspSRv6EndSid
	// HasSid checks if Sid has been set in IsisLspSRv6EndSid
	HasSid() bool
	// EndpointBehavior returns IsisLspSRv6EndSidEndpointBehaviorEnum, set in IsisLspSRv6EndSid
	EndpointBehavior() IsisLspSRv6EndSidEndpointBehaviorEnum
	// SetEndpointBehavior assigns IsisLspSRv6EndSidEndpointBehaviorEnum provided by user to IsisLspSRv6EndSid
	SetEndpointBehavior(value IsisLspSRv6EndSidEndpointBehaviorEnum) IsisLspSRv6EndSid
	// HasEndpointBehavior checks if EndpointBehavior has been set in IsisLspSRv6EndSid
	HasEndpointBehavior() bool
	// SidStructure returns IsisLspSRv6SidStructure, set in IsisLspSRv6EndSid.
	// IsisLspSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1) learned from a received LSP. Describes the internal bit-field decomposition of an SRv6 SID, enabling routers to independently interpret each component of the 128-bit SID value. The sum of all four length fields must not exceed 128 bits. Reference: RFC 9352 Section 9.
	SidStructure() IsisLspSRv6SidStructure
	// SetSidStructure assigns IsisLspSRv6SidStructure provided by user to IsisLspSRv6EndSid.
	// IsisLspSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1) learned from a received LSP. Describes the internal bit-field decomposition of an SRv6 SID, enabling routers to independently interpret each component of the 128-bit SID value. The sum of all four length fields must not exceed 128 bits. Reference: RFC 9352 Section 9.
	SetSidStructure(value IsisLspSRv6SidStructure) IsisLspSRv6EndSid
	// HasSidStructure checks if SidStructure has been set in IsisLspSRv6EndSid
	HasSidStructure() bool
	setNil()
}

// The 128-bit SRv6 SID value in IPv6 address format.
// Sid returns a string
func (obj *isisLspSRv6EndSid) Sid() string {

	return *obj.obj.Sid

}

// The 128-bit SRv6 SID value in IPv6 address format.
// Sid returns a string
func (obj *isisLspSRv6EndSid) HasSid() bool {
	return obj.obj.Sid != nil
}

// The 128-bit SRv6 SID value in IPv6 address format.
// SetSid sets the string value in the IsisLspSRv6EndSid object
func (obj *isisLspSRv6EndSid) SetSid(value string) IsisLspSRv6EndSid {

	obj.obj.Sid = &value
	return obj
}

type IsisLspSRv6EndSidEndpointBehaviorEnum string

// Enum of EndpointBehavior on IsisLspSRv6EndSid
var IsisLspSRv6EndSidEndpointBehavior = struct {
	END                  IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP         IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_USP         IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USP     IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_USD         IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USD     IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_USP_USD     IsisLspSRv6EndSidEndpointBehaviorEnum
	END_WITH_PSP_USP_USD IsisLspSRv6EndSidEndpointBehaviorEnum
	END_DT4              IsisLspSRv6EndSidEndpointBehaviorEnum
	END_DT6              IsisLspSRv6EndSidEndpointBehaviorEnum
	END_DT46             IsisLspSRv6EndSidEndpointBehaviorEnum
}{
	END:                  IsisLspSRv6EndSidEndpointBehaviorEnum("end"),
	END_WITH_PSP:         IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_psp"),
	END_WITH_USP:         IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_usp"),
	END_WITH_PSP_USP:     IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_psp_usp"),
	END_WITH_USD:         IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_usd"),
	END_WITH_PSP_USD:     IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_psp_usd"),
	END_WITH_USP_USD:     IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_usp_usd"),
	END_WITH_PSP_USP_USD: IsisLspSRv6EndSidEndpointBehaviorEnum("end_with_psp_usp_usd"),
	END_DT4:              IsisLspSRv6EndSidEndpointBehaviorEnum("end_dt4"),
	END_DT6:              IsisLspSRv6EndSidEndpointBehaviorEnum("end_dt6"),
	END_DT46:             IsisLspSRv6EndSidEndpointBehaviorEnum("end_dt46"),
}

func (obj *isisLspSRv6EndSid) EndpointBehavior() IsisLspSRv6EndSidEndpointBehaviorEnum {
	return IsisLspSRv6EndSidEndpointBehaviorEnum(obj.obj.EndpointBehavior.Enum().String())
}

// Endpoint behavior codepoint for this End SID (RFC 8986 Section 7.4). Indicates how the originating router processes packets arriving with this SID as the Active Segment.
// EndpointBehavior returns a string
func (obj *isisLspSRv6EndSid) HasEndpointBehavior() bool {
	return obj.obj.EndpointBehavior != nil
}

func (obj *isisLspSRv6EndSid) SetEndpointBehavior(value IsisLspSRv6EndSidEndpointBehaviorEnum) IsisLspSRv6EndSid {
	intValue, ok := otg.IsisLspSRv6EndSid_EndpointBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspSRv6EndSidEndpointBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspSRv6EndSid_EndpointBehavior_Enum(intValue)
	obj.obj.EndpointBehavior = &enumValue

	return obj
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the bit-field decomposition of this SID value. Specifies Locator Block, Locator Node, Function, and Argument lengths in bits. Reference: RFC 9352 Section 9.
// SidStructure returns a IsisLspSRv6SidStructure
func (obj *isisLspSRv6EndSid) SidStructure() IsisLspSRv6SidStructure {
	if obj.obj.SidStructure == nil {
		obj.obj.SidStructure = NewIsisLspSRv6SidStructure().msg()
	}
	if obj.sidStructureHolder == nil {
		obj.sidStructureHolder = &isisLspSRv6SidStructure{obj: obj.obj.SidStructure}
	}
	return obj.sidStructureHolder
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the bit-field decomposition of this SID value. Specifies Locator Block, Locator Node, Function, and Argument lengths in bits. Reference: RFC 9352 Section 9.
// SidStructure returns a IsisLspSRv6SidStructure
func (obj *isisLspSRv6EndSid) HasSidStructure() bool {
	return obj.obj.SidStructure != nil
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the bit-field decomposition of this SID value. Specifies Locator Block, Locator Node, Function, and Argument lengths in bits. Reference: RFC 9352 Section 9.
// SetSidStructure sets the IsisLspSRv6SidStructure value in the IsisLspSRv6EndSid object
func (obj *isisLspSRv6EndSid) SetSidStructure(value IsisLspSRv6SidStructure) IsisLspSRv6EndSid {

	obj.sidStructureHolder = nil
	obj.obj.SidStructure = value.msg()

	return obj
}

func (obj *isisLspSRv6EndSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Sid != nil {

		err := obj.validateIpv6(obj.Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspSRv6EndSid.Sid"))
		}

	}

	if obj.obj.SidStructure != nil {

		obj.SidStructure().validateObj(vObj, set_default)
	}

}

func (obj *isisLspSRv6EndSid) setDefault() {

}
