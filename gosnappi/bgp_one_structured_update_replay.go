package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpOneStructuredUpdateReplay *****
type bgpOneStructuredUpdateReplay struct {
	validation
	obj                           *otg.BgpOneStructuredUpdateReplay
	marshaller                    marshalBgpOneStructuredUpdateReplay
	unMarshaller                  unMarshalBgpOneStructuredUpdateReplay
	pathAttributesHolder          BgpAttributes
	traditionalUnreachNlrisHolder BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	traditionalReachNlrisHolder   BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
}

func NewBgpOneStructuredUpdateReplay() BgpOneStructuredUpdateReplay {
	obj := bgpOneStructuredUpdateReplay{obj: &otg.BgpOneStructuredUpdateReplay{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpOneStructuredUpdateReplay) msg() *otg.BgpOneStructuredUpdateReplay {
	return obj.obj
}

func (obj *bgpOneStructuredUpdateReplay) setMsg(msg *otg.BgpOneStructuredUpdateReplay) BgpOneStructuredUpdateReplay {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpOneStructuredUpdateReplay struct {
	obj *bgpOneStructuredUpdateReplay
}

type marshalBgpOneStructuredUpdateReplay interface {
	// ToProto marshals BgpOneStructuredUpdateReplay to protobuf object *otg.BgpOneStructuredUpdateReplay
	ToProto() (*otg.BgpOneStructuredUpdateReplay, error)
	// ToPbText marshals BgpOneStructuredUpdateReplay to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpOneStructuredUpdateReplay to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpOneStructuredUpdateReplay to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpOneStructuredUpdateReplay to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpOneStructuredUpdateReplay struct {
	obj *bgpOneStructuredUpdateReplay
}

type unMarshalBgpOneStructuredUpdateReplay interface {
	// FromProto unmarshals BgpOneStructuredUpdateReplay from protobuf object *otg.BgpOneStructuredUpdateReplay
	FromProto(msg *otg.BgpOneStructuredUpdateReplay) (BgpOneStructuredUpdateReplay, error)
	// FromPbText unmarshals BgpOneStructuredUpdateReplay from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpOneStructuredUpdateReplay from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpOneStructuredUpdateReplay from JSON text
	FromJson(value string) error
}

func (obj *bgpOneStructuredUpdateReplay) Marshal() marshalBgpOneStructuredUpdateReplay {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpOneStructuredUpdateReplay{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpOneStructuredUpdateReplay) Unmarshal() unMarshalBgpOneStructuredUpdateReplay {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpOneStructuredUpdateReplay{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpOneStructuredUpdateReplay) ToProto() (*otg.BgpOneStructuredUpdateReplay, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpOneStructuredUpdateReplay) FromProto(msg *otg.BgpOneStructuredUpdateReplay) (BgpOneStructuredUpdateReplay, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpOneStructuredUpdateReplay) ToPbText() (string, error) {
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

func (m *unMarshalbgpOneStructuredUpdateReplay) FromPbText(value string) error {
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

func (m *marshalbgpOneStructuredUpdateReplay) ToYaml() (string, error) {
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

func (m *unMarshalbgpOneStructuredUpdateReplay) FromYaml(value string) error {
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

func (m *marshalbgpOneStructuredUpdateReplay) ToJsonRaw() (string, error) {
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

func (m *marshalbgpOneStructuredUpdateReplay) ToJson() (string, error) {
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

func (m *unMarshalbgpOneStructuredUpdateReplay) FromJson(value string) error {
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

func (obj *bgpOneStructuredUpdateReplay) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpOneStructuredUpdateReplay) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpOneStructuredUpdateReplay) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpOneStructuredUpdateReplay) Clone() (BgpOneStructuredUpdateReplay, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpOneStructuredUpdateReplay()
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

func (obj *bgpOneStructuredUpdateReplay) setNil() {
	obj.pathAttributesHolder = nil
	obj.traditionalUnreachNlrisHolder = nil
	obj.traditionalReachNlrisHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpOneStructuredUpdateReplay is one structured BGP Update.
type BgpOneStructuredUpdateReplay interface {
	Validation
	// msg marshals BgpOneStructuredUpdateReplay to protobuf object *otg.BgpOneStructuredUpdateReplay
	// and doesn't set defaults
	msg() *otg.BgpOneStructuredUpdateReplay
	// setMsg unmarshals BgpOneStructuredUpdateReplay from protobuf object *otg.BgpOneStructuredUpdateReplay
	// and doesn't set defaults
	setMsg(*otg.BgpOneStructuredUpdateReplay) BgpOneStructuredUpdateReplay
	// provides marshal interface
	Marshal() marshalBgpOneStructuredUpdateReplay
	// provides unmarshal interface
	Unmarshal() unMarshalBgpOneStructuredUpdateReplay
	// validate validates BgpOneStructuredUpdateReplay
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpOneStructuredUpdateReplay, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TimeGap returns uint32, set in BgpOneStructuredUpdateReplay.
	TimeGap() uint32
	// SetTimeGap assigns uint32 provided by user to BgpOneStructuredUpdateReplay
	SetTimeGap(value uint32) BgpOneStructuredUpdateReplay
	// HasTimeGap checks if TimeGap has been set in BgpOneStructuredUpdateReplay
	HasTimeGap() bool
	// PathAttributes returns BgpAttributes, set in BgpOneStructuredUpdateReplay.
	// BgpAttributes is attributes carried in the Update packet alongwith the reach/unreach prefixes.
	PathAttributes() BgpAttributes
	// SetPathAttributes assigns BgpAttributes provided by user to BgpOneStructuredUpdateReplay.
	// BgpAttributes is attributes carried in the Update packet alongwith the reach/unreach prefixes.
	SetPathAttributes(value BgpAttributes) BgpOneStructuredUpdateReplay
	// HasPathAttributes checks if PathAttributes has been set in BgpOneStructuredUpdateReplay
	HasPathAttributes() bool
	// TraditionalUnreachNlris returns BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIterIter, set in BgpOneStructuredUpdateReplay
	TraditionalUnreachNlris() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	// TraditionalReachNlris returns BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIterIter, set in BgpOneStructuredUpdateReplay
	TraditionalReachNlris() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	setNil()
}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// TimeGap returns a uint32
func (obj *bgpOneStructuredUpdateReplay) TimeGap() uint32 {

	return *obj.obj.TimeGap

}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// TimeGap returns a uint32
func (obj *bgpOneStructuredUpdateReplay) HasTimeGap() bool {
	return obj.obj.TimeGap != nil
}

// Minimum time interval in milliseconds from previous Update from the sequence of BGP Updates to be replayed.
// SetTimeGap sets the uint32 value in the BgpOneStructuredUpdateReplay object
func (obj *bgpOneStructuredUpdateReplay) SetTimeGap(value uint32) BgpOneStructuredUpdateReplay {

	obj.obj.TimeGap = &value
	return obj
}

// Attributes carried in the Update packet alongwith the reach/unreach prefixes.
// PathAttributes returns a BgpAttributes
func (obj *bgpOneStructuredUpdateReplay) PathAttributes() BgpAttributes {
	if obj.obj.PathAttributes == nil {
		obj.obj.PathAttributes = NewBgpAttributes().msg()
	}
	if obj.pathAttributesHolder == nil {
		obj.pathAttributesHolder = &bgpAttributes{obj: obj.obj.PathAttributes}
	}
	return obj.pathAttributesHolder
}

// Attributes carried in the Update packet alongwith the reach/unreach prefixes.
// PathAttributes returns a BgpAttributes
func (obj *bgpOneStructuredUpdateReplay) HasPathAttributes() bool {
	return obj.obj.PathAttributes != nil
}

// Attributes carried in the Update packet alongwith the reach/unreach prefixes.
// SetPathAttributes sets the BgpAttributes value in the BgpOneStructuredUpdateReplay object
func (obj *bgpOneStructuredUpdateReplay) SetPathAttributes(value BgpAttributes) BgpOneStructuredUpdateReplay {

	obj.pathAttributesHolder = nil
	obj.obj.PathAttributes = value.msg()

	return obj
}

// The IPv4 prefixes to be included in the traditional UNREACH_NLRI.
// TraditionalUnreachNlris returns a []BgpOneTraditionalNlriPrefix
func (obj *bgpOneStructuredUpdateReplay) TraditionalUnreachNlris() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	if len(obj.obj.TraditionalUnreachNlris) == 0 {
		obj.obj.TraditionalUnreachNlris = []*otg.BgpOneTraditionalNlriPrefix{}
	}
	if obj.traditionalUnreachNlrisHolder == nil {
		obj.traditionalUnreachNlrisHolder = newBgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter(&obj.obj.TraditionalUnreachNlris).setMsg(obj)
	}
	return obj.traditionalUnreachNlrisHolder
}

type bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter struct {
	obj                              *bgpOneStructuredUpdateReplay
	bgpOneTraditionalNlriPrefixSlice []BgpOneTraditionalNlriPrefix
	fieldPtr                         *[]*otg.BgpOneTraditionalNlriPrefix
}

func newBgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter(ptr *[]*otg.BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	return &bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter{fieldPtr: ptr}
}

type BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter interface {
	setMsg(*bgpOneStructuredUpdateReplay) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	Items() []BgpOneTraditionalNlriPrefix
	Add() BgpOneTraditionalNlriPrefix
	Append(items ...BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	Set(index int, newObj BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	Clear() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	clearHolderSlice() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
	appendHolderSlice(item BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter
}

func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) setMsg(msg *bgpOneStructuredUpdateReplay) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneTraditionalNlriPrefix{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) Items() []BgpOneTraditionalNlriPrefix {
	return obj.bgpOneTraditionalNlriPrefixSlice
}

func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) Add() BgpOneTraditionalNlriPrefix {
	newObj := &otg.BgpOneTraditionalNlriPrefix{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneTraditionalNlriPrefix{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneTraditionalNlriPrefixSlice = append(obj.bgpOneTraditionalNlriPrefixSlice, newLibObj)
	return newLibObj
}

func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) Append(items ...BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneTraditionalNlriPrefixSlice = append(obj.bgpOneTraditionalNlriPrefixSlice, item)
	}
	return obj
}

func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) Set(index int, newObj BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneTraditionalNlriPrefixSlice[index] = newObj
	return obj
}
func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) Clear() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneTraditionalNlriPrefix{}
		obj.bgpOneTraditionalNlriPrefixSlice = []BgpOneTraditionalNlriPrefix{}
	}
	return obj
}
func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) clearHolderSlice() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	if len(obj.bgpOneTraditionalNlriPrefixSlice) > 0 {
		obj.bgpOneTraditionalNlriPrefixSlice = []BgpOneTraditionalNlriPrefix{}
	}
	return obj
}
func (obj *bgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter) appendHolderSlice(item BgpOneTraditionalNlriPrefix) BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	obj.bgpOneTraditionalNlriPrefixSlice = append(obj.bgpOneTraditionalNlriPrefixSlice, item)
	return obj
}

// The IPv4 prefixes to be included in the traditional REACH_NLRI.
// TraditionalReachNlris returns a []BgpOneTraditionalNlriPrefix
func (obj *bgpOneStructuredUpdateReplay) TraditionalReachNlris() BgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter {
	if len(obj.obj.TraditionalReachNlris) == 0 {
		obj.obj.TraditionalReachNlris = []*otg.BgpOneTraditionalNlriPrefix{}
	}
	if obj.traditionalReachNlrisHolder == nil {
		obj.traditionalReachNlrisHolder = newBgpOneStructuredUpdateReplayBgpOneTraditionalNlriPrefixIter(&obj.obj.TraditionalReachNlris).setMsg(obj)
	}
	return obj.traditionalReachNlrisHolder
}

func (obj *bgpOneStructuredUpdateReplay) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.PathAttributes != nil {

		obj.PathAttributes().validateObj(vObj, set_default)
	}

	if len(obj.obj.TraditionalUnreachNlris) != 0 {

		if set_default {
			obj.TraditionalUnreachNlris().clearHolderSlice()
			for _, item := range obj.obj.TraditionalUnreachNlris {
				obj.TraditionalUnreachNlris().appendHolderSlice(&bgpOneTraditionalNlriPrefix{obj: item})
			}
		}
		for _, item := range obj.TraditionalUnreachNlris().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.TraditionalReachNlris) != 0 {

		if set_default {
			obj.TraditionalReachNlris().clearHolderSlice()
			for _, item := range obj.obj.TraditionalReachNlris {
				obj.TraditionalReachNlris().appendHolderSlice(&bgpOneTraditionalNlriPrefix{obj: item})
			}
		}
		for _, item := range obj.TraditionalReachNlris().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpOneStructuredUpdateReplay) setDefault() {
	if obj.obj.TimeGap == nil {
		obj.SetTimeGap(0)
	}

}
