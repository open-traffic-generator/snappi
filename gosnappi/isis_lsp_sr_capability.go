package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisLspSRCapability *****
type isisLspSRCapability struct {
	validation
	obj              *otg.IsisLspSRCapability
	marshaller       marshalIsisLspSRCapability
	unMarshaller     unMarshalIsisLspSRCapability
	flagsHolder      IsisLspCapasFlags
	srgbRangesHolder IsisLspSRCapabilityIsisLspSrgbIter
}

func NewIsisLspSRCapability() IsisLspSRCapability {
	obj := isisLspSRCapability{obj: &otg.IsisLspSRCapability{}}
	obj.setDefault()
	return &obj
}

func (obj *isisLspSRCapability) msg() *otg.IsisLspSRCapability {
	return obj.obj
}

func (obj *isisLspSRCapability) setMsg(msg *otg.IsisLspSRCapability) IsisLspSRCapability {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisLspSRCapability struct {
	obj *isisLspSRCapability
}

type marshalIsisLspSRCapability interface {
	// ToProto marshals IsisLspSRCapability to protobuf object *otg.IsisLspSRCapability
	ToProto() (*otg.IsisLspSRCapability, error)
	// ToPbText marshals IsisLspSRCapability to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisLspSRCapability to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisLspSRCapability to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals IsisLspSRCapability to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalisisLspSRCapability struct {
	obj *isisLspSRCapability
}

type unMarshalIsisLspSRCapability interface {
	// FromProto unmarshals IsisLspSRCapability from protobuf object *otg.IsisLspSRCapability
	FromProto(msg *otg.IsisLspSRCapability) (IsisLspSRCapability, error)
	// FromPbText unmarshals IsisLspSRCapability from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisLspSRCapability from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisLspSRCapability from JSON text
	FromJson(value string) error
}

func (obj *isisLspSRCapability) Marshal() marshalIsisLspSRCapability {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisLspSRCapability{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisLspSRCapability) Unmarshal() unMarshalIsisLspSRCapability {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisLspSRCapability{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisLspSRCapability) ToProto() (*otg.IsisLspSRCapability, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisLspSRCapability) FromProto(msg *otg.IsisLspSRCapability) (IsisLspSRCapability, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisLspSRCapability) ToPbText() (string, error) {
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

func (m *unMarshalisisLspSRCapability) FromPbText(value string) error {
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

func (m *marshalisisLspSRCapability) ToYaml() (string, error) {
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

func (m *unMarshalisisLspSRCapability) FromYaml(value string) error {
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

func (m *marshalisisLspSRCapability) ToJsonRaw() (string, error) {
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

func (m *marshalisisLspSRCapability) ToJson() (string, error) {
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

func (m *unMarshalisisLspSRCapability) FromJson(value string) error {
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

func (obj *isisLspSRCapability) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisLspSRCapability) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisLspSRCapability) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisLspSRCapability) Clone() (IsisLspSRCapability, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisLspSRCapability()
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

func (obj *isisLspSRCapability) setNil() {
	obj.flagsHolder = nil
	obj.srgbRangesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// IsisLspSRCapability is container of IS-IS SR-CAPABILITY Sub-Tlv.
type IsisLspSRCapability interface {
	Validation
	// msg marshals IsisLspSRCapability to protobuf object *otg.IsisLspSRCapability
	// and doesn't set defaults
	msg() *otg.IsisLspSRCapability
	// setMsg unmarshals IsisLspSRCapability from protobuf object *otg.IsisLspSRCapability
	// and doesn't set defaults
	setMsg(*otg.IsisLspSRCapability) IsisLspSRCapability
	// provides marshal interface
	Marshal() marshalIsisLspSRCapability
	// provides unmarshal interface
	Unmarshal() unMarshalIsisLspSRCapability
	// validate validates IsisLspSRCapability
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisLspSRCapability, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Flags returns IsisLspCapasFlags, set in IsisLspSRCapability.
	// IsisLspCapasFlags is container for the configuration of IS-IS SR-CAPABILITY flags.
	Flags() IsisLspCapasFlags
	// SetFlags assigns IsisLspCapasFlags provided by user to IsisLspSRCapability.
	// IsisLspCapasFlags is container for the configuration of IS-IS SR-CAPABILITY flags.
	SetFlags(value IsisLspCapasFlags) IsisLspSRCapability
	// HasFlags checks if Flags has been set in IsisLspSRCapability
	HasFlags() bool
	// SrgbRanges returns IsisLspSRCapabilityIsisLspSrgbIterIter, set in IsisLspSRCapability
	SrgbRanges() IsisLspSRCapabilityIsisLspSrgbIter
	setNil()
}

// 1 octet of flags.
// Flags returns a IsisLspCapasFlags
func (obj *isisLspSRCapability) Flags() IsisLspCapasFlags {
	if obj.obj.Flags == nil {
		obj.obj.Flags = NewIsisLspCapasFlags().msg()
	}
	if obj.flagsHolder == nil {
		obj.flagsHolder = &isisLspCapasFlags{obj: obj.obj.Flags}
	}
	return obj.flagsHolder
}

// 1 octet of flags.
// Flags returns a IsisLspCapasFlags
func (obj *isisLspSRCapability) HasFlags() bool {
	return obj.obj.Flags != nil
}

// 1 octet of flags.
// SetFlags sets the IsisLspCapasFlags value in the IsisLspSRCapability object
func (obj *isisLspSRCapability) SetFlags(value IsisLspCapasFlags) IsisLspSRCapability {

	obj.flagsHolder = nil
	obj.obj.Flags = value.msg()

	return obj
}

// This contains the list of SRGB.
// SrgbRanges returns a []IsisLspSrgb
func (obj *isisLspSRCapability) SrgbRanges() IsisLspSRCapabilityIsisLspSrgbIter {
	if len(obj.obj.SrgbRanges) == 0 {
		obj.obj.SrgbRanges = []*otg.IsisLspSrgb{}
	}
	if obj.srgbRangesHolder == nil {
		obj.srgbRangesHolder = newIsisLspSRCapabilityIsisLspSrgbIter(&obj.obj.SrgbRanges).setMsg(obj)
	}
	return obj.srgbRangesHolder
}

type isisLspSRCapabilityIsisLspSrgbIter struct {
	obj              *isisLspSRCapability
	isisLspSrgbSlice []IsisLspSrgb
	fieldPtr         *[]*otg.IsisLspSrgb
}

func newIsisLspSRCapabilityIsisLspSrgbIter(ptr *[]*otg.IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter {
	return &isisLspSRCapabilityIsisLspSrgbIter{fieldPtr: ptr}
}

type IsisLspSRCapabilityIsisLspSrgbIter interface {
	setMsg(*isisLspSRCapability) IsisLspSRCapabilityIsisLspSrgbIter
	Items() []IsisLspSrgb
	Add() IsisLspSrgb
	Append(items ...IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter
	Set(index int, newObj IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter
	Clear() IsisLspSRCapabilityIsisLspSrgbIter
	clearHolderSlice() IsisLspSRCapabilityIsisLspSrgbIter
	appendHolderSlice(item IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter
}

func (obj *isisLspSRCapabilityIsisLspSrgbIter) setMsg(msg *isisLspSRCapability) IsisLspSRCapabilityIsisLspSrgbIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&isisLspSrgb{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *isisLspSRCapabilityIsisLspSrgbIter) Items() []IsisLspSrgb {
	return obj.isisLspSrgbSlice
}

func (obj *isisLspSRCapabilityIsisLspSrgbIter) Add() IsisLspSrgb {
	newObj := &otg.IsisLspSrgb{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &isisLspSrgb{obj: newObj}
	newLibObj.setDefault()
	obj.isisLspSrgbSlice = append(obj.isisLspSrgbSlice, newLibObj)
	return newLibObj
}

func (obj *isisLspSRCapabilityIsisLspSrgbIter) Append(items ...IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.isisLspSrgbSlice = append(obj.isisLspSrgbSlice, item)
	}
	return obj
}

func (obj *isisLspSRCapabilityIsisLspSrgbIter) Set(index int, newObj IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.isisLspSrgbSlice[index] = newObj
	return obj
}
func (obj *isisLspSRCapabilityIsisLspSrgbIter) Clear() IsisLspSRCapabilityIsisLspSrgbIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.IsisLspSrgb{}
		obj.isisLspSrgbSlice = []IsisLspSrgb{}
	}
	return obj
}
func (obj *isisLspSRCapabilityIsisLspSrgbIter) clearHolderSlice() IsisLspSRCapabilityIsisLspSrgbIter {
	if len(obj.isisLspSrgbSlice) > 0 {
		obj.isisLspSrgbSlice = []IsisLspSrgb{}
	}
	return obj
}
func (obj *isisLspSRCapabilityIsisLspSrgbIter) appendHolderSlice(item IsisLspSrgb) IsisLspSRCapabilityIsisLspSrgbIter {
	obj.isisLspSrgbSlice = append(obj.isisLspSrgbSlice, item)
	return obj
}

func (obj *isisLspSRCapability) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Flags != nil {

		obj.Flags().validateObj(vObj, set_default)
	}

	if len(obj.obj.SrgbRanges) != 0 {

		if set_default {
			obj.SrgbRanges().clearHolderSlice()
			for _, item := range obj.obj.SrgbRanges {
				obj.SrgbRanges().appendHolderSlice(&isisLspSrgb{obj: item})
			}
		}
		for _, item := range obj.SrgbRanges().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *isisLspSRCapability) setDefault() {

}
