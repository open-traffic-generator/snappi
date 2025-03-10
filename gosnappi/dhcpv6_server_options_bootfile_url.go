package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ServerOptionsBootfileUrl *****
type dhcpv6ServerOptionsBootfileUrl struct {
	validation
	obj                          *otg.Dhcpv6ServerOptionsBootfileUrl
	marshaller                   marshalDhcpv6ServerOptionsBootfileUrl
	unMarshaller                 unMarshalDhcpv6ServerOptionsBootfileUrl
	bootfileParamsHolder         Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	associatedDhcpMessagesHolder Dhcpv6ServerOptionsIncludedMessages
}

func NewDhcpv6ServerOptionsBootfileUrl() Dhcpv6ServerOptionsBootfileUrl {
	obj := dhcpv6ServerOptionsBootfileUrl{obj: &otg.Dhcpv6ServerOptionsBootfileUrl{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ServerOptionsBootfileUrl) msg() *otg.Dhcpv6ServerOptionsBootfileUrl {
	return obj.obj
}

func (obj *dhcpv6ServerOptionsBootfileUrl) setMsg(msg *otg.Dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptionsBootfileUrl {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ServerOptionsBootfileUrl struct {
	obj *dhcpv6ServerOptionsBootfileUrl
}

type marshalDhcpv6ServerOptionsBootfileUrl interface {
	// ToProto marshals Dhcpv6ServerOptionsBootfileUrl to protobuf object *otg.Dhcpv6ServerOptionsBootfileUrl
	ToProto() (*otg.Dhcpv6ServerOptionsBootfileUrl, error)
	// ToPbText marshals Dhcpv6ServerOptionsBootfileUrl to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ServerOptionsBootfileUrl to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ServerOptionsBootfileUrl to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Dhcpv6ServerOptionsBootfileUrl to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshaldhcpv6ServerOptionsBootfileUrl struct {
	obj *dhcpv6ServerOptionsBootfileUrl
}

type unMarshalDhcpv6ServerOptionsBootfileUrl interface {
	// FromProto unmarshals Dhcpv6ServerOptionsBootfileUrl from protobuf object *otg.Dhcpv6ServerOptionsBootfileUrl
	FromProto(msg *otg.Dhcpv6ServerOptionsBootfileUrl) (Dhcpv6ServerOptionsBootfileUrl, error)
	// FromPbText unmarshals Dhcpv6ServerOptionsBootfileUrl from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ServerOptionsBootfileUrl from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ServerOptionsBootfileUrl from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ServerOptionsBootfileUrl) Marshal() marshalDhcpv6ServerOptionsBootfileUrl {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ServerOptionsBootfileUrl{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ServerOptionsBootfileUrl) Unmarshal() unMarshalDhcpv6ServerOptionsBootfileUrl {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ServerOptionsBootfileUrl{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ServerOptionsBootfileUrl) ToProto() (*otg.Dhcpv6ServerOptionsBootfileUrl, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ServerOptionsBootfileUrl) FromProto(msg *otg.Dhcpv6ServerOptionsBootfileUrl) (Dhcpv6ServerOptionsBootfileUrl, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ServerOptionsBootfileUrl) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootfileUrl) FromPbText(value string) error {
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

func (m *marshaldhcpv6ServerOptionsBootfileUrl) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootfileUrl) FromYaml(value string) error {
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

func (m *marshaldhcpv6ServerOptionsBootfileUrl) ToJsonRaw() (string, error) {
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

func (m *marshaldhcpv6ServerOptionsBootfileUrl) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ServerOptionsBootfileUrl) FromJson(value string) error {
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

func (obj *dhcpv6ServerOptionsBootfileUrl) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsBootfileUrl) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ServerOptionsBootfileUrl) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ServerOptionsBootfileUrl) Clone() (Dhcpv6ServerOptionsBootfileUrl, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ServerOptionsBootfileUrl()
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

func (obj *dhcpv6ServerOptionsBootfileUrl) setNil() {
	obj.bootfileParamsHolder = nil
	obj.associatedDhcpMessagesHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ServerOptionsBootfileUrl is the server sends this option to inform the client about a URL to a boot file. This information is required for booting  over the network includes the details about the server on which the boot files can be found, the protocol to be used for  the download (for example,HTTP or TFTP, and the path and name of the boot file on the server. The option code is 59. The URL will contain the network communication protocol, a subdomain, a domain name, and its extension. If the host in the URL  is expressed using an IPv6 address rather than a domain name, the address in the URL then must be enclosed in "[" and "]"  characters, conforming to [RFC3986]. Eg of a boot file url can be "tftp://[xxxx:xxxx:xxxx:xxxx::xxxx]/mboot.efi".
type Dhcpv6ServerOptionsBootfileUrl interface {
	Validation
	// msg marshals Dhcpv6ServerOptionsBootfileUrl to protobuf object *otg.Dhcpv6ServerOptionsBootfileUrl
	// and doesn't set defaults
	msg() *otg.Dhcpv6ServerOptionsBootfileUrl
	// setMsg unmarshals Dhcpv6ServerOptionsBootfileUrl from protobuf object *otg.Dhcpv6ServerOptionsBootfileUrl
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptionsBootfileUrl
	// provides marshal interface
	Marshal() marshalDhcpv6ServerOptionsBootfileUrl
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ServerOptionsBootfileUrl
	// validate validates Dhcpv6ServerOptionsBootfileUrl
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ServerOptionsBootfileUrl, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Url returns string, set in Dhcpv6ServerOptionsBootfileUrl.
	Url() string
	// SetUrl assigns string provided by user to Dhcpv6ServerOptionsBootfileUrl
	SetUrl(value string) Dhcpv6ServerOptionsBootfileUrl
	// BootfileParams returns Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIterIter, set in Dhcpv6ServerOptionsBootfileUrl
	BootfileParams() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	// AssociatedDhcpMessages returns Dhcpv6ServerOptionsIncludedMessages, set in Dhcpv6ServerOptionsBootfileUrl.
	// Dhcpv6ServerOptionsIncludedMessages is the dhcpv6 server messages where the option will be included. If all is selected the selected option will be added  in the all the Dhcpv6 server messages, else based on the selection in particular Dhcpv6 server messages the option  will be included.
	AssociatedDhcpMessages() Dhcpv6ServerOptionsIncludedMessages
	// SetAssociatedDhcpMessages assigns Dhcpv6ServerOptionsIncludedMessages provided by user to Dhcpv6ServerOptionsBootfileUrl.
	// Dhcpv6ServerOptionsIncludedMessages is the dhcpv6 server messages where the option will be included. If all is selected the selected option will be added  in the all the Dhcpv6 server messages, else based on the selection in particular Dhcpv6 server messages the option  will be included.
	SetAssociatedDhcpMessages(value Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsBootfileUrl
	// HasAssociatedDhcpMessages checks if AssociatedDhcpMessages has been set in Dhcpv6ServerOptionsBootfileUrl
	HasAssociatedDhcpMessages() bool
	setNil()
}

// The URL for the boot file. It must comply with STD 66 format.
// Url returns a string
func (obj *dhcpv6ServerOptionsBootfileUrl) Url() string {

	return *obj.obj.Url

}

// The URL for the boot file. It must comply with STD 66 format.
// SetUrl sets the string value in the Dhcpv6ServerOptionsBootfileUrl object
func (obj *dhcpv6ServerOptionsBootfileUrl) SetUrl(value string) Dhcpv6ServerOptionsBootfileUrl {

	obj.obj.Url = &value
	return obj
}

// They are used to specify parameters for the boot file (similar to the command line arguments in most modern operating  systems). For example, these parameters could be used to specify the root file system of the OS kernel, or the location  from which a second-stage boot-loader program can download its configuration file.
// BootfileParams returns a []Dhcpv6ServerOptionsBootFileParams
func (obj *dhcpv6ServerOptionsBootfileUrl) BootfileParams() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	if len(obj.obj.BootfileParams) == 0 {
		obj.obj.BootfileParams = []*otg.Dhcpv6ServerOptionsBootFileParams{}
	}
	if obj.bootfileParamsHolder == nil {
		obj.bootfileParamsHolder = newDhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter(&obj.obj.BootfileParams).setMsg(obj)
	}
	return obj.bootfileParamsHolder
}

type dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter struct {
	obj                                    *dhcpv6ServerOptionsBootfileUrl
	dhcpv6ServerOptionsBootFileParamsSlice []Dhcpv6ServerOptionsBootFileParams
	fieldPtr                               *[]*otg.Dhcpv6ServerOptionsBootFileParams
}

func newDhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter(ptr *[]*otg.Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	return &dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter{fieldPtr: ptr}
}

type Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter interface {
	setMsg(*dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	Items() []Dhcpv6ServerOptionsBootFileParams
	Add() Dhcpv6ServerOptionsBootFileParams
	Append(items ...Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	Set(index int, newObj Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	Clear() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	clearHolderSlice() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
	appendHolderSlice(item Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter
}

func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) setMsg(msg *dhcpv6ServerOptionsBootfileUrl) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&dhcpv6ServerOptionsBootFileParams{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) Items() []Dhcpv6ServerOptionsBootFileParams {
	return obj.dhcpv6ServerOptionsBootFileParamsSlice
}

func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) Add() Dhcpv6ServerOptionsBootFileParams {
	newObj := &otg.Dhcpv6ServerOptionsBootFileParams{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &dhcpv6ServerOptionsBootFileParams{obj: newObj}
	newLibObj.setDefault()
	obj.dhcpv6ServerOptionsBootFileParamsSlice = append(obj.dhcpv6ServerOptionsBootFileParamsSlice, newLibObj)
	return newLibObj
}

func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) Append(items ...Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.dhcpv6ServerOptionsBootFileParamsSlice = append(obj.dhcpv6ServerOptionsBootFileParamsSlice, item)
	}
	return obj
}

func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) Set(index int, newObj Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.dhcpv6ServerOptionsBootFileParamsSlice[index] = newObj
	return obj
}
func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) Clear() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.Dhcpv6ServerOptionsBootFileParams{}
		obj.dhcpv6ServerOptionsBootFileParamsSlice = []Dhcpv6ServerOptionsBootFileParams{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) clearHolderSlice() Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	if len(obj.dhcpv6ServerOptionsBootFileParamsSlice) > 0 {
		obj.dhcpv6ServerOptionsBootFileParamsSlice = []Dhcpv6ServerOptionsBootFileParams{}
	}
	return obj
}
func (obj *dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter) appendHolderSlice(item Dhcpv6ServerOptionsBootFileParams) Dhcpv6ServerOptionsBootfileUrlDhcpv6ServerOptionsBootFileParamsIter {
	obj.dhcpv6ServerOptionsBootFileParamsSlice = append(obj.dhcpv6ServerOptionsBootFileParamsSlice, item)
	return obj
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ServerOptionsIncludedMessages
func (obj *dhcpv6ServerOptionsBootfileUrl) AssociatedDhcpMessages() Dhcpv6ServerOptionsIncludedMessages {
	if obj.obj.AssociatedDhcpMessages == nil {
		obj.obj.AssociatedDhcpMessages = NewDhcpv6ServerOptionsIncludedMessages().msg()
	}
	if obj.associatedDhcpMessagesHolder == nil {
		obj.associatedDhcpMessagesHolder = &dhcpv6ServerOptionsIncludedMessages{obj: obj.obj.AssociatedDhcpMessages}
	}
	return obj.associatedDhcpMessagesHolder
}

// The list of dhcpv6 client messages where this option is included.
// AssociatedDhcpMessages returns a Dhcpv6ServerOptionsIncludedMessages
func (obj *dhcpv6ServerOptionsBootfileUrl) HasAssociatedDhcpMessages() bool {
	return obj.obj.AssociatedDhcpMessages != nil
}

// The list of dhcpv6 client messages where this option is included.
// SetAssociatedDhcpMessages sets the Dhcpv6ServerOptionsIncludedMessages value in the Dhcpv6ServerOptionsBootfileUrl object
func (obj *dhcpv6ServerOptionsBootfileUrl) SetAssociatedDhcpMessages(value Dhcpv6ServerOptionsIncludedMessages) Dhcpv6ServerOptionsBootfileUrl {

	obj.associatedDhcpMessagesHolder = nil
	obj.obj.AssociatedDhcpMessages = value.msg()

	return obj
}

func (obj *dhcpv6ServerOptionsBootfileUrl) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Url is required
	if obj.obj.Url == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Url is required field on interface Dhcpv6ServerOptionsBootfileUrl")
	}

	if len(obj.obj.BootfileParams) != 0 {

		if set_default {
			obj.BootfileParams().clearHolderSlice()
			for _, item := range obj.obj.BootfileParams {
				obj.BootfileParams().appendHolderSlice(&dhcpv6ServerOptionsBootFileParams{obj: item})
			}
		}
		for _, item := range obj.BootfileParams().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if obj.obj.AssociatedDhcpMessages != nil {

		obj.AssociatedDhcpMessages().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ServerOptionsBootfileUrl) setDefault() {

}
