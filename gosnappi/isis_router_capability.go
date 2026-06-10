package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisRouterCapability *****
type isisRouterCapability struct {
	validation
	obj                *otg.IsisRouterCapability
	marshaller         marshalIsisRouterCapability
	unMarshaller       unMarshalIsisRouterCapability
	srCapabilityHolder IsisSRCapability
	srlbRangesHolder   IsisRouterCapabilityIsisSRSrlbIter
}

func NewIsisRouterCapability() IsisRouterCapability {
	obj := isisRouterCapability{obj: &otg.IsisRouterCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisRouterCapability) msg() *otg.IsisRouterCapability {
	return obj.obj
}

func (obj *isisRouterCapability) setMsg(msg *otg.IsisRouterCapability) IsisRouterCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisRouterCapability struct {
	obj *isisRouterCapability
}

type marshalIsisRouterCapability interface {
	// ToProto marshals IsisRouterCapability to protobuf object *otg.IsisRouterCapability
	ToProto() (*otg.IsisRouterCapability, error)
	// ToPbText marshals IsisRouterCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisRouterCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisRouterCapability to JSON text
	ToJson() (string, error)
}

type unMarshalisisRouterCapability struct {
	obj *isisRouterCapability
}

type unMarshalIsisRouterCapability interface {
	// FromProto unmarshals IsisRouterCapability from protobuf object *otg.IsisRouterCapability
	FromProto(msg *otg.IsisRouterCapability) (IsisRouterCapability, error)
	// FromPbText unmarshals IsisRouterCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisRouterCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisRouterCapability from JSON text
	FromJson(value string) error
}

func (obj *isisRouterCapability) Marshal() marshalIsisRouterCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisRouterCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisRouterCapability) Unmarshal() unMarshalIsisRouterCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisRouterCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisRouterCapability) ToProto() (*otg.IsisRouterCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisRouterCapability) FromProto(msg *otg.IsisRouterCapability) (IsisRouterCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisRouterCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisRouterCapability) FromPbText(value string) error {
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

func (m *marshalisisRouterCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisRouterCapability) FromYaml(value string) error {
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

func (m *marshalisisRouterCapability) ToJson() (string, error) {
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

func (m *unMarshalisisRouterCapability) FromJson(value string) error {
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

func (obj *isisRouterCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisRouterCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisRouterCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisRouterCapability) Clone() (IsisRouterCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisRouterCapability()
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

func (obj *isisRouterCapability) setNil() {
	obj.srCapabilityHolder = nil
	obj.srlbRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisRouterCapability is container for the configuration of IS-IS Router CAPABILITY TLV.
// https://datatracker.ietf.org/doc/html/rfc7981#section-2.
// An implementation should set default values appropriately if any mandatory item is not configured by a user.
type IsisRouterCapability interface {
	Validation
	// msg marshals IsisRouterCapability to protobuf object *otg.IsisRouterCapability
	// and doesn't set defaults
	msg() *otg.IsisRouterCapability
	// setMsg unmarshals IsisRouterCapability from protobuf object *otg.IsisRouterCapability
	// and doesn't set defaults
	setMsg(*otg.IsisRouterCapability) IsisRouterCapability
	// provides marshal interface
	Marshal() marshalIsisRouterCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisRouterCapability
	// validate validates IsisRouterCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisRouterCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisRouterCapabilityChoiceEnum, set in IsisRouterCapability
	Choice() IsisRouterCapabilityChoiceEnum
	// setChoice assigns IsisRouterCapabilityChoiceEnum provided by user to IsisRouterCapability
	setChoice(value IsisRouterCapabilityChoiceEnum) IsisRouterCapability
	// HasChoice checks if Choice has been set in IsisRouterCapability
	HasChoice() bool
	// getter for Ipv4TeRouterId to set choice.
	Ipv4TeRouterId()
	// getter for InterfaceIp to set choice.
	InterfaceIp()
	// CustomRouterCapId returns string, set in IsisRouterCapability.
	CustomRouterCapId() string
	// SetCustomRouterCapId assigns string provided by user to IsisRouterCapability
	SetCustomRouterCapId(value string) IsisRouterCapability
	// HasCustomRouterCapId checks if CustomRouterCapId has been set in IsisRouterCapability
	HasCustomRouterCapId() bool
	// SBit returns IsisRouterCapabilitySBitEnum, set in IsisRouterCapability
	SBit() IsisRouterCapabilitySBitEnum
	// SetSBit assigns IsisRouterCapabilitySBitEnum provided by user to IsisRouterCapability
	SetSBit(value IsisRouterCapabilitySBitEnum) IsisRouterCapability
	// HasSBit checks if SBit has been set in IsisRouterCapability
	HasSBit() bool
	// DBit returns IsisRouterCapabilityDBitEnum, set in IsisRouterCapability
	DBit() IsisRouterCapabilityDBitEnum
	// SetDBit assigns IsisRouterCapabilityDBitEnum provided by user to IsisRouterCapability
	SetDBit(value IsisRouterCapabilityDBitEnum) IsisRouterCapability
	// HasDBit checks if DBit has been set in IsisRouterCapability
	HasDBit() bool
	// SrCapability returns IsisSRCapability, set in IsisRouterCapability.
	// IsisSRCapability is container for the configuration of IS-IS SR-CAPABILITY TLV that Segment Routing requires
	// each router to advertise its SR data plane capability and the range of MPLS label values
	// it uses for Segment Routing in the case where global SIDs are allocated (i.e., global indexes).
	// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-capabilities-sub-tlv.
	// An implementation should set default values appropriately if any mandatory item is not configured by a user.
	SrCapability() IsisSRCapability
	// SetSrCapability assigns IsisSRCapability provided by user to IsisRouterCapability.
	// IsisSRCapability is container for the configuration of IS-IS SR-CAPABILITY TLV that Segment Routing requires
	// each router to advertise its SR data plane capability and the range of MPLS label values
	// it uses for Segment Routing in the case where global SIDs are allocated (i.e., global indexes).
	// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-capabilities-sub-tlv.
	// An implementation should set default values appropriately if any mandatory item is not configured by a user.
	SetSrCapability(value IsisSRCapability) IsisRouterCapability
	// HasSrCapability checks if SrCapability has been set in IsisRouterCapability
	HasSrCapability() bool
	// Algorithms returns []uint32, set in IsisRouterCapability.
	Algorithms() []uint32
	// SetAlgorithms assigns []uint32 provided by user to IsisRouterCapability
	SetAlgorithms(value []uint32) IsisRouterCapability
	// SrlbRanges returns IsisRouterCapabilityIsisSRSrlbIterIter, set in IsisRouterCapability
	SrlbRanges() IsisRouterCapabilityIsisSRSrlbIter
	setNil()
}

type IsisRouterCapabilityChoiceEnum string

// Enum of Choice on IsisRouterCapability
var IsisRouterCapabilityChoice = struct {
	IPV4_TE_ROUTER_ID    IsisRouterCapabilityChoiceEnum
	INTERFACE_IP         IsisRouterCapabilityChoiceEnum
	CUSTOM_ROUTER_CAP_ID IsisRouterCapabilityChoiceEnum
}{
	IPV4_TE_ROUTER_ID:    IsisRouterCapabilityChoiceEnum("ipv4_te_router_id"),
	INTERFACE_IP:         IsisRouterCapabilityChoiceEnum("interface_ip"),
	CUSTOM_ROUTER_CAP_ID: IsisRouterCapabilityChoiceEnum("custom_router_cap_id"),
}

func (obj *isisRouterCapability) Choice() IsisRouterCapabilityChoiceEnum {
	return IsisRouterCapabilityChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Ipv4TeRouterId to set choice
func (obj *isisRouterCapability) Ipv4TeRouterId() {
	obj.setChoice(IsisRouterCapabilityChoice.IPV4_TE_ROUTER_ID)
}

// getter for InterfaceIp to set choice
func (obj *isisRouterCapability) InterfaceIp() {
	obj.setChoice(IsisRouterCapabilityChoice.INTERFACE_IP)
}

// The Router Capability ID SHOULD be identical to the value advertised in the Traffic Engineering Router ID TLV [RFC5305].
// If no Traffic Engineering Router ID is assigned, the Router ID SHOULD be identical to an IP Interface Address [RFC1195]
// advertised by the originating IS.
// If the originating node does not support IPv4, then the reserved value 0.0.0.0 MUST be used in the Router ID field,
// and the IPv6 TE Router ID sub-TLV [RFC5316] MUST be present in the TLV.
// - ipv4_te_router_id: IPv4 Traffic Engineering(TE) router id (defined in isis.basic.ipv4_te_router_id) to be used as Router Capability ID.
// - interface_ip: When this is chosen the first IPv4 address of the first eth interface associated with this isis router instance should be assigned as the Router Capability ID.
// - custom_router_cap_id: This option should be chosen when Router Capability ID needs to be configured different from above two options.
// Choice returns a string
func (obj *isisRouterCapability) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisRouterCapability) setChoice(value IsisRouterCapabilityChoiceEnum) IsisRouterCapability {
	intValue, ok := otg.IsisRouterCapability_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisRouterCapabilityChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisRouterCapability_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.CustomRouterCapId = nil

	if value == IsisRouterCapabilityChoice.CUSTOM_ROUTER_CAP_ID {
		defaultValue := "0.0.0.0"
		obj.obj.CustomRouterCapId = &defaultValue
	}

	return obj
}

// Router Capability ID in IPv4 address format.
// CustomRouterCapId returns a string
func (obj *isisRouterCapability) CustomRouterCapId() string {

	if obj.obj.CustomRouterCapId == nil {
		obj.setChoice(IsisRouterCapabilityChoice.CUSTOM_ROUTER_CAP_ID)
	}

	return *obj.obj.CustomRouterCapId

}

// Router Capability ID in IPv4 address format.
// CustomRouterCapId returns a string
func (obj *isisRouterCapability) HasCustomRouterCapId() bool {
	return obj.obj.CustomRouterCapId != nil
}

// Router Capability ID in IPv4 address format.
// SetCustomRouterCapId sets the string value in the IsisRouterCapability object
func (obj *isisRouterCapability) SetCustomRouterCapId(value string) IsisRouterCapability {
	obj.setChoice(IsisRouterCapabilityChoice.CUSTOM_ROUTER_CAP_ID)
	obj.obj.CustomRouterCapId = &value
	return obj
}

type IsisRouterCapabilitySBitEnum string

// Enum of SBit on IsisRouterCapability
var IsisRouterCapabilitySBit = struct {
	FLOOD     IsisRouterCapabilitySBitEnum
	NOT_FLOOD IsisRouterCapabilitySBitEnum
}{
	FLOOD:     IsisRouterCapabilitySBitEnum("flood"),
	NOT_FLOOD: IsisRouterCapabilitySBitEnum("not_flood"),
}

func (obj *isisRouterCapability) SBit() IsisRouterCapabilitySBitEnum {
	return IsisRouterCapabilitySBitEnum(obj.obj.SBit.Enum().String())
}

// S bit (0x01): If the S bit is set(1), the IS-IS Router CAPABILITY TLV
// MUST be flooded across the entire routing domain.  If the S bit is
// not set(0), the TLV MUST NOT be leaked between levels.  This bit MUST
// NOT be altered during the TLV leaking.
// SBit returns a string
func (obj *isisRouterCapability) HasSBit() bool {
	return obj.obj.SBit != nil
}

func (obj *isisRouterCapability) SetSBit(value IsisRouterCapabilitySBitEnum) IsisRouterCapability {
	intValue, ok := otg.IsisRouterCapability_SBit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisRouterCapabilitySBitEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisRouterCapability_SBit_Enum(intValue)
	obj.obj.SBit = &enumValue

	return obj
}

type IsisRouterCapabilityDBitEnum string

// Enum of DBit on IsisRouterCapability
var IsisRouterCapabilityDBit = struct {
	DOWN     IsisRouterCapabilityDBitEnum
	NOT_DOWN IsisRouterCapabilityDBitEnum
}{
	DOWN:     IsisRouterCapabilityDBitEnum("down"),
	NOT_DOWN: IsisRouterCapabilityDBitEnum("not_down"),
}

func (obj *isisRouterCapability) DBit() IsisRouterCapabilityDBitEnum {
	return IsisRouterCapabilityDBitEnum(obj.obj.DBit.Enum().String())
}

// D bit (0x02): When the IS-IS Router CAPABILITY TLV is leaked from
// Level 2 (L2) to Level 1 (L1), the D bit MUST be set.  Otherwise, this
// bit MUST be clear.  IS-IS Router CAPABILITY TLVs with the D bit set
// MUST NOT be leaked from Level 1 to Level 2.  This is to prevent TLV looping.
// DBit returns a string
func (obj *isisRouterCapability) HasDBit() bool {
	return obj.obj.DBit != nil
}

func (obj *isisRouterCapability) SetDBit(value IsisRouterCapabilityDBitEnum) IsisRouterCapability {
	intValue, ok := otg.IsisRouterCapability_DBit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisRouterCapabilityDBitEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisRouterCapability_DBit_Enum(intValue)
	obj.obj.DBit = &enumValue

	return obj
}

// SR-Capabilities.
// SrCapability returns a IsisSRCapability
func (obj *isisRouterCapability) SrCapability() IsisSRCapability {
	if obj.obj.SrCapability == nil {
		obj.obj.SrCapability = NewIsisSRCapability().msg()
	}
	if obj.srCapabilityHolder == nil {
		obj.srCapabilityHolder = &isisSRCapability{obj: obj.obj.SrCapability}
	}
	return obj.srCapabilityHolder
}

// SR-Capabilities.
// SrCapability returns a IsisSRCapability
func (obj *isisRouterCapability) HasSrCapability() bool {
	return obj.obj.SrCapability != nil
}

// SR-Capabilities.
// SetSrCapability sets the IsisSRCapability value in the IsisRouterCapability object
func (obj *isisRouterCapability) SetSrCapability(value IsisSRCapability) IsisRouterCapability {

	obj.srCapabilityHolder = nil
	obj.obj.SrCapability = value.msg()

	return obj
}

// This contains one or more Segment Routing Algorithm that a router may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these nodes.
// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these
// nodes. Examples of these algorithms are metric-based SPF, various sorts of Constrained SPF, etc.
// - 0: SPF algorithm based on link metric.
// - 1: Strict SPF algorithm based on link metric.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-igp-algorithm-types-registr.
// When the originating router does not advertise the SR-Algorithm sub-TLV, it implies that algorithm 0 is the only algorithm supported by the routers.
// When the originating router does advertise the SR-Algorithm sub-TLV, then algorithm 0 MUST be present while non-zero algorithms MAY be present.
// Algorithms returns a []uint32
func (obj *isisRouterCapability) Algorithms() []uint32 {
	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	return obj.obj.Algorithms
}

// This contains one or more Segment Routing Algorithm that a router may use various algorithms when calculating
// reachability to other nodes or to prefixes attached to these nodes.
// The Isis may use various algorithms when calculating reachability to other nodes or to prefixes attached to these
// nodes. Examples of these algorithms are metric-based SPF, various sorts of Constrained SPF, etc.
// - 0: SPF algorithm based on link metric.
// - 1: Strict SPF algorithm based on link metric.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665#name-igp-algorithm-types-registr.
// When the originating router does not advertise the SR-Algorithm sub-TLV, it implies that algorithm 0 is the only algorithm supported by the routers.
// When the originating router does advertise the SR-Algorithm sub-TLV, then algorithm 0 MUST be present while non-zero algorithms MAY be present.
// SetAlgorithms sets the []uint32 value in the IsisRouterCapability object
func (obj *isisRouterCapability) SetAlgorithms(value []uint32) IsisRouterCapability {

	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	obj.obj.Algorithms = value

	return obj
}

// This contains the list of SR Local Block (SRLB)
// SrlbRanges returns a []IsisSRSrlb
func (obj *isisRouterCapability) SrlbRanges() IsisRouterCapabilityIsisSRSrlbIter {
	if len(obj.obj.SrlbRanges) == 0 {
		obj.obj.SrlbRanges = []*otg.IsisSRSrlb{}
	}
	if obj.srlbRangesHolder == nil {
		obj.srlbRangesHolder = newIsisRouterCapabilityIsisSRSrlbIter(&obj.obj.SrlbRanges).setMsg(obj)
	}
	return obj.srlbRangesHolder
}

type isisRouterCapabilityIsisSRSrlbIter struct {
	obj             *isisRouterCapability
	isisSRSrlbSlice []IsisSRSrlb
	fieldPtr        *[]*otg.IsisSRSrlb
}

func newIsisRouterCapabilityIsisSRSrlbIter(ptr *[]*otg.IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter {
	return &isisRouterCapabilityIsisSRSrlbIter{fieldPtr: ptr}
}

type IsisRouterCapabilityIsisSRSrlbIter interface {
	setMsg(*isisRouterCapability) IsisRouterCapabilityIsisSRSrlbIter
	Items() []IsisSRSrlb
	Add() IsisSRSrlb
	Append(items ...IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter
	Set(index int, newObj IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter
	Clear() IsisRouterCapabilityIsisSRSrlbIter
	clearHolderSlice() IsisRouterCapabilityIsisSRSrlbIter
	appendHolderSlice(item IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter
}

func (obj *isisRouterCapabilityIsisSRSrlbIter) setMsg(msg *isisRouterCapability) IsisRouterCapabilityIsisSRSrlbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRSrlb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisRouterCapabilityIsisSRSrlbIter) Items() []IsisSRSrlb {
	return obj.isisSRSrlbSlice
}

func (obj *isisRouterCapabilityIsisSRSrlbIter) Add() IsisSRSrlb {
	newObj := &otg.IsisSRSrlb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRSrlb{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRSrlbSlice = append(obj.isisSRSrlbSlice, newLibObj)
	return newLibObj
}

func (obj *isisRouterCapabilityIsisSRSrlbIter) Append(items ...IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRSrlbSlice = append(obj.isisSRSrlbSlice, item)
	}
	return obj
}

func (obj *isisRouterCapabilityIsisSRSrlbIter) Set(index int, newObj IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRSrlbSlice[index] = newObj
	return obj
}
func (obj *isisRouterCapabilityIsisSRSrlbIter) Clear() IsisRouterCapabilityIsisSRSrlbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRSrlb{}
		obj.isisSRSrlbSlice = []IsisSRSrlb{}
	}
	return obj
}
func (obj *isisRouterCapabilityIsisSRSrlbIter) clearHolderSlice() IsisRouterCapabilityIsisSRSrlbIter {
	if len(obj.isisSRSrlbSlice) > 0 {
		obj.isisSRSrlbSlice = []IsisSRSrlb{}
	}
	return obj
}
func (obj *isisRouterCapabilityIsisSRSrlbIter) appendHolderSlice(item IsisSRSrlb) IsisRouterCapabilityIsisSRSrlbIter {
	obj.isisSRSrlbSlice = append(obj.isisSRSrlbSlice, item)
	return obj
}

func (obj *isisRouterCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.CustomRouterCapId != nil {

		err := obj.validateIpv4(obj.CustomRouterCapId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisRouterCapability.CustomRouterCapId"))
		}

	}

	if obj.obj.SrCapability != nil {

		obj.SrCapability().validateObj(vObj, set_default)
	}

	if obj.obj.Algorithms != nil {

		for _, item := range obj.obj.Algorithms {
			if item > 255 {
				vObj.validationErrors = append(
					vObj.validationErrors,
					fmt.Sprintf("0 <= IsisRouterCapability.Algorithms <= 255 but Got %d", item))
			}

		}

	}

	if len(obj.obj.SrlbRanges) != 0 {

		if set_default {
			obj.SrlbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrlbRanges {
				obj.SrlbRanges().appendHolderSlice(&isisSRSrlb{obj: item})
			}
		}
		for _, item := range obj.SrlbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisRouterCapability) setDefault() {
	var choices_set int = 0
	var choice IsisRouterCapabilityChoiceEnum

	if obj.obj.CustomRouterCapId != nil {
		choices_set += 1
		choice = IsisRouterCapabilityChoice.CUSTOM_ROUTER_CAP_ID
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisRouterCapabilityChoice.IPV4_TE_ROUTER_ID)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisRouterCapability")
			}
		} else {
			intVal := otg.IsisRouterCapability_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisRouterCapability_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.SBit == nil {
		obj.SetSBit(IsisRouterCapabilitySBit.NOT_FLOOD)

	}
	if obj.obj.DBit == nil {
		obj.SetDBit(IsisRouterCapabilityDBit.NOT_DOWN)

	}

}
