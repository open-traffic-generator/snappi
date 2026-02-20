package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Dhcpv6ClientOptionsOptionsRequest *****
type dhcpv6ClientOptionsOptionsRequest struct {
	validation
	obj          *otg.Dhcpv6ClientOptionsOptionsRequest
	marshaller   marshalDhcpv6ClientOptionsOptionsRequest
	unMarshaller unMarshalDhcpv6ClientOptionsOptionsRequest
	customHolder Dhcpv6ClientOptionsCustom
}

func NewDhcpv6ClientOptionsOptionsRequest() Dhcpv6ClientOptionsOptionsRequest {
	obj := dhcpv6ClientOptionsOptionsRequest{obj: &otg.Dhcpv6ClientOptionsOptionsRequest{}}
	obj.setDefault()
	return &obj
}

func (obj *dhcpv6ClientOptionsOptionsRequest) msg() *otg.Dhcpv6ClientOptionsOptionsRequest {
	return obj.obj
}

func (obj *dhcpv6ClientOptionsOptionsRequest) setMsg(msg *otg.Dhcpv6ClientOptionsOptionsRequest) Dhcpv6ClientOptionsOptionsRequest {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldhcpv6ClientOptionsOptionsRequest struct {
	obj *dhcpv6ClientOptionsOptionsRequest
}

type marshalDhcpv6ClientOptionsOptionsRequest interface {
	// ToProto marshals Dhcpv6ClientOptionsOptionsRequest to protobuf object *otg.Dhcpv6ClientOptionsOptionsRequest
	ToProto() (*otg.Dhcpv6ClientOptionsOptionsRequest, error)
	// ToPbText marshals Dhcpv6ClientOptionsOptionsRequest to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Dhcpv6ClientOptionsOptionsRequest to YAML text
	ToYaml() (string, error)
	// ToJson marshals Dhcpv6ClientOptionsOptionsRequest to JSON text
	ToJson() (string, error)
}

type unMarshaldhcpv6ClientOptionsOptionsRequest struct {
	obj *dhcpv6ClientOptionsOptionsRequest
}

type unMarshalDhcpv6ClientOptionsOptionsRequest interface {
	// FromProto unmarshals Dhcpv6ClientOptionsOptionsRequest from protobuf object *otg.Dhcpv6ClientOptionsOptionsRequest
	FromProto(msg *otg.Dhcpv6ClientOptionsOptionsRequest) (Dhcpv6ClientOptionsOptionsRequest, error)
	// FromPbText unmarshals Dhcpv6ClientOptionsOptionsRequest from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Dhcpv6ClientOptionsOptionsRequest from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Dhcpv6ClientOptionsOptionsRequest from JSON text
	FromJson(value string) error
}

func (obj *dhcpv6ClientOptionsOptionsRequest) Marshal() marshalDhcpv6ClientOptionsOptionsRequest {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldhcpv6ClientOptionsOptionsRequest{obj: obj}
	}
	return obj.marshaller
}

