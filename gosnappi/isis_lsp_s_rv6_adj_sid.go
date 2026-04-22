package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRv6AdjSid *****
type isisLspSRv6AdjSid struct {
	validation
	obj                *otg.IsisLspSRv6AdjSid
	marshaller         marshalIsisLspSRv6AdjSid
	unMarshaller       unMarshalIsisLspSRv6AdjSid
	sidStructureHolder IsisLspSRv6SidStructure
}

func NewIsisLspSRv6AdjSid() IsisLspSRv6AdjSid {
	obj := isisLspSRv6AdjSid{obj: &otg.IsisLspSRv6AdjSid{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRv6AdjSid) msg() *otg.IsisLspSRv6AdjSid {
	return obj.obj
}

func (obj *isisLspSRv6AdjSid) setMsg(msg *otg.IsisLspSRv6AdjSid) IsisLspSRv6AdjSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRv6AdjSid struct {
	obj *isisLspSRv6AdjSid
}

type marshalIsisLspSRv6AdjSid interface {
	// ToProto marshals IsisLspSRv6AdjSid to protobuf object *otg.IsisLspSRv6AdjSid
	ToProto() (*otg.IsisLspSRv6AdjSid, error)
	// ToPbText marshals IsisLspSRv6AdjSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRv6AdjSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRv6AdjSid to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspSRv6AdjSid struct {
	obj *isisLspSRv6AdjSid
}

type unMarshalIsisLspSRv6AdjSid interface {
	// FromProto unmarshals IsisLspSRv6AdjSid from protobuf object *otg.IsisLspSRv6AdjSid
	FromProto(msg *otg.IsisLspSRv6AdjSid) (IsisLspSRv6AdjSid, error)
	// FromPbText unmarshals IsisLspSRv6AdjSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRv6AdjSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRv6AdjSid from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRv6AdjSid) Marshal() marshalIsisLspSRv6AdjSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRv6AdjSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRv6AdjSid) Unmarshal() unMarshalIsisLspSRv6AdjSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRv6AdjSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRv6AdjSid) ToProto() (*otg.IsisLspSRv6AdjSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRv6AdjSid) FromProto(msg *otg.IsisLspSRv6AdjSid) (IsisLspSRv6AdjSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRv6AdjSid) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRv6AdjSid) FromPbText(value string) error {
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

func (m *marshalisisLspSRv6AdjSid) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRv6AdjSid) FromYaml(value string) error {
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

func (m *marshalisisLspSRv6AdjSid) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRv6AdjSid) FromJson(value string) error {
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

func (obj *isisLspSRv6AdjSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRv6AdjSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRv6AdjSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRv6AdjSid) Clone() (IsisLspSRv6AdjSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRv6AdjSid()
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

func (obj *isisLspSRv6AdjSid) setNil() {
	obj.sidStructureHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRv6AdjSid is sRv6 Adjacency SID Sub-TLV learned from the Extended IS Reachability TLV (type 22) or MT-ISN TLV (type 222) in a received LSP. Point-to-point End.X SID Sub-TLV has sub-TLV type 43; LAN End.X SID Sub-TLV has sub-TLV type 44 and includes the neighbor System ID. Reference: RFC 9352 Sections 8.1-8.2, RFC 8986.
type IsisLspSRv6AdjSid interface {
	Validation
	// msg marshals IsisLspSRv6AdjSid to protobuf object *otg.IsisLspSRv6AdjSid
	// and doesn't set defaults
	msg() *otg.IsisLspSRv6AdjSid
	// setMsg unmarshals IsisLspSRv6AdjSid from protobuf object *otg.IsisLspSRv6AdjSid
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRv6AdjSid) IsisLspSRv6AdjSid
	// provides marshal interface
	Marshal() marshalIsisLspSRv6AdjSid
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRv6AdjSid
	// validate validates IsisLspSRv6AdjSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRv6AdjSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Type returns IsisLspSRv6AdjSidTypeEnum, set in IsisLspSRv6AdjSid
	Type() IsisLspSRv6AdjSidTypeEnum
	// SetType assigns IsisLspSRv6AdjSidTypeEnum provided by user to IsisLspSRv6AdjSid
	SetType(value IsisLspSRv6AdjSidTypeEnum) IsisLspSRv6AdjSid
	// HasType checks if Type has been set in IsisLspSRv6AdjSid
	HasType() bool
	// NeighborId returns string, set in IsisLspSRv6AdjSid.
	NeighborId() string
	// SetNeighborId assigns string provided by user to IsisLspSRv6AdjSid
	SetNeighborId(value string) IsisLspSRv6AdjSid
	// HasNeighborId checks if NeighborId has been set in IsisLspSRv6AdjSid
	HasNeighborId() bool
	// Sid returns string, set in IsisLspSRv6AdjSid.
	Sid() string
	// SetSid assigns string provided by user to IsisLspSRv6AdjSid
	SetSid(value string) IsisLspSRv6AdjSid
	// HasSid checks if Sid has been set in IsisLspSRv6AdjSid
	HasSid() bool
	// EndpointBehavior returns IsisLspSRv6AdjSidEndpointBehaviorEnum, set in IsisLspSRv6AdjSid
	EndpointBehavior() IsisLspSRv6AdjSidEndpointBehaviorEnum
	// SetEndpointBehavior assigns IsisLspSRv6AdjSidEndpointBehaviorEnum provided by user to IsisLspSRv6AdjSid
	SetEndpointBehavior(value IsisLspSRv6AdjSidEndpointBehaviorEnum) IsisLspSRv6AdjSid
	// HasEndpointBehavior checks if EndpointBehavior has been set in IsisLspSRv6AdjSid
	HasEndpointBehavior() bool
	// BFlag returns bool, set in IsisLspSRv6AdjSid.
	BFlag() bool
	// SetBFlag assigns bool provided by user to IsisLspSRv6AdjSid
	SetBFlag(value bool) IsisLspSRv6AdjSid
	// HasBFlag checks if BFlag has been set in IsisLspSRv6AdjSid
	HasBFlag() bool
	// SFlag returns bool, set in IsisLspSRv6AdjSid.
	SFlag() bool
	// SetSFlag assigns bool provided by user to IsisLspSRv6AdjSid
	SetSFlag(value bool) IsisLspSRv6AdjSid
	// HasSFlag checks if SFlag has been set in IsisLspSRv6AdjSid
	HasSFlag() bool
	// PFlag returns bool, set in IsisLspSRv6AdjSid.
	PFlag() bool
	// SetPFlag assigns bool provided by user to IsisLspSRv6AdjSid
	SetPFlag(value bool) IsisLspSRv6AdjSid
	// HasPFlag checks if PFlag has been set in IsisLspSRv6AdjSid
	HasPFlag() bool
	// Weight returns uint32, set in IsisLspSRv6AdjSid.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to IsisLspSRv6AdjSid
	SetWeight(value uint32) IsisLspSRv6AdjSid
	// HasWeight checks if Weight has been set in IsisLspSRv6AdjSid
	HasWeight() bool
	// Algorithm returns uint32, set in IsisLspSRv6AdjSid.
	Algorithm() uint32
	// SetAlgorithm assigns uint32 provided by user to IsisLspSRv6AdjSid
	SetAlgorithm(value uint32) IsisLspSRv6AdjSid
	// HasAlgorithm checks if Algorithm has been set in IsisLspSRv6AdjSid
	HasAlgorithm() bool
	// SidStructure returns IsisLspSRv6SidStructure, set in IsisLspSRv6AdjSid.
	// IsisLspSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1) learned from a received LSP. Describes the internal bit-field decomposition of an SRv6 SID, enabling routers to independently interpret each component of the 128-bit SID value. The sum of all four length fields must not exceed 128 bits. Reference: RFC 9352 Section 9.
	SidStructure() IsisLspSRv6SidStructure
	// SetSidStructure assigns IsisLspSRv6SidStructure provided by user to IsisLspSRv6AdjSid.
	// IsisLspSRv6SidStructure is sRv6 SID Structure Sub-Sub-TLV (type 1) learned from a received LSP. Describes the internal bit-field decomposition of an SRv6 SID, enabling routers to independently interpret each component of the 128-bit SID value. The sum of all four length fields must not exceed 128 bits. Reference: RFC 9352 Section 9.
	SetSidStructure(value IsisLspSRv6SidStructure) IsisLspSRv6AdjSid
	// HasSidStructure checks if SidStructure has been set in IsisLspSRv6AdjSid
	HasSidStructure() bool
	setNil()
}

type IsisLspSRv6AdjSidTypeEnum string

// Enum of Type on IsisLspSRv6AdjSid
var IsisLspSRv6AdjSidType = struct {
	END_X_SID     IsisLspSRv6AdjSidTypeEnum
	LAN_END_X_SID IsisLspSRv6AdjSidTypeEnum
}{
	END_X_SID:     IsisLspSRv6AdjSidTypeEnum("end_x_sid"),
	LAN_END_X_SID: IsisLspSRv6AdjSidTypeEnum("lan_end_x_sid"),
}

func (obj *isisLspSRv6AdjSid) Type() IsisLspSRv6AdjSidTypeEnum {
	return IsisLspSRv6AdjSidTypeEnum(obj.obj.Type.Enum().String())
}

// The sub-TLV type that carried this adjacency SID. 'end_x_sid' = point-to-point (sub-TLV type 43); 'lan_end_x_sid' = LAN adjacency (sub-TLV type 44, includes neighbor System ID in neighbor_id field).
// Type returns a string
func (obj *isisLspSRv6AdjSid) HasType() bool {
	return obj.obj.Type != nil
}

func (obj *isisLspSRv6AdjSid) SetType(value IsisLspSRv6AdjSidTypeEnum) IsisLspSRv6AdjSid {
	intValue, ok := otg.IsisLspSRv6AdjSid_Type_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspSRv6AdjSidTypeEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspSRv6AdjSid_Type_Enum(intValue)
	obj.obj.Type = &enumValue

	return obj
}

// IS-IS System ID of the LAN neighbor from which this End.X SID was received. Present only when type is 'lan_end_x_sid' (sub-TLV type 44, RFC 9352 Section 8.2).
// NeighborId returns a string
func (obj *isisLspSRv6AdjSid) NeighborId() string {

	return *obj.obj.NeighborId

}

// IS-IS System ID of the LAN neighbor from which this End.X SID was received. Present only when type is 'lan_end_x_sid' (sub-TLV type 44, RFC 9352 Section 8.2).
// NeighborId returns a string
func (obj *isisLspSRv6AdjSid) HasNeighborId() bool {
	return obj.obj.NeighborId != nil
}

// IS-IS System ID of the LAN neighbor from which this End.X SID was received. Present only when type is 'lan_end_x_sid' (sub-TLV type 44, RFC 9352 Section 8.2).
// SetNeighborId sets the string value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetNeighborId(value string) IsisLspSRv6AdjSid {

	obj.obj.NeighborId = &value
	return obj
}

// The 128-bit SRv6 End.X SID value in IPv6 address format.
// Sid returns a string
func (obj *isisLspSRv6AdjSid) Sid() string {

	return *obj.obj.Sid

}

// The 128-bit SRv6 End.X SID value in IPv6 address format.
// Sid returns a string
func (obj *isisLspSRv6AdjSid) HasSid() bool {
	return obj.obj.Sid != nil
}

// The 128-bit SRv6 End.X SID value in IPv6 address format.
// SetSid sets the string value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetSid(value string) IsisLspSRv6AdjSid {

	obj.obj.Sid = &value
	return obj
}

type IsisLspSRv6AdjSidEndpointBehaviorEnum string

// Enum of EndpointBehavior on IsisLspSRv6AdjSid
var IsisLspSRv6AdjSidEndpointBehavior = struct {
	END_X                  IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP         IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USP         IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USP     IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USD         IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USD     IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_USP_USD     IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_X_WITH_PSP_USP_USD IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_DX4                IsisLspSRv6AdjSidEndpointBehaviorEnum
	END_DX6                IsisLspSRv6AdjSidEndpointBehaviorEnum
}{
	END_X:                  IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x"),
	END_X_WITH_PSP:         IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp"),
	END_X_WITH_USP:         IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_usp"),
	END_X_WITH_PSP_USP:     IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usp"),
	END_X_WITH_USD:         IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_usd"),
	END_X_WITH_PSP_USD:     IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usd"),
	END_X_WITH_USP_USD:     IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_usp_usd"),
	END_X_WITH_PSP_USP_USD: IsisLspSRv6AdjSidEndpointBehaviorEnum("end_x_with_psp_usp_usd"),
	END_DX4:                IsisLspSRv6AdjSidEndpointBehaviorEnum("end_dx4"),
	END_DX6:                IsisLspSRv6AdjSidEndpointBehaviorEnum("end_dx6"),
}

func (obj *isisLspSRv6AdjSid) EndpointBehavior() IsisLspSRv6AdjSidEndpointBehaviorEnum {
	return IsisLspSRv6AdjSidEndpointBehaviorEnum(obj.obj.EndpointBehavior.Enum().String())
}

// Endpoint behavior codepoint for this adjacency SID (RFC 8986 Section 7.4).
// EndpointBehavior returns a string
func (obj *isisLspSRv6AdjSid) HasEndpointBehavior() bool {
	return obj.obj.EndpointBehavior != nil
}

func (obj *isisLspSRv6AdjSid) SetEndpointBehavior(value IsisLspSRv6AdjSidEndpointBehaviorEnum) IsisLspSRv6AdjSid {
	intValue, ok := otg.IsisLspSRv6AdjSid_EndpointBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspSRv6AdjSidEndpointBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspSRv6AdjSid_EndpointBehavior_Enum(intValue)
	obj.obj.EndpointBehavior = &enumValue

	return obj
}

// Backup flag. When set, this End.X SID is eligible for IP-FRR protection (RFC 9352 Section 8.1).
// BFlag returns a bool
func (obj *isisLspSRv6AdjSid) BFlag() bool {

	return *obj.obj.BFlag

}

// Backup flag. When set, this End.X SID is eligible for IP-FRR protection (RFC 9352 Section 8.1).
// BFlag returns a bool
func (obj *isisLspSRv6AdjSid) HasBFlag() bool {
	return obj.obj.BFlag != nil
}

// Backup flag. When set, this End.X SID is eligible for IP-FRR protection (RFC 9352 Section 8.1).
// SetBFlag sets the bool value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetBFlag(value bool) IsisLspSRv6AdjSid {

	obj.obj.BFlag = &value
	return obj
}

// Set flag. When set, this End.X SID refers to a set of adjacencies (RFC 9352 Section 8.1).
// SFlag returns a bool
func (obj *isisLspSRv6AdjSid) SFlag() bool {

	return *obj.obj.SFlag

}

// Set flag. When set, this End.X SID refers to a set of adjacencies (RFC 9352 Section 8.1).
// SFlag returns a bool
func (obj *isisLspSRv6AdjSid) HasSFlag() bool {
	return obj.obj.SFlag != nil
}

// Set flag. When set, this End.X SID refers to a set of adjacencies (RFC 9352 Section 8.1).
// SetSFlag sets the bool value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetSFlag(value bool) IsisLspSRv6AdjSid {

	obj.obj.SFlag = &value
	return obj
}

