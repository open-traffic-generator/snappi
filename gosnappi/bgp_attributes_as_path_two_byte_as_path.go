package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAsPathTwoByteAsPath *****
type bgpAttributesAsPathTwoByteAsPath struct {
	validation
	obj            *otg.BgpAttributesAsPathTwoByteAsPath
	marshaller     marshalBgpAttributesAsPathTwoByteAsPath
	unMarshaller   unMarshalBgpAttributesAsPathTwoByteAsPath
	segmentsHolder BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
}

func NewBgpAttributesAsPathTwoByteAsPath() BgpAttributesAsPathTwoByteAsPath {
	obj := bgpAttributesAsPathTwoByteAsPath{obj: &otg.BgpAttributesAsPathTwoByteAsPath{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAsPathTwoByteAsPath) msg() *otg.BgpAttributesAsPathTwoByteAsPath {
	return obj.obj
}

func (obj *bgpAttributesAsPathTwoByteAsPath) setMsg(msg *otg.BgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPathTwoByteAsPath {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAsPathTwoByteAsPath struct {
	obj *bgpAttributesAsPathTwoByteAsPath
}

type marshalBgpAttributesAsPathTwoByteAsPath interface {
	// ToProto marshals BgpAttributesAsPathTwoByteAsPath to protobuf object *otg.BgpAttributesAsPathTwoByteAsPath
	ToProto() (*otg.BgpAttributesAsPathTwoByteAsPath, error)
	// ToPbText marshals BgpAttributesAsPathTwoByteAsPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAsPathTwoByteAsPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAsPathTwoByteAsPath to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpAttributesAsPathTwoByteAsPath to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpAttributesAsPathTwoByteAsPath struct {
	obj *bgpAttributesAsPathTwoByteAsPath
}

type unMarshalBgpAttributesAsPathTwoByteAsPath interface {
	// FromProto unmarshals BgpAttributesAsPathTwoByteAsPath from protobuf object *otg.BgpAttributesAsPathTwoByteAsPath
	FromProto(msg *otg.BgpAttributesAsPathTwoByteAsPath) (BgpAttributesAsPathTwoByteAsPath, error)
	// FromPbText unmarshals BgpAttributesAsPathTwoByteAsPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAsPathTwoByteAsPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAsPathTwoByteAsPath from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAsPathTwoByteAsPath) Marshal() marshalBgpAttributesAsPathTwoByteAsPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAsPathTwoByteAsPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAsPathTwoByteAsPath) Unmarshal() unMarshalBgpAttributesAsPathTwoByteAsPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAsPathTwoByteAsPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAsPathTwoByteAsPath) ToProto() (*otg.BgpAttributesAsPathTwoByteAsPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAsPathTwoByteAsPath) FromProto(msg *otg.BgpAttributesAsPathTwoByteAsPath) (BgpAttributesAsPathTwoByteAsPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAsPathTwoByteAsPath) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathTwoByteAsPath) FromPbText(value string) error {
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

func (m *marshalbgpAttributesAsPathTwoByteAsPath) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathTwoByteAsPath) FromYaml(value string) error {
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

func (m *marshalbgpAttributesAsPathTwoByteAsPath) ToJsonRaw() (string, error) {
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

func (m *marshalbgpAttributesAsPathTwoByteAsPath) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathTwoByteAsPath) FromJson(value string) error {
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

func (obj *bgpAttributesAsPathTwoByteAsPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPathTwoByteAsPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPathTwoByteAsPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAsPathTwoByteAsPath) Clone() (BgpAttributesAsPathTwoByteAsPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAsPathTwoByteAsPath()
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

func (obj *bgpAttributesAsPathTwoByteAsPath) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesAsPathTwoByteAsPath is aS Paths with 2 byte AS numbers is used when any of the two scenarios occur :
// - An old BGP speaker and new BGP speaker are sending BGP Updates to one another.
// - Two old BGP speakers are sending BGP Updates to one another.
type BgpAttributesAsPathTwoByteAsPath interface {
	Validation
	// msg marshals BgpAttributesAsPathTwoByteAsPath to protobuf object *otg.BgpAttributesAsPathTwoByteAsPath
	// and doesn't set defaults
	msg() *otg.BgpAttributesAsPathTwoByteAsPath
	// setMsg unmarshals BgpAttributesAsPathTwoByteAsPath from protobuf object *otg.BgpAttributesAsPathTwoByteAsPath
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPathTwoByteAsPath
	// provides marshal interface
	Marshal() marshalBgpAttributesAsPathTwoByteAsPath
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAsPathTwoByteAsPath
	// validate validates BgpAttributesAsPathTwoByteAsPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAsPathTwoByteAsPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segments returns BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIterIter, set in BgpAttributesAsPathTwoByteAsPath
	Segments() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	setNil()
}

// The AS path segments containing 2 byte AS numbers to be added in the AS Path attribute.  By default, an empty AS path should always be included and for EBGP the sender's AS number should be prepended to the AS Path.
// Segments returns a []BgpAttributesTwoByteAsPathSegment
func (obj *bgpAttributesAsPathTwoByteAsPath) Segments() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpAttributesTwoByteAsPathSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter struct {
	obj                                    *bgpAttributesAsPathTwoByteAsPath
	bgpAttributesTwoByteAsPathSegmentSlice []BgpAttributesTwoByteAsPathSegment
	fieldPtr                               *[]*otg.BgpAttributesTwoByteAsPathSegment
}

func newBgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter(ptr *[]*otg.BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	return &bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter{fieldPtr: ptr}
}

type BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter interface {
	setMsg(*bgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	Items() []BgpAttributesTwoByteAsPathSegment
	Add() BgpAttributesTwoByteAsPathSegment
	Append(items ...BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	Set(index int, newObj BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	Clear() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	clearHolderSlice() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
	appendHolderSlice(item BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter
}

func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) setMsg(msg *bgpAttributesAsPathTwoByteAsPath) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesTwoByteAsPathSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) Items() []BgpAttributesTwoByteAsPathSegment {
	return obj.bgpAttributesTwoByteAsPathSegmentSlice
}

func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) Add() BgpAttributesTwoByteAsPathSegment {
	newObj := &otg.BgpAttributesTwoByteAsPathSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesTwoByteAsPathSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesTwoByteAsPathSegmentSlice = append(obj.bgpAttributesTwoByteAsPathSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) Append(items ...BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesTwoByteAsPathSegmentSlice = append(obj.bgpAttributesTwoByteAsPathSegmentSlice, item)
	}
	return obj
}

func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) Set(index int, newObj BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesTwoByteAsPathSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) Clear() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesTwoByteAsPathSegment{}
		obj.bgpAttributesTwoByteAsPathSegmentSlice = []BgpAttributesTwoByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) clearHolderSlice() BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	if len(obj.bgpAttributesTwoByteAsPathSegmentSlice) > 0 {
		obj.bgpAttributesTwoByteAsPathSegmentSlice = []BgpAttributesTwoByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter) appendHolderSlice(item BgpAttributesTwoByteAsPathSegment) BgpAttributesAsPathTwoByteAsPathBgpAttributesTwoByteAsPathSegmentIter {
	obj.bgpAttributesTwoByteAsPathSegmentSlice = append(obj.bgpAttributesTwoByteAsPathSegmentSlice, item)
	return obj
}

func (obj *bgpAttributesAsPathTwoByteAsPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&bgpAttributesTwoByteAsPathSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpAttributesAsPathTwoByteAsPath) setDefault() {

}
