package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ResultBgpAsPath *****
type resultBgpAsPath struct {
	validation
	obj            *otg.ResultBgpAsPath
	marshaller     marshalResultBgpAsPath
	unMarshaller   unMarshalResultBgpAsPath
	segmentsHolder ResultBgpAsPathResultBgpAsPathSegmentIter
}

func NewResultBgpAsPath() ResultBgpAsPath {
	obj := resultBgpAsPath{obj: &otg.ResultBgpAsPath{}}
	obj.setDefault()
	return &obj
}

func (obj *resultBgpAsPath) msg() *otg.ResultBgpAsPath {
	return obj.obj
}

func (obj *resultBgpAsPath) setMsg(msg *otg.ResultBgpAsPath) ResultBgpAsPath {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalresultBgpAsPath struct {
	obj *resultBgpAsPath
}

type marshalResultBgpAsPath interface {
	// ToProto marshals ResultBgpAsPath to protobuf object *otg.ResultBgpAsPath
	ToProto() (*otg.ResultBgpAsPath, error)
	// ToPbText marshals ResultBgpAsPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ResultBgpAsPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals ResultBgpAsPath to JSON text
	ToJson() (string, error)
}

type unMarshalresultBgpAsPath struct {
	obj *resultBgpAsPath
}

type unMarshalResultBgpAsPath interface {
	// FromProto unmarshals ResultBgpAsPath from protobuf object *otg.ResultBgpAsPath
	FromProto(msg *otg.ResultBgpAsPath) (ResultBgpAsPath, error)
	// FromPbText unmarshals ResultBgpAsPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ResultBgpAsPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ResultBgpAsPath from JSON text
	FromJson(value string) error
}

func (obj *resultBgpAsPath) Marshal() marshalResultBgpAsPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalresultBgpAsPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *resultBgpAsPath) Unmarshal() unMarshalResultBgpAsPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalresultBgpAsPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalresultBgpAsPath) ToProto() (*otg.ResultBgpAsPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalresultBgpAsPath) FromProto(msg *otg.ResultBgpAsPath) (ResultBgpAsPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalresultBgpAsPath) ToPbText() (string, error) {
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

func (m *unMarshalresultBgpAsPath) FromPbText(value string) error {
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

func (m *marshalresultBgpAsPath) ToYaml() (string, error) {
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

func (m *unMarshalresultBgpAsPath) FromYaml(value string) error {
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

func (m *marshalresultBgpAsPath) ToJson() (string, error) {
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

func (m *unMarshalresultBgpAsPath) FromJson(value string) error {
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

func (obj *resultBgpAsPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *resultBgpAsPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *resultBgpAsPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *resultBgpAsPath) Clone() (ResultBgpAsPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewResultBgpAsPath()
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

func (obj *resultBgpAsPath) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ResultBgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed.
type ResultBgpAsPath interface {
	Validation
	// msg marshals ResultBgpAsPath to protobuf object *otg.ResultBgpAsPath
	// and doesn't set defaults
	msg() *otg.ResultBgpAsPath
	// setMsg unmarshals ResultBgpAsPath from protobuf object *otg.ResultBgpAsPath
	// and doesn't set defaults
	setMsg(*otg.ResultBgpAsPath) ResultBgpAsPath
	// provides marshal interface
	Marshal() marshalResultBgpAsPath
	// provides unmarshal interface
	Unmarshal() unMarshalResultBgpAsPath
	// validate validates ResultBgpAsPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ResultBgpAsPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segments returns ResultBgpAsPathResultBgpAsPathSegmentIterIter, set in ResultBgpAsPath
	Segments() ResultBgpAsPathResultBgpAsPathSegmentIter
	setNil()
}

// AS Path segments present in the received AS Path attribute.
// Segments returns a []ResultBgpAsPathSegment
func (obj *resultBgpAsPath) Segments() ResultBgpAsPathResultBgpAsPathSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.ResultBgpAsPathSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newResultBgpAsPathResultBgpAsPathSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type resultBgpAsPathResultBgpAsPathSegmentIter struct {
	obj                         *resultBgpAsPath
	resultBgpAsPathSegmentSlice []ResultBgpAsPathSegment
	fieldPtr                    *[]*otg.ResultBgpAsPathSegment
}

func newResultBgpAsPathResultBgpAsPathSegmentIter(ptr *[]*otg.ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter {
	return &resultBgpAsPathResultBgpAsPathSegmentIter{fieldPtr: ptr}
}

type ResultBgpAsPathResultBgpAsPathSegmentIter interface {
	setMsg(*resultBgpAsPath) ResultBgpAsPathResultBgpAsPathSegmentIter
	Items() []ResultBgpAsPathSegment
	Add() ResultBgpAsPathSegment
	Append(items ...ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter
	Set(index int, newObj ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter
	Clear() ResultBgpAsPathResultBgpAsPathSegmentIter
	clearHolderSlice() ResultBgpAsPathResultBgpAsPathSegmentIter
	appendHolderSlice(item ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter
}

func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) setMsg(msg *resultBgpAsPath) ResultBgpAsPathResultBgpAsPathSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&resultBgpAsPathSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) Items() []ResultBgpAsPathSegment {
	return obj.resultBgpAsPathSegmentSlice
}

func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) Add() ResultBgpAsPathSegment {
	newObj := &otg.ResultBgpAsPathSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &resultBgpAsPathSegment{obj: newObj}
	newLibObj.setDefault()
	obj.resultBgpAsPathSegmentSlice = append(obj.resultBgpAsPathSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) Append(items ...ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.resultBgpAsPathSegmentSlice = append(obj.resultBgpAsPathSegmentSlice, item)
	}
	return obj
}

func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) Set(index int, newObj ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.resultBgpAsPathSegmentSlice[index] = newObj
	return obj
}
func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) Clear() ResultBgpAsPathResultBgpAsPathSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.ResultBgpAsPathSegment{}
		obj.resultBgpAsPathSegmentSlice = []ResultBgpAsPathSegment{}
	}
	return obj
}
func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) clearHolderSlice() ResultBgpAsPathResultBgpAsPathSegmentIter {
	if len(obj.resultBgpAsPathSegmentSlice) > 0 {
		obj.resultBgpAsPathSegmentSlice = []ResultBgpAsPathSegment{}
	}
	return obj
}
func (obj *resultBgpAsPathResultBgpAsPathSegmentIter) appendHolderSlice(item ResultBgpAsPathSegment) ResultBgpAsPathResultBgpAsPathSegmentIter {
	obj.resultBgpAsPathSegmentSlice = append(obj.resultBgpAsPathSegmentSlice, item)
	return obj
}

func (obj *resultBgpAsPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&resultBgpAsPathSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *resultBgpAsPath) setDefault() {

}
