package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowIpv6SegmentRoutingUsidSegment *****
type flowIpv6SegmentRoutingUsidSegment struct {
	validation
	obj                 *otg.FlowIpv6SegmentRoutingUsidSegment
	marshaller          marshalFlowIpv6SegmentRoutingUsidSegment
	unMarshaller        unMarshalFlowIpv6SegmentRoutingUsidSegment
	locatorHolder       PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	locatorLengthHolder PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	usidsHolder         FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
}

func NewFlowIpv6SegmentRoutingUsidSegment() FlowIpv6SegmentRoutingUsidSegment {
	obj := flowIpv6SegmentRoutingUsidSegment{obj: &otg.FlowIpv6SegmentRoutingUsidSegment{}}
	obj.setDefault()
	return &obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) msg() *otg.FlowIpv6SegmentRoutingUsidSegment {
	return obj.obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) setMsg(msg *otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type marshalFlowIpv6SegmentRoutingUsidSegment interface {
	// ToProto marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error)
	// ToPbText marshals FlowIpv6SegmentRoutingUsidSegment to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowIpv6SegmentRoutingUsidSegment to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowIpv6SegmentRoutingUsidSegment to JSON text
	ToJson() (string, error)
}

type unMarshalflowIpv6SegmentRoutingUsidSegment struct {
	obj *flowIpv6SegmentRoutingUsidSegment
}

type unMarshalFlowIpv6SegmentRoutingUsidSegment interface {
	// FromProto unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error)
	// FromPbText unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowIpv6SegmentRoutingUsidSegment from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowIpv6SegmentRoutingUsidSegment from JSON text
	FromJson(value string) error
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Marshal() marshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowIpv6SegmentRoutingUsidSegment{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToProto() (*otg.FlowIpv6SegmentRoutingUsidSegment, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromProto(msg *otg.FlowIpv6SegmentRoutingUsidSegment) (FlowIpv6SegmentRoutingUsidSegment, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToPbText() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromPbText(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToYaml() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromYaml(value string) error {
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

func (m *marshalflowIpv6SegmentRoutingUsidSegment) ToJson() (string, error) {
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

func (m *unMarshalflowIpv6SegmentRoutingUsidSegment) FromJson(value string) error {
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

func (obj *flowIpv6SegmentRoutingUsidSegment) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowIpv6SegmentRoutingUsidSegment) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowIpv6SegmentRoutingUsidSegment) Clone() (FlowIpv6SegmentRoutingUsidSegment, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowIpv6SegmentRoutingUsidSegment()
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

func (obj *flowIpv6SegmentRoutingUsidSegment) setNil() {
	obj.locatorHolder = nil
	obj.locatorLengthHolder = nil
	obj.usidsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowIpv6SegmentRoutingUsidSegment is one compressed uSID container entry in the SRH segment list (RFC 9800 Section 4,
// RFC 8754 Section 2.1). The implementation assembles the 128-bit wire
// value by packing the locator block followed by the uSID values in order,
// zero-padding the remainder to 128 bits (End-of-Container marker).
//
// For F3216 format (RFC 9800 Section 3): Locator Block = 32 bits, each uSID = 16 bits,
// up to 6 uSIDs per container.
// Example - locator fc00::/32 (locator_length 32), usids ["0001","0002","0003"]
// assembles to the 128-bit SRH entry fc00:0:1:2:3::
type FlowIpv6SegmentRoutingUsidSegment interface {
	Validation
	// msg marshals FlowIpv6SegmentRoutingUsidSegment to protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	msg() *otg.FlowIpv6SegmentRoutingUsidSegment
	// setMsg unmarshals FlowIpv6SegmentRoutingUsidSegment from protobuf object *otg.FlowIpv6SegmentRoutingUsidSegment
	// and doesn't set defaults
	setMsg(*otg.FlowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegment
	// provides marshal interface
	Marshal() marshalFlowIpv6SegmentRoutingUsidSegment
	// provides unmarshal interface
	Unmarshal() unMarshalFlowIpv6SegmentRoutingUsidSegment
	// validate validates FlowIpv6SegmentRoutingUsidSegment
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowIpv6SegmentRoutingUsidSegment, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Locator returns PatternFlowIpv6SegmentRoutingUsidSegmentLocator, set in FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits of every uSID assembled from this entry. For F3216, this is a /32 prefix (e.g., fc00::). The prefix length is given by locator_length.
	Locator() PatternFlowIpv6SegmentRoutingUsidSegmentLocator
	// SetLocator assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocator provided by user to FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocator is the Locator Block (LB) IPv6 prefix shared by all uSIDs in this container (RFC 9800 Section 3). Defines the common high-order bits of every uSID assembled from this entry. For F3216, this is a /32 prefix (e.g., fc00::). The prefix length is given by locator_length.
	SetLocator(value PatternFlowIpv6SegmentRoutingUsidSegmentLocator) FlowIpv6SegmentRoutingUsidSegment
	// HasLocator checks if Locator has been set in FlowIpv6SegmentRoutingUsidSegment
	HasLocator() bool
	// LocatorLength returns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength, set in FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for uSID packing. For F3216: 32. For F3208: 32. Valid range: 1-112.
	LocatorLength() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
	// SetLocatorLength assigns PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength provided by user to FlowIpv6SegmentRoutingUsidSegment.
	// PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength is length of the Locator Block in bits (RFC 9800 Section 3). Determines how many high-order bits of locator are used as the LB and how many bits remain for uSID packing. For F3216: 32. For F3208: 32. Valid range: 1-112.
	SetLocatorLength(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FlowIpv6SegmentRoutingUsidSegment
	// HasLocatorLength checks if LocatorLength has been set in FlowIpv6SegmentRoutingUsidSegment
	HasLocatorLength() bool
	// Usids returns FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIterIter, set in FlowIpv6SegmentRoutingUsidSegment
	Usids() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	setNil()
}

// description is TBD
// Locator returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocator
func (obj *flowIpv6SegmentRoutingUsidSegment) Locator() PatternFlowIpv6SegmentRoutingUsidSegmentLocator {
	if obj.obj.Locator == nil {
		obj.obj.Locator = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocator().msg()
	}
	if obj.locatorHolder == nil {
		obj.locatorHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocator{obj: obj.obj.Locator}
	}
	return obj.locatorHolder
}

// description is TBD
// Locator returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocator
func (obj *flowIpv6SegmentRoutingUsidSegment) HasLocator() bool {
	return obj.obj.Locator != nil
}

// description is TBD
// SetLocator sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocator value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetLocator(value PatternFlowIpv6SegmentRoutingUsidSegmentLocator) FlowIpv6SegmentRoutingUsidSegment {

	obj.locatorHolder = nil
	obj.obj.Locator = value.msg()

	return obj
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
func (obj *flowIpv6SegmentRoutingUsidSegment) LocatorLength() PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength {
	if obj.obj.LocatorLength == nil {
		obj.obj.LocatorLength = NewPatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength().msg()
	}
	if obj.locatorLengthHolder == nil {
		obj.locatorLengthHolder = &patternFlowIpv6SegmentRoutingUsidSegmentLocatorLength{obj: obj.obj.LocatorLength}
	}
	return obj.locatorLengthHolder
}

// description is TBD
// LocatorLength returns a PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength
func (obj *flowIpv6SegmentRoutingUsidSegment) HasLocatorLength() bool {
	return obj.obj.LocatorLength != nil
}

// description is TBD
// SetLocatorLength sets the PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength value in the FlowIpv6SegmentRoutingUsidSegment object
func (obj *flowIpv6SegmentRoutingUsidSegment) SetLocatorLength(value PatternFlowIpv6SegmentRoutingUsidSegmentLocatorLength) FlowIpv6SegmentRoutingUsidSegment {

	obj.locatorLengthHolder = nil
	obj.obj.LocatorLength = value.msg()

	return obj
}

// Ordered list of uSID values to pack into this container after the Locator Block (RFC 9800 Section 3). Each uSID occupies ((128 - locator_length) / no of uSIDs) bits. For F3216 each uSID is 16 bits (4 hex chars) and up to 6 uSIDs fit per container. The implementation appends the End-of-Container zero-pad automatically. Example for F3216: ["0001","0002","0003"] with locator fc00::/32 assembles to SRH entry fc00:0:1:2:3::
// Usids returns a []FlowIpv6SegmentRoutingUsiduSid
func (obj *flowIpv6SegmentRoutingUsidSegment) Usids() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	if len(obj.obj.Usids) == 0 {
		obj.obj.Usids = []*otg.FlowIpv6SegmentRoutingUsiduSid{}
	}
	if obj.usidsHolder == nil {
		obj.usidsHolder = newFlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter(&obj.obj.Usids).setMsg(obj)
	}
	return obj.usidsHolder
}

type flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter struct {
	obj                                 *flowIpv6SegmentRoutingUsidSegment
	flowIpv6SegmentRoutingUsiduSidSlice []FlowIpv6SegmentRoutingUsiduSid
	fieldPtr                            *[]*otg.FlowIpv6SegmentRoutingUsiduSid
}

func newFlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter(ptr *[]*otg.FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	return &flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter{fieldPtr: ptr}
}

type FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter interface {
	setMsg(*flowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	Items() []FlowIpv6SegmentRoutingUsiduSid
	Add() FlowIpv6SegmentRoutingUsiduSid
	Append(items ...FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	Set(index int, newObj FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	Clear() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	clearHolderSlice() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
	appendHolderSlice(item FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter
}

func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) setMsg(msg *flowIpv6SegmentRoutingUsidSegment) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&flowIpv6SegmentRoutingUsiduSid{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) Items() []FlowIpv6SegmentRoutingUsiduSid {
	return obj.flowIpv6SegmentRoutingUsiduSidSlice
}

func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) Add() FlowIpv6SegmentRoutingUsiduSid {
	newObj := &otg.FlowIpv6SegmentRoutingUsiduSid{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &flowIpv6SegmentRoutingUsiduSid{obj: newObj}
	newLibObj.setDefault()
	obj.flowIpv6SegmentRoutingUsiduSidSlice = append(obj.flowIpv6SegmentRoutingUsiduSidSlice, newLibObj)
	return newLibObj
}

func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) Append(items ...FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.flowIpv6SegmentRoutingUsiduSidSlice = append(obj.flowIpv6SegmentRoutingUsiduSidSlice, item)
	}
	return obj
}

func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) Set(index int, newObj FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.flowIpv6SegmentRoutingUsiduSidSlice[index] = newObj
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) Clear() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.FlowIpv6SegmentRoutingUsiduSid{}
		obj.flowIpv6SegmentRoutingUsiduSidSlice = []FlowIpv6SegmentRoutingUsiduSid{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) clearHolderSlice() FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	if len(obj.flowIpv6SegmentRoutingUsiduSidSlice) > 0 {
		obj.flowIpv6SegmentRoutingUsiduSidSlice = []FlowIpv6SegmentRoutingUsiduSid{}
	}
	return obj
}
func (obj *flowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter) appendHolderSlice(item FlowIpv6SegmentRoutingUsiduSid) FlowIpv6SegmentRoutingUsidSegmentFlowIpv6SegmentRoutingUsiduSidIter {
	obj.flowIpv6SegmentRoutingUsiduSidSlice = append(obj.flowIpv6SegmentRoutingUsiduSidSlice, item)
	return obj
}

func (obj *flowIpv6SegmentRoutingUsidSegment) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
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
				obj.Usids().appendHolderSlice(&flowIpv6SegmentRoutingUsiduSid{obj: item})
			}
		}
		for _, item := range obj.Usids().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *flowIpv6SegmentRoutingUsidSegment) setDefault() {

}
