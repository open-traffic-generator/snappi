package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsVendorInfo *****
type dhcpv6ClientOptionsVendorInfo struct {
	validation
	obj                          *otg.Dhcpv6ClientOptionsVendorInfo
	marshaller                   marshalDhcpv6ClientOptionsVendorInfo
	unMarshaller                 unMarshalDhcpv6ClientOptionsVendorInfo
	optionDataHolder             Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	associatedDhcpMessagesHolder Dhcpv6ClientOptionsIncludedMessages
}

func NewDhcpv6ClientOptionsVendorInfo() Dhcpv6ClientOptionsVendorInfo {
	obj := dhcpv6ClientOptionsVendorInfo{obj: &otg.Dhcpv6ClientOptionsVendorInfo{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsVendorInfo) msg() *otg.Dhcpv6ClientOptionsVendorInfo {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsVendorInfo) setMsg(msg *otg.Dhcpv6ClientOptionsVendorInfo) Dhcpv6ClientOptionsVendorInfo {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsVendorInfo struct {
	obj *dhcpv6ClientOptionsVendorInfo
}

type marshalDhcpv6ClientOptionsVendorInfo interface {
	// ToProto marshals Dhcpv6ClientOptionsVendorInfo to protobuf object *otg.Dhcpv6ClientOptionsVendorInfo
	ToProto() (*otg.Dhcpv6ClientOptionsVendorInfo, error)
	// ToPbText marshals Dhcpv6ClientOptionsVendorInfo to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsVendorInfo to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsVendorInfo to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ClientOptionsVendorInfo to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ClientOptionsVendorInfo struct {
	obj *dhcpv6ClientOptionsVendorInfo
}

type unMarshalDhcpv6ClientOptionsVendorInfo interface {
	// FromProto unmarshals Dhcpv6ClientOptionsVendorInfo from protobuf object *otg.Dhcpv6ClientOptionsVendorInfo
	FromProto(msg *otg.Dhcpv6ClientOptionsVendorInfo) (Dhcpv6ClientOptionsVendorInfo, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsVendorInfo from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsVendorInfo from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsVendorInfo from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsVendorInfo) Marshal() marshalDhcpv6ClientOptionsVendorInfo {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsVendorInfo{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsVendorInfo) Unmarshal() unMarshalDhcpv6ClientOptionsVendorInfo {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsVendorInfo{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsVendorInfo) ToProto() (*otg.Dhcpv6ClientOptionsVendorInfo, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsVendorInfo) FromProto(msg *otg.Dhcpv6ClientOptionsVendorInfo) (Dhcpv6ClientOptionsVendorInfo, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsVendorInfo) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorInfo) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsVendorInfo) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorInfo) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsVendorInfo) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ClientOptionsVendorInfo) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsVendorInfo) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsVendorInfo) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsVendorInfo) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsVendorInfo) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsVendorInfo) Clone() (Dhcpv6ClientOptionsVendorInfo, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsVendorInfo()
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

func (obj *dhcpv6ClientOptionsVendorInfo) setNil() {
	obj.optionDataHolder = nil
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsVendorInfo is this option is used by clients to exchange vendor-specific information. The option code is 17.
type Dhcpv6ClientOptionsVendorInfo interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsVendorInfo to protobuf object *otg.Dhcpv6ClientOptionsVendorInfo
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsVendorInfo
	// setMsg unmarshals Dhcpv6ClientOptionsVendorInfo from protobuf object *otg.Dhcpv6ClientOptionsVendorInfo
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsVendorInfo) Dhcpv6ClientOptionsVendorInfo
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsVendorInfo
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsVendorInfo
	// validate validates Dhcpv6ClientOptionsVendorInfo
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsVendorInfo, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// EnterpriseNumber returns uint32, set in Dhcpv6ClientOptionsVendorInfo.
	EnterpriseNumber() uint32
	// SetEnterpriseNumber assigns uint32 provided by user to Dhcpv6ClientOptionsVendorInfo
	SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsVendorInfo
	// OptionData returns Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIterIter, set in Dhcpv6ClientOptionsVendorInfo
	OptionData() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	// AssociatedDhcpMessages returns Dhcpv6ClientOptionsIncludedMessages, set in Dhcpv6ClientOptionsVendorInfo.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ClientOptionsIncludedMessages provided by user to Dhcpv6ClientOptionsVendorInfo.
	// Dhcpv6ClientOptionsIncludedMessages is the dhcpv6 client messages where the option will be included. If all is selected the selected option will be added in the all the Dhcpv6 client messages, else based on the selection in particular Dhcpv6 client messages the option will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsVendorInfo
	setNil()
}

