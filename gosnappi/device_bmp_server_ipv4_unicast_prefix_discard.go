package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerIpv4UnicastPrefixDiscard *****
type deviceBmpServerIpv4UnicastPrefixDiscard struct {
	validation
	obj              *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	marshaller       marshalDeviceBmpServerIpv4UnicastPrefixDiscard
	unMarshaller     unMarshalDeviceBmpServerIpv4UnicastPrefixDiscard
	exceptionsHolder DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
}

func NewDeviceBmpServerIpv4UnicastPrefixDiscard() DeviceBmpServerIpv4UnicastPrefixDiscard {
	obj := deviceBmpServerIpv4UnicastPrefixDiscard{obj: &otg.DeviceBmpServerIpv4UnicastPrefixDiscard{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) msg() *otg.DeviceBmpServerIpv4UnicastPrefixDiscard {
	return obj.obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) setMsg(msg *otg.DeviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixDiscard {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerIpv4UnicastPrefixDiscard struct {
	obj *deviceBmpServerIpv4UnicastPrefixDiscard
}

type marshalDeviceBmpServerIpv4UnicastPrefixDiscard interface {
	// ToProto marshals DeviceBmpServerIpv4UnicastPrefixDiscard to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixDiscard, error)
	// ToPbText marshals DeviceBmpServerIpv4UnicastPrefixDiscard to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerIpv4UnicastPrefixDiscard to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerIpv4UnicastPrefixDiscard to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard struct {
	obj *deviceBmpServerIpv4UnicastPrefixDiscard
}

type unMarshalDeviceBmpServerIpv4UnicastPrefixDiscard interface {
	// FromProto unmarshals DeviceBmpServerIpv4UnicastPrefixDiscard from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixDiscard) (DeviceBmpServerIpv4UnicastPrefixDiscard, error)
	// FromPbText unmarshals DeviceBmpServerIpv4UnicastPrefixDiscard from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerIpv4UnicastPrefixDiscard from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerIpv4UnicastPrefixDiscard from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) Marshal() marshalDeviceBmpServerIpv4UnicastPrefixDiscard {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerIpv4UnicastPrefixDiscard{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixDiscard {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixDiscard) ToProto() (*otg.DeviceBmpServerIpv4UnicastPrefixDiscard, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard) FromProto(msg *otg.DeviceBmpServerIpv4UnicastPrefixDiscard) (DeviceBmpServerIpv4UnicastPrefixDiscard, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerIpv4UnicastPrefixDiscard) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixDiscard) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerIpv4UnicastPrefixDiscard) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerIpv4UnicastPrefixDiscard) FromJson(value string) error {
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

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) Clone() (DeviceBmpServerIpv4UnicastPrefixDiscard, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerIpv4UnicastPrefixDiscard()
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

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) setNil() {
	obj.exceptionsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerIpv4UnicastPrefixDiscard is the exception list can be used to specify exceptions to the specification to discard all IPv4 unicast prefixes. It is expected that when required, there would normally be a limited number of exceptions.
type DeviceBmpServerIpv4UnicastPrefixDiscard interface {
	Validation
	// msg marshals DeviceBmpServerIpv4UnicastPrefixDiscard to protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	// setMsg unmarshals DeviceBmpServerIpv4UnicastPrefixDiscard from protobuf object *otg.DeviceBmpServerIpv4UnicastPrefixDiscard
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixDiscard
	// provides marshal interface
	Marshal() marshalDeviceBmpServerIpv4UnicastPrefixDiscard
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerIpv4UnicastPrefixDiscard
	// validate validates DeviceBmpServerIpv4UnicastPrefixDiscard
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerIpv4UnicastPrefixDiscard, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Exceptions returns DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIterIter, set in DeviceBmpServerIpv4UnicastPrefixDiscard
	Exceptions() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	setNil()
}

// This contains an array of exceptions to the discard IPv4 unicast prefixes specification i.e. prefixes matching the exceptions would be stored instead. One exception to the specification is specified by a combination of an IPv4 prefix and a prefix length.  If the received prefix masked upto the exception's prefix length matches the prefix specified in the exception, the received prefix is deemed as having matched the specified exception e.g. received prefix 172.16.1.0/24 and 172.16.2.0/24 would match specified exception of 172.16.0.0/16 but 172.0.0.0/8 or 192.16.2.0/24 would not.
// Exceptions returns a []DeviceBmpServerIpv4UnicastPrefixException
func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) Exceptions() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	if len(obj.obj.Exceptions) == 0 {
		obj.obj.Exceptions = []*otg.DeviceBmpServerIpv4UnicastPrefixException{}
	}
	if obj.exceptionsHolder == nil {
		obj.exceptionsHolder = newDeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter(&obj.obj.Exceptions).setMsg(obj)
	}
	return obj.exceptionsHolder
}

type deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter struct {
	obj                                            *deviceBmpServerIpv4UnicastPrefixDiscard
	deviceBmpServerIpv4UnicastPrefixExceptionSlice []DeviceBmpServerIpv4UnicastPrefixException
	fieldPtr                                       *[]*otg.DeviceBmpServerIpv4UnicastPrefixException
}

func newDeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter(ptr *[]*otg.DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	return &deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter{fieldPtr: ptr}
}

type DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter interface {
	setMsg(*deviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	Items() []DeviceBmpServerIpv4UnicastPrefixException
	Add() DeviceBmpServerIpv4UnicastPrefixException
	Append(items ...DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	Set(index int, newObj DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	Clear() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	clearHolderSlice() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
	appendHolderSlice(item DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) setMsg(msg *deviceBmpServerIpv4UnicastPrefixDiscard) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&deviceBmpServerIpv4UnicastPrefixException{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) Items() []DeviceBmpServerIpv4UnicastPrefixException {
	return obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) Add() DeviceBmpServerIpv4UnicastPrefixException {
	newObj := &otg.DeviceBmpServerIpv4UnicastPrefixException{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &deviceBmpServerIpv4UnicastPrefixException{obj: newObj}
	newLibObj.setDefault()
	obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice, newLibObj)
	return newLibObj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) Append(items ...DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice, item)
	}
	return obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) Set(index int, newObj DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice[index] = newObj
	return obj
}
func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) Clear() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.DeviceBmpServerIpv4UnicastPrefixException{}
		obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice = []DeviceBmpServerIpv4UnicastPrefixException{}
	}
	return obj
}
func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) clearHolderSlice() DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	if len(obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice) > 0 {
		obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice = []DeviceBmpServerIpv4UnicastPrefixException{}
	}
	return obj
}
func (obj *deviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter) appendHolderSlice(item DeviceBmpServerIpv4UnicastPrefixException) DeviceBmpServerIpv4UnicastPrefixDiscardDeviceBmpServerIpv4UnicastPrefixExceptionIter {
	obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice = append(obj.deviceBmpServerIpv4UnicastPrefixExceptionSlice, item)
	return obj
}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Exceptions) != 0 {

		if set_default {
			obj.Exceptions().clearHolderSlice()
			for _, item := range obj.obj.Exceptions {
				obj.Exceptions().appendHolderSlice(&deviceBmpServerIpv4UnicastPrefixException{obj: item})
			}
		}
		for _, item := range obj.Exceptions().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *deviceBmpServerIpv4UnicastPrefixDiscard) setDefault() {

}