// Persistent flag. When set, the SID value is stable across router restarts and interface flaps (RFC 9352 Section 8.1).
// PFlag returns a bool
func (obj *isisLspSRv6AdjSid) PFlag() bool {

	return *obj.obj.PFlag

}

// Persistent flag. When set, the SID value is stable across router restarts and interface flaps (RFC 9352 Section 8.1).
// PFlag returns a bool
func (obj *isisLspSRv6AdjSid) HasPFlag() bool {
	return obj.obj.PFlag != nil
}

// Persistent flag. When set, the SID value is stable across router restarts and interface flaps (RFC 9352 Section 8.1).
// SetPFlag sets the bool value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetPFlag(value bool) IsisLspSRv6AdjSid {

	obj.obj.PFlag = &value
	return obj
}

// Weight for load balancing across the adjacency set when S-flag is set (RFC 9352 Section 8.1).
// Weight returns a uint32
func (obj *isisLspSRv6AdjSid) Weight() uint32 {

	return *obj.obj.Weight

}

// Weight for load balancing across the adjacency set when S-flag is set (RFC 9352 Section 8.1).
// Weight returns a uint32
func (obj *isisLspSRv6AdjSid) HasWeight() bool {
	return obj.obj.Weight != nil
}

// Weight for load balancing across the adjacency set when S-flag is set (RFC 9352 Section 8.1).
// SetWeight sets the uint32 value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetWeight(value uint32) IsisLspSRv6AdjSid {

	obj.obj.Weight = &value
	return obj
}

