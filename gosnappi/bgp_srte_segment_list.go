package gosnappi

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpSrteSegmentList *****
type bgpSrteSegmentList struct {
	validation
	obj            *otg.BgpSrteSegmentList
	marshaller     marshalBgpSrteSegmentList
	unMarshaller   unMarshalBgpSrteSegmentList
	segmentsHolder BgpSrteSegmentListBgpSrteSegmentIter
}

func NewBgpSrteSegmentList() BgpSrteSegmentList {
	obj := bgpSrteSegmentList{obj: &otg.BgpSrteSegmentList{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpSrteSegmentList) msg() *otg.BgpSrteSegmentList {
	return obj.obj
}

func (obj *bgpSrteSegmentList) setMsg(msg *otg.BgpSrteSegmentList) BgpSrteSegmentList {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpSrteSegmentList struct {
	obj *bgpSrteSegmentList
}

type marshalBgpSrteSegmentList interface {
	// ToProto marshals BgpSrteSegmentList to protobuf object *otg.BgpSrteSegmentList
	ToProto() (*otg.BgpSrteSegmentList, error)
	// ToPbText marshals BgpSrteSegmentList to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpSrteSegmentList to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpSrteSegmentList to JSON text
	ToJson() (string, error)
}

type unMarshalbgpSrteSegmentList struct {
	obj *bgpSrteSegmentList
}

type unMarshalBgpSrteSegmentList interface {
	// FromProto unmarshals BgpSrteSegmentList from protobuf object *otg.BgpSrteSegmentList
	FromProto(msg *otg.BgpSrteSegmentList) (BgpSrteSegmentList, error)
	// FromPbText unmarshals BgpSrteSegmentList from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpSrteSegmentList from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpSrteSegmentList from JSON text
	FromJson(value string) error
}

func (obj *bgpSrteSegmentList) Marshal() marshalBgpSrteSegmentList {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpSrteSegmentList{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpSrteSegmentList) Unmarshal() unMarshalBgpSrteSegmentList {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpSrteSegmentList{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpSrteSegmentList) ToProto() (*otg.BgpSrteSegmentList, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpSrteSegmentList) FromProto(msg *otg.BgpSrteSegmentList) (BgpSrteSegmentList, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpSrteSegmentList) ToPbText() (string, error) {
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

func (m *unMarshalbgpSrteSegmentList) FromPbText(value string) error {
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

func (m *marshalbgpSrteSegmentList) ToYaml() (string, error) {
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

func (m *unMarshalbgpSrteSegmentList) FromYaml(value string) error {
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

func (m *marshalbgpSrteSegmentList) ToJson() (string, error) {
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

func (m *unMarshalbgpSrteSegmentList) FromJson(value string) error {
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

func (obj *bgpSrteSegmentList) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentList) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpSrteSegmentList) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpSrteSegmentList) Clone() (BgpSrteSegmentList, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpSrteSegmentList()
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

func (obj *bgpSrteSegmentList) setNil() {
	obj.segmentsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpSrteSegmentList is optional configuration for BGP SR TE Policy segment list. The Segment List sub-TLV encodes a single explicit path towards the Endpoint.
type BgpSrteSegmentList interface {
	Validation
	// msg marshals BgpSrteSegmentList to protobuf object *otg.BgpSrteSegmentList
	// and doesn't set defaults
	msg() *otg.BgpSrteSegmentList
	// setMsg unmarshals BgpSrteSegmentList from protobuf object *otg.BgpSrteSegmentList
	// and doesn't set defaults
	setMsg(*otg.BgpSrteSegmentList) BgpSrteSegmentList
	// provides marshal interface
	Marshal() marshalBgpSrteSegmentList
	// provides unmarshal interface
	Unmarshal() unMarshalBgpSrteSegmentList
	// validate validates BgpSrteSegmentList
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpSrteSegmentList, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Weight returns uint32, set in BgpSrteSegmentList.
	Weight() uint32
	// SetWeight assigns uint32 provided by user to BgpSrteSegmentList
	SetWeight(value uint32) BgpSrteSegmentList
	// HasWeight checks if Weight has been set in BgpSrteSegmentList
	HasWeight() bool
	// Segments returns BgpSrteSegmentListBgpSrteSegmentIterIter, set in BgpSrteSegmentList
	Segments() BgpSrteSegmentListBgpSrteSegmentIter
	// Name returns string, set in BgpSrteSegmentList.
	Name() string
	// SetName assigns string provided by user to BgpSrteSegmentList
	SetName(value string) BgpSrteSegmentList
	// Active returns bool, set in BgpSrteSegmentList.
	Active() bool
	// SetActive assigns bool provided by user to BgpSrteSegmentList
	SetActive(value bool) BgpSrteSegmentList
	// HasActive checks if Active has been set in BgpSrteSegmentList
	HasActive() bool
	setNil()
}

// The Weight associated with a given path and the sub-TLV is optional.
// Weight returns a uint32
func (obj *bgpSrteSegmentList) Weight() uint32 {

	return *obj.obj.Weight

}

// The Weight associated with a given path and the sub-TLV is optional.
// Weight returns a uint32
func (obj *bgpSrteSegmentList) HasWeight() bool {
	return obj.obj.Weight != nil
}

// The Weight associated with a given path and the sub-TLV is optional.
// SetWeight sets the uint32 value in the BgpSrteSegmentList object
func (obj *bgpSrteSegmentList) SetWeight(value uint32) BgpSrteSegmentList {

	obj.obj.Weight = &value
	return obj
}

// description is TBD
// Segments returns a []BgpSrteSegment
func (obj *bgpSrteSegmentList) Segments() BgpSrteSegmentListBgpSrteSegmentIter {
	if len(obj.obj.Segments) == 0 {
		obj.obj.Segments = []*otg.BgpSrteSegment{}
	}
	if obj.segmentsHolder == nil {
		obj.segmentsHolder = newBgpSrteSegmentListBgpSrteSegmentIter(&obj.obj.Segments).setMsg(obj)
	}
	return obj.segmentsHolder
}

type bgpSrteSegmentListBgpSrteSegmentIter struct {
	obj                 *bgpSrteSegmentList
	bgpSrteSegmentSlice []BgpSrteSegment
	fieldPtr            *[]*otg.BgpSrteSegment
}

func newBgpSrteSegmentListBgpSrteSegmentIter(ptr *[]*otg.BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter {
	return &bgpSrteSegmentListBgpSrteSegmentIter{fieldPtr: ptr}
}

type BgpSrteSegmentListBgpSrteSegmentIter interface {
	setMsg(*bgpSrteSegmentList) BgpSrteSegmentListBgpSrteSegmentIter
	Items() []BgpSrteSegment
	Add() BgpSrteSegment
	Append(items ...BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter
	Set(index int, newObj BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter
	Clear() BgpSrteSegmentListBgpSrteSegmentIter
	clearHolderSlice() BgpSrteSegmentListBgpSrteSegmentIter
	appendHolderSlice(item BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter
}

func (obj *bgpSrteSegmentListBgpSrteSegmentIter) setMsg(msg *bgpSrteSegmentList) BgpSrteSegmentListBgpSrteSegmentIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpSrteSegment{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpSrteSegmentListBgpSrteSegmentIter) Items() []BgpSrteSegment {
	return obj.bgpSrteSegmentSlice
}

func (obj *bgpSrteSegmentListBgpSrteSegmentIter) Add() BgpSrteSegment {
	newObj := &otg.BgpSrteSegment{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpSrteSegment{obj: newObj}
	newLibObj.setDefault()
	obj.bgpSrteSegmentSlice = append(obj.bgpSrteSegmentSlice, newLibObj)
	return newLibObj
}

func (obj *bgpSrteSegmentListBgpSrteSegmentIter) Append(items ...BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpSrteSegmentSlice = append(obj.bgpSrteSegmentSlice, item)
	}
	return obj
}

func (obj *bgpSrteSegmentListBgpSrteSegmentIter) Set(index int, newObj BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpSrteSegmentSlice[index] = newObj
	return obj
}
func (obj *bgpSrteSegmentListBgpSrteSegmentIter) Clear() BgpSrteSegmentListBgpSrteSegmentIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpSrteSegment{}
		obj.bgpSrteSegmentSlice = []BgpSrteSegment{}
	}
	return obj
}
func (obj *bgpSrteSegmentListBgpSrteSegmentIter) clearHolderSlice() BgpSrteSegmentListBgpSrteSegmentIter {
	if len(obj.bgpSrteSegmentSlice) > 0 {
		obj.bgpSrteSegmentSlice = []BgpSrteSegment{}
	}
	return obj
}
func (obj *bgpSrteSegmentListBgpSrteSegmentIter) appendHolderSlice(item BgpSrteSegment) BgpSrteSegmentListBgpSrteSegmentIter {
	obj.bgpSrteSegmentSlice = append(obj.bgpSrteSegmentSlice, item)
	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *bgpSrteSegmentList) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the BgpSrteSegmentList object
func (obj *bgpSrteSegmentList) SetName(value string) BgpSrteSegmentList {

	obj.obj.Name = &value
	return obj
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteSegmentList) Active() bool {

	return *obj.obj.Active

}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// Active returns a bool
func (obj *bgpSrteSegmentList) HasActive() bool {
	return obj.obj.Active != nil
}

// If enabled means that this part of the configuration including any active 'children' nodes will be advertised to peer.  If disabled, this means that though config is present, it is not taking any part of the test but can be activated at run-time to advertise just this part of the configuration to the peer.
// SetActive sets the bool value in the BgpSrteSegmentList object
func (obj *bgpSrteSegmentList) SetActive(value bool) BgpSrteSegmentList {

	obj.obj.Active = &value
	return obj
}

func (obj *bgpSrteSegmentList) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Segments) != 0 {

		if set_default {
			obj.Segments().clearHolderSlice()
			for _, item := range obj.obj.Segments {
				obj.Segments().appendHolderSlice(&bgpSrteSegment{obj: item})
			}
		}
		for _, item := range obj.Segments().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface BgpSrteSegmentList")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"BgpSrteSegmentList.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

}

func (obj *bgpSrteSegmentList) setDefault() {
	if obj.obj.Weight == nil {
		obj.SetWeight(0)
	}
	if obj.obj.Active == nil {
		obj.SetActive(true)
	}

}
