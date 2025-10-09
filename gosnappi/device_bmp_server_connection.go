package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** DeviceBmpServerConnection *****
type deviceBmpServerConnection struct {
	validation
	obj           *otg.DeviceBmpServerConnection
	marshaller    marshalDeviceBmpServerConnection
	unMarshaller  unMarshalDeviceBmpServerConnection
	passiveHolder DeviceBmpServerPassiveConnection
	activeHolder  DeviceBmpServerActiveConnection
}

func NewDeviceBmpServerConnection() DeviceBmpServerConnection {
	obj := deviceBmpServerConnection{obj: &otg.DeviceBmpServerConnection{}}
	obj.setDefault()
	return &obj
}

func (obj *deviceBmpServerConnection) msg() *otg.DeviceBmpServerConnection {
	return obj.obj
}

func (obj *deviceBmpServerConnection) setMsg(msg *otg.DeviceBmpServerConnection) DeviceBmpServerConnection {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshaldeviceBmpServerConnection struct {
	obj *deviceBmpServerConnection
}

type marshalDeviceBmpServerConnection interface {
	// ToProto marshals DeviceBmpServerConnection to protobuf object *otg.DeviceBmpServerConnection
	ToProto() (*otg.DeviceBmpServerConnection, error)
	// ToPbText marshals DeviceBmpServerConnection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals DeviceBmpServerConnection to YAML text
	ToYaml() (string, error)
	// ToJson marshals DeviceBmpServerConnection to JSON text
	ToJson() (string, error)
}

type unMarshaldeviceBmpServerConnection struct {
	obj *deviceBmpServerConnection
}

type unMarshalDeviceBmpServerConnection interface {
	// FromProto unmarshals DeviceBmpServerConnection from protobuf object *otg.DeviceBmpServerConnection
	FromProto(msg *otg.DeviceBmpServerConnection) (DeviceBmpServerConnection, error)
	// FromPbText unmarshals DeviceBmpServerConnection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals DeviceBmpServerConnection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals DeviceBmpServerConnection from JSON text
	FromJson(value string) error
}

func (obj *deviceBmpServerConnection) Marshal() marshalDeviceBmpServerConnection {
	if obj.marshaller == nil {
		obj.marshaller = &marshaldeviceBmpServerConnection{obj: obj}
	}
	return obj.marshaller
}

func (obj *deviceBmpServerConnection) Unmarshal() unMarshalDeviceBmpServerConnection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshaldeviceBmpServerConnection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshaldeviceBmpServerConnection) ToProto() (*otg.DeviceBmpServerConnection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshaldeviceBmpServerConnection) FromProto(msg *otg.DeviceBmpServerConnection) (DeviceBmpServerConnection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshaldeviceBmpServerConnection) ToPbText() (string, error) {
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

func (m *unMarshaldeviceBmpServerConnection) FromPbText(value string) error {
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

func (m *marshaldeviceBmpServerConnection) ToYaml() (string, error) {
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

func (m *unMarshaldeviceBmpServerConnection) FromYaml(value string) error {
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

func (m *marshaldeviceBmpServerConnection) ToJson() (string, error) {
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

func (m *unMarshaldeviceBmpServerConnection) FromJson(value string) error {
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

func (obj *deviceBmpServerConnection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *deviceBmpServerConnection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *deviceBmpServerConnection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *deviceBmpServerConnection) Clone() (DeviceBmpServerConnection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewDeviceBmpServerConnection()
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

func (obj *deviceBmpServerConnection) setNil() {
	obj.passiveHolder = nil
	obj.activeHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// DeviceBmpServerConnection is container of information about whether the BMP Server should operate in passive or active mode and corresponding information depending on the mode.
type DeviceBmpServerConnection interface {
	Validation
	// msg marshals DeviceBmpServerConnection to protobuf object *otg.DeviceBmpServerConnection
	// and doesn't set defaults
	msg() *otg.DeviceBmpServerConnection
	// setMsg unmarshals DeviceBmpServerConnection from protobuf object *otg.DeviceBmpServerConnection
	// and doesn't set defaults
	setMsg(*otg.DeviceBmpServerConnection) DeviceBmpServerConnection
	// provides marshal interface
	Marshal() marshalDeviceBmpServerConnection
	// provides unmarshal interface
	Unmarshal() unMarshalDeviceBmpServerConnection
	// validate validates DeviceBmpServerConnection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (DeviceBmpServerConnection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns DeviceBmpServerConnectionChoiceEnum, set in DeviceBmpServerConnection
	Choice() DeviceBmpServerConnectionChoiceEnum
	// setChoice assigns DeviceBmpServerConnectionChoiceEnum provided by user to DeviceBmpServerConnection
	setChoice(value DeviceBmpServerConnectionChoiceEnum) DeviceBmpServerConnection
	// HasChoice checks if Choice has been set in DeviceBmpServerConnection
	HasChoice() bool
	// Passive returns DeviceBmpServerPassiveConnection, set in DeviceBmpServerConnection.
	// DeviceBmpServerPassiveConnection is container of information when BMP Server is configured in passive mode. This means that BMP Server will not initiate the TCP connection but wait for the BMP client it is configured to accept connection from  to initiate the connection.  Note that in this case it is required to configure the BMP client in active mode otherwise BMP connection will not be intiated by either end.
	Passive() DeviceBmpServerPassiveConnection
	// SetPassive assigns DeviceBmpServerPassiveConnection provided by user to DeviceBmpServerConnection.
	// DeviceBmpServerPassiveConnection is container of information when BMP Server is configured in passive mode. This means that BMP Server will not initiate the TCP connection but wait for the BMP client it is configured to accept connection from  to initiate the connection.  Note that in this case it is required to configure the BMP client in active mode otherwise BMP connection will not be intiated by either end.
	SetPassive(value DeviceBmpServerPassiveConnection) DeviceBmpServerConnection
	// HasPassive checks if Passive has been set in DeviceBmpServerConnection
	HasPassive() bool
	// Active returns DeviceBmpServerActiveConnection, set in DeviceBmpServerConnection.
	// DeviceBmpServerActiveConnection is container of information when BMP Server is configured in active mode. This means that BMP Server will initiate the TCP connection to the remote BMP client .         Note that in this case it is required to configure the BMP client in passive mode for the BMP session to not be rejected from both ends.
	Active() DeviceBmpServerActiveConnection
	// SetActive assigns DeviceBmpServerActiveConnection provided by user to DeviceBmpServerConnection.
	// DeviceBmpServerActiveConnection is container of information when BMP Server is configured in active mode. This means that BMP Server will initiate the TCP connection to the remote BMP client .         Note that in this case it is required to configure the BMP client in passive mode for the BMP session to not be rejected from both ends.
	SetActive(value DeviceBmpServerActiveConnection) DeviceBmpServerConnection
	// HasActive checks if Active has been set in DeviceBmpServerConnection
	HasActive() bool
	setNil()
}

type DeviceBmpServerConnectionChoiceEnum string

// Enum of Choice on DeviceBmpServerConnection
var DeviceBmpServerConnectionChoice = struct {
	PASSIVE DeviceBmpServerConnectionChoiceEnum
	ACTIVE  DeviceBmpServerConnectionChoiceEnum
}{
	PASSIVE: DeviceBmpServerConnectionChoiceEnum("passive"),
	ACTIVE:  DeviceBmpServerConnectionChoiceEnum("active"),
}

func (obj *deviceBmpServerConnection) Choice() DeviceBmpServerConnectionChoiceEnum {
	return DeviceBmpServerConnectionChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *deviceBmpServerConnection) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *deviceBmpServerConnection) setChoice(value DeviceBmpServerConnectionChoiceEnum) DeviceBmpServerConnection {
	intValue, ok := otg.DeviceBmpServerConnection_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on DeviceBmpServerConnectionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.DeviceBmpServerConnection_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Active = nil
	obj.activeHolder = nil
	obj.obj.Passive = nil
	obj.passiveHolder = nil

	if value == DeviceBmpServerConnectionChoice.PASSIVE {
		obj.obj.Passive = NewDeviceBmpServerPassiveConnection().msg()
	}

	if value == DeviceBmpServerConnectionChoice.ACTIVE {
		obj.obj.Active = NewDeviceBmpServerActiveConnection().msg()
	}

	return obj
}

// description is TBD
// Passive returns a DeviceBmpServerPassiveConnection
func (obj *deviceBmpServerConnection) Passive() DeviceBmpServerPassiveConnection {
	if obj.obj.Passive == nil {
		obj.setChoice(DeviceBmpServerConnectionChoice.PASSIVE)
	}
	if obj.passiveHolder == nil {
		obj.passiveHolder = &deviceBmpServerPassiveConnection{obj: obj.obj.Passive}
	}
	return obj.passiveHolder
}

// description is TBD
// Passive returns a DeviceBmpServerPassiveConnection
func (obj *deviceBmpServerConnection) HasPassive() bool {
	return obj.obj.Passive != nil
}

// description is TBD
// SetPassive sets the DeviceBmpServerPassiveConnection value in the DeviceBmpServerConnection object
func (obj *deviceBmpServerConnection) SetPassive(value DeviceBmpServerPassiveConnection) DeviceBmpServerConnection {
	obj.setChoice(DeviceBmpServerConnectionChoice.PASSIVE)
	obj.passiveHolder = nil
	obj.obj.Passive = value.msg()

	return obj
}

// description is TBD
// Active returns a DeviceBmpServerActiveConnection
func (obj *deviceBmpServerConnection) Active() DeviceBmpServerActiveConnection {
	if obj.obj.Active == nil {
		obj.setChoice(DeviceBmpServerConnectionChoice.ACTIVE)
	}
	if obj.activeHolder == nil {
		obj.activeHolder = &deviceBmpServerActiveConnection{obj: obj.obj.Active}
	}
	return obj.activeHolder
}

// description is TBD
// Active returns a DeviceBmpServerActiveConnection
func (obj *deviceBmpServerConnection) HasActive() bool {
	return obj.obj.Active != nil
}

// description is TBD
// SetActive sets the DeviceBmpServerActiveConnection value in the DeviceBmpServerConnection object
func (obj *deviceBmpServerConnection) SetActive(value DeviceBmpServerActiveConnection) DeviceBmpServerConnection {
	obj.setChoice(DeviceBmpServerConnectionChoice.ACTIVE)
	obj.activeHolder = nil
	obj.obj.Active = value.msg()

	return obj
}

func (obj *deviceBmpServerConnection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Passive != nil {

		obj.Passive().validateObj(vObj, set_default)
	}

	if obj.obj.Active != nil {

		obj.Active().validateObj(vObj, set_default)
	}

}

func (obj *deviceBmpServerConnection) setDefault() {
	var choices_set int = 0
	var choice DeviceBmpServerConnectionChoiceEnum

	if obj.obj.Passive != nil {
		choices_set += 1
		choice = DeviceBmpServerConnectionChoice.PASSIVE
	}

	if obj.obj.Active != nil {
		choices_set += 1
		choice = DeviceBmpServerConnectionChoice.ACTIVE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(DeviceBmpServerConnectionChoice.PASSIVE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in DeviceBmpServerConnection")
			}
		} else {
			intVal := otg.DeviceBmpServerConnection_Choice_Enum_value[string(choice)]
			enumValue := otg.DeviceBmpServerConnection_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
