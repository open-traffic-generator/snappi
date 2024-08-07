package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpAsPath *****
type bgpAsPath struct {
	validation
	obj            *otg.BgpAsPath
	marshaller     marshalBgpAsPath
	unMarshaller   unMarshalBgpAsPath
	segmentsHolder BgpAsPathBgpAsPathSegmentIter
}

func NewBgpAsPath() BgpAsPath {
	obj := bgpAsPath{obj: &otg.BgpAsPath{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpAsPath) msg() *otg.BgpAsPath {
	return obj.obj
}

func (obj *bgpAsPath) setMsg(msg *otg.BgpAsPath) BgpAsPath {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpAsPath struct {
	obj *bgpAsPath
}

type marshalBgpAsPath interface {
	// ToProto marshals BgpAsPath to protobuf object *otg.BgpAsPath
	ToProto() (*otg.BgpAsPath, error)
	// ToPbText marshals BgpAsPath to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpAsPath to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpAsPath to JSON text
	ToJson() (string, error)
}

type unMarshalbgpAsPath struct {
	obj *bgpAsPath
}

type unMarshalBgpAsPath interface {
	// FromProto unmarshals BgpAsPath from protobuf object *otg.BgpAsPath
	FromProto(msg *otg.BgpAsPath) (BgpAsPath, error)
	// FromPbText unmarshals BgpAsPath from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpAsPath from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpAsPath from JSON text
	FromJson(value string) error
}

func (obj *bgpAsPath) Marshal() marshalBgpAsPath {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpAsPath{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpAsPath) Unmarshal() unMarshalBgpAsPath {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpAsPath{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpAsPath) ToProto() (*otg.BgpAsPath, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpAsPath) FromProto(msg *otg.BgpAsPath) (BgpAsPath, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpAsPath) ToPbText() (string, error) {
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

func (m *unMarshalbgpAsPath) FromPbText(value string) error {
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

func (m *marshalbgpAsPath) ToYaml() (string, error) {
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

func (m *unMarshalbgpAsPath) FromYaml(value string) error {
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

func (m *marshalbgpAsPath) ToJson() (string, error) {
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

func (m *unMarshalbgpAsPath) FromJson(value string) error {
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

func (obj *bgpAsPath) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpAsPath) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpAsPath) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpAsPath) Clone() (BgpAsPath, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpAsPath()
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

func (obj *bgpAsPath) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpAsPath is this attribute identifies the autonomous systems through  which routing information carried in this UPDATE message has passed. This contains the configuration of how to include the Local AS in the AS path attribute of the MP REACH NLRI. It also contains optional configuration of additional AS Path Segments that can be included in the AS Path attribute. The AS Path consists of a Set or Sequence of Autonomous Systems (AS) numbers  that a routing information passes through to reach the destination.
type BgpAsPath interface {
	Validation
	// msg marshals BgpAsPath to protobuf object *otg.BgpAsPath
	// and doesn't set defaults
	msg() *otg.BgpAsPath
	// setMsg unmarshals BgpAsPath from protobuf object *otg.BgpAsPath
	// and doesn't set defaults
	setMsg(*otg.BgpAsPath) BgpAsPath
	// provides marshal interface
	Marshal() marshalBgpAsPath
	// provides unmarshal interface
	Unmarshal() unMarshalBgpAsPath
	// validate validates BgpAsPath
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpAsPath, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// AsSetMode returns BgpAsPathAsSetModeEnum, set in BgpAsPath
	AsSetMode() BgpAsPathAsSetModeEnum
	// SetAsSetMode assigns BgpAsPathAsSetModeEnum provided by user to BgpAsPath
	SetAsSetMode(value BgpAsPathAsSetModeEnum) BgpAsPath
	// HasAsSetMode checks if AsSetMode has been set in BgpAsPath
	HasAsSetMode() bool
	// Segments returns BgpAsPathBgpAsPathSegmentIterIter, set in BgpAsPath
	Segments() BgpAsPathBgpAsPathSegmentIter
	setNil()
}

type BgpAsPathAsSetModeEnum string

// Enum of AsSetMode on BgpAsPath
var BgpAsPathAsSetMode = struct {
	DO_NOT_INCLUDE_LOCAL_AS  BgpAsPathAsSetModeEnum
	INCLUDE_AS_SEQ           BgpAsPathAsSetModeEnum
	INCLUDE_AS_SET           BgpAsPathAsSetModeEnum
	INCLUDE_AS_CONFED_SEQ    BgpAsPathAsSetModeEnum
	INCLUDE_AS_CONFED_SET    BgpAsPathAsSetModeEnum
	PREPEND_TO_FIRST_SEGMENT BgpAsPathAsSetModeEnum
}{
	DO_NOT_INCLUDE_LOCAL_AS:  BgpAsPathAsSetModeEnum("do_not_include_local_as"),
	INCLUDE_AS_SEQ:           BgpAsPathAsSetModeEnum("include_as_seq"),
	INCLUDE_AS_SET:           BgpAsPathAsSetModeEnum("include_as_set"),
	INCLUDE_AS_CONFED_SEQ:    BgpAsPathAsSetModeEnum("include_as_confed_seq"),
	INCLUDE_AS_CONFED_SET:    BgpAsPathAsSetModeEnum("include_as_confed_set"),
	PREPEND_TO_FIRST_SEGMENT: BgpAsPathAsSetModeEnum("prepend_to_first_segment"),
}

func (obj *bgpAsPath) AsSetMode() BgpAsPathAsSetModeEnum {
	return BgpAsPathAsSetModeEnum(obj.obj.AsSetMode.Enum().String())
}

// Defines how the Local AS should be included in the MP REACH NLRI. For iBGP sessions, "Do Not Include Local AS" must be chosen. For eBGP sessions, any choice other than "Do Not Include Local AS" can be chosen.
// AsSetMode returns a string
func (obj *bgpAsPath) HasAsSetMode() bool {
	return obj.obj.AsSetMode != nil
}

func (obj *bgpAsPath) SetAsSetMode(value BgpAsPathAsSetModeEnum) BgpAsPath {
	intValue, ok := otg.BgpAsPath_AsSetMode_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on BgpAsPathAsSetModeEnum", string(value)))
		return obj
	}
	enumValue := otg.BgpAsPath_AsSetMode_Enum(intValue)
	obj.obj.AsSetMode = &enumValue

	return obj
}

// The additional AS path segments to be added in the NLRI.  By default, an empty AS path is always included and  the local AS is added to it as per the value of 'as_set_mode' attribute.
// Segments returns a []BgpAsPathSegment
func (obj *bgpAsPath) Segments() BgpAsPathBgpAsPathSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpAsPathSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpAsPathBgpAsPathSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpAsPathBgpAsPathSegmentIter struct {
	obj                   *bgpAsPath
	bgpAsPathSegmentSlice []BgpAsPathSegment
	fieldPtr              *[]*otg.BgpAsPathSegment
}

func newBgpAsPathBgpAsPathSegmentIter(ptr *[]*otg.BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter {
	return &bgpAsPathBgpAsPathSegmentIter{fieldPtr: ptr}
}

type BgpAsPathBgpAsPathSegmentIter interface {
	setMsg(*bgpAsPath) BgpAsPathBgpAsPathSegmentIter
	Items() []BgpAsPathSegment
	Add() BgpAsPathSegment
	Append(items ...BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter
	Set(index int, newObj BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter
	Clear() BgpAsPathBgpAsPathSegmentIter
	clearHolderSlice() BgpAsPathBgpAsPathSegmentIter
	appendHolderSlice(item BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter
}

func (obj *bgpAsPathBgpAsPathSegmentIter) setMsg(msg *bgpAsPath) BgpAsPathBgpAsPathSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpAsPathSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpAsPathBgpAsPathSegmentIter) Items() []BgpAsPathSegment {
	return obj.bgpAsPathSegmentSlice
}

func (obj *bgpAsPathBgpAsPathSegmentIter) Add() BgpAsPathSegment {
	newObj := &otg.BgpAsPathSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpAsPathSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpAsPathSegmentSlice = append(obj.bgpAsPathSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpAsPathBgpAsPathSegmentIter) Append(items ...BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpAsPathSegmentSlice = append(obj.bgpAsPathSegmentSlice, item)
	}
	return obj
}

func (obj *bgpAsPathBgpAsPathSegmentIter) Set(index int, newObj BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpAsPathSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpAsPathBgpAsPathSegmentIter) Clear() BgpAsPathBgpAsPathSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpAsPathSegment{}
		obj.bgpAsPathSegmentSlice = []BgpAsPathSegment{}
	}
	return obj
}
func (obj *bgpAsPathBgpAsPathSegmentIter) clearHolderSlice() BgpAsPathBgpAsPathSegmentIter {
	if len(obj.bgpAsPathSegmentSlice) > 0 {
		obj.bgpAsPathSegmentSlice = []BgpAsPathSegment{}
	}
	return obj
}
func (obj *bgpAsPathBgpAsPathSegmentIter) appendHolderSlice(item BgpAsPathSegment) BgpAsPathBgpAsPathSegmentIter {
	obj.bgpAsPathSegmentSlice = append(obj.bgpAsPathSegmentSlice, item)
	return obj
}

func (obj *bgpAsPath) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&bgpAsPathSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpAsPath) setDefault() {
	if obj.obj.AsSetMode == nil {
		obj.SetAsSetMode(BgpAsPathAsSetMode.DO_NOT_INCLUDE_LOCAL_AS)

	}

}