// IGP algorithm bound to this adjacency SID. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// Algorithm returns a uint32
func (obj *isisLspSRv6AdjSid) Algorithm() uint32 {

	return *obj.obj.Algorithm

}

// IGP algorithm bound to this adjacency SID. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// Algorithm returns a uint32
func (obj *isisLspSRv6AdjSid) HasAlgorithm() bool {
	return obj.obj.Algorithm != nil
}

// IGP algorithm bound to this adjacency SID. 0 = SPF, 1 = Strict SPF, 128-255 = Flex-Algo (RFC 9350, RFC 8665).
// SetAlgorithm sets the uint32 value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetAlgorithm(value uint32) IsisLspSRv6AdjSid {

	obj.obj.Algorithm = &value
	return obj
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the internal bit-field decomposition of this adjacency SID. Reference: RFC 9352 Section 9.
// SidStructure returns a IsisLspSRv6SidStructure
func (obj *isisLspSRv6AdjSid) SidStructure() IsisLspSRv6SidStructure {
	if obj.obj.SidStructure == nil {
		obj.obj.SidStructure = NewIsisLspSRv6SidStructure().msg()
	}
	if obj.sidStructureHolder == nil {
		obj.sidStructureHolder = &isisLspSRv6SidStructure{obj: obj.obj.SidStructure}
	}
	return obj.sidStructureHolder
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the internal bit-field decomposition of this adjacency SID. Reference: RFC 9352 Section 9.
// SidStructure returns a IsisLspSRv6SidStructure
func (obj *isisLspSRv6AdjSid) HasSidStructure() bool {
	return obj.obj.SidStructure != nil
}

// SRv6 SID Structure Sub-Sub-TLV (type 1) if present, describing the internal bit-field decomposition of this adjacency SID. Reference: RFC 9352 Section 9.
// SetSidStructure sets the IsisLspSRv6SidStructure value in the IsisLspSRv6AdjSid object
func (obj *isisLspSRv6AdjSid) SetSidStructure(value IsisLspSRv6SidStructure) IsisLspSRv6AdjSid {

	obj.sidStructureHolder = nil
	obj.obj.SidStructure = value.msg()

	return obj
}

func (obj *isisLspSRv6AdjSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.NeighborId != nil {

		err := obj.validateHex(obj.NeighborId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspSRv6AdjSid.NeighborId"))
		}

	}

	if obj.obj.Sid != nil {

		err := obj.validateIpv6(obj.Sid())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspSRv6AdjSid.Sid"))
		}

	}

	if obj.obj.Weight != nil {

		if *obj.obj.Weight > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6AdjSid.Weight <= 255 but Got %d", *obj.obj.Weight))
		}

	}

	if obj.obj.Algorithm != nil {

		if *obj.obj.Algorithm > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisLspSRv6AdjSid.Algorithm <= 255 but Got %d", *obj.obj.Algorithm))
		}

	}

	if obj.obj.SidStructure != nil {

		obj.SidStructure().validateObj(vObj, set_default)
	}

}

func (obj *isisLspSRv6AdjSid) setDefault() {

}
