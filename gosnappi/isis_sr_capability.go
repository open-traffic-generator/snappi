package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisSRCapability *****
type isisSRCapability struct {
	validation
	obj              *otg.IsisSRCapability
	marshaller       marshalIsisSRCapability
	unMarshaller     unMarshalIsisSRCapability
	srgbRangesHolder IsisSRCapabilityIsisSRSrgbIter
}

func NewIsisSRCapability() IsisSRCapability {
	obj := isisSRCapability{obj: &otg.IsisSRCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisSRCapability) msg() *otg.IsisSRCapability {
	return obj.obj
}

func (obj *isisSRCapability) setMsg(msg *otg.IsisSRCapability) IsisSRCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisSRCapability struct {
	obj *isisSRCapability
}

type marshalIsisSRCapability interface {
	// ToProto marshals IsisSRCapability to protobuf object *otg.IsisSRCapability
	ToProto() (*otg.IsisSRCapability, error)
	// ToPbText marshals IsisSRCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisSRCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisSRCapability to JSON text
	ToJson() (string, error)
}

type unMarshalisisSRCapability struct {
	obj *isisSRCapability
}

type unMarshalIsisSRCapability interface {
	// FromProto unmarshals IsisSRCapability from protobuf object *otg.IsisSRCapability
	FromProto(msg *otg.IsisSRCapability) (IsisSRCapability, error)
	// FromPbText unmarshals IsisSRCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisSRCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisSRCapability from JSON text
	FromJson(value string) error
}

func (obj *isisSRCapability) Marshal() marshalIsisSRCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisSRCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisSRCapability) Unmarshal() unMarshalIsisSRCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisSRCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisSRCapability) ToProto() (*otg.IsisSRCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisSRCapability) FromProto(msg *otg.IsisSRCapability) (IsisSRCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisSRCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisSRCapability) FromPbText(value string) error {
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

func (m *marshalisisSRCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisSRCapability) FromYaml(value string) error {
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

func (m *marshalisisSRCapability) ToJson() (string, error) {
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

func (m *unMarshalisisSRCapability) FromJson(value string) error {
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

func (obj *isisSRCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisSRCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisSRCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisSRCapability) Clone() (IsisSRCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisSRCapability()
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

func (obj *isisSRCapability) setNil() {
	obj.srgbRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisSRCapability is container for the configuration of IS-IS SR-CAPABILITY TLV that Segment Routing requires
// each router to advertise its SR data plane capability and the range of MPLS label values
// it uses for Segment Routing in the case where global SIDs are allocated (i.e., global indexes).
//
// Reference: https://datatracker.ietf.org/doc/html/rfc8667#name-sr-capabilities-sub-tlv.
// The Router Capability TLV specifies flags that control its advertisement.
// The SR-Capabilities sub-TLV MUST be propagated throughout the level and MUST NOT be advertised across level boundaries.
// Therefore, Router Capability TLV distribution flags are set accordingly, i.e.,
// the S-Flag in the Router Capability TLV [RFC7981] MUST be unset.
type IsisSRCapability interface {
	Validation
	// msg marshals IsisSRCapability to protobuf object *otg.IsisSRCapability
	// and doesn't set defaults
	msg() *otg.IsisSRCapability
	// setMsg unmarshals IsisSRCapability from protobuf object *otg.IsisSRCapability
	// and doesn't set defaults
	setMsg(*otg.IsisSRCapability) IsisSRCapability
	// provides marshal interface
	Marshal() marshalIsisSRCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisSRCapability
	// validate validates IsisSRCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisSRCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// IFlag returns bool, set in IsisSRCapability.
	IFlag() bool
	// SetIFlag assigns bool provided by user to IsisSRCapability
	SetIFlag(value bool) IsisSRCapability
	// HasIFlag checks if IFlag has been set in IsisSRCapability
	HasIFlag() bool
	// VFlag returns bool, set in IsisSRCapability.
	VFlag() bool
	// SetVFlag assigns bool provided by user to IsisSRCapability
	SetVFlag(value bool) IsisSRCapability
	// HasVFlag checks if VFlag has been set in IsisSRCapability
	HasVFlag() bool
	// SrgbRanges returns IsisSRCapabilityIsisSRSrgbIterIter, set in IsisSRCapability
	SrgbRanges() IsisSRCapabilityIsisSRSrgbIter
	setNil()
}

// MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// IFlag returns a bool
func (obj *isisSRCapability) IFlag() bool {

	return *obj.obj.IFlag

}

// MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// IFlag returns a bool
func (obj *isisSRCapability) HasIFlag() bool {
	return obj.obj.IFlag != nil
}

// MPLS IPv4 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv4 packets on all interfaces.
// SetIFlag sets the bool value in the IsisSRCapability object
func (obj *isisSRCapability) SetIFlag(value bool) IsisSRCapability {

	obj.obj.IFlag = &value
	return obj
}

// MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// VFlag returns a bool
func (obj *isisSRCapability) VFlag() bool {

	return *obj.obj.VFlag

}

// MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// VFlag returns a bool
func (obj *isisSRCapability) HasVFlag() bool {
	return obj.obj.VFlag != nil
}

// MPLS IPv6 Flag. If set, then the router is capable of processing SR-MPLS-encapsulated IPv6 packets on all interfaces.
// SetVFlag sets the bool value in the IsisSRCapability object
func (obj *isisSRCapability) SetVFlag(value bool) IsisSRCapability {

	obj.obj.VFlag = &value
	return obj
}

// This contains the list of SRGB.
// SrgbRanges returns a []IsisSRSrgb
func (obj *isisSRCapability) SrgbRanges() IsisSRCapabilityIsisSRSrgbIter {
	if len(obj.obj.SrgbRanges) == 0 {
		obj.obj.SrgbRanges = []*otg.IsisSRSrgb{}
	}
	if obj.srgbRangesHolder == nil {
		obj.srgbRangesHolder = newIsisSRCapabilityIsisSRSrgbIter(&obj.obj.SrgbRanges).setMsg(obj)
	}
	return obj.srgbRangesHolder
}

type isisSRCapabilityIsisSRSrgbIter struct {
	obj             *isisSRCapability
	isisSRSrgbSlice []IsisSRSrgb
	fieldPtr        *[]*otg.IsisSRSrgb
}

func newIsisSRCapabilityIsisSRSrgbIter(ptr *[]*otg.IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter {
	return &isisSRCapabilityIsisSRSrgbIter{fieldPtr: ptr}
}

type IsisSRCapabilityIsisSRSrgbIter interface {
	setMsg(*isisSRCapability) IsisSRCapabilityIsisSRSrgbIter
	Items() []IsisSRSrgb
	Add() IsisSRSrgb
	Append(items ...IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter
	Set(index int, newObj IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter
	Clear() IsisSRCapabilityIsisSRSrgbIter
	clearHolderSlice() IsisSRCapabilityIsisSRSrgbIter
	appendHolderSlice(item IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter
}

func (obj *isisSRCapabilityIsisSRSrgbIter) setMsg(msg *isisSRCapability) IsisSRCapabilityIsisSRSrgbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisSRSrgb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisSRCapabilityIsisSRSrgbIter) Items() []IsisSRSrgb {
	return obj.isisSRSrgbSlice
}

func (obj *isisSRCapabilityIsisSRSrgbIter) Add() IsisSRSrgb {
	newObj := &otg.IsisSRSrgb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisSRSrgb{obj: newObj}
	newLibObj.setDefault()
	obj.isisSRSrgbSlice = append(obj.isisSRSrgbSlice, newLibObj)
	return newLibObj
}

func (obj *isisSRCapabilityIsisSRSrgbIter) Append(items ...IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisSRSrgbSlice = append(obj.isisSRSrgbSlice, item)
	}
	return obj
}

func (obj *isisSRCapabilityIsisSRSrgbIter) Set(index int, newObj IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisSRSrgbSlice[index] = newObj
	return obj
}
func (obj *isisSRCapabilityIsisSRSrgbIter) Clear() IsisSRCapabilityIsisSRSrgbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisSRSrgb{}
		obj.isisSRSrgbSlice = []IsisSRSrgb{}
	}
	return obj
}
func (obj *isisSRCapabilityIsisSRSrgbIter) clearHolderSlice() IsisSRCapabilityIsisSRSrgbIter {
	if len(obj.isisSRSrgbSlice) > 0 {
		obj.isisSRSrgbSlice = []IsisSRSrgb{}
	}
	return obj
}
func (obj *isisSRCapabilityIsisSRSrgbIter) appendHolderSlice(item IsisSRSrgb) IsisSRCapabilityIsisSRSrgbIter {
	obj.isisSRSrgbSlice = append(obj.isisSRSrgbSlice, item)
	return obj
}

func (obj *isisSRCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SrgbRanges) != 0 {

		if set_default {
			obj.SrgbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrgbRanges {
				obj.SrgbRanges().appendHolderSlice(&isisSRSrgb{obj: item})
			}
		}
		for _, item := range obj.SrgbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisSRCapability) setDefault() {
	if obj.obj.IFlag == nil {
		obj.SetIFlag(true)
	}
	if obj.obj.VFlag == nil {
		obj.SetVFlag(true)
	}

}
