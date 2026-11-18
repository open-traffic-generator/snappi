package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Ospfv2LsaSrCapability *****
type ospfv2LsaSrCapability struct {
	validation
	obj              *otg.Ospfv2LsaSrCapability
	marshaller       marshalOspfv2LsaSrCapability
	unMarshaller     unMarshalOspfv2LsaSrCapability
	srgbRangesHolder Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	srlbRangesHolder Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
}

func NewOspfv2LsaSrCapability() Ospfv2LsaSrCapability {
	obj := ospfv2LsaSrCapability{obj: &otg.Ospfv2LsaSrCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *ospfv2LsaSrCapability) msg() *otg.Ospfv2LsaSrCapability {
	return obj.obj
}

func (obj *ospfv2LsaSrCapability) setMsg(msg *otg.Ospfv2LsaSrCapability) Ospfv2LsaSrCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalospfv2LsaSrCapability struct {
	obj *ospfv2LsaSrCapability
}

type marshalOspfv2LsaSrCapability interface {
	// ToProto marshals Ospfv2LsaSrCapability to protobuf object *otg.Ospfv2LsaSrCapability
	ToProto() (*otg.Ospfv2LsaSrCapability, error)
	// ToPbText marshals Ospfv2LsaSrCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Ospfv2LsaSrCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals Ospfv2LsaSrCapability to JSON text
	ToJson() (string, error)
}

type unMarshalospfv2LsaSrCapability struct {
	obj *ospfv2LsaSrCapability
}

type unMarshalOspfv2LsaSrCapability interface {
	// FromProto unmarshals Ospfv2LsaSrCapability from protobuf object *otg.Ospfv2LsaSrCapability
	FromProto(msg *otg.Ospfv2LsaSrCapability) (Ospfv2LsaSrCapability, error)
	// FromPbText unmarshals Ospfv2LsaSrCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Ospfv2LsaSrCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Ospfv2LsaSrCapability from JSON text
	FromJson(value string) error
}

func (obj *ospfv2LsaSrCapability) Marshal() marshalOspfv2LsaSrCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalospfv2LsaSrCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *ospfv2LsaSrCapability) Unmarshal() unMarshalOspfv2LsaSrCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalospfv2LsaSrCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalospfv2LsaSrCapability) ToProto() (*otg.Ospfv2LsaSrCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalospfv2LsaSrCapability) FromProto(msg *otg.Ospfv2LsaSrCapability) (Ospfv2LsaSrCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalospfv2LsaSrCapability) ToPbText() (string, error) {
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

func (m *unMarshalospfv2LsaSrCapability) FromPbText(value string) error {
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

func (m *marshalospfv2LsaSrCapability) ToYaml() (string, error) {
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

func (m *unMarshalospfv2LsaSrCapability) FromYaml(value string) error {
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

func (m *marshalospfv2LsaSrCapability) ToJson() (string, error) {
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

func (m *unMarshalospfv2LsaSrCapability) FromJson(value string) error {
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

func (obj *ospfv2LsaSrCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ospfv2LsaSrCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ospfv2LsaSrCapability) Clone() (Ospfv2LsaSrCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewOspfv2LsaSrCapability()
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

func (obj *ospfv2LsaSrCapability) setNil() {
	obj.srgbRangesHolder = nil
	obj.srlbRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Ospfv2LsaSrCapability is the Segment Routing capability learned from the Router Information (RI) Opaque LSA:
// the SR-Algorithm TLV, SID/Label Range (SRGB) TLV and SR Local Block (SRLB) TLV.
// Reference: https://datatracker.ietf.org/doc/html/rfc8665.
type Ospfv2LsaSrCapability interface {
	Validation
	// msg marshals Ospfv2LsaSrCapability to protobuf object *otg.Ospfv2LsaSrCapability
	// and doesn't set defaults
	msg() *otg.Ospfv2LsaSrCapability
	// setMsg unmarshals Ospfv2LsaSrCapability from protobuf object *otg.Ospfv2LsaSrCapability
	// and doesn't set defaults
	setMsg(*otg.Ospfv2LsaSrCapability) Ospfv2LsaSrCapability
	// provides marshal interface
	Marshal() marshalOspfv2LsaSrCapability
	// provides unmarshal interface
	Unmarshal() unMarshalOspfv2LsaSrCapability
	// validate validates Ospfv2LsaSrCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Ospfv2LsaSrCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Algorithms returns []uint32, set in Ospfv2LsaSrCapability.
	Algorithms() []uint32
	// SetAlgorithms assigns []uint32 provided by user to Ospfv2LsaSrCapability
	SetAlgorithms(value []uint32) Ospfv2LsaSrCapability
	// SrgbRanges returns Ospfv2LsaSrCapabilityOspfv2LsaSrgbIterIter, set in Ospfv2LsaSrCapability
	SrgbRanges() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	// SrlbRanges returns Ospfv2LsaSrCapabilityOspfv2LsaSrlbIterIter, set in Ospfv2LsaSrCapability
	SrlbRanges() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	setNil()
}

// The Segment Routing algorithms advertised in the SR-Algorithm TLV (0 = SPF, 1 = Strict SPF).
// Algorithms returns a []uint32
func (obj *ospfv2LsaSrCapability) Algorithms() []uint32 {
	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	return obj.obj.Algorithms
}

// The Segment Routing algorithms advertised in the SR-Algorithm TLV (0 = SPF, 1 = Strict SPF).
// SetAlgorithms sets the []uint32 value in the Ospfv2LsaSrCapability object
func (obj *ospfv2LsaSrCapability) SetAlgorithms(value []uint32) Ospfv2LsaSrCapability {

	if obj.obj.Algorithms == nil {
		obj.obj.Algorithms = make([]uint32, 0)
	}
	obj.obj.Algorithms = value

	return obj
}

// The learned Segment Routing Global Block (SRGB) ranges from the SID/Label Range TLV.
// SrgbRanges returns a []Ospfv2LsaSrgb
func (obj *ospfv2LsaSrCapability) SrgbRanges() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	if len(obj.obj.SrgbRanges) == 0 {
		obj.obj.SrgbRanges = []*otg.Ospfv2LsaSrgb{}
	}
	if obj.srgbRangesHolder == nil {
		obj.srgbRangesHolder = newOspfv2LsaSrCapabilityOspfv2LsaSrgbIter(&obj.obj.SrgbRanges).setMsg(obj)
	}
	return obj.srgbRangesHolder
}

type ospfv2LsaSrCapabilityOspfv2LsaSrgbIter struct {
	obj                *ospfv2LsaSrCapability
	ospfv2LsaSrgbSlice []Ospfv2LsaSrgb
	fieldPtr           *[]*otg.Ospfv2LsaSrgb
}

func newOspfv2LsaSrCapabilityOspfv2LsaSrgbIter(ptr *[]*otg.Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	return &ospfv2LsaSrCapabilityOspfv2LsaSrgbIter{fieldPtr: ptr}
}

type Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter interface {
	setMsg(*ospfv2LsaSrCapability) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	Items() []Ospfv2LsaSrgb
	Add() Ospfv2LsaSrgb
	Append(items ...Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	Set(index int, newObj Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	Clear() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	clearHolderSlice() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
	appendHolderSlice(item Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) setMsg(msg *ospfv2LsaSrCapability) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaSrgb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) Items() []Ospfv2LsaSrgb {
	return obj.ospfv2LsaSrgbSlice
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) Add() Ospfv2LsaSrgb {
	newObj := &otg.Ospfv2LsaSrgb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaSrgb{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaSrgbSlice = append(obj.ospfv2LsaSrgbSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) Append(items ...Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaSrgbSlice = append(obj.ospfv2LsaSrgbSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) Set(index int, newObj Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaSrgbSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) Clear() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaSrgb{}
		obj.ospfv2LsaSrgbSlice = []Ospfv2LsaSrgb{}
	}
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) clearHolderSlice() Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	if len(obj.ospfv2LsaSrgbSlice) > 0 {
		obj.ospfv2LsaSrgbSlice = []Ospfv2LsaSrgb{}
	}
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrgbIter) appendHolderSlice(item Ospfv2LsaSrgb) Ospfv2LsaSrCapabilityOspfv2LsaSrgbIter {
	obj.ospfv2LsaSrgbSlice = append(obj.ospfv2LsaSrgbSlice, item)
	return obj
}

// The learned SR Local Block (SRLB) ranges from the SR Local Block TLV.
// SrlbRanges returns a []Ospfv2LsaSrlb
func (obj *ospfv2LsaSrCapability) SrlbRanges() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	if len(obj.obj.SrlbRanges) == 0 {
		obj.obj.SrlbRanges = []*otg.Ospfv2LsaSrlb{}
	}
	if obj.srlbRangesHolder == nil {
		obj.srlbRangesHolder = newOspfv2LsaSrCapabilityOspfv2LsaSrlbIter(&obj.obj.SrlbRanges).setMsg(obj)
	}
	return obj.srlbRangesHolder
}

type ospfv2LsaSrCapabilityOspfv2LsaSrlbIter struct {
	obj                *ospfv2LsaSrCapability
	ospfv2LsaSrlbSlice []Ospfv2LsaSrlb
	fieldPtr           *[]*otg.Ospfv2LsaSrlb
}

func newOspfv2LsaSrCapabilityOspfv2LsaSrlbIter(ptr *[]*otg.Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	return &ospfv2LsaSrCapabilityOspfv2LsaSrlbIter{fieldPtr: ptr}
}

type Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter interface {
	setMsg(*ospfv2LsaSrCapability) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	Items() []Ospfv2LsaSrlb
	Add() Ospfv2LsaSrlb
	Append(items ...Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	Set(index int, newObj Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	Clear() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	clearHolderSlice() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
	appendHolderSlice(item Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) setMsg(msg *ospfv2LsaSrCapability) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ospfv2LsaSrlb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) Items() []Ospfv2LsaSrlb {
	return obj.ospfv2LsaSrlbSlice
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) Add() Ospfv2LsaSrlb {
	newObj := &otg.Ospfv2LsaSrlb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ospfv2LsaSrlb{obj: newObj}
	newLibObj.setDefault()
	obj.ospfv2LsaSrlbSlice = append(obj.ospfv2LsaSrlbSlice, newLibObj)
	return newLibObj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) Append(items ...Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ospfv2LsaSrlbSlice = append(obj.ospfv2LsaSrlbSlice, item)
	}
	return obj
}

func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) Set(index int, newObj Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ospfv2LsaSrlbSlice[index] = newObj
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) Clear() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Ospfv2LsaSrlb{}
		obj.ospfv2LsaSrlbSlice = []Ospfv2LsaSrlb{}
	}
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) clearHolderSlice() Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	if len(obj.ospfv2LsaSrlbSlice) > 0 {
		obj.ospfv2LsaSrlbSlice = []Ospfv2LsaSrlb{}
	}
	return obj
}
func (obj *ospfv2LsaSrCapabilityOspfv2LsaSrlbIter) appendHolderSlice(item Ospfv2LsaSrlb) Ospfv2LsaSrCapabilityOspfv2LsaSrlbIter {
	obj.ospfv2LsaSrlbSlice = append(obj.ospfv2LsaSrlbSlice, item)
	return obj
}

func (obj *ospfv2LsaSrCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.SrgbRanges) != 0 {

		if set_default {
			obj.SrgbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrgbRanges {
				obj.SrgbRanges().appendHolderSlice(&ospfv2LsaSrgb{obj: item})
			}
		}
		for _, item := range obj.SrgbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.SrlbRanges) != 0 {

		if set_default {
			obj.SrlbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrlbRanges {
				obj.SrlbRanges().appendHolderSlice(&ospfv2LsaSrlb{obj: item})
			}
		}
		for _, item := range obj.SrlbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ospfv2LsaSrCapability) setDefault() {

}