// The vendor's registered Enterprise Number as registered with IANA.
// EnterpriseNumber returns a uint32
func (obj *dhcpv6ClientOptionsVendorInfo) EnterpriseNumber() uint32 {

	return *obj.obj.EnterpriseNumber

}

// The vendor's registered Enterprise Number as registered with IANA.
// SetEnterpriseNumber sets the uint32 value in the Dhcpv6ClientOptionsVendorInfo object
func (obj *dhcpv6ClientOptionsVendorInfo) SetEnterpriseNumber(value uint32) Dhcpv6ClientOptionsVendorInfo {

	obj.obj.EnterpriseNumber = &value
	return obj
}

// An opaque object of octets,interpreted by vendor-specific code on the clients and servers.
// OptionData returns a []Dhcpv6OptionsVendorSpecificOptions
func (obj *dhcpv6ClientOptionsVendorInfo) OptionData() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(obj.obj.OptionData) == 0 {
		obj.obj.OptionData = []*otg.Dhcpv6OptionsVendorSpecificOptions{}
	}
	if obj.optionDataHolder == nil {
		obj.optionDataHolder = newDhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter(&obj.obj.OptionData).setMsg(obj)
	}
	return obj.optionDataHolder
}

type dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter struct {
	obj                                     *dhcpv6ClientOptionsVendorInfo
	dhcpv6OptionsVendorSpecificOptionsSlice []Dhcpv6OptionsVendorSpecificOptions
	fieldPtr                                *[]*otg.Dhcpv6OptionsVendorSpecificOptions
}

func newDhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter(ptr *[]*otg.Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	return &dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter{fieldPtr: ptr}
}

type Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter interface {
	setMsg(*dhcpv6ClientOptionsVendorInfo) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Items() []Dhcpv6OptionsVendorSpecificOptions
	Add() Dhcpv6OptionsVendorSpecificOptions
	Append(items ...Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Set(index int, newObj Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	Clear() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	clearHolderSlice() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
	appendHolderSlice(item Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter
}

func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) setMsg(msg *dhcpv6ClientOptionsVendorInfo) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6OptionsVendorSpecificOptions{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Items() []Dhcpv6OptionsVendorSpecificOptions {
	return obj.dhcpv6OptionsVendorSpecificOptionsSlice
}

func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Add() Dhcpv6OptionsVendorSpecificOptions {
	newObj := &otg.Dhcpv6OptionsVendorSpecificOptions{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6OptionsVendorSpecificOptions{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Append(items ...Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, item)
	}
	return obj
}

func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Set(index int, newObj Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6OptionsVendorSpecificOptionsSlice[index] = newObj
	return obj
}
func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) Clear() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6OptionsVendorSpecificOptions{}
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = []Dhcpv6OptionsVendorSpecificOptions{}
	}
	return obj
}
func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) clearHolderSlice() Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	if len(obj.dhcpv6OptionsVendorSpecificOptionsSlice) > 0 {
		obj.dhcpv6OptionsVendorSpecificOptionsSlice = []Dhcpv6OptionsVendorSpecificOptions{}
	}
	return obj
}
func (obj *dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter) appendHolderSlice(item Dhcpv6OptionsVendorSpecificOptions) Dhcpv6ClientOptionsVendorInfoDhcpv6OptionsVendorSpecificOptionsIter {
	obj.dhcpv6OptionsVendorSpecificOptionsSlice = append(obj.dhcpv6OptionsVendorSpecificOptionsSlice, item)
	return obj
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ClientOptionsIncludedMessages
func (obj *dhcpv6ClientOptionsVendorInfo) AssociatedDhcpMessages() Dhcpv6ClientOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ClientOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ClientOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The list of dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ClientOptionsIncludedMessages value in the Dhcpv6ClientOptionsVendorInfo object
func (obj *dhcpv6ClientOptionsVendorInfo) SetAssociatedDhcpMessages(value Dhcpv6ClientOptionsIncludedMessages) Dhcpv6ClientOptionsVendorInfo {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsVendorInfo) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// EnterpriseNumber is required
	if obj.obj.EnterpriseNumber == nil {
		vObj.validationErrors = append(vObj.validationErrors, "EnterpriseNumber is required field on interface Dhcpv6ClientOptionsVendorInfo")
	}
	if obj.obj.EnterpriseNumber != nil {

		if *obj.obj.EnterpriseNumber > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= Dhcpv6ClientOptionsVendorInfo.EnterpriseNumber <= 4294967295 but Got %d", *obj.obj.EnterpriseNumber))
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
		vObj.validationErrors = append(vObj.validationErrors, "AssociatedDhcpMessages is required field on interface Dhcpv6ClientOptionsVendorInfo")
	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsVendorInfo) setDefault() {

}