func (obj *dhcpv6ClientOptionsOptionsRequest) Unmarshal() unMarshalDhcpv6ClientOptionsOptionsRequest {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldhcpv6ClientOptionsOptionsRequest{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldhcpv6ClientOptionsOptionsRequest) ToProto() (*otg.Dhcpv6ClientOptionsOptionsRequest, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldhcpv6ClientOptionsOptionsRequest) FromProto(msg *otg.Dhcpv6ClientOptionsOptionsRequest) (Dhcpv6ClientOptionsOptionsRequest, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldhcpv6ClientOptionsOptionsRequest) ToPbText() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequest) FromPbText(value string) error {
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

func (m *marshaldhcpv6ClientOptionsOptionsRequest) ToYaml() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequest) FromYaml(value string) error {
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

func (m *marshaldhcpv6ClientOptionsOptionsRequest) ToJson() (string, error) {
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

func (m *unMarshaldhcpv6ClientOptionsOptionsRequest) FromJson(value string) error {
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

func (obj *dhcpv6ClientOptionsOptionsRequest) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsOptionsRequest) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *dhcpv6ClientOptionsOptionsRequest) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *dhcpv6ClientOptionsOptionsRequest) Clone() (Dhcpv6ClientOptionsOptionsRequest, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDhcpv6ClientOptionsOptionsRequest()
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

func (obj *dhcpv6ClientOptionsOptionsRequest) setNil() {
	obj.customHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Dhcpv6ClientOptionsOptionsRequest is description is TBD
type Dhcpv6ClientOptionsOptionsRequest interface {
	Validation
	// msg marshals Dhcpv6ClientOptionsOptionsRequest to protobuf object *otg.Dhcpv6ClientOptionsOptionsRequest
	// and doesn't set defaults
	msg() *otg.Dhcpv6ClientOptionsOptionsRequest
	// setMsg unmarshals Dhcpv6ClientOptionsOptionsRequest from protobuf object *otg.Dhcpv6ClientOptionsOptionsRequest
	// and doesn't set defaults
	setMsg(*otg.Dhcpv6ClientOptionsOptionsRequest) Dhcpv6ClientOptionsOptionsRequest
	// provides marshal interface
	Marshal() marshalDhcpv6ClientOptionsOptionsRequest
	// provides unmarshal interface
	Unmarshal() unMarshalDhcpv6ClientOptionsOptionsRequest
	// validate validates Dhcpv6ClientOptionsOptionsRequest
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Dhcpv6ClientOptionsOptionsRequest, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Dhcpv6ClientOptionsOptionsRequestChoiceEnum, set in Dhcpv6ClientOptionsOptionsRequest
	Choice() Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	// setChoice assigns Dhcpv6ClientOptionsOptionsRequestChoiceEnum provided by user to Dhcpv6ClientOptionsOptionsRequest
	setChoice(value Dhcpv6ClientOptionsOptionsRequestChoiceEnum) Dhcpv6ClientOptionsOptionsRequest
	// HasChoice checks if Choice has been set in Dhcpv6ClientOptionsOptionsRequest
	HasChoice() bool
	// getter for NameServers to set choice.
	NameServers()
	// getter for Fqdn to set choice.
	Fqdn()
	// getter for BootfileUrl to set choice.
	BootfileUrl()
	// getter for Sztp to set choice.
	Sztp()
	// getter for VendorInformation to set choice.
	VendorInformation()
	// Custom returns Dhcpv6ClientOptionsCustom, set in Dhcpv6ClientOptionsOptionsRequest.
	// Dhcpv6ClientOptionsCustom is the Custom option is used to provide a not so well known option in the message between a client and a server.
	Custom() Dhcpv6ClientOptionsCustom
	// SetCustom assigns Dhcpv6ClientOptionsCustom provided by user to Dhcpv6ClientOptionsOptionsRequest.
	// Dhcpv6ClientOptionsCustom is the Custom option is used to provide a not so well known option in the message between a client and a server.
	SetCustom(value Dhcpv6ClientOptionsCustom) Dhcpv6ClientOptionsOptionsRequest
	// HasCustom checks if Custom has been set in Dhcpv6ClientOptionsOptionsRequest
	HasCustom() bool
	setNil()
}

type Dhcpv6ClientOptionsOptionsRequestChoiceEnum string

// Enum of Choice on Dhcpv6ClientOptionsOptionsRequest
var Dhcpv6ClientOptionsOptionsRequestChoice = struct {
	VENDOR_INFORMATION Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	NAME_SERVERS       Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	FQDN               Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	BOOTFILE_URL       Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	SZTP               Dhcpv6ClientOptionsOptionsRequestChoiceEnum
	CUSTOM             Dhcpv6ClientOptionsOptionsRequestChoiceEnum
}{
	VENDOR_INFORMATION: Dhcpv6ClientOptionsOptionsRequestChoiceEnum("vendor_information"),
	NAME_SERVERS:       Dhcpv6ClientOptionsOptionsRequestChoiceEnum("name_servers"),
	FQDN:               Dhcpv6ClientOptionsOptionsRequestChoiceEnum("fqdn"),
	BOOTFILE_URL:       Dhcpv6ClientOptionsOptionsRequestChoiceEnum("bootfile_url"),
	SZTP:               Dhcpv6ClientOptionsOptionsRequestChoiceEnum("sztp"),
	CUSTOM:             Dhcpv6ClientOptionsOptionsRequestChoiceEnum("custom"),
}

func (obj *dhcpv6ClientOptionsOptionsRequest) Choice() Dhcpv6ClientOptionsOptionsRequestChoiceEnum {
	return Dhcpv6ClientOptionsOptionsRequestChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for NameServers to set choice
func (obj *dhcpv6ClientOptionsOptionsRequest) NameServers() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.NAME_SERVERS)
}

// getter for Fqdn to set choice
func (obj *dhcpv6ClientOptionsOptionsRequest) Fqdn() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.FQDN)
}

