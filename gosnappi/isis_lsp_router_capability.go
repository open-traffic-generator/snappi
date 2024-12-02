package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspRouterCapability *****
type isisLspRouterCapability struct {
	validation
	obj                *otg.IsisLspRouterCapability
	marshaller         marshalIsisLspRouterCapability
	unMarshaller       unMarshalIsisLspRouterCapability
	srCapabilityHolder IsisLspSRCapability
	algorithmsHolder   IsisLspRouterCapabilityIsisLspAlgorithmIter
	srlbsHolder        IsisLspRouterCapabilityIsisLspSrlbIter
}

func NewIsisLspRouterCapability() IsisLspRouterCapability {
	obj := isisLspRouterCapability{obj: &otg.IsisLspRouterCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspRouterCapability) msg() *otg.IsisLspRouterCapability {
	return obj.obj
}

func (obj *isisLspRouterCapability) setMsg(msg *otg.IsisLspRouterCapability) IsisLspRouterCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspRouterCapability struct {
	obj *isisLspRouterCapability
}

type marshalIsisLspRouterCapability interface {
	// ToProto marshals IsisLspRouterCapability to protobuf object *otg.IsisLspRouterCapability
	ToProto() (*otg.IsisLspRouterCapability, error)
	// ToPbText marshals IsisLspRouterCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspRouterCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspRouterCapability to JSON text
	ToJson() (string, error)
}

type unMarshalisisLspRouterCapability struct {
	obj *isisLspRouterCapability
}

type unMarshalIsisLspRouterCapability interface {
	// FromProto unmarshals IsisLspRouterCapability from protobuf object *otg.IsisLspRouterCapability
	FromProto(msg *otg.IsisLspRouterCapability) (IsisLspRouterCapability, error)
	// FromPbText unmarshals IsisLspRouterCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspRouterCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspRouterCapability from JSON text
	FromJson(value string) error
}

func (obj *isisLspRouterCapability) Marshal() marshalIsisLspRouterCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspRouterCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspRouterCapability) Unmarshal() unMarshalIsisLspRouterCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspRouterCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspRouterCapability) ToProto() (*otg.IsisLspRouterCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspRouterCapability) FromProto(msg *otg.IsisLspRouterCapability) (IsisLspRouterCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspRouterCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisLspRouterCapability) FromPbText(value string) error {
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

func (m *marshalisisLspRouterCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisLspRouterCapability) FromYaml(value string) error {
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

func (m *marshalisisLspRouterCapability) ToJson() (string, error) {
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

func (m *unMarshalisisLspRouterCapability) FromJson(value string) error {
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

func (obj *isisLspRouterCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspRouterCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspRouterCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspRouterCapability) Clone() (IsisLspRouterCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspRouterCapability()
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

func (obj *isisLspRouterCapability) setNil() {
	obj.srCapabilityHolder = nil
	obj.algorithmsHolder = nil
	obj.srlbsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspRouterCapability is container of IS-IS Router CAPABILITY TLV.
type IsisLspRouterCapability interface {
	Validation
	// msg marshals IsisLspRouterCapability to protobuf object *otg.IsisLspRouterCapability
	// and doesn't set defaults
	msg() *otg.IsisLspRouterCapability
	// setMsg unmarshals IsisLspRouterCapability from protobuf object *otg.IsisLspRouterCapability
	// and doesn't set defaults
	setMsg(*otg.IsisLspRouterCapability) IsisLspRouterCapability
	// provides marshal interface
	Marshal() marshalIsisLspRouterCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspRouterCapability
	// validate validates IsisLspRouterCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspRouterCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterCapId returns string, set in IsisLspRouterCapability.
	RouterCapId() string
	// SetRouterCapId assigns string provided by user to IsisLspRouterCapability
	SetRouterCapId(value string) IsisLspRouterCapability
	// HasRouterCapId checks if RouterCapId has been set in IsisLspRouterCapability
	HasRouterCapId() bool
	// SBit returns bool, set in IsisLspRouterCapability.
	SBit() bool
	// SetSBit assigns bool provided by user to IsisLspRouterCapability
	SetSBit(value bool) IsisLspRouterCapability
	// HasSBit checks if SBit has been set in IsisLspRouterCapability
	HasSBit() bool
	// DBit returns bool, set in IsisLspRouterCapability.
	DBit() bool
	// SetDBit assigns bool provided by user to IsisLspRouterCapability
	SetDBit(value bool) IsisLspRouterCapability
	// HasDBit checks if DBit has been set in IsisLspRouterCapability
	HasDBit() bool
	// SrCapability returns IsisLspSRCapability, set in IsisLspRouterCapability.
	// IsisLspSRCapability is container of IS-IS SR-CAPABILITY Sub-Tlv.
	SrCapability() IsisLspSRCapability
	// SetSrCapability assigns IsisLspSRCapability provided by user to IsisLspRouterCapability.
	// IsisLspSRCapability is container of IS-IS SR-CAPABILITY Sub-Tlv.
	SetSrCapability(value IsisLspSRCapability) IsisLspRouterCapability
	// HasSrCapability checks if SrCapability has been set in IsisLspRouterCapability
	HasSrCapability() bool
	// Algorithms returns IsisLspRouterCapabilityIsisLspAlgorithmIterIter, set in IsisLspRouterCapability
	Algorithms() IsisLspRouterCapabilityIsisLspAlgorithmIter
	// Srlbs returns IsisLspRouterCapabilityIsisLspSrlbIterIter, set in IsisLspRouterCapability
	Srlbs() IsisLspRouterCapabilityIsisLspSrlbIter
	setNil()
}

// Router CapabilityID in IPv4 address format.
// RouterCapId returns a string
func (obj *isisLspRouterCapability) RouterCapId() string {

	return *obj.obj.RouterCapId

}

// Router CapabilityID in IPv4 address format.
// RouterCapId returns a string
func (obj *isisLspRouterCapability) HasRouterCapId() bool {
	return obj.obj.RouterCapId != nil
}

// Router CapabilityID in IPv4 address format.
// SetRouterCapId sets the string value in the IsisLspRouterCapability object
func (obj *isisLspRouterCapability) SetRouterCapId(value string) IsisLspRouterCapability {

	obj.obj.RouterCapId = &value
	return obj
}

// S bit (0x01): If the S bit is set(1), the IS-IS Router CAPABILITY TLV
// MUST be flooded across the entire routing domain.  If the S bit is
// not set(0), the TLV MUST NOT be leaked between levels.  This bit MUST
// NOT be altered during the TLV leaking.
// SBit returns a bool
func (obj *isisLspRouterCapability) SBit() bool {

	return *obj.obj.SBit

}

// S bit (0x01): If the S bit is set(1), the IS-IS Router CAPABILITY TLV
// MUST be flooded across the entire routing domain.  If the S bit is
// not set(0), the TLV MUST NOT be leaked between levels.  This bit MUST
// NOT be altered during the TLV leaking.
// SBit returns a bool
func (obj *isisLspRouterCapability) HasSBit() bool {
	return obj.obj.SBit != nil
}

// S bit (0x01): If the S bit is set(1), the IS-IS Router CAPABILITY TLV
// MUST be flooded across the entire routing domain.  If the S bit is
// not set(0), the TLV MUST NOT be leaked between levels.  This bit MUST
// NOT be altered during the TLV leaking.
// SetSBit sets the bool value in the IsisLspRouterCapability object
func (obj *isisLspRouterCapability) SetSBit(value bool) IsisLspRouterCapability {

	obj.obj.SBit = &value
	return obj
}

// D bit (0x02): When the IS-IS Router CAPABILITY TLV is leaked from
// Level 2 (L2) to Level 1 (L1), the D bit MUST be set.  Otherwise, this
// bit MUST be clear.  IS-IS Router CAPABILITY TLVs with the D bit set
// MUST NOT be leaked from Level 1 to Level 2.  This is to prevent TLV looping.
// DBit returns a bool
func (obj *isisLspRouterCapability) DBit() bool {

	return *obj.obj.DBit

}

// D bit (0x02): When the IS-IS Router CAPABILITY TLV is leaked from
// Level 2 (L2) to Level 1 (L1), the D bit MUST be set.  Otherwise, this
// bit MUST be clear.  IS-IS Router CAPABILITY TLVs with the D bit set
// MUST NOT be leaked from Level 1 to Level 2.  This is to prevent TLV looping.
// DBit returns a bool
func (obj *isisLspRouterCapability) HasDBit() bool {
	return obj.obj.DBit != nil
}

// D bit (0x02): When the IS-IS Router CAPABILITY TLV is leaked from
// Level 2 (L2) to Level 1 (L1), the D bit MUST be set.  Otherwise, this
// bit MUST be clear.  IS-IS Router CAPABILITY TLVs with the D bit set
// MUST NOT be leaked from Level 1 to Level 2.  This is to prevent TLV looping.
// SetDBit sets the bool value in the IsisLspRouterCapability object
func (obj *isisLspRouterCapability) SetDBit(value bool) IsisLspRouterCapability {

	obj.obj.DBit = &value
	return obj
}

// SR-Capabilities.
// SrCapability returns a IsisLspSRCapability
func (obj *isisLspRouterCapability) SrCapability() IsisLspSRCapability {
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
func (obj *isisLspRouterCapability) HasSrCapability() bool {
	return obj.obj.SrCapability != nil
}

// SR-Capabilities.
// SetSrCapability sets the IsisLspSRCapability value in the IsisLspRouterCapability object
func (obj *isisLspRouterCapability) SetSrCapability(value IsisLspSRCapability) IsisLspRouterCapability {

	obj.srCapabilityHolder = nil
	obj.obj.SrCapability = value.msg()

	return obj
}

// This contains one or more SR-Algorithm.
// Algorithms returns a []IsisLspAlgorithm
func (obj *isisLspRouterCapability) Algorithms() IsisLspRouterCapabilityIsisLspAlgorithmIter {
	if len(obj.obj.Algorithms) == 0 {
		obj.obj.Algorithms = []*otg.IsisLspAlgorithm{}
	}
	if obj.algorithmsHolder == nil {
		obj.algorithmsHolder = newIsisLspRouterCapabilityIsisLspAlgorithmIter(&obj.obj.Algorithms).setMsg(obj)
	}
	return obj.algorithmsHolder
}

type isisLspRouterCapabilityIsisLspAlgorithmIter struct {
	obj                   *isisLspRouterCapability
	isisLspAlgorithmSlice []IsisLspAlgorithm
	fieldPtr              *[]*otg.IsisLspAlgorithm
}

func newIsisLspRouterCapabilityIsisLspAlgorithmIter(ptr *[]*otg.IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter {
	return &isisLspRouterCapabilityIsisLspAlgorithmIter{fieldPtr: ptr}
}

type IsisLspRouterCapabilityIsisLspAlgorithmIter interface {
	setMsg(*isisLspRouterCapability) IsisLspRouterCapabilityIsisLspAlgorithmIter
	Items() []IsisLspAlgorithm
	Add() IsisLspAlgorithm
	Append(items ...IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter
	Set(index int, newObj IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter
	Clear() IsisLspRouterCapabilityIsisLspAlgorithmIter
	clearHolderSlice() IsisLspRouterCapabilityIsisLspAlgorithmIter
	appendHolderSlice(item IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter
}

func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) setMsg(msg *isisLspRouterCapability) IsisLspRouterCapabilityIsisLspAlgorithmIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspAlgorithm{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) Items() []IsisLspAlgorithm {
	return obj.isisLspAlgorithmSlice
}

func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) Add() IsisLspAlgorithm {
	newObj := &otg.IsisLspAlgorithm{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspAlgorithm{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspAlgorithmSlice = append(obj.isisLspAlgorithmSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) Append(items ...IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspAlgorithmSlice = append(obj.isisLspAlgorithmSlice, item)
	}
	return obj
}

func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) Set(index int, newObj IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspAlgorithmSlice[index] = newObj
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) Clear() IsisLspRouterCapabilityIsisLspAlgorithmIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspAlgorithm{}
		obj.isisLspAlgorithmSlice = []IsisLspAlgorithm{}
	}
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) clearHolderSlice() IsisLspRouterCapabilityIsisLspAlgorithmIter {
	if len(obj.isisLspAlgorithmSlice) > 0 {
		obj.isisLspAlgorithmSlice = []IsisLspAlgorithm{}
	}
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspAlgorithmIter) appendHolderSlice(item IsisLspAlgorithm) IsisLspRouterCapabilityIsisLspAlgorithmIter {
	obj.isisLspAlgorithmSlice = append(obj.isisLspAlgorithmSlice, item)
	return obj
}

// This contains the list of SR Local Block (SRLB)
// Srlbs returns a []IsisLspSrlb
func (obj *isisLspRouterCapability) Srlbs() IsisLspRouterCapabilityIsisLspSrlbIter {
	if len(obj.obj.Srlbs) == 0 {
		obj.obj.Srlbs = []*otg.IsisLspSrlb{}
	}
	if obj.srlbsHolder == nil {
		obj.srlbsHolder = newIsisLspRouterCapabilityIsisLspSrlbIter(&obj.obj.Srlbs).setMsg(obj)
	}
	return obj.srlbsHolder
}

type isisLspRouterCapabilityIsisLspSrlbIter struct {
	obj              *isisLspRouterCapability
	isisLspSrlbSlice []IsisLspSrlb
	fieldPtr         *[]*otg.IsisLspSrlb
}

func newIsisLspRouterCapabilityIsisLspSrlbIter(ptr *[]*otg.IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter {
	return &isisLspRouterCapabilityIsisLspSrlbIter{fieldPtr: ptr}
}

type IsisLspRouterCapabilityIsisLspSrlbIter interface {
	setMsg(*isisLspRouterCapability) IsisLspRouterCapabilityIsisLspSrlbIter
	Items() []IsisLspSrlb
	Add() IsisLspSrlb
	Append(items ...IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter
	Set(index int, newObj IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter
	Clear() IsisLspRouterCapabilityIsisLspSrlbIter
	clearHolderSlice() IsisLspRouterCapabilityIsisLspSrlbIter
	appendHolderSlice(item IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter
}

func (obj *isisLspRouterCapabilityIsisLspSrlbIter) setMsg(msg *isisLspRouterCapability) IsisLspRouterCapabilityIsisLspSrlbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspSrlb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspRouterCapabilityIsisLspSrlbIter) Items() []IsisLspSrlb {
	return obj.isisLspSrlbSlice
}

func (obj *isisLspRouterCapabilityIsisLspSrlbIter) Add() IsisLspSrlb {
	newObj := &otg.IsisLspSrlb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspSrlb{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspRouterCapabilityIsisLspSrlbIter) Append(items ...IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, item)
	}
	return obj
}

func (obj *isisLspRouterCapabilityIsisLspSrlbIter) Set(index int, newObj IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspSrlbSlice[index] = newObj
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspSrlbIter) Clear() IsisLspRouterCapabilityIsisLspSrlbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspSrlb{}
		obj.isisLspSrlbSlice = []IsisLspSrlb{}
	}
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspSrlbIter) clearHolderSlice() IsisLspRouterCapabilityIsisLspSrlbIter {
	if len(obj.isisLspSrlbSlice) > 0 {
		obj.isisLspSrlbSlice = []IsisLspSrlb{}
	}
	return obj
}
func (obj *isisLspRouterCapabilityIsisLspSrlbIter) appendHolderSlice(item IsisLspSrlb) IsisLspRouterCapabilityIsisLspSrlbIter {
	obj.isisLspSrlbSlice = append(obj.isisLspSrlbSlice, item)
	return obj
}

func (obj *isisLspRouterCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RouterCapId != nil {

		err := obj.validateIpv4(obj.RouterCapId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on IsisLspRouterCapability.RouterCapId"))
		}

	}

	if obj.obj.SrCapability != nil {

		obj.SrCapability().validateObj(vObj, set_default)
	}

	if len(obj.obj.Algorithms) != 0 {

		if set_default {
			obj.Algorithms().clearHolderSlice()
			for _, item := range obj.obj.Algorithms {
				obj.Algorithms().appendHolderSlice(&isisLspAlgorithm{obj: item})
			}
		}
		for _, item := range obj.Algorithms().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.Srlbs) != 0 {

		if set_default {
			obj.Srlbs().clearHolderSlice()
			for _, item := range obj.obj.Srlbs {
				obj.Srlbs().appendHolderSlice(&isisLspSrlb{obj: item})
			}
		}
		for _, item := range obj.Srlbs().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspRouterCapability) setDefault() {

}
