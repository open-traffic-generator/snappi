package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAsPathFourByteAsPath *****
type bgpAttributesAsPathFourByteAsPath struct {
	validation
	obj            *otg.BgpAttributesAsPathFourByteAsPath
	marshaller     marshalBgpAttributesAsPathFourByteAsPath
	unMarshaller   unMarshalBgpAttributesAsPathFourByteAsPath
	segmentsHolder BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
}

func NewBgpAttributesAsPathFourByteAsPath() BgpAttributesAsPathFourByteAsPath {
	obj := bgpAttributesAsPathFourByteAsPath{obj: &otg.BgpAttributesAsPathFourByteAsPath{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAsPathFourByteAsPath) msg() *otg.BgpAttributesAsPathFourByteAsPath {
	return obj.obj
}

func (obj *bgpAttributesAsPathFourByteAsPath) setMsg(msg *otg.BgpAttributesAsPathFourByteAsPath) BgpAttributesAsPathFourByteAsPath {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAsPathFourByteAsPath struct {
	obj *bgpAttributesAsPathFourByteAsPath
}

type marshalBgpAttributesAsPathFourByteAsPath interface {
	// ToProto marshals BgpAttributesAsPathFourByteAsPath to protobuf object *otg.BgpAttributesAsPathFourByteAsPath
	ToProto() (*otg.BgpAttributesAsPathFourByteAsPath, error)
	// ToPbText marshals BgpAttributesAsPathFourByteAsPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAsPathFourByteAsPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAsPathFourByteAsPath to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesAsPathFourByteAsPath struct {
	obj *bgpAttributesAsPathFourByteAsPath
}

type unMarshalBgpAttributesAsPathFourByteAsPath interface {
	// FromProto unmarshals BgpAttributesAsPathFourByteAsPath from protobuf object *otg.BgpAttributesAsPathFourByteAsPath
	FromProto(msg *otg.BgpAttributesAsPathFourByteAsPath) (BgpAttributesAsPathFourByteAsPath, error)
	// FromPbText unmarshals BgpAttributesAsPathFourByteAsPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAsPathFourByteAsPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAsPathFourByteAsPath from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAsPathFourByteAsPath) Marshal() marshalBgpAttributesAsPathFourByteAsPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAsPathFourByteAsPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAsPathFourByteAsPath) Unmarshal() unMarshalBgpAttributesAsPathFourByteAsPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAsPathFourByteAsPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAsPathFourByteAsPath) ToProto() (*otg.BgpAttributesAsPathFourByteAsPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAsPathFourByteAsPath) FromProto(msg *otg.BgpAttributesAsPathFourByteAsPath) (BgpAttributesAsPathFourByteAsPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAsPathFourByteAsPath) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathFourByteAsPath) FromPbText(value string) error {
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

func (m *marshalbgpAttributesAsPathFourByteAsPath) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathFourByteAsPath) FromYaml(value string) error {
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

func (m *marshalbgpAttributesAsPathFourByteAsPath) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAsPathFourByteAsPath) FromJson(value string) error {
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

func (obj *bgpAttributesAsPathFourByteAsPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPathFourByteAsPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAsPathFourByteAsPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAsPathFourByteAsPath) Clone() (BgpAttributesAsPathFourByteAsPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAsPathFourByteAsPath()
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

func (obj *bgpAttributesAsPathFourByteAsPath) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesAsPathFourByteAsPath is aS Paths with 4 byte AS numbers can be exchanged only if both BGP speakers support 4 byte AS number extensions.
type BgpAttributesAsPathFourByteAsPath interface {
	Validation
	// msg marshals BgpAttributesAsPathFourByteAsPath to protobuf object *otg.BgpAttributesAsPathFourByteAsPath
	// and doesn't set defaults
	msg() *otg.BgpAttributesAsPathFourByteAsPath
	// setMsg unmarshals BgpAttributesAsPathFourByteAsPath from protobuf object *otg.BgpAttributesAsPathFourByteAsPath
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAsPathFourByteAsPath) BgpAttributesAsPathFourByteAsPath
	// provides marshal interface
	Marshal() marshalBgpAttributesAsPathFourByteAsPath
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAsPathFourByteAsPath
	// validate validates BgpAttributesAsPathFourByteAsPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAsPathFourByteAsPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segments returns BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIterIter, set in BgpAttributesAsPathFourByteAsPath
	Segments() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	setNil()
}

// The AS path segments containing 4 byte AS numbers to be added in the AS Path attribute.  By default, an empty AS path should always be included and for EBGP at minimum the local AS number should be present in the AS Path.
// Segments returns a []BgpAttributesFourByteAsPathSegment
func (obj *bgpAttributesAsPathFourByteAsPath) Segments() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpAttributesFourByteAsPathSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter struct {
	obj                                     *bgpAttributesAsPathFourByteAsPath
	bgpAttributesFourByteAsPathSegmentSlice []BgpAttributesFourByteAsPathSegment
	fieldPtr                                *[]*otg.BgpAttributesFourByteAsPathSegment
}

func newBgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter(ptr *[]*otg.BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	return &bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter{fieldPtr: ptr}
}

type BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter interface {
	setMsg(*bgpAttributesAsPathFourByteAsPath) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	Items() []BgpAttributesFourByteAsPathSegment
	Add() BgpAttributesFourByteAsPathSegment
	Append(items ...BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	Set(index int, newObj BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	Clear() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	clearHolderSlice() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
	appendHolderSlice(item BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter
}

func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) setMsg(msg *bgpAttributesAsPathFourByteAsPath) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesFourByteAsPathSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) Items() []BgpAttributesFourByteAsPathSegment {
	return obj.bgpAttributesFourByteAsPathSegmentSlice
}

func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) Add() BgpAttributesFourByteAsPathSegment {
	newObj := &otg.BgpAttributesFourByteAsPathSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesFourByteAsPathSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) Append(items ...BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, item)
	}
	return obj
}

func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) Set(index int, newObj BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesFourByteAsPathSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) Clear() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesFourByteAsPathSegment{}
		obj.bgpAttributesFourByteAsPathSegmentSlice = []BgpAttributesFourByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) clearHolderSlice() BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	if len(obj.bgpAttributesFourByteAsPathSegmentSlice) > 0 {
		obj.bgpAttributesFourByteAsPathSegmentSlice = []BgpAttributesFourByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter) appendHolderSlice(item BgpAttributesFourByteAsPathSegment) BgpAttributesAsPathFourByteAsPathBgpAttributesFourByteAsPathSegmentIter {
	obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, item)
	return obj
}

func (obj *bgpAttributesAsPathFourByteAsPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&bgpAttributesFourByteAsPathSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpAttributesAsPathFourByteAsPath) setDefault() {

}
