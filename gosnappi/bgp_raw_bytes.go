package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpRawBytes *****
type bgpRawBytes struct {
	validation
	obj           *otg.BgpRawBytes
	marshaller    marshalBgpRawBytes
	unMarshaller  unMarshalBgpRawBytes
	updatesHolder BgpRawBytesBgpOneUpdateReplayIter
}

func NewBgpRawBytes() BgpRawBytes {
	obj := bgpRawBytes{obj: &otg.BgpRawBytes{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpRawBytes) msg() *otg.BgpRawBytes {
	return obj.obj
}

func (obj *bgpRawBytes) setMsg(msg *otg.BgpRawBytes) BgpRawBytes {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpRawBytes struct {
	obj *bgpRawBytes
}

type marshalBgpRawBytes interface {
	// ToProto marshals BgpRawBytes to protobuf object *otg.BgpRawBytes
	ToProto() (*otg.BgpRawBytes, error)
	// ToPbText marshals BgpRawBytes to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpRawBytes to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpRawBytes to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpRawBytes to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpRawBytes struct {
	obj *bgpRawBytes
}

type unMarshalBgpRawBytes interface {
	// FromProto unmarshals BgpRawBytes from protobuf object *otg.BgpRawBytes
	FromProto(msg *otg.BgpRawBytes) (BgpRawBytes, error)
	// FromPbText unmarshals BgpRawBytes from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpRawBytes from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpRawBytes from JSON text
	FromJson(value string) error
}

func (obj *bgpRawBytes) Marshal() marshalBgpRawBytes {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpRawBytes{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpRawBytes) Unmarshal() unMarshalBgpRawBytes {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpRawBytes{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpRawBytes) ToProto() (*otg.BgpRawBytes, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpRawBytes) FromProto(msg *otg.BgpRawBytes) (BgpRawBytes, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpRawBytes) ToPbText() (string, error) {
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

func (m *unMarshalbgpRawBytes) FromPbText(value string) error {
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

func (m *marshalbgpRawBytes) ToYaml() (string, error) {
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

func (m *unMarshalbgpRawBytes) FromYaml(value string) error {
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

func (m *marshalbgpRawBytes) ToJsonRaw() (string, error) {
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

func (m *marshalbgpRawBytes) ToJson() (string, error) {
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

func (m *unMarshalbgpRawBytes) FromJson(value string) error {
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

func (obj *bgpRawBytes) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpRawBytes) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpRawBytes) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpRawBytes) Clone() (BgpRawBytes, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpRawBytes()
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

func (obj *bgpRawBytes) setNil() {
	obj.updatesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpRawBytes is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
type BgpRawBytes interface {
	Validation
	// msg marshals BgpRawBytes to protobuf object *otg.BgpRawBytes
	// and doesn't set defaults
	msg() *otg.BgpRawBytes
	// setMsg unmarshals BgpRawBytes from protobuf object *otg.BgpRawBytes
	// and doesn't set defaults
	setMsg(*otg.BgpRawBytes) BgpRawBytes
	// provides marshal interface
	Marshal() marshalBgpRawBytes
	// provides unmarshal interface
	Unmarshal() unMarshalBgpRawBytes
	// validate validates BgpRawBytes
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpRawBytes, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Updates returns BgpRawBytesBgpOneUpdateReplayIterIter, set in BgpRawBytes
	Updates() BgpRawBytesBgpOneUpdateReplayIter
	setNil()
}

// Array of ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
// Updates returns a []BgpOneUpdateReplay
func (obj *bgpRawBytes) Updates() BgpRawBytesBgpOneUpdateReplayIter {
	if len(obj.obj.Updates) == 0 {
		obj.obj.Updates = []*otg.BgpOneUpdateReplay{}
	}
	if obj.updatesHolder == nil {
		obj.updatesHolder = newBgpRawBytesBgpOneUpdateReplayIter(&obj.obj.Updates).setMsg(obj)
	}
	return obj.updatesHolder
}

type bgpRawBytesBgpOneUpdateReplayIter struct {
	obj                     *bgpRawBytes
	bgpOneUpdateReplaySlice []BgpOneUpdateReplay
	fieldPtr                *[]*otg.BgpOneUpdateReplay
}

func newBgpRawBytesBgpOneUpdateReplayIter(ptr *[]*otg.BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter {
	return &bgpRawBytesBgpOneUpdateReplayIter{fieldPtr: ptr}
}

type BgpRawBytesBgpOneUpdateReplayIter interface {
	setMsg(*bgpRawBytes) BgpRawBytesBgpOneUpdateReplayIter
	Items() []BgpOneUpdateReplay
	Add() BgpOneUpdateReplay
	Append(items ...BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter
	Set(index int, newObj BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter
	Clear() BgpRawBytesBgpOneUpdateReplayIter
	clearHolderSlice() BgpRawBytesBgpOneUpdateReplayIter
	appendHolderSlice(item BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter
}

func (obj *bgpRawBytesBgpOneUpdateReplayIter) setMsg(msg *bgpRawBytes) BgpRawBytesBgpOneUpdateReplayIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneUpdateReplay{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpRawBytesBgpOneUpdateReplayIter) Items() []BgpOneUpdateReplay {
	return obj.bgpOneUpdateReplaySlice
}

func (obj *bgpRawBytesBgpOneUpdateReplayIter) Add() BgpOneUpdateReplay {
	newObj := &otg.BgpOneUpdateReplay{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneUpdateReplay{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneUpdateReplaySlice = append(obj.bgpOneUpdateReplaySlice, newLibObj)
	return newLibObj
}

func (obj *bgpRawBytesBgpOneUpdateReplayIter) Append(items ...BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneUpdateReplaySlice = append(obj.bgpOneUpdateReplaySlice, item)
	}
	return obj
}

func (obj *bgpRawBytesBgpOneUpdateReplayIter) Set(index int, newObj BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneUpdateReplaySlice[index] = newObj
	return obj
}
func (obj *bgpRawBytesBgpOneUpdateReplayIter) Clear() BgpRawBytesBgpOneUpdateReplayIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneUpdateReplay{}
		obj.bgpOneUpdateReplaySlice = []BgpOneUpdateReplay{}
	}
	return obj
}
func (obj *bgpRawBytesBgpOneUpdateReplayIter) clearHolderSlice() BgpRawBytesBgpOneUpdateReplayIter {
	if len(obj.bgpOneUpdateReplaySlice) > 0 {
		obj.bgpOneUpdateReplaySlice = []BgpOneUpdateReplay{}
	}
	return obj
}
func (obj *bgpRawBytesBgpOneUpdateReplayIter) appendHolderSlice(item BgpOneUpdateReplay) BgpRawBytesBgpOneUpdateReplayIter {
	obj.bgpOneUpdateReplaySlice = append(obj.bgpOneUpdateReplaySlice, item)
	return obj
}

func (obj *bgpRawBytes) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Updates) != 0 {

		if set_default {
			obj.Updates().clearHolderSlice()
			for _, item := range obj.obj.Updates {
				obj.Updates().appendHolderSlice(&bgpOneUpdateReplay{obj: item})
			}
		}
		for _, item := range obj.Updates().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpRawBytes) setDefault() {

}
