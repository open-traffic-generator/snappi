package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceDhcpv6ClientOptionsRequest *****
type deviceDhcpv6ClientOptionsRequest struct {
	validation
	obj                          *otg.DeviceDhcpv6ClientOptionsRequest
	marshaller                   marshalDeviceDhcpv6ClientOptionsRequest
	unMarshaller                 unMarshalDeviceDhcpv6ClientOptionsRequest
	requestHolder                DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	associatedDhcpMessagesHolder Dhcpv6ClientOptionsIncludedMessages
}

func NewDeviceDhcpv6ClientOptionsRequest() DeviceDhcpv6ClientOptionsRequest {
	obj := deviceDhcpv6ClientOptionsRequest{obj: &otg.DeviceDhcpv6ClientOptionsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv6ClientOptionsRequest) msg() *otg.DeviceDhcpv6ClientOptionsRequest {
	return obj.obj
}

func (obj *deviceDhcpv6ClientOptionsRequest) setMsg(msg *otg.DeviceDhcpv6ClientOptionsRequest) DeviceDhcpv6ClientOptionsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv6ClientOptionsRequest struct {
	obj *deviceDhcpv6ClientOptionsRequest
}

type marshalDeviceDhcpv6ClientOptionsRequest interface {
	// ToProto marshals DeviceDhcpv6ClientOptionsRequest to protobuf object *otg.DeviceDhcpv6ClientOptionsRequest
	ToProto() (*otg.DeviceDhcpv6ClientOptionsRequest, error)
	// ToPbText marshals DeviceDhcpv6ClientOptionsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv6ClientOptionsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv6ClientOptionsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv6ClientOptionsRequest struct {
	obj *deviceDhcpv6ClientOptionsRequest
}

type unMarshalDeviceDhcpv6ClientOptionsRequest interface {
	// FromProto unmarshals DeviceDhcpv6ClientOptionsRequest from protobuf object *otg.DeviceDhcpv6ClientOptionsRequest
	FromProto(msg *otg.DeviceDhcpv6ClientOptionsRequest) (DeviceDhcpv6ClientOptionsRequest, error)
	// FromPbText unmarshals DeviceDhcpv6ClientOptionsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv6ClientOptionsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv6ClientOptionsRequest from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv6ClientOptionsRequest) Marshal() marshalDeviceDhcpv6ClientOptionsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv6ClientOptionsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv6ClientOptionsRequest) Unmarshal() unMarshalDeviceDhcpv6ClientOptionsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv6ClientOptionsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv6ClientOptionsRequest) ToProto() (*otg.DeviceDhcpv6ClientOptionsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv6ClientOptionsRequest) FromProto(msg *otg.DeviceDhcpv6ClientOptionsRequest) (DeviceDhcpv6ClientOptionsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv6ClientOptionsRequest) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptionsRequest) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv6ClientOptionsRequest) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptionsRequest) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv6ClientOptionsRequest) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv6ClientOptionsRequest) FromJson(value string) error {
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

func (obj *deviceDhcpv6ClientOptionsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientOptionsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv6ClientOptionsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv6ClientOptionsRequest) Clone() (DeviceDhcpv6ClientOptionsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv6ClientOptionsRequest()
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

func (obj *deviceDhcpv6ClientOptionsRequest) setNil() {
	obj.requestHolder = nil
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv6ClientOptionsRequest is dHCP client options, these configured options are sent in Dhcp client messages.
type DeviceDhcpv6ClientOptionsRequest interface {
	Validation
	// msg marshals DeviceDhcpv6ClientOptionsRequest to protobuf object *otg.DeviceDhcpv6ClientOptionsRequest
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv6ClientOptionsRequest
	// setMsg unmarshals DeviceDhcpv6ClientOptionsRequest from protobuf object *otg.DeviceDhcpv6ClientOptionsRequest
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv6ClientOptionsRequest) DeviceDhcpv6ClientOptionsRequest
	// provides marshal interface
	Marshal() marshalDeviceDhcpv6ClientOptionsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv6ClientOptionsRequest
	// validate validates DeviceDhcpv6ClientOptionsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv6ClientOptionsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Request returns DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIterIter, set in DeviceDhcpv6ClientOptionsRequest
	Request() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	// AssociatedDhcpMessages returns Dhcpv6ClientOptionsIncludedMessages, set in DeviceDhcpv6ClientOptionsRequest.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ClientOptionsIncludedMessages provided by user to DeviceDhcpv6ClientOptionsRequest.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) DeviceDhcpv6ClientOptionsRequest
	setNil()
}

// List of options requested by a client from a server.
// Request returns a []Dhcpv6ClientOptionsOptionsRequest
func (obj *deviceDhcpv6ClientOptionsRequest) Request() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	if len(obj.obj.Request) == 0 {
		obj.obj.Request = []*otg.Dhcpv6ClientOptionsOptionsRequest{}
	}
	if obj.requestHolder == nil {
		obj.requestHolder = newDeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter(&obj.obj.Request).setMsg(obj)
	}
	return obj.requestHolder
}

type deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter struct {
	obj                                    *deviceDhcpv6ClientOptionsRequest
	dhcpv6ClientOptionsOptionsRequestSlice []Dhcpv6ClientOptionsOptionsRequest
	fieldPtr                               *[]*otg.Dhcpv6ClientOptionsOptionsRequest
}

func newDeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter(ptr *[]*otg.Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	return &deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter{fieldPtr: ptr}
}

type DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter interface {
	setMsg(*deviceDhcpv6ClientOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	Items() []Dhcpv6ClientOptionsOptionsRequest
	Add() Dhcpv6ClientOptionsOptionsRequest
	Append(items ...Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	Set(index int, newObj Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	Clear() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	clearHolderSlice() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
	appendHolderSlice(item Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter
}

func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) setMsg(msg *deviceDhcpv6ClientOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ClientOptionsOptionsRequest{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) Items() []Dhcpv6ClientOptionsOptionsRequest {
	return obj.dhcpv6ClientOptionsOptionsRequestSlice
}

func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) Add() Dhcpv6ClientOptionsOptionsRequest {
	newObj := &otg.Dhcpv6ClientOptionsOptionsRequest{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ClientOptionsOptionsRequest{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, newLibObj)
	return newLibObj
}

func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) Append(items ...Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, item)
	}
	return obj
}

func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) Set(index int, newObj Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ClientOptionsOptionsRequestSlice[index] = newObj
	return obj
}
func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) Clear() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ClientOptionsOptionsRequest{}
		obj.dhcpv6ClientOptionsOptionsRequestSlice = []Dhcpv6ClientOptionsOptionsRequest{}
	}
	return obj
}
func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) clearHolderSlice() DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	if len(obj.dhcpv6ClientOptionsOptionsRequestSlice) > 0 {
		obj.dhcpv6ClientOptionsOptionsRequestSlice = []Dhcpv6ClientOptionsOptionsRequest{}
	}
	return obj
}
func (obj *deviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter) appendHolderSlice(item Dhcpv6ClientOptionsOptionsRequest) DeviceDhcpv6ClientOptionsRequestDhcpv6ClientOptionsOptionsRequestIter {
	obj.dhcpv6ClientOptionsOptionsRequestSlice = append(obj.dhcpv6ClientOptionsOptionsRequestSlice, item)
	return obj
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ClientOptionsIncludedMessages
func (obj *deviceDhcpv6ClientOptionsRequest) AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ClientOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ClientOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The list of dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ClientOptionsIncludedMessages value in the DeviceDhcpv6ClientOptionsRequest object
func (obj *deviceDhcpv6ClientOptionsRequest) SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) DeviceDhcpv6ClientOptionsRequest {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *deviceDhcpv6ClientOptionsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.Request) != 0 {

		if set_default {
			obj.Request().clearHolderSlice()
			for _, item := range obj.obj.Request {
				obj.Request().appendHolderSlice(&dhcpv6ClientOptionsOptionsRequest{obj: item})
			}
		}
		for _, item := range obj.Request().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// AssociatedDhcpMessages is required
	if obj.obj.AssociatedDhcpMessages == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AssociatedDhcpMessages is required field on interface DeviceDhcpv6ClientOptionsRequest")
	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv6ClientOptionsRequest) setDefault() {

}
