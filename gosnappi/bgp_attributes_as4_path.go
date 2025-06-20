package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAttributesAs4Path *****
type bgpAttributesAs4Path struct {
	validation
	obj            *otg.BgpAttributesAs4Path
	marshaller     marshalBgpAttributesAs4Path
	unMarshaller   unMarshalBgpAttributesAs4Path
	segmentsHolder BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
}

func NewBgpAttributesAs4Path() BgpAttributesAs4Path {
	obj := bgpAttributesAs4Path{obj: &otg.BgpAttributesAs4Path{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAttributesAs4Path) msg() *otg.BgpAttributesAs4Path {
	return obj.obj
}

func (obj *bgpAttributesAs4Path) setMsg(msg *otg.BgpAttributesAs4Path) BgpAttributesAs4Path {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAttributesAs4Path struct {
	obj *bgpAttributesAs4Path
}

type marshalBgpAttributesAs4Path interface {
	// ToProto marshals BgpAttributesAs4Path to protobuf object *otg.BgpAttributesAs4Path
	ToProto() (*otg.BgpAttributesAs4Path, error)
	// ToPbText marshals BgpAttributesAs4Path to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAttributesAs4Path to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAttributesAs4Path to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAttributesAs4Path struct {
	obj *bgpAttributesAs4Path
}

type unMarshalBgpAttributesAs4Path interface {
	// FromProto unmarshals BgpAttributesAs4Path from protobuf object *otg.BgpAttributesAs4Path
	FromProto(msg *otg.BgpAttributesAs4Path) (BgpAttributesAs4Path, error)
	// FromPbText unmarshals BgpAttributesAs4Path from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAttributesAs4Path from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAttributesAs4Path from JSON text
	FromJson(value string) error
}

func (obj *bgpAttributesAs4Path) Marshal() marshalBgpAttributesAs4Path {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAttributesAs4Path{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAttributesAs4Path) Unmarshal() unMarshalBgpAttributesAs4Path {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAttributesAs4Path{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAttributesAs4Path) ToProto() (*otg.BgpAttributesAs4Path, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAttributesAs4Path) FromProto(msg *otg.BgpAttributesAs4Path) (BgpAttributesAs4Path, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAttributesAs4Path) ToPbText() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Path) FromPbText(value string) error {
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

func (m *marshalbgpAttributesAs4Path) ToYaml() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Path) FromYaml(value string) error {
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

func (m *marshalbgpAttributesAs4Path) ToJson() (string, error) {
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

func (m *unMarshalbgpAttributesAs4Path) FromJson(value string) error {
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

func (obj *bgpAttributesAs4Path) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAttributesAs4Path) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAttributesAs4Path) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAttributesAs4Path) Clone() (BgpAttributesAs4Path, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAttributesAs4Path()
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

func (obj *bgpAttributesAs4Path) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAttributesAs4Path is the AS4_PATH attribute identifies the autonomous systems through  which routing information
// carried in this UPDATE message has passed.
// This contains the configuration of how to include the Local AS in the AS path
// attribute of the MP REACH NLRI. It also contains optional configuration of
// additional AS Path Segments that can be included in the AS Path attribute.
// The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that
// a routing information passes through to reach the destination.
// AS4_PATH is only exchanged in two scenarios:
// - When an old BGP speaker has to forward a received AS4_PATH containing 4 byte AS numbers to new BGP speaker.
// - When a new BGP speaker is connected to an old BGP speaker and has to propagate 4 byte AS numbers via the old BGP speaker.
// Its usage is described in RFC4893.
type BgpAttributesAs4Path interface {
	Validation
	// msg marshals BgpAttributesAs4Path to protobuf object *otg.BgpAttributesAs4Path
	// and doesn't set defaults
	msg() *otg.BgpAttributesAs4Path
	// setMsg unmarshals BgpAttributesAs4Path from protobuf object *otg.BgpAttributesAs4Path
	// and doesn't set defaults
	setMsg(*otg.BgpAttributesAs4Path) BgpAttributesAs4Path
	// provides marshal interface
	Marshal() marshalBgpAttributesAs4Path
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAttributesAs4Path
	// validate validates BgpAttributesAs4Path
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAttributesAs4Path, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Segments returns BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIterIter, set in BgpAttributesAs4Path
	Segments() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	setNil()
}

// The AS path segments containing 4 byte AS numbers to be added in the AS4_PATH attribute.
// Segments returns a []BgpAttributesFourByteAsPathSegment
func (obj *bgpAttributesAs4Path) Segments() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpAttributesFourByteAsPathSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter struct {
	obj                                     *bgpAttributesAs4Path
	bgpAttributesFourByteAsPathSegmentSlice []BgpAttributesFourByteAsPathSegment
	fieldPtr                                *[]*otg.BgpAttributesFourByteAsPathSegment
}

func newBgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter(ptr *[]*otg.BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	return &bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter{fieldPtr: ptr}
}

type BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter interface {
	setMsg(*bgpAttributesAs4Path) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	Items() []BgpAttributesFourByteAsPathSegment
	Add() BgpAttributesFourByteAsPathSegment
	Append(items ...BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	Set(index int, newObj BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	Clear() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	clearHolderSlice() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
	appendHolderSlice(item BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter
}

func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) setMsg(msg *bgpAttributesAs4Path) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAttributesFourByteAsPathSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) Items() []BgpAttributesFourByteAsPathSegment {
	return obj.bgpAttributesFourByteAsPathSegmentSlice
}

func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) Add() BgpAttributesFourByteAsPathSegment {
	newObj := &otg.BgpAttributesFourByteAsPathSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAttributesFourByteAsPathSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) Append(items ...BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, item)
	}
	return obj
}

func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) Set(index int, newObj BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAttributesFourByteAsPathSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) Clear() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAttributesFourByteAsPathSegment{}
		obj.bgpAttributesFourByteAsPathSegmentSlice = []BgpAttributesFourByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) clearHolderSlice() BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	if len(obj.bgpAttributesFourByteAsPathSegmentSlice) > 0 {
		obj.bgpAttributesFourByteAsPathSegmentSlice = []BgpAttributesFourByteAsPathSegment{}
	}
	return obj
}
func (obj *bgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter) appendHolderSlice(item BgpAttributesFourByteAsPathSegment) BgpAttributesAs4PathBgpAttributesFourByteAsPathSegmentIter {
	obj.bgpAttributesFourByteAsPathSegmentSlice = append(obj.bgpAttributesFourByteAsPathSegmentSlice, item)
	return obj
}

func (obj *bgpAttributesAs4Path) validateObj(vObj *validation, set_default bool) {
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

func (obj *bgpAttributesAs4Path) setDefault() {

}
