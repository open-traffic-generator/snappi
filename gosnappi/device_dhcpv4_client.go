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

// ***** DeviceDhcpv4Client *****
type deviceDhcpv4Client struct {
	validation
	obj                         *otg.DeviceDhcpv4Client
	marshaller                  marshalDeviceDhcpv4Client
	unMarshaller                unMarshalDeviceDhcpv4Client
	parametersRequestListHolder Dhcpv4ClientParams
}

func NewDeviceDhcpv4Client() DeviceDhcpv4Client {
	obj := deviceDhcpv4Client{obj: &otg.DeviceDhcpv4Client{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceDhcpv4Client) msg() *otg.DeviceDhcpv4Client {
	return obj.obj
}

func (obj *deviceDhcpv4Client) setMsg(msg *otg.DeviceDhcpv4Client) DeviceDhcpv4Client {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceDhcpv4Client struct {
	obj *deviceDhcpv4Client
}

type marshalDeviceDhcpv4Client interface {
	// ToProto marshals DeviceDhcpv4Client to protobuf object *otg.DeviceDhcpv4Client
	ToProto() (*otg.DeviceDhcpv4Client, error)
	// ToPbText marshals DeviceDhcpv4Client to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceDhcpv4Client to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceDhcpv4Client to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceDhcpv4Client struct {
	obj *deviceDhcpv4Client
}

type unMarshalDeviceDhcpv4Client interface {
	// FromProto unmarshals DeviceDhcpv4Client from protobuf object *otg.DeviceDhcpv4Client
	FromProto(msg *otg.DeviceDhcpv4Client) (DeviceDhcpv4Client, error)
	// FromPbText unmarshals DeviceDhcpv4Client from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceDhcpv4Client from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceDhcpv4Client from JSON text
	FromJson(value string) error
}

func (obj *deviceDhcpv4Client) Marshal() marshalDeviceDhcpv4Client {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceDhcpv4Client{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceDhcpv4Client) Unmarshal() unMarshalDeviceDhcpv4Client {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceDhcpv4Client{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceDhcpv4Client) ToProto() (*otg.DeviceDhcpv4Client, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceDhcpv4Client) FromProto(msg *otg.DeviceDhcpv4Client) (DeviceDhcpv4Client, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceDhcpv4Client) ToPbText() (string, error) {
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

func (m *unMarshaldeviceDhcpv4Client) FromPbText(value string) error {
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

func (m *marshaldeviceDhcpv4Client) ToYaml() (string, error) {
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

func (m *unMarshaldeviceDhcpv4Client) FromYaml(value string) error {
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

func (m *marshaldeviceDhcpv4Client) ToJson() (string, error) {
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

func (m *unMarshaldeviceDhcpv4Client) FromJson(value string) error {
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

func (obj *deviceDhcpv4Client) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceDhcpv4Client) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceDhcpv4Client) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceDhcpv4Client) Clone() (DeviceDhcpv4Client, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceDhcpv4Client()
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

func (obj *deviceDhcpv4Client) setNil() {
	obj.parametersRequestListHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceDhcpv4Client is configuration for emulated DHCPv4 Client on a single Interface. https://www.rfc-editor.org/rfc/rfc2131.html
type DeviceDhcpv4Client interface {
	Validation
	// msg marshals DeviceDhcpv4Client to protobuf object *otg.DeviceDhcpv4Client
	// and doesn't set defaults
	msg() *otg.DeviceDhcpv4Client
	// setMsg unmarshals DeviceDhcpv4Client from protobuf object *otg.DeviceDhcpv4Client
	// and doesn't set defaults
	setMsg(*otg.DeviceDhcpv4Client) DeviceDhcpv4Client
	// provides marshal interface
	Marshal() marshalDeviceDhcpv4Client
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceDhcpv4Client
	// validate validates DeviceDhcpv4Client
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceDhcpv4Client, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in DeviceDhcpv4Client.
	Name() string
	// SetName assigns string provided by user to DeviceDhcpv4Client
	SetName(value string) DeviceDhcpv4Client
	// Choice returns DeviceDhcpv4ClientChoiceEnum, set in DeviceDhcpv4Client
	Choice() DeviceDhcpv4ClientChoiceEnum
	// setChoice assigns DeviceDhcpv4ClientChoiceEnum provided by user to DeviceDhcpv4Client
	setChoice(value DeviceDhcpv4ClientChoiceEnum) DeviceDhcpv4Client
	// HasChoice checks if Choice has been set in DeviceDhcpv4Client
	HasChoice() bool
	// getter for FirstServer to set choice.
	FirstServer()
	// ServerAddress returns string, set in DeviceDhcpv4Client.
	ServerAddress() string
	// SetServerAddress assigns string provided by user to DeviceDhcpv4Client
	SetServerAddress(value string) DeviceDhcpv4Client
	// HasServerAddress checks if ServerAddress has been set in DeviceDhcpv4Client
	HasServerAddress() bool
	// Broadcast returns bool, set in DeviceDhcpv4Client.
	Broadcast() bool
	// SetBroadcast assigns bool provided by user to DeviceDhcpv4Client
	SetBroadcast(value bool) DeviceDhcpv4Client
	// HasBroadcast checks if Broadcast has been set in DeviceDhcpv4Client
	HasBroadcast() bool
	// ParametersRequestList returns Dhcpv4ClientParams, set in DeviceDhcpv4Client.
	// Dhcpv4ClientParams is configuration Parameter request list by emulated DHCPv4 Client.
	ParametersRequestList() Dhcpv4ClientParams
	// SetParametersRequestList assigns Dhcpv4ClientParams provided by user to DeviceDhcpv4Client.
	// Dhcpv4ClientParams is configuration Parameter request list by emulated DHCPv4 Client.
	SetParametersRequestList(value Dhcpv4ClientParams) DeviceDhcpv4Client
	// HasParametersRequestList checks if ParametersRequestList has been set in DeviceDhcpv4Client
	HasParametersRequestList() bool
	setNil()
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *deviceDhcpv4Client) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the DeviceDhcpv4Client object
func (obj *deviceDhcpv4Client) SetName(value string) DeviceDhcpv4Client {

	obj.obj.Name = &value
	return obj
}

type DeviceDhcpv4ClientChoiceEnum string

// Enum of Choice on DeviceDhcpv4Client
var DeviceDhcpv4ClientChoice = struct {
	FIRST_SERVER   DeviceDhcpv4ClientChoiceEnum
	SERVER_ADDRESS DeviceDhcpv4ClientChoiceEnum
}{
	FIRST_SERVER:   DeviceDhcpv4ClientChoiceEnum("first_server"),
	SERVER_ADDRESS: DeviceDhcpv4ClientChoiceEnum("server_address"),
}

func (obj *deviceDhcpv4Client) Choice() DeviceDhcpv4ClientChoiceEnum {
	return DeviceDhcpv4ClientChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for FirstServer to set choice
func (obj *deviceDhcpv4Client) FirstServer() {
	obj.setChoice(DeviceDhcpv4ClientChoice.FIRST_SERVER)
}

// The client receives one or more DHCPOFFER messages from one or more servers and client may choose to wait for multiple responses.
// The client chooses one server from which to request configuration
// parameters, based on the configuration parameters offered in the DHCPOFFER messages.
// - first_server: if selected, the subnet accepts the IP addresses offered by the first server to respond with an offer of IP addresses.
// - server_address: The address of the DHCP server from which the subnet will accept IP addresses.
// Choice returns a string
func (obj *deviceDhcpv4Client) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceDhcpv4Client) setChoice(value DeviceDhcpv4ClientChoiceEnum) DeviceDhcpv4Client {
	intValue, ok := otg.DeviceDhcpv4Client_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceDhcpv4ClientChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceDhcpv4Client_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.ServerAddress = nil
	return obj
}

// The address of the DHCP server.
// ServerAddress returns a string
func (obj *deviceDhcpv4Client) ServerAddress() string {

	if obj.obj.ServerAddress == nil {
		obj.setChoice(DeviceDhcpv4ClientChoice.SERVER_ADDRESS)
	}

	return *obj.obj.ServerAddress

}

// The address of the DHCP server.
// ServerAddress returns a string
func (obj *deviceDhcpv4Client) HasServerAddress() bool {
	return obj.obj.ServerAddress != nil
}

// The address of the DHCP server.
// SetServerAddress sets the string value in the DeviceDhcpv4Client object
func (obj *deviceDhcpv4Client) SetServerAddress(value string) DeviceDhcpv4Client {
	obj.setChoice(DeviceDhcpv4ClientChoice.SERVER_ADDRESS)
	obj.obj.ServerAddress = &value
	return obj
}

// If the broadcast bit is set, then the server and relay agent broadcast DHCPOFFER and DHCPACK messages.
// Broadcast returns a bool
func (obj *deviceDhcpv4Client) Broadcast() bool {

	return *obj.obj.Broadcast

}

// If the broadcast bit is set, then the server and relay agent broadcast DHCPOFFER and DHCPACK messages.
// Broadcast returns a bool
func (obj *deviceDhcpv4Client) HasBroadcast() bool {
	return obj.obj.Broadcast != nil
}

// If the broadcast bit is set, then the server and relay agent broadcast DHCPOFFER and DHCPACK messages.
// SetBroadcast sets the bool value in the DeviceDhcpv4Client object
func (obj *deviceDhcpv4Client) SetBroadcast(value bool) DeviceDhcpv4Client {

	obj.obj.Broadcast = &value
	return obj
}

// Optional parameters field request list of DHCPv4 Client.
// ParametersRequestList returns a Dhcpv4ClientParams
func (obj *deviceDhcpv4Client) ParametersRequestList() Dhcpv4ClientParams {
	if obj.obj.ParametersRequestList == nil {
		obj.obj.ParametersRequestList = NewDhcpv4ClientParams().msg()
	}
	if obj.parametersRequestListHolder == nil {
		obj.parametersRequestListHolder = &dhcpv4ClientParams{obj: obj.obj.ParametersRequestList}
	}
	return obj.parametersRequestListHolder
}

// Optional parameters field request list of DHCPv4 Client.
// ParametersRequestList returns a Dhcpv4ClientParams
func (obj *deviceDhcpv4Client) HasParametersRequestList() bool {
	return obj.obj.ParametersRequestList != nil
}

// Optional parameters field request list of DHCPv4 Client.
// SetParametersRequestList sets the Dhcpv4ClientParams value in the DeviceDhcpv4Client object
func (obj *deviceDhcpv4Client) SetParametersRequestList(value Dhcpv4ClientParams) DeviceDhcpv4Client {

	obj.parametersRequestListHolder = nil
	obj.obj.ParametersRequestList = value.msg()

	return obj
}

func (obj *deviceDhcpv4Client) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface DeviceDhcpv4Client")
	}
	if obj.obj.Name != nil {

		if !regexp.MustCompile(`^[\sa-zA-Z0-9-_()><\[\]]+$`).MatchString(*obj.obj.Name) {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf(
					"DeviceDhcpv4Client.Name should adhere to this regex pattern '%s', but Got %s", `^[\sa-zA-Z0-9-_()><\[\]]+$`, *obj.obj.Name))
		}

	}

	if obj.obj.ServerAddress != nil {

		err := obj.validateIpv4(obj.ServerAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on DeviceDhcpv4Client.ServerAddress"))
		}

	}

	if obj.obj.ParametersRequestList != nil {

		obj.ParametersRequestList().validateObj(vObj, set_default)
	}

}

func (obj *deviceDhcpv4Client) setDefault() {
	var choices_set int = 0
	var choice DeviceDhcpv4ClientChoiceEnum

	if obj.obj.ServerAddress != nil {
		choices_set += 1
		choice = DeviceDhcpv4ClientChoice.SERVER_ADDRESS
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceDhcpv4ClientChoice.FIRST_SERVER)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceDhcpv4Client")
			}
		} else {
			intVal := otg.DeviceDhcpv4Client_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceDhcpv4Client_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.Broadcast == nil {
		obj.SetBroadcast(false)
	}

}
