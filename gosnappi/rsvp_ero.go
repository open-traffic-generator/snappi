package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpEro *****
type rsvpEro struct {
	validation
	obj              *otg.RsvpEro
	marshaller       marshalRsvpEro
	unMarshaller     unMarshalRsvpEro
	subobjectsHolder RsvpEroRsvpEroSubobjectIter
}

func NewRsvpEro() RsvpEro {
	obj := rsvpEro{obj: &otg.RsvpEro{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpEro) msg() *otg.RsvpEro {
	return obj.obj
}

func (obj *rsvpEro) setMsg(msg *otg.RsvpEro) RsvpEro {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpEro struct {
	obj *rsvpEro
}

type marshalRsvpEro interface {
	// ToProto marshals RsvpEro to protobuf object *otg.RsvpEro
	ToProto() (*otg.RsvpEro, error)
	// ToPbText marshals RsvpEro to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpEro to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpEro to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals RsvpEro to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrsvpEro struct {
	obj *rsvpEro
}

type unMarshalRsvpEro interface {
	// FromProto unmarshals RsvpEro from protobuf object *otg.RsvpEro
	FromProto(msg *otg.RsvpEro) (RsvpEro, error)
	// FromPbText unmarshals RsvpEro from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpEro from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpEro from JSON text
	FromJson(value string) error
}

func (obj *rsvpEro) Marshal() marshalRsvpEro {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpEro{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpEro) Unmarshal() unMarshalRsvpEro {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpEro{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpEro) ToProto() (*otg.RsvpEro, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpEro) FromProto(msg *otg.RsvpEro) (RsvpEro, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpEro) ToPbText() (string, error) {
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

func (m *unMarshalrsvpEro) FromPbText(value string) error {
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

func (m *marshalrsvpEro) ToYaml() (string, error) {
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

func (m *unMarshalrsvpEro) FromYaml(value string) error {
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

func (m *marshalrsvpEro) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalrsvpEro) ToJson() (string, error) {
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

func (m *unMarshalrsvpEro) FromJson(value string) error {
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

func (obj *rsvpEro) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpEro) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpEro) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpEro) Clone() (RsvpEro, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpEro()
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

func (obj *rsvpEro) setNil() {
	obj.subobjectsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// RsvpEro is configuration for the optional RSVP-TE explicit route object(ERO) object included in Path Messages.
type RsvpEro interface {
	Validation
	// msg marshals RsvpEro to protobuf object *otg.RsvpEro
	// and doesn't set defaults
	msg() *otg.RsvpEro
	// setMsg unmarshals RsvpEro from protobuf object *otg.RsvpEro
	// and doesn't set defaults
	setMsg(*otg.RsvpEro) RsvpEro
	// provides marshal interface
	Marshal() marshalRsvpEro
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpEro
	// validate validates RsvpEro
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpEro, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PrependNeighborIp returns RsvpEroPrependNeighborIpEnum, set in RsvpEro
	PrependNeighborIp() RsvpEroPrependNeighborIpEnum
	// SetPrependNeighborIp assigns RsvpEroPrependNeighborIpEnum provided by user to RsvpEro
	SetPrependNeighborIp(value RsvpEroPrependNeighborIpEnum) RsvpEro
	// HasPrependNeighborIp checks if PrependNeighborIp has been set in RsvpEro
	HasPrependNeighborIp() bool
	// PrefixLength returns uint32, set in RsvpEro.
	PrefixLength() uint32
	// SetPrefixLength assigns uint32 provided by user to RsvpEro
	SetPrefixLength(value uint32) RsvpEro
	// HasPrefixLength checks if PrefixLength has been set in RsvpEro
	HasPrefixLength() bool
	// Subobjects returns RsvpEroRsvpEroSubobjectIterIter, set in RsvpEro
	Subobjects() RsvpEroRsvpEroSubobjectIter
	setNil()
}

type RsvpEroPrependNeighborIpEnum string

// Enum of PrependNeighborIp on RsvpEro
var RsvpEroPrependNeighborIp = struct {
	DONT_PREPEND   RsvpEroPrependNeighborIpEnum
	PREPEND_LOOSE  RsvpEroPrependNeighborIpEnum
	PREPEND_STRICT RsvpEroPrependNeighborIpEnum
}{
	DONT_PREPEND:   RsvpEroPrependNeighborIpEnum("dont_prepend"),
	PREPEND_LOOSE:  RsvpEroPrependNeighborIpEnum("prepend_loose"),
	PREPEND_STRICT: RsvpEroPrependNeighborIpEnum("prepend_strict"),
}

func (obj *rsvpEro) PrependNeighborIp() RsvpEroPrependNeighborIpEnum {
	return RsvpEroPrependNeighborIpEnum(obj.obj.PrependNeighborIp.Enum().String())
}

// Determines whether the IP address of the RSVP neighbor should be added as an ERO sub-object. If it is to be included, it can be included as a Loose hop or as a Strict hop.
// PrependNeighborIp returns a string
func (obj *rsvpEro) HasPrependNeighborIp() bool {
	return obj.obj.PrependNeighborIp != nil
}

func (obj *rsvpEro) SetPrependNeighborIp(value RsvpEroPrependNeighborIpEnum) RsvpEro {
	intValue, ok := otg.RsvpEro_PrependNeighborIp_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpEroPrependNeighborIpEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpEro_PrependNeighborIp_Enum(intValue)
	obj.obj.PrependNeighborIp = &enumValue

	return obj
}

// If prepend_egress_ip is set to one of 'prepend_loose' or 'prepend_strict', then set this value as the prefix length of the ERO sub-object containing egress IP address.
// PrefixLength returns a uint32
func (obj *rsvpEro) PrefixLength() uint32 {

	return *obj.obj.PrefixLength

}

// If prepend_egress_ip is set to one of 'prepend_loose' or 'prepend_strict', then set this value as the prefix length of the ERO sub-object containing egress IP address.
// PrefixLength returns a uint32
func (obj *rsvpEro) HasPrefixLength() bool {
	return obj.obj.PrefixLength != nil
}

// If prepend_egress_ip is set to one of 'prepend_loose' or 'prepend_strict', then set this value as the prefix length of the ERO sub-object containing egress IP address.
// SetPrefixLength sets the uint32 value in the RsvpEro object
func (obj *rsvpEro) SetPrefixLength(value uint32) RsvpEro {

	obj.obj.PrefixLength = &value
	return obj
}

// Array of sub-objects to be included in the ERO. These sub-objects contain the intermediate hops to be traversed by the LSP while being forwarded towards the egress endpoint. These sub-objects are included after the optional sub-object containing IP address of egress endpoint of the LSP (when present).
// Subobjects returns a []RsvpEroSubobject
func (obj *rsvpEro) Subobjects() RsvpEroRsvpEroSubobjectIter {
	if len(obj.obj.Subobjects) == 0 {
		obj.obj.Subobjects = []*otg.RsvpEroSubobject{}
	}
	if obj.subobjectsHolder == nil {
		obj.subobjectsHolder = newRsvpEroRsvpEroSubobjectIter(&obj.obj.Subobjects).setMsg(obj)
	}
	return obj.subobjectsHolder
}

type rsvpEroRsvpEroSubobjectIter struct {
	obj                   *rsvpEro
	rsvpEroSubobjectSlice []RsvpEroSubobject
	fieldPtr              *[]*otg.RsvpEroSubobject
}

func newRsvpEroRsvpEroSubobjectIter(ptr *[]*otg.RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter {
	return &rsvpEroRsvpEroSubobjectIter{fieldPtr: ptr}
}

type RsvpEroRsvpEroSubobjectIter interface {
	setMsg(*rsvpEro) RsvpEroRsvpEroSubobjectIter
	Items() []RsvpEroSubobject
	Add() RsvpEroSubobject
	Append(items ...RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter
	Set(index int, newObj RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter
	Clear() RsvpEroRsvpEroSubobjectIter
	clearHolderSlice() RsvpEroRsvpEroSubobjectIter
	appendHolderSlice(item RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter
}

func (obj *rsvpEroRsvpEroSubobjectIter) setMsg(msg *rsvpEro) RsvpEroRsvpEroSubobjectIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&rsvpEroSubobject{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *rsvpEroRsvpEroSubobjectIter) Items() []RsvpEroSubobject {
	return obj.rsvpEroSubobjectSlice
}

func (obj *rsvpEroRsvpEroSubobjectIter) Add() RsvpEroSubobject {
	newObj := &otg.RsvpEroSubobject{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &rsvpEroSubobject{obj: newObj}
	newLibObj.setDefault()
	obj.rsvpEroSubobjectSlice = append(obj.rsvpEroSubobjectSlice, newLibObj)
	return newLibObj
}

func (obj *rsvpEroRsvpEroSubobjectIter) Append(items ...RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.rsvpEroSubobjectSlice = append(obj.rsvpEroSubobjectSlice, item)
	}
	return obj
}

func (obj *rsvpEroRsvpEroSubobjectIter) Set(index int, newObj RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.rsvpEroSubobjectSlice[index] = newObj
	return obj
}
func (obj *rsvpEroRsvpEroSubobjectIter) Clear() RsvpEroRsvpEroSubobjectIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.RsvpEroSubobject{}
		obj.rsvpEroSubobjectSlice = []RsvpEroSubobject{}
	}
	return obj
}
func (obj *rsvpEroRsvpEroSubobjectIter) clearHolderSlice() RsvpEroRsvpEroSubobjectIter {
	if len(obj.rsvpEroSubobjectSlice) > 0 {
		obj.rsvpEroSubobjectSlice = []RsvpEroSubobject{}
	}
	return obj
}
func (obj *rsvpEroRsvpEroSubobjectIter) appendHolderSlice(item RsvpEroSubobject) RsvpEroRsvpEroSubobjectIter {
	obj.rsvpEroSubobjectSlice = append(obj.rsvpEroSubobjectSlice, item)
	return obj
}

func (obj *rsvpEro) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PrefixLength != nil {

		if *obj.obj.PrefixLength > 32 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= RsvpEro.PrefixLength <= 32 but Got %d", *obj.obj.PrefixLength))
		}

	}

	if len(obj.obj.Subobjects) != 0 {

		if set_default {
			obj.Subobjects().clearHolderSlice()
			for _, item := range obj.obj.Subobjects {
				obj.Subobjects().appendHolderSlice(&rsvpEroSubobject{obj: item})
			}
		}
		for _, item := range obj.Subobjects().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *rsvpEro) setDefault() {
	if obj.obj.PrependNeighborIp == nil {
		obj.SetPrependNeighborIp(RsvpEroPrependNeighborIp.PREPEND_LOOSE)

	}
	if obj.obj.PrefixLength == nil {
		obj.SetPrefixLength(32)
	}

}
