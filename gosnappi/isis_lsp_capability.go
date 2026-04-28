package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspCapability *****
type isisLspCapability struct {
	validation
	obj                *otg.IsisLspCapability
	marshaller         marshalIsisLspCapability
	unMarshaller       unMarshalIsisLspCapability
	srCapabilityHolder IsisLspSRCapability
	srlbRangesHolder   IsisLspCapabilityIsisLspSrlbIter
}

func NewIsisLspCapability() IsisLspCapability {
	obj := isisLspCapability{obj: &otg.IsisLspCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspCapability) msg() *otg.IsisLspCapability {
	return obj.obj
}

func (obj *isisLspCapability) setMsg(msg *otg.IsisLspCapability) IsisLspCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspCapability struct {
	obj *isisLspCapability
}

type marshalIsisLspCapability interface {
	// ToProto marshals IsisLspCapability to protobuf object *otg.IsisLspCapability
	ToProto() (*otg.IsisLspCapability, error)
	// ToPbText marshals IsisLspCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspCapability to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspCapability struct {
	obj *isisLspCapability
}

type unMarshalIsisLspCapability interface {
	// FromProto unmarshals IsisLspCapability from protobuf object *otg.IsisLspCapability
	FromProto(msg *otg.IsisLspCapability) (IsisLspCapability, error)
	// FromPbText unmarshals IsisLspCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspCapability from JSON text
	FromJson(value string) error
}

func (obj *isisLspCapability) Marshal() marshalIsisLspCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspCapability) Unmarshal() unMarshalIsisLspCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspCapability) ToProto() (*otg.IsisLspCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspCapability) FromProto(msg *otg.IsisLspCapability) (IsisLspCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisLspCapability) FromPbText(value string) error {
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

func (m *marshalisisLspCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisLspCapability) FromYaml(value string) error {
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

func (m *marshalisisLspCapability) ToJson() (string, error) {
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

func (m *unMarshalisisLspCapability) FromJson(value string) error {
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

func (obj *isisLspCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspCapability) Clone() (IsisLspCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspCapability()
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

func (obj *isisLspCapability) setNil() {
	obj.srCapabilityHolder = nil
	obj.srlbRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspCapability is container of IS-IS Router CAPABILITY TLV.
type IsisLspCapability interface {
	Validation
	// msg marshals IsisLspCapability to protobuf object *otg.IsisLspCapability
	// and doesn't set defaults
	msg() *otg.IsisLspCapability
	// setMsg unmarshals IsisLspCapability from protobuf object *otg.IsisLspCapability
	// and doesn't set defaults
	setMsg(*otg.IsisLspCapability) IsisLspCapability
	// provides marshal interface
	Marshal() marshalIsisLspCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspCapability
	// validate validates IsisLspCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterCapId returns string, set in IsisLspCapability.
	RouterCapId() string
	// SetRouterCapId assigns string provided by user to IsisLspCapability
	SetRouterCapId(value string) IsisLspCapability
	// HasRouterCapId checks if RouterCapId has been set in IsisLspCapability
	HasRouterCapId() bool
	// SBit returns IsisLspCapabilitySBitEnum, set in IsisLspCapability
	SBit() IsisLspCapabilitySBitEnum
	// SetSBit assigns IsisLspCapabilitySBitEnum provided by user to IsisLspCapability
	SetSBit(value IsisLspCapabilitySBitEnum) IsisLspCapability
	// HasSBit checks if SBit has been set in IsisLspCapability
	HasSBit() bool
	// DBit returns IsisLspCapabilityDBitEnum, set in IsisLspCapability
	DBit() IsisLspCapabilityDBitEnum
	// SetDBit assigns IsisLspCapabilityDBitEnum provided by user to IsisLspCapability
	SetDBit(value IsisLspCapabilityDBitEnum) IsisLspCapability
	// HasDBit checks if DBit has been set in IsisLspCapability
	HasDBit() bool
	// SrCapability returns IsisLspSRCapability, set in IsisLspCapability.
	// IsisLspSRCapability is container of IS-IS SR-CAPABILITY Sub-Tlv.
	SrCapability() IsisLspSRCapability
	// SetSrCapability assigns IsisLspSRCapability provided by user to IsisLspCapability.
	// IsisLspSRCapability is container of IS-IS SR-CAPABILITY Sub-Tlv.
	SetSrCapability(value IsisLspSRCapability) IsisLspCapability
	// HasSrCapability checks if SrCapability has been set in IsisLspCapability
	HasSrCapability() bool
	// Algorithms returns []uint32, set in IsisLspCapability.
	Algorithms() []uint32
	// SetAlgorithms assigns []uint32 provided by user to IsisLspCapability
	SetAlgorithms(value []uint32) IsisLspCapability
	// SrlbRanges returns IsisLspCapabilityIsisLspSrlbIterIter, set in IsisLspCapability
	SrlbRanges() IsisLspCapabilityIsisLspSrlbIter
	setNil()
}

// Router CapabilityID in IPv4 address format.
// RouterCapId returns a string
func (obj *isisLspCapability) RouterCapId() string {

	return *obj.obj.RouterCapId

}

// Router CapabilityID in IPv4 address format.
// RouterCapId returns a string
func (obj *isisLspCapability) HasRouterCapId() bool {
	return obj.obj.RouterCapId != nil
}

// Router CapabilityID in IPv4 address format.
// SetRouterCapId sets the string value in the IsisLspCapability object
func (obj *isisLspCapability) SetRouterCapId(value string) IsisLspCapability {

	obj.obj.RouterCapId = &value
	return obj
}

type IsisLspCapabilitySBitEnum string

// Enum of SBit on IsisLspCapability
var IsisLspCapabilitySBit = struct {
	FLOOD     IsisLspCapabilitySBitEnum
	NOT_FLOOD IsisLspCapabilitySBitEnum
}{
	FLOOD:     IsisLspCapabilitySBitEnum("flood"),
	NOT_FLOOD: IsisLspCapabilitySBitEnum("not_flood"),
}

func (obj *isisLspCapability) SBit() IsisLspCapabilitySBitEnum {
	return IsisLspCapabilitySBitEnum(obj.obj.SBit.Enum().String())
}

// S bit (0x01): If the S bit is set(1), the IS-IS Router CAPABILITY TLV
// MUST be flooded across the entire routing domain.  If the S bit is
// not set(0), the TLV MUST NOT be leaked between levels.  This bit MUST
// NOT be altered during the TLV leaking.
// SBit returns a string
func (obj *isisLspCapability) HasSBit() bool {
	return obj.obj.SBit != nil
}

func (obj *isisLspCapability) SetSBit(value IsisLspCapabilitySBitEnum) IsisLspCapability {
	intValue, ok := otg.IsisLspCapability_SBit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspCapabilitySBitEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspCapability_SBit_Enum(intValue)
	obj.obj.SBit = &enumValue

	return obj
}

type IsisLspCapabilityDBitEnum string

// Enum of DBit on IsisLspCapability
var IsisLspCapabilityDBit = struct {
	DOWN     IsisLspCapabilityDBitEnum
	NOT_DOWN IsisLspCapabilityDBitEnum
}{
	DOWN:     IsisLspCapabilityDBitEnum("down"),
	NOT_DOWN: IsisLspCapabilityDBitEnum("not_down"),
}

func (obj *isisLspCapability) DBit() IsisLspCapabilityDBitEnum {
	return IsisLspCapabilityDBitEnum(obj.obj.DBit.Enum().String())
}

// D bit (0x02): When the IS-IS Router CAPABILITY TLV is leaked from
// Level 2 (L2) to Level 1 (L1), the D bit MUST be set.  Otherwise, this
// bit MUST be clear.  IS-IS Router CAPABILITY TLVs with the D bit set
// MUST NOT be leaked from Level 1 to Level 2.  This is to prevent TLV looping.
// DBit returns a string
func (obj *isisLspCapability) HasDBit() bool {
	return obj.obj.DBit != nil
}

func (obj *isisLspCapability) SetDBit(value IsisLspCapabilityDBitEnum) IsisLspCapability {
	intValue, ok := otg.IsisLspCapability_DBit_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisLspCapabilityDBitEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisLspCapability_DBit_Enum(intValue)
	obj.obj.DBit = &enumValue

	return obj
}

// SR-Capabilities.
// SrCapability returns a IsisLspSRCapability
func (obj *isisLspCapability) SrCapability() IsisLspSRCapability {
	if obj.obj.SrCapability == nil {
		obj.obj.SrCapability = NewIsisLspSRCapability().msg()
	}
	if obj.srCapabilityHolder == nil {
		obj.srCapabilityHolder = &isisLspSRCapability{obj: obj.obj.SrCapability}
	}
	return obj.srCapabilityHolder
}

// SR-Capabilities.
// SrCapability returns a IsisLspSRCapability
func (obj *isisLspCapability) HasSrCapability() bool {
	return obj.obj.SrCapability != nil
}

// SR-Capabilities.
// SetSrCapability sets the IsisLspSRCapability value in the IsisLspCapability object
func (obj *isisLspCapability) SetSrCapability(value IsisLspSRCapability) IsisLspCapability {

	obj.srCapabilityHolder = nil
	obj.obj.SrCapability = value.msg()

	return obj
}

// This contains one or more SR-Algorithm.
// Algorithms returns a []uint32
func (obj *isisLspCapability) Algorithms() []uint32 {
	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	return obj.obj.Algorithms
}

// This contains one or more SR-Algorithm.
// SetAlgorithms sets the []uint32 value in the IsisLspCapability object
func (obj *isisLspCapability) SetAlgorithms(value []uint32) IsisLspCapability {

	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	obj.obj.Algorithms = value

	return obj
}

// This contains the list of SR Local Block (SRLB)
// SrlbRanges returns a []IsisLspSrlb
func (obj *isisLspCapability) SrlbRanges() IsisLspCapabilityIsisLspSrlbIter {
	if len(obj.obj.SrlbRanges) == 0 {
		obj.obj.SrlbRanges = []*otg.IsisLspSrlb{}
	}
	if obj.srlbRangesHolder == nil {
		obj.srlbRangesHolder = newIsisLspCapabilityIsisLspSrlbIter(&obj.obj.SrlbRanges).setMsg(obj)
	}
	return obj.srlbRangesHolder
}

type isisLspCapabilityIsisLspSrlbIter struct {
	obj              *isisLspCapability
	isisLspSrlbSlice []IsisLspSrlb
	fieldPtr         *[]*otg.IsisLspSrlb
}

func newIsisLspCapabilityIsisLspSrlbIter(ptr *[]*otg.IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter {
	return &isisLspCapabilityIsisLspSrlbIter{fieldPtr: ptr}
}

type IsisLspCapabilityIsisLspSrlbIter interface {
	setMsg(*isisLspCapability) IsisLspCapabilityIsisLspSrlbIter
	Items() []IsisLspSrlb
	Add() IsisLspSrlb
	Append(items ...IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter
	Set(index int, newObj IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter
	Clear() IsisLspCapabilityIsisLspSrlbIter
	clearHolderSlice() IsisLspCapabilityIsisLspSrlbIter
	appendHolderSlice(item IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter
}

func (obj *isisLspCapabilityIsisLspSrlbIter) setMsg(msg *isisLspCapability) IsisLspCapabilityIsisLspSrlbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspSrlb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspCapabilityIsisLspSrlbIter) Items() []IsisLspSrlb {
	return obj.isisLspSrlbSlice
}

func (obj *isisLspCapabilityIsisLspSrlbIter) Add() IsisLspSrlb {
	newObj := &otg.IsisLspSrlb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspSrlb{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspCapabilityIsisLspSrlbIter) Append(items ...IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, item)
	}
	return obj
}

func (obj *isisLspCapabilityIsisLspSrlbIter) Set(index int, newObj IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspSrlbSlice[index] = newObj
	return obj
}
func (obj *isisLspCapabilityIsisLspSrlbIter) Clear() IsisLspCapabilityIsisLspSrlbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspSrlb{}
		obj.isisLspSrlbSlice = []IsisLspSrlb{}
	}
	return obj
}
func (obj *isisLspCapabilityIsisLspSrlbIter) clearHolderSlice() IsisLspCapabilityIsisLspSrlbIter {
	if len(obj.isisLspSrlbSlice) > 0 {
		obj.isisLspSrlbSlice = []IsisLspSrlb{}
	}
	return obj
}
func (obj *isisLspCapabilityIsisLspSrlbIter) appendHolderSlice(item IsisLspSrlb) IsisLspCapabilityIsisLspSrlbIter {
	obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, item)
	return obj
}

func (obj *isisLspCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterCapId != nil {

		err := obj.validateIpv4(obj.RouterCapId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspCapability.RouterCapId"))
		}

	}

	if obj.obj.SrCapability != nil {

		obj.SrCapability().validateObj(vObj, set_default)
	}

	if len(obj.obj.SrlbRanges) != 0 {

		if set_default {
			obj.SrlbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrlbRanges {
				obj.SrlbRanges().appendHolderSlice(&isisLspSrlb{obj: item})
			}
		}
		for _, item := range obj.SrlbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspCapability) setDefault() {

}
