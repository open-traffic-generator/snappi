package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** LldpConnection *****
type lldpConnection struct {
	validation
	obj          *otg.LldpConnection
	marshaller   marshalLldpConnection
	unMarshaller unMarshalLldpConnection
}

func NewLldpConnection() LldpConnection {
	obj := lldpConnection{obj: &otg.LldpConnection{}}
	obj.setDefault()
	return &obj
}

func (obj *lldpConnection) msg() *otg.LldpConnection {
	return obj.obj
}

func (obj *lldpConnection) setMsg(msg *otg.LldpConnection) LldpConnection {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshallldpConnection struct {
	obj *lldpConnection
}

type marshalLldpConnection interface {
	// ToProto marshals LldpConnection to protobuf object *otg.LldpConnection
	ToProto() (*otg.LldpConnection, error)
	// ToPbText marshals LldpConnection to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals LldpConnection to YAML text
	ToYaml() (string, error)
	// ToJson marshals LldpConnection to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals LldpConnection to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshallldpConnection struct {
	obj *lldpConnection
}

type unMarshalLldpConnection interface {
	// FromProto unmarshals LldpConnection from protobuf object *otg.LldpConnection
	FromProto(msg *otg.LldpConnection) (LldpConnection, error)
	// FromPbText unmarshals LldpConnection from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals LldpConnection from YAML text
	FromYaml(value string) error
	// FromJson unmarshals LldpConnection from JSON text
	FromJson(value string) error
}

func (obj *lldpConnection) Marshal() marshalLldpConnection {
	if obj.marshaller == nil {
		obj.marshaller = &marshallldpConnection{obj: obj}
	}
	return obj.marshaller
}

func (obj *lldpConnection) Unmarshal() unMarshalLldpConnection {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallldpConnection{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallldpConnection) ToProto() (*otg.LldpConnection, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallldpConnection) FromProto(msg *otg.LldpConnection) (LldpConnection, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallldpConnection) ToPbText() (string, error) {
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

func (m *unMarshallldpConnection) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshallldpConnection) ToYaml() (string, error) {
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

func (m *unMarshallldpConnection) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshallldpConnection) ToJsonRaw() (string, error) {
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

func (m *marshallldpConnection) ToJson() (string, error) {
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

func (m *unMarshallldpConnection) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *lldpConnection) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *lldpConnection) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *lldpConnection) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *lldpConnection) Clone() (LldpConnection, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLldpConnection()
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

// LldpConnection is lLDP connection to a test port. In future if more connection options arise  LLDP connection object will be enhanced.
type LldpConnection interface {
	Validation
	// msg marshals LldpConnection to protobuf object *otg.LldpConnection
	// and doesn't set defaults
	msg() *otg.LldpConnection
	// setMsg unmarshals LldpConnection from protobuf object *otg.LldpConnection
	// and doesn't set defaults
	setMsg(*otg.LldpConnection) LldpConnection
	// provides marshal interface
	Marshal() marshalLldpConnection
	// provides unmarshal interface
	Unmarshal() unMarshalLldpConnection
	// validate validates LldpConnection
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (LldpConnection, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns LldpConnectionChoiceEnum, set in LldpConnection
	Choice() LldpConnectionChoiceEnum
	// setChoice assigns LldpConnectionChoiceEnum provided by user to LldpConnection
	setChoice(value LldpConnectionChoiceEnum) LldpConnection
	// HasChoice checks if Choice has been set in LldpConnection
	HasChoice() bool
	// PortName returns string, set in LldpConnection.
	PortName() string
	// SetPortName assigns string provided by user to LldpConnection
	SetPortName(value string) LldpConnection
	// HasPortName checks if PortName has been set in LldpConnection
	HasPortName() bool
}

type LldpConnectionChoiceEnum string

// Enum of Choice on LldpConnection
var LldpConnectionChoice = struct {
	PORT_NAME LldpConnectionChoiceEnum
}{
	PORT_NAME: LldpConnectionChoiceEnum("port_name"),
}

func (obj *lldpConnection) Choice() LldpConnectionChoiceEnum {
	return LldpConnectionChoiceEnum(obj.obj.Choice.Enum().String())
}

// The name of the test port or other connection objects on which LLDP is configured.
// Choice returns a string
func (obj *lldpConnection) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *lldpConnection) setChoice(value LldpConnectionChoiceEnum) LldpConnection {
	intValue, ok := otg.LldpConnection_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on LldpConnectionChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.LldpConnection_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.PortName = nil
	return obj
}

// Name of the test port on which LLDP is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *lldpConnection) PortName() string {

	if obj.obj.PortName == nil {
		obj.setChoice(LldpConnectionChoice.PORT_NAME)
	}

	return *obj.obj.PortName

}

// Name of the test port on which LLDP is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortName returns a string
func (obj *lldpConnection) HasPortName() bool {
	return obj.obj.PortName != nil
}

// Name of the test port on which LLDP is configured on.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortName sets the string value in the LldpConnection object
func (obj *lldpConnection) SetPortName(value string) LldpConnection {
	obj.setChoice(LldpConnectionChoice.PORT_NAME)
	obj.obj.PortName = &value
	return obj
}

func (obj *lldpConnection) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *lldpConnection) setDefault() {
	var choices_set int = 0
	var choice LldpConnectionChoiceEnum

	if obj.obj.PortName != nil {
		choices_set += 1
		choice = LldpConnectionChoice.PORT_NAME
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in LldpConnection")
			}
		} else {
			intVal := otg.LldpConnection_Choice_Enum_value[string(choice)]
			enumValue := otg.LldpConnection_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
