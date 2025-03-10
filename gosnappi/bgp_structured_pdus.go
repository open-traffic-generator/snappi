package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** BgpStructuredPdus *****
type bgpStructuredPdus struct {
	validation
	obj           *otg.BgpStructuredPdus
	marshaller    marshalBgpStructuredPdus
	unMarshaller  unMarshalBgpStructuredPdus
	updatesHolder BgpStructuredPdusBgpOneStructuredUpdateReplayIter
}

func NewBgpStructuredPdus() BgpStructuredPdus {
	obj := bgpStructuredPdus{obj: &otg.BgpStructuredPdus{}}
	obj.setDefault()
	return &obj
}

func (obj *bgpStructuredPdus) msg() *otg.BgpStructuredPdus {
	return obj.obj
}

func (obj *bgpStructuredPdus) setMsg(msg *otg.BgpStructuredPdus) BgpStructuredPdus {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalbgpStructuredPdus struct {
	obj *bgpStructuredPdus
}

type marshalBgpStructuredPdus interface {
	// ToProto marshals BgpStructuredPdus to protobuf object *otg.BgpStructuredPdus
	ToProto() (*otg.BgpStructuredPdus, error)
	// ToPbText marshals BgpStructuredPdus to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals BgpStructuredPdus to YAML text
	ToYaml() (string, error)
	// ToJson marshals BgpStructuredPdus to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals BgpStructuredPdus to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalbgpStructuredPdus struct {
	obj *bgpStructuredPdus
}

type unMarshalBgpStructuredPdus interface {
	// FromProto unmarshals BgpStructuredPdus from protobuf object *otg.BgpStructuredPdus
	FromProto(msg *otg.BgpStructuredPdus) (BgpStructuredPdus, error)
	// FromPbText unmarshals BgpStructuredPdus from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals BgpStructuredPdus from YAML text
	FromYaml(value string) error
	// FromJson unmarshals BgpStructuredPdus from JSON text
	FromJson(value string) error
}

func (obj *bgpStructuredPdus) Marshal() marshalBgpStructuredPdus {
	if obj.marshaller == nil {
		obj.marshaller = &marshalbgpStructuredPdus{obj: obj}
	}
	return obj.marshaller
}

func (obj *bgpStructuredPdus) Unmarshal() unMarshalBgpStructuredPdus {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalbgpStructuredPdus{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalbgpStructuredPdus) ToProto() (*otg.BgpStructuredPdus, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalbgpStructuredPdus) FromProto(msg *otg.BgpStructuredPdus) (BgpStructuredPdus, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalbgpStructuredPdus) ToPbText() (string, error) {
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

func (m *unMarshalbgpStructuredPdus) FromPbText(value string) error {
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

func (m *marshalbgpStructuredPdus) ToYaml() (string, error) {
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

func (m *unMarshalbgpStructuredPdus) FromYaml(value string) error {
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

func (m *marshalbgpStructuredPdus) ToJsonRaw() (string, error) {
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

func (m *marshalbgpStructuredPdus) ToJson() (string, error) {
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

func (m *unMarshalbgpStructuredPdus) FromJson(value string) error {
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

func (obj *bgpStructuredPdus) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *bgpStructuredPdus) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *bgpStructuredPdus) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *bgpStructuredPdus) Clone() (BgpStructuredPdus, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewBgpStructuredPdus()
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

func (obj *bgpStructuredPdus) setNil() {
	obj.updatesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// BgpStructuredPdus is ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
type BgpStructuredPdus interface {
	Validation
	// msg marshals BgpStructuredPdus to protobuf object *otg.BgpStructuredPdus
	// and doesn't set defaults
	msg() *otg.BgpStructuredPdus
	// setMsg unmarshals BgpStructuredPdus from protobuf object *otg.BgpStructuredPdus
	// and doesn't set defaults
	setMsg(*otg.BgpStructuredPdus) BgpStructuredPdus
	// provides marshal interface
	Marshal() marshalBgpStructuredPdus
	// provides unmarshal interface
	Unmarshal() unMarshalBgpStructuredPdus
	// validate validates BgpStructuredPdus
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (BgpStructuredPdus, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Updates returns BgpStructuredPdusBgpOneStructuredUpdateReplayIterIter, set in BgpStructuredPdus
	Updates() BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	setNil()
}

// Array of ordered BGP Updates ( including both Advertise and Withdraws ) to be sent in the order given in the input to the peer after the BGP session is established.
// Updates returns a []BgpOneStructuredUpdateReplay
func (obj *bgpStructuredPdus) Updates() BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	if len(obj.obj.Updates) == 0 {
		obj.obj.Updates = []*otg.BgpOneStructuredUpdateReplay{}
	}
	if obj.updatesHolder == nil {
		obj.updatesHolder = newBgpStructuredPdusBgpOneStructuredUpdateReplayIter(&obj.obj.Updates).setMsg(obj)
	}
	return obj.updatesHolder
}

type bgpStructuredPdusBgpOneStructuredUpdateReplayIter struct {
	obj                               *bgpStructuredPdus
	bgpOneStructuredUpdateReplaySlice []BgpOneStructuredUpdateReplay
	fieldPtr                          *[]*otg.BgpOneStructuredUpdateReplay
}

func newBgpStructuredPdusBgpOneStructuredUpdateReplayIter(ptr *[]*otg.BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	return &bgpStructuredPdusBgpOneStructuredUpdateReplayIter{fieldPtr: ptr}
}

type BgpStructuredPdusBgpOneStructuredUpdateReplayIter interface {
	setMsg(*bgpStructuredPdus) BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	Items() []BgpOneStructuredUpdateReplay
	Add() BgpOneStructuredUpdateReplay
	Append(items ...BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	Set(index int, newObj BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	Clear() BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	clearHolderSlice() BgpStructuredPdusBgpOneStructuredUpdateReplayIter
	appendHolderSlice(item BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter
}

func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) setMsg(msg *bgpStructuredPdus) BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&bgpOneStructuredUpdateReplay{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) Items() []BgpOneStructuredUpdateReplay {
	return obj.bgpOneStructuredUpdateReplaySlice
}

func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) Add() BgpOneStructuredUpdateReplay {
	newObj := &otg.BgpOneStructuredUpdateReplay{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &bgpOneStructuredUpdateReplay{obj: newObj}
	newLibObj.setDefault()
	obj.bgpOneStructuredUpdateReplaySlice = append(obj.bgpOneStructuredUpdateReplaySlice, newLibObj)
	return newLibObj
}

func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) Append(items ...BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.bgpOneStructuredUpdateReplaySlice = append(obj.bgpOneStructuredUpdateReplaySlice, item)
	}
	return obj
}

func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) Set(index int, newObj BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.bgpOneStructuredUpdateReplaySlice[index] = newObj
	return obj
}
func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) Clear() BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.BgpOneStructuredUpdateReplay{}
		obj.bgpOneStructuredUpdateReplaySlice = []BgpOneStructuredUpdateReplay{}
	}
	return obj
}
func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) clearHolderSlice() BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	if len(obj.bgpOneStructuredUpdateReplaySlice) > 0 {
		obj.bgpOneStructuredUpdateReplaySlice = []BgpOneStructuredUpdateReplay{}
	}
	return obj
}
func (obj *bgpStructuredPdusBgpOneStructuredUpdateReplayIter) appendHolderSlice(item BgpOneStructuredUpdateReplay) BgpStructuredPdusBgpOneStructuredUpdateReplayIter {
	obj.bgpOneStructuredUpdateReplaySlice = append(obj.bgpOneStructuredUpdateReplaySlice, item)
	return obj
}

func (obj *bgpStructuredPdus) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Updates) != 0 {

		if set_default {
			obj.Updates().clearHolderSlice()
			for _, item := range obj.obj.Updates {
				obj.Updates().appendHolderSlice(&bgpOneStructuredUpdateReplay{obj: item})
			}
		}
		for _, item := range obj.Updates().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *bgpStructuredPdus) setDefault() {

}
