package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv6UnicastPrefixDiscard *****
type deviceBmpServerIpv6UnicastPrefixDiscard struct {
	validation
	obj              *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	marshaller       marshalDeviceBmpServerIpv6UnicastPrefixDiscard
	unMarshaller     unMarshalDeviceBmpServerIpv6UnicastPrefixDiscard
	exceptionsHolder DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
}

func NewDeviceBmpServerIpv6UnicastPrefixDiscard() DeviceBmpServerIpv6UnicastPrefixDiscard {
	obj := deviceBmpServerIpv6UnicastPrefixDiscard{obj: &otg.DeviceBmpServerIpv6UnicastPrefixDiscard{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) msg() *otg.DeviceBmpServerIpv6UnicastPrefixDiscard {
	return obj.obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) setMsg(msg *otg.DeviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixDiscard {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv6UnicastPrefixDiscard struct {
	obj *deviceBmpServerIpv6UnicastPrefixDiscard
}

type marshalDeviceBmpServerIpv6UnicastPrefixDiscard interface {
	// ToProto marshals DeviceBmpServerIpv6UnicastPrefixDiscard to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixDiscard, error)
	// ToPbText marshals DeviceBmpServerIpv6UnicastPrefixDiscard to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv6UnicastPrefixDiscard to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv6UnicastPrefixDiscard to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard struct {
	obj *deviceBmpServerIpv6UnicastPrefixDiscard
}

type unMarshalDeviceBmpServerIpv6UnicastPrefixDiscard interface {
	// FromProto unmarshals DeviceBmpServerIpv6UnicastPrefixDiscard from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixDiscard) (DeviceBmpServerIpv6UnicastPrefixDiscard, error)
	// FromPbText unmarshals DeviceBmpServerIpv6UnicastPrefixDiscard from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv6UnicastPrefixDiscard from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv6UnicastPrefixDiscard from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) Marshal() marshalDeviceBmpServerIpv6UnicastPrefixDiscard {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv6UnicastPrefixDiscard{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixDiscard {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixDiscard) ToProto() (*otg.DeviceBmpServerIpv6UnicastPrefixDiscard, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard) FromProto(msg *otg.DeviceBmpServerIpv6UnicastPrefixDiscard) (DeviceBmpServerIpv6UnicastPrefixDiscard, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv6UnicastPrefixDiscard) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixDiscard) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv6UnicastPrefixDiscard) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv6UnicastPrefixDiscard) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) Clone() (DeviceBmpServerIpv6UnicastPrefixDiscard, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv6UnicastPrefixDiscard()
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

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) setNil() {
	obj.exceptionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerIpv6UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv6 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
type DeviceBmpServerIpv6UnicastPrefixDiscard interface {
	Validation
	// msg marshals DeviceBmpServerIpv6UnicastPrefixDiscard to protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	// setMsg unmarshals DeviceBmpServerIpv6UnicastPrefixDiscard from protobuf object *otg.DeviceBmpServerIpv6UnicastPrefixDiscard
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixDiscard
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv6UnicastPrefixDiscard
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv6UnicastPrefixDiscard
	// validate validates DeviceBmpServerIpv6UnicastPrefixDiscard
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv6UnicastPrefixDiscard, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Exceptions returns DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIterIter, set in DeviceBmpServerIpv6UnicastPrefixDiscard
	Exceptions() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	setNil()
}

// This contains an array of exceptions to the discard IPv6 unicast prefixes specification i.e. prefixes matching the exceptions would be stored instead. One exception to the specification is specified by a combination of an IPv6 prefix and a prefix length.  If the received prefix masked upto the exception's prefix length matches the prefix specified in the exception, the received prefix is deemed as having matched the specified exception e.g. received prefix 1:1:1:1::/64  and 1:1:1:2::/64 would match specified exception of 1:1:1::/48 but 1:1::/16 or 2:2:2:2::/64 would not.
// Exceptions returns a []DeviceBmpServerIpv6UnicastPrefixException
func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) Exceptions() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	if len(obj.obj.Exceptions) == 0 {
		obj.obj.Exceptions = []*otg.DeviceBmpServerIpv6UnicastPrefixException{}
	}
	if obj.exceptionsHolder == nil {
		obj.exceptionsHolder = newDeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter(&obj.obj.Exceptions).setMsg(obj)
	}
	return obj.exceptionsHolder
}

type deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter struct {
	obj                                            *deviceBmpServerIpv6UnicastPrefixDiscard
	deviceBmpServerIpv6UnicastPrefixExceptionSlice []DeviceBmpServerIpv6UnicastPrefixException
	fieldPtr                                       *[]*otg.DeviceBmpServerIpv6UnicastPrefixException
}

func newDeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter(ptr *[]*otg.DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	return &deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter{fieldPtr: ptr}
}

type DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter interface {
	setMsg(*deviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	Items() []DeviceBmpServerIpv6UnicastPrefixException
	Add() DeviceBmpServerIpv6UnicastPrefixException
	Append(items ...DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	Set(index int, newObj DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	Clear() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	clearHolderSlice() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
	appendHolderSlice(item DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) setMsg(msg *deviceBmpServerIpv6UnicastPrefixDiscard) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpServerIpv6UnicastPrefixException{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) Items() []DeviceBmpServerIpv6UnicastPrefixException {
	return obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) Add() DeviceBmpServerIpv6UnicastPrefixException {
	newObj := &otg.DeviceBmpServerIpv6UnicastPrefixException{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpServerIpv6UnicastPrefixException{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) Append(items ...DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice, item)
	}
	return obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) Set(index int, newObj DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice[index] = newObj
	return obj
}
func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) Clear() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpServerIpv6UnicastPrefixException{}
		obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice = []DeviceBmpServerIpv6UnicastPrefixException{}
	}
	return obj
}
func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) clearHolderSlice() DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	if len(obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice) > 0 {
		obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice = []DeviceBmpServerIpv6UnicastPrefixException{}
	}
	return obj
}
func (obj *deviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter) appendHolderSlice(item DeviceBmpServerIpv6UnicastPrefixException) DeviceBmpServerIpv6UnicastPrefixDiscardDeviceBmpServerIpv6UnicastPrefixExceptionIter {
	obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv6UnicastPrefixExceptionSlice, item)
	return obj
}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Exceptions) != 0 {

		if set_default {
			obj.Exceptions().clearHolderSlice()
			for _, item := range obj.obj.Exceptions {
				obj.Exceptions().appendHolderSlice(&deviceBmpServerIpv6UnicastPrefixException{obj: item})
			}
		}
		for _, item := range obj.Exceptions().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBmpServerIpv6UnicastPrefixDiscard) setDefault() {

}
