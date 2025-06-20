package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesSrPolicySegmentList *****
type bgpAttributesSrPolicySegmentList struct {
	validation
	obj            *otg.BgpAttributesSrPolicySegmentList
	marshaller     marshalBgpAttributesSrPolicySegmentList
	unMarshaller   unMarshalBgpAttributesSrPolicySegmentList
	weightHolder   BgpAttributesSegmentRoutingPolicySegmentListWeight
	segmentsHolder BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
}

func NewBgpAttributesSrPolicySegmentList() BgpAttributesSrPolicySegmentList {
	obj := bgpAttributesSrPolicySegmentList{obj: &otg.BgpAttributesSrPolicySegmentList{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesSrPolicySegmentList) msg() *otg.BgpAttributesSrPolicySegmentList {
	return obj.obj
}

func (obj *bgpAttributesSrPolicySegmentList) setMsg(msg *otg.BgpAttributesSrPolicySegmentList) BgpAttributesSrPolicySegmentList {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesSrPolicySegmentList struct {
	obj *bgpAttributesSrPolicySegmentList
}

type marshalBgpAttributesSrPolicySegmentList interface {
	// ToProto marshals BgpAttributesSrPolicySegmentList to protobuf object *otg.BgpAttributesSrPolicySegmentList
	ToProto() (*otg.BgpAttributesSrPolicySegmentList, error)
	// ToPbText marshals BgpAttributesSrPolicySegmentList to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesSrPolicySegmentList to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesSrPolicySegmentList to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesSrPolicySegmentList struct {
	obj *bgpAttributesSrPolicySegmentList
}

type unMarshalBgpAttributesSrPolicySegmentList interface {
	// FromProto unmarshals BgpAttributesSrPolicySegmentList from protobuf object *otg.BgpAttributesSrPolicySegmentList
	FromProto(msg *otg.BgpAttributesSrPolicySegmentList) (BgpAttributesSrPolicySegmentList, error)
	// FromPbText unmarshals BgpAttributesSrPolicySegmentList from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesSrPolicySegmentList from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesSrPolicySegmentList from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesSrPolicySegmentList) Marshal() marshalBgpAttributesSrPolicySegmentList {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesSrPolicySegmentList{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesSrPolicySegmentList) Unmarshal() unMarshalBgpAttributesSrPolicySegmentList {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesSrPolicySegmentList{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesSrPolicySegmentList) ToProto() (*otg.BgpAttributesSrPolicySegmentList, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesSrPolicySegmentList) FromProto(msg *otg.BgpAttributesSrPolicySegmentList) (BgpAttributesSrPolicySegmentList, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesSrPolicySegmentList) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicySegmentList) FromPbText(value string) error {
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

func (m *marshalbgpAttributesSrPolicySegmentList) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicySegmentList) FromYaml(value string) error {
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

func (m *marshalbgpAttributesSrPolicySegmentList) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesSrPolicySegmentList) FromJson(value string) error {
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

func (obj *bgpAttributesSrPolicySegmentList) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicySegmentList) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesSrPolicySegmentList) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesSrPolicySegmentList) Clone() (BgpAttributesSrPolicySegmentList, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesSrPolicySegmentList()
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

func (obj *bgpAttributesSrPolicySegmentList) setNil() {
	obj.weightHolder = nil
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesSrPolicySegmentList is one optional SEGMENT_LIST sub-tlv encoded with type of 128.
// One sub-tlv (Type 128) encodes a single explicit path towards the endpoint as described in
// section 5.1 of [RFC9256].
// The Segment List sub-TLV includes the elements of the paths (i.e., segments) as well
// as an optional Weight sub-TLV.
type BgpAttributesSrPolicySegmentList interface {
	Validation
	// msg marshals BgpAttributesSrPolicySegmentList to protobuf object *otg.BgpAttributesSrPolicySegmentList
	// and doesn't set defaults
	msg() *otg.BgpAttributesSrPolicySegmentList
	// setMsg unmarshals BgpAttributesSrPolicySegmentList from protobuf object *otg.BgpAttributesSrPolicySegmentList
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesSrPolicySegmentList) BgpAttributesSrPolicySegmentList
	// provides marshal interface
	Marshal() marshalBgpAttributesSrPolicySegmentList
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesSrPolicySegmentList
	// validate validates BgpAttributesSrPolicySegmentList
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesSrPolicySegmentList, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Weight returns BgpAttributesSegmentRoutingPolicySegmentListWeight, set in BgpAttributesSrPolicySegmentList.
	// BgpAttributesSegmentRoutingPolicySegmentListWeight is the optional Weight sub-TLV (Type 9) specifies the weight associated with a given segment list. The weight is used for weighted multipath.
	Weight() BgpAttributesSegmentRoutingPolicySegmentListWeight
	// SetWeight assigns BgpAttributesSegmentRoutingPolicySegmentListWeight provided by user to BgpAttributesSrPolicySegmentList.
	// BgpAttributesSegmentRoutingPolicySegmentListWeight is the optional Weight sub-TLV (Type 9) specifies the weight associated with a given segment list. The weight is used for weighted multipath.
	SetWeight(value BgpAttributesSegmentRoutingPolicySegmentListWeight) BgpAttributesSrPolicySegmentList
	// HasWeight checks if Weight has been set in BgpAttributesSrPolicySegmentList
	HasWeight() bool
	// Segments returns BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIterIter, set in BgpAttributesSrPolicySegmentList
	Segments() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	setNil()
}

// description is TBD
// Weight returns a BgpAttributesSegmentRoutingPolicySegmentListWeight
func (obj *bgpAttributesSrPolicySegmentList) Weight() BgpAttributesSegmentRoutingPolicySegmentListWeight {
	if obj.obj.Weight == nil {
		obj.obj.Weight = NewBgpAttributesSegmentRoutingPolicySegmentListWeight().msg()
	}
	if obj.weightHolder == nil {
		obj.weightHolder = &bgpAttributesSegmentRoutingPolicySegmentListWeight{obj: obj.obj.Weight}
	}
	return obj.weightHolder
}

// description is TBD
// Weight returns a BgpAttributesSegmentRoutingPolicySegmentListWeight
func (obj *bgpAttributesSrPolicySegmentList) HasWeight() bool {
	return obj.obj.Weight != nil
}

// description is TBD
// SetWeight sets the BgpAttributesSegmentRoutingPolicySegmentListWeight value in the BgpAttributesSrPolicySegmentList object
func (obj *bgpAttributesSrPolicySegmentList) SetWeight(value BgpAttributesSegmentRoutingPolicySegmentListWeight) BgpAttributesSrPolicySegmentList {

	obj.weightHolder = nil
	obj.obj.Weight = value.msg()

	return obj
}

// description is TBD
// Segments returns a []BgpAttributesSegmentRoutingPolicySegmentListSegment
func (obj *bgpAttributesSrPolicySegmentList) Segments() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter struct {
	obj                                                      *bgpAttributesSrPolicySegmentList
	bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice []BgpAttributesSegmentRoutingPolicySegmentListSegment
	fieldPtr                                                 *[]*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment
}

func newBgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter(ptr *[]*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	return &bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter{fieldPtr: ptr}
}

type BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter interface {
	setMsg(*bgpAttributesSrPolicySegmentList) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	Items() []BgpAttributesSegmentRoutingPolicySegmentListSegment
	Add() BgpAttributesSegmentRoutingPolicySegmentListSegment
	Append(items ...BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	Set(index int, newObj BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	Clear() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	clearHolderSlice() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
	appendHolderSlice(item BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter
}

func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) setMsg(msg *bgpAttributesSrPolicySegmentList) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesSegmentRoutingPolicySegmentListSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) Items() []BgpAttributesSegmentRoutingPolicySegmentListSegment {
	return obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice
}

func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) Add() BgpAttributesSegmentRoutingPolicySegmentListSegment {
	newObj := &otg.BgpAttributesSegmentRoutingPolicySegmentListSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesSegmentRoutingPolicySegmentListSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice = append(obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) Append(items ...BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice = append(obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice, item)
	}
	return obj
}

func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) Set(index int, newObj BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) Clear() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesSegmentRoutingPolicySegmentListSegment{}
		obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice = []BgpAttributesSegmentRoutingPolicySegmentListSegment{}
	}
	return obj
}
func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) clearHolderSlice() BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	if len(obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice) > 0 {
		obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice = []BgpAttributesSegmentRoutingPolicySegmentListSegment{}
	}
	return obj
}
func (obj *bgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter) appendHolderSlice(item BgpAttributesSegmentRoutingPolicySegmentListSegment) BgpAttributesSrPolicySegmentListBgpAttributesSegmentRoutingPolicySegmentListSegmentIter {
	obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice = append(obj.bgpAttributesSegmentRoutingPolicySegmentListSegmentSlice, item)
	return obj
}

func (obj *bgpAttributesSrPolicySegmentList) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Weight != nil {

		obj.Weight().validateObj(vObj, set_default)
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&bgpAttributesSegmentRoutingPolicySegmentListSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpAttributesSrPolicySegmentList) setDefault() {

}