// getter for BootfileUrl to set choice
func (obj *dhcpv6ClientOptionsOptionsRequest) BootfileUrl() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.BOOTFILE_URL)
}

// getter for Sztp to set choice
func (obj *dhcpv6ClientOptionsOptionsRequest) Sztp() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.SZTP)
}

// getter for VendorInformation to set choice
func (obj *dhcpv6ClientOptionsOptionsRequest) VendorInformation() {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.VENDOR_INFORMATION)
}

// The Option Request option is used to identify a list of options in a message between a client and a server. The  option code is 6. - Vendor_specific information option, requested by clients for vendor-specific informations from servers. - DNS Recursive Name Server Option, requested by clients to get the list ofIPv6 addresses of DNS recursive name
// servers to which DNS queries may be sent by the client resolver in order of preference.
// - Client FQDN option - indicates whether the client or the DHCP server should update DNS with the AAAA record
// corresponding to the assigned IPv6 address and the FQDN provided in this option. The DHCP server always updates
// the PTR record.
// - bootfile_url, if client is configured for network booting then the client must use this option to obtain the boot
// file url from the server.
// - sztp. Securely provision a networking device when it is booting in a factory-default state.
// Choice returns a string
func (obj *dhcpv6ClientOptionsOptionsRequest) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *dhcpv6ClientOptionsOptionsRequest) setChoice(value Dhcpv6ClientOptionsOptionsRequestChoiceEnum) Dhcpv6ClientOptionsOptionsRequest {
	intValue, ok := otg.Dhcpv6ClientOptionsOptionsRequest_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Dhcpv6ClientOptionsOptionsRequestChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Dhcpv6ClientOptionsOptionsRequest_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Custom = nil
	obj.customHolder = nil

	if value == Dhcpv6ClientOptionsOptionsRequestChoice.CUSTOM {
		obj.obj.Custom = NewDhcpv6ClientOptionsCustom().msg()
	}

	return obj
}

// description is TBD
// Custom returns a Dhcpv6ClientOptionsCustom
func (obj *dhcpv6ClientOptionsOptionsRequest) Custom() Dhcpv6ClientOptionsCustom {
	if obj.obj.Custom == nil {
		obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.CUSTOM)
	}
	if obj.customHolder == nil {
		obj.customHolder = &dhcpv6ClientOptionsCustom{obj: obj.obj.Custom}
	}
	return obj.customHolder
}

// description is TBD
// Custom returns a Dhcpv6ClientOptionsCustom
func (obj *dhcpv6ClientOptionsOptionsRequest) HasCustom() bool {
	return obj.obj.Custom != nil
}

// description is TBD
// SetCustom sets the Dhcpv6ClientOptionsCustom value in the Dhcpv6ClientOptionsOptionsRequest object
func (obj *dhcpv6ClientOptionsOptionsRequest) SetCustom(value Dhcpv6ClientOptionsCustom) Dhcpv6ClientOptionsOptionsRequest {
	obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.CUSTOM)
	obj.customHolder = nil
	obj.obj.Custom = value.msg()

	return obj
}

func (obj *dhcpv6ClientOptionsOptionsRequest) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Custom != nil {

		obj.Custom().validateObj(vObj, set_default)
	}

}

func (obj *dhcpv6ClientOptionsOptionsRequest) setDefault() {
	var choices_set int = 0
	var choice Dhcpv6ClientOptionsOptionsRequestChoiceEnum

	if obj.obj.Custom != nil {
		choices_set += 1
		choice = Dhcpv6ClientOptionsOptionsRequestChoice.CUSTOM
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Dhcpv6ClientOptionsOptionsRequestChoice.VENDOR_INFORMATION)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Dhcpv6ClientOptionsOptionsRequest")
			}
		} else {
			intVal := otg.Dhcpv6ClientOptionsOptionsRequest_Choice_Enum_value[string(choice)]
			enumValue := otg.Dhcpv6ClientOptionsOptionsRequest_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
