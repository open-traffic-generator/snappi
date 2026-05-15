package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6UsidDst *****
type flowIpv6UsidDst struct {
	validation
	obj                 *otg.FlowIpv6UsidDst
	marshaller          marshalFlowIpv6UsidDst
	unMarshaller        unMarshalFlowIpv6UsidDst
	locatorHolder       PatternFlowIpv6UsidDstLocator
	locatorLengthHolder PatternFlowIpv6UsidDstLocatorLength
	usidsHolder         FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
}

func NewFlowIpv6UsidDst() FlowIpv6UsidDst {
	obj := flowIpv6UsidDst{obj: &otg.FlowIpv6UsidDst{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6UsidDst) msg() *otg.FlowIpv6UsidDst {
	return obj.obj
}

func (obj *flowIpv6UsidDst) setMsg(msg *otg.FlowIpv6UsidDst) FlowIpv6UsidDst {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6UsidDst struct {
	obj *flowIpv6UsidDst
}

type marshalFlowIpv6UsidDst interface {
	// ToProto marshals FlowIpv6UsidDst to protobuf object *otg.FlowIpv6UsidDst
	ToProto() (*otg.FlowIpv6UsidDst, error)
	// ToPbText marshals FlowIpv6UsidDst to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6UsidDst to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6UsidDst to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6UsidDst struct {
	obj *flowIpv6UsidDst
}

type unMarshalFlowIpv6UsidDst interface {
	// FromProto unmarshals FlowIpv6UsidDst from protobuf object *otg.FlowIpv6UsidDst
	FromProto(msg *otg.FlowIpv6UsidDst) (FlowIpv6UsidDst, error)
	// FromPbText unmarshals FlowIpv6UsidDst from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6UsidDst from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6UsidDst from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6UsidDst) Marshal() marshalFlowIpv6UsidDst {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6UsidDst{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6UsidDst) Unmarshal() unMarshalFlowIpv6UsidDst {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6UsidDst{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6UsidDst) ToProto() (*otg.FlowIpv6UsidDst, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6UsidDst) FromProto(msg *otg.FlowIpv6UsidDst) (FlowIpv6UsidDst, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6UsidDst) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromPbText(value string) error {
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

func (m *marshalflowIpv6UsidDst) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromYaml(value string) error {
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

func (m *marshalflowIpv6UsidDst) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6UsidDst) FromJson(value string) error {
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

func (obj *flowIpv6UsidDst) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDst) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6UsidDst) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6UsidDst) Clone() (FlowIpv6UsidDst, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6UsidDst()
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

func (obj *flowIpv6UsidDst) setNil() {
	obj.locatorHolder = nil
	obj.locatorLengthHolder = nil
	obj.usidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6UsidDst is structured input for the SRv6 uSID reduced encapsulation case (RFC 9800 Section 4) where the entire SR path fits in a single 128-bit uSID container placed directly in the outer IPv6 dst, with no Segment Routing Header. The implementation packs the fields as: LB (locator_length high-order bits of locator) || uSID-1 || uSID-2 || ... || EoC (zero-pad to 128 bits).
// For F3216 format (RFC 9800 Section 3): LB = 32 bits, each uSID = 16 bits, up to 6 uSIDs per container. Example - locator fc00::/32, locator_length 32, usids ["0001","0002","0003"]: assembled dst = fc00:0:1:2:3::
// When this field is present, ipv6.dst should be ignored by the implementation.
type FlowIpv6UsidDst interface {
	Validation
	// msg marshals FlowIpv6UsidDst to protobuf object *otg.FlowIpv6UsidDst
	// and doesn't set defaults
	msg() *otg.FlowIpv6UsidDst
	// setMsg unmarshals FlowIpv6UsidDst from protobuf object *otg.FlowIpv6UsidDst
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6UsidDst) FlowIpv6UsidDst
	// provides marshal interface
	Marshal() marshalFlowIpv6UsidDst
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6UsidDst
	// validate validates FlowIpv6UsidDst
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6UsidDst, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns PatternFlowIpv6UsidDstLocator, set in FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits packed into the 128-bit dst. For F3216 this is a /32 prefix (e.g., fc00::). The number of bits used is given by locator_length.
	Locator() PatternFlowIpv6UsidDstLocator
	// SetLocator assigns PatternFlowIpv6UsidDstLocator provided by user to FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits packed into the 128-bit dst. For F3216 this is a /32 prefix (e.g., fc00::). The number of bits used is given by locator_length.
	SetLocator(value PatternFlowIpv6UsidDstLocator) FlowIpv6UsidDst
	// LocatorLength returns PatternFlowIpv6UsidDstLocatorLength, set in FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for uSID packing. For F3216: 32. For F3208: 32. Valid range: 1-112.
	LocatorLength() PatternFlowIpv6UsidDstLocatorLength
	// SetLocatorLength assigns PatternFlowIpv6UsidDstLocatorLength provided by user to FlowIpv6UsidDst.
	// PatternFlowIpv6UsidDstLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for uSID packing. For F3216: 32. For F3208: 32. Valid range: 1-112.
	SetLocatorLength(value PatternFlowIpv6UsidDstLocatorLength) FlowIpv6UsidDst
	// HasLocatorLength checks if LocatorLength has been set in FlowIpv6UsidDst
	HasLocatorLength() bool
	// Usids returns FlowIpv6UsidDstFlowIpv6UsidDstuSidIterIter, set in FlowIpv6UsidDst
	Usids() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	setNil()
}

// description is TBD
// Locator returns a PatternFlowIpv6UsidDstLocator
func (obj *flowIpv6UsidDst) Locator() PatternFlowIpv6UsidDstLocator {
	if obj.obj.Locator == nil {
		obj.obj.Locator = NewPatternFlowIpv6UsidDstLocator().msg()
	}
	if obj.locatorHolder == nil {
		obj.locatorHolder = &patternFlowIpv6UsidDstLocator{obj: obj.obj.Locator}
	}
	return obj.locatorHolder
}

// description is TBD
// SetLocator sets the PatternFlowIpv6UsidDstLocator value in the FlowIpv6UsidDst object
func (obj *flowIpv6UsidDst) SetLocator(value PatternFlowIpv6UsidDstLocator) FlowIpv6UsidDst {

	obj.locatorHolder = nil
	obj.obj.Locator = value.msg()

	return obj
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6UsidDstLocatorLength
func (obj *flowIpv6UsidDst) LocatorLength() PatternFlowIpv6UsidDstLocatorLength {
	if obj.obj.LocatorLength == nil {
		obj.obj.LocatorLength = NewPatternFlowIpv6UsidDstLocatorLength().msg()
	}
	if obj.locatorLengthHolder == nil {
		obj.locatorLengthHolder = &patternFlowIpv6UsidDstLocatorLength{obj: obj.obj.LocatorLength}
	}
	return obj.locatorLengthHolder
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6UsidDstLocatorLength
func (obj *flowIpv6UsidDst) HasLocatorLength() bool {
	return obj.obj.LocatorLength != nil
}

// description is TBD
// SetLocatorLength sets the PatternFlowIpv6UsidDstLocatorLength value in the FlowIpv6UsidDst object
func (obj *flowIpv6UsidDst) SetLocatorLength(value PatternFlowIpv6UsidDstLocatorLength) FlowIpv6UsidDst {

	obj.locatorLengthHolder = nil
	obj.obj.LocatorLength = value.msg()

	return obj
}

// Ordered list of uSID values to pack after the Locator Block (RFC 9800 Section 3). For F3216 each uSID is 16 bits (4 hex chars); up to 6 uSIDs fit per container. The End-of-Container zero-pad is appended automatically by the implementation. Example: ["0001","0002","0003"] with locator fc00::/32 assembles to fc00:0:1:2:3::
// Usids returns a []FlowIpv6UsidDstuSid
func (obj *flowIpv6UsidDst) Usids() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	if len(obj.obj.Usids) == 0 {
		obj.obj.Usids = []*otg.FlowIpv6UsidDstuSid{}
	}
	if obj.usidsHolder == nil {
		obj.usidsHolder = newFlowIpv6UsidDstFlowIpv6UsidDstuSidIter(&obj.obj.Usids).setMsg(obj)
	}
	return obj.usidsHolder
}

type flowIpv6UsidDstFlowIpv6UsidDstuSidIter struct {
	obj                      *flowIpv6UsidDst
	flowIpv6UsidDstuSidSlice []FlowIpv6UsidDstuSid
	fieldPtr                 *[]*otg.FlowIpv6UsidDstuSid
}

func newFlowIpv6UsidDstFlowIpv6UsidDstuSidIter(ptr *[]*otg.FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	return &flowIpv6UsidDstFlowIpv6UsidDstuSidIter{fieldPtr: ptr}
}

type FlowIpv6UsidDstFlowIpv6UsidDstuSidIter interface {
	setMsg(*flowIpv6UsidDst) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	Items() []FlowIpv6UsidDstuSid
	Add() FlowIpv6UsidDstuSid
	Append(items ...FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	Set(index int, newObj FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	Clear() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	clearHolderSlice() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
	appendHolderSlice(item FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter
}

func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) setMsg(msg *flowIpv6UsidDst) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv6UsidDstuSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) Items() []FlowIpv6UsidDstuSid {
	return obj.flowIpv6UsidDstuSidSlice
}

func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) Add() FlowIpv6UsidDstuSid {
	newObj := &otg.FlowIpv6UsidDstuSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv6UsidDstuSid{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv6UsidDstuSidSlice = append(obj.flowIpv6UsidDstuSidSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) Append(items ...FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv6UsidDstuSidSlice = append(obj.flowIpv6UsidDstuSidSlice, item)
	}
	return obj
}

func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) Set(index int, newObj FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv6UsidDstuSidSlice[index] = newObj
	return obj
}
func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) Clear() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv6UsidDstuSid{}
		obj.flowIpv6UsidDstuSidSlice = []FlowIpv6UsidDstuSid{}
	}
	return obj
}
func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) clearHolderSlice() FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	if len(obj.flowIpv6UsidDstuSidSlice) > 0 {
		obj.flowIpv6UsidDstuSidSlice = []FlowIpv6UsidDstuSid{}
	}
	return obj
}
func (obj *flowIpv6UsidDstFlowIpv6UsidDstuSidIter) appendHolderSlice(item FlowIpv6UsidDstuSid) FlowIpv6UsidDstFlowIpv6UsidDstuSidIter {
	obj.flowIpv6UsidDstuSidSlice = append(obj.flowIpv6UsidDstuSidSlice, item)
	return obj
}

func (obj *flowIpv6UsidDst) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Locator is required
	if obj.obj.Locator == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Locator is required field on interface FlowIpv6UsidDst")
	}

	if obj.obj.Locator != nil {

		obj.Locator().validateObj(vObj, set_default)
	}

	if obj.obj.LocatorLength != nil {

		obj.LocatorLength().validateObj(vObj, set_default)
	}

	if len(obj.obj.Usids) != 0 {

		if set_default {
			obj.Usids().clearHolderSlice()
			for _, item := range obj.obj.Usids {
				obj.Usids().appendHolderSlice(&flowIpv6UsidDstuSid{obj: item})
			}
		}
		for _, item := range obj.Usids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv6UsidDst) setDefault() {

}
