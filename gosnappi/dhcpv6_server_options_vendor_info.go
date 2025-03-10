package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptionsVendorInfo *****
type dhcpv6ServerOptionsVendorInfo struct {
	validation
	obj                          *otg.Dhcpv6ServerOptionsVendorInfo
	marshaller                   marshalDhcpv6ServerOptionsVendorInfo
	unMarshaller                 unMarshalDhcpv6ServerOptionsVendorInfo
	optionDataHolder             Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	associatedDhcpMessagesHolder Dhcpv6ServerOptionsIncludedMessages
}

func NewDhcpv6ServerOptionsVendorInfo() Dhcpv6ServerOptionsVendorInfo {
	obj := dhcpv6ServerOptionsVendorInfo{obj: &otg.Dhcpv6ServerOptionsVendorInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptionsVendorInfo) msg() *otg.Dhcpv6ServerOptionsVendorInfo {
	return obj.obj
}

func (obj *dhcpv6ServerOptionsVendorInfo) setMsg(msg *otg.Dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptionsVendorInfo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptionsVendorInfo struct {
	obj *dhcpv6ServerOptionsVendorInfo
}

type marshalDhcpv6ServerOptionsVendorInfo interface {
	// ToProto marshals Dhcpv6ServerOptionsVendorInfo to protobuf object *otg.Dhcpv6ServerOptionsVendorInfo
	ToProto() (*otg.Dhcpv6ServerOptionsVendorInfo, error)
	// ToPbText marshals Dhcpv6ServerOptionsVendorInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptionsVendorInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptionsVendorInfo to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerOptionsVendorInfo to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerOptionsVendorInfo struct {
	obj *dhcpv6ServerOptionsVendorInfo
}

type unMarshalDhcpv6ServerOptionsVendorInfo interface {
	// FromProto unmarshals Dhcpv6ServerOptionsVendorInfo from protobuf object *otg.Dhcpv6ServerOptionsVendorInfo
	FromProto(msg *otg.Dhcpv6ServerOptionsVendorInfo) (Dhcpv6ServerOptionsVendorInfo, error)
	// FromPbText unmarshals Dhcpv6ServerOptionsVendorInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptionsVendorInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptionsVendorInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptionsVendorInfo) Marshal() marshalDhcpv6ServerOptionsVendorInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptionsVendorInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptionsVendorInfo) Unmarshal() unMarshalDhcpv6ServerOptionsVendorInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptionsVendorInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptionsVendorInfo) ToProto() (*otg.Dhcpv6ServerOptionsVendorInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptionsVendorInfo) FromProto(msg *otg.Dhcpv6ServerOptionsVendorInfo) (Dhcpv6ServerOptionsVendorInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptionsVendorInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsVendorInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerOptionsVendorInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsVendorInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerOptionsVendorInfo) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerOptionsVendorInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsVendorInfo) FromJson(value string) error {
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

func (obj *dhcpv6ServerOptionsVendorInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsVendorInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsVendorInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptionsVendorInfo) Clone() (Dhcpv6ServerOptionsVendorInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptionsVendorInfo()
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

func (obj *dhcpv6ServerOptionsVendorInfo) setNil() {
	obj.optionDataHolder = nil
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerOptionsVendorInfo is this option is used by servers to exchange vendor-specific information. The option code is 17.
type Dhcpv6ServerOptionsVendorInfo interface {
	Validation
	// msg marshals Dhcpv6ServerOptionsVendorInfo to protobuf object *otg.Dhcpv6ServerOptionsVendorInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptionsVendorInfo
	// setMsg unmarshals Dhcpv6ServerOptionsVendorInfo from protobuf object *otg.Dhcpv6ServerOptionsVendorInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptionsVendorInfo
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptionsVendorInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptionsVendorInfo
	// validate validates Dhcpv6ServerOptionsVendorInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptionsVendorInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnterpriseNumber returns uint32, set in Dhcpv6ServerOptionsVendorInfo.
	EnterpriseNumber() uint32
	// SetEnterpriseNumber assigns uint32 provided by user to Dhcpv6ServerOptionsVendorInfo
	SetEnterpriseNumber(value uint32) Dhcpv6ServerOptionsVendorInfo
	// OptionData returns Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIterIter, set in Dhcpv6ServerOptionsVendorInfo
	OptionData() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	// AssociatedDhcpMessages returns Dhcpv6ServerOptionsIncludedMessages, set in Dhcpv6ServerOptionsVendorInfo.
	// Dhcpv6ServerOptionsIncludedMessages is the dhcpv6 server messages where the option will be included. If all is selected the selected option will be added  in the all the Dhcpv6 server messages, else based on the selection in particular Dhcpv6 server messages the option  will be included.
	AssociatedDhcpMessages() Dhcpv6ServerOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ServerOptionsIncludedMessages provided by user to Dhcpv6ServerOptionsVendorInfo.
	// Dhcpv6ServerOptionsIncludedMessages is the dhcpv6 server messages where the option will be included. If all is selected the selected option will be added  in the all the Dhcpv6 server messages, else based on the selection in particular Dhcpv6 server messages the option  will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsVendorInfo
	setNil()
}

// The vendor's registered Enterprise Number as registered with IANA.
// EnterpriseNumber returns a uint32
func (obj *dhcpv6ServerOptionsVendorInfo) EnterpriseNumber() uint32 {

	return *obj.obj.EnterpriseNumber

}

// The vendor's registered Enterprise Number as registered with IANA.
// SetEnterpriseNumber sets the uint32 value in the Dhcpv6ServerOptionsVendorInfo object
func (obj *dhcpv6ServerOptionsVendorInfo) SetEnterpriseNumber(value uint32) Dhcpv6ServerOptionsVendorInfo {

	obj.obj.EnterpriseNumber = &value
	return obj
}

// An opaque object of octets,interpreted by vendor-specific code on the clients and servers.
// OptionData returns a []Dhcpv6OptionsVendorSpecificOptions
func (obj *dhcpv6ServerOptionsVendorInfo) OptionData() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(obj.obj.OptionData) == 0 {
		obj.obj.OptionData = []*otg.Dhcpv6OptionsVendorSpecificOptions{}
	}
	if obj.optionDataHolder == nil {
		obj.optionDataHolder = newDhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter(&obj.obj.OptionData).setMsg(obj)
	}
	return obj.optionDataHolder
}

type dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter struct {
	obj                                     *dhcpv6ServerOptionsVendorInfo
	dhcpv6OptionsVendorSpecificOptionsSlice []Dhcpv6OptionsVendorSpecificOptions
	fieldPtr                                *[]*otg.Dhcpv6OptionsVendorSpecificOptions
}

func newDhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter(ptr *[]*otg.Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	return &dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter{fieldPtr: ptr}
}

type Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter interface {
	setMsg(*dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Items() []Dhcpv6OptionsVendorSpecificOptions
	Add() Dhcpv6OptionsVendorSpecificOptions
	Append(items ...Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Set(index int, newObj Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Clear() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	clearHolderSlice() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	appendHolderSlice(item Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
}

func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) setMsg(msg *dhcpv6ServerOptionsVendorInfo) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6OptionsVendorSpecificOptions{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Items() []Dhcpv6OptionsVendorSpecificOptions {
	return obj.dhcpv6OptionsVendorSpecificOptionsSlice
}

func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Add() Dhcpv6OptionsVendorSpecificOptions {
	newObj := &otg.Dhcpv6OptionsVendorSpecificOptions{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6OptionsVendorSpecificOptions{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Append(items ...Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, item)
	}
	return obj
}

func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Set(index int, newObj Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6OptionsVendorSpecificOptionsSlice[index] = newObj
	return obj
}
func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Clear() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6OptionsVendorSpecificOptions{}
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = []Dhcpv6OptionsVendorSpecificOptions{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) clearHolderSlice() Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(obj.dhcpv6OptionsVendorSpecificOptionsSlice) > 0 {
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = []Dhcpv6OptionsVendorSpecificOptions{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) appendHolderSlice(item Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ServerOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, item)
	return obj
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ServerOptionsIncludedMessages
func (obj *dhcpv6ServerOptionsVendorInfo) AssociatedDhcpMessages() Dhcpv6ServerOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ServerOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ServerOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The list of dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ServerOptionsIncludedMessages value in the Dhcpv6ServerOptionsVendorInfo object
func (obj *dhcpv6ServerOptionsVendorInfo) SetAssociatedDhcpMessages(value Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsVendorInfo {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *dhcpv6ServerOptionsVendorInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EnterpriseNumber is required
	if obj.obj.EnterpriseNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EnterpriseNumber is required field on interface Dhcpv6ServerOptionsVendorInfo")
	}
	if obj.obj.EnterpriseNumber != nil {

		if *obj.obj.EnterpriseNumber > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ServerOptionsVendorInfo.EnterpriseNumber <= 4294967295 but Got %d", *obj.obj.EnterpriseNumber))
		}

	}

	if len(obj.obj.OptionData) != 0 {

		if set_default {
			obj.OptionData().clearHolderSlice()
			for _, item := range obj.obj.OptionData {
				obj.OptionData().appendHolderSlice(&dhcpv6OptionsVendorSpecificOptions{obj: item})
			}
		}
		for _, item := range obj.OptionData().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	// AssociatedDhcpMessages is required
	if obj.obj.AssociatedDhcpMessages == nil {
		vObj.validationErrors = append(vObj.validationErrors, "AssociatedDhcpMessages is required field on interface Dhcpv6ServerOptionsVendorInfo")
	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ServerOptionsVendorInfo) setDefault() {

}
