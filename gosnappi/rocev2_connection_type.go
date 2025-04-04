package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2ConnectionType *****
type rocev2ConnectionType struct {
	validation
	obj                      *otg.Rocev2ConnectionType
	marshaller               marshalRocev2ConnectionType
	unMarshaller             unMarshalRocev2ConnectionType
	reliableConnectionHolder Rocev2QPParameters
}

func NewRocev2ConnectionType() Rocev2ConnectionType {
	obj := rocev2ConnectionType{obj: &otg.Rocev2ConnectionType{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2ConnectionType) msg() *otg.Rocev2ConnectionType {
	return obj.obj
}

func (obj *rocev2ConnectionType) setMsg(msg *otg.Rocev2ConnectionType) Rocev2ConnectionType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2ConnectionType struct {
	obj *rocev2ConnectionType
}

type marshalRocev2ConnectionType interface {
	// ToProto marshals Rocev2ConnectionType to protobuf object *otg.Rocev2ConnectionType
	ToProto() (*otg.Rocev2ConnectionType, error)
	// ToPbText marshals Rocev2ConnectionType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2ConnectionType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2ConnectionType to JSON text
	ToJson() (string, error)
}

type unMarshalrocev2ConnectionType struct {
	obj *rocev2ConnectionType
}

type unMarshalRocev2ConnectionType interface {
	// FromProto unmarshals Rocev2ConnectionType from protobuf object *otg.Rocev2ConnectionType
	FromProto(msg *otg.Rocev2ConnectionType) (Rocev2ConnectionType, error)
	// FromPbText unmarshals Rocev2ConnectionType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2ConnectionType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2ConnectionType from JSON text
	FromJson(value string) error
}

func (obj *rocev2ConnectionType) Marshal() marshalRocev2ConnectionType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2ConnectionType{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2ConnectionType) Unmarshal() unMarshalRocev2ConnectionType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2ConnectionType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2ConnectionType) ToProto() (*otg.Rocev2ConnectionType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2ConnectionType) FromProto(msg *otg.Rocev2ConnectionType) (Rocev2ConnectionType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2ConnectionType) ToPbText() (string, error) {
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

func (m *unMarshalrocev2ConnectionType) FromPbText(value string) error {
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

func (m *marshalrocev2ConnectionType) ToYaml() (string, error) {
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

func (m *unMarshalrocev2ConnectionType) FromYaml(value string) error {
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

func (m *marshalrocev2ConnectionType) ToJson() (string, error) {
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

func (m *unMarshalrocev2ConnectionType) FromJson(value string) error {
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

func (obj *rocev2ConnectionType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2ConnectionType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2ConnectionType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2ConnectionType) Clone() (Rocev2ConnectionType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2ConnectionType()
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

func (obj *rocev2ConnectionType) setNil() {
	obj.reliableConnectionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2ConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
type Rocev2ConnectionType interface {
	Validation
	// msg marshals Rocev2ConnectionType to protobuf object *otg.Rocev2ConnectionType
	// and doesn't set defaults
	msg() *otg.Rocev2ConnectionType
	// setMsg unmarshals Rocev2ConnectionType from protobuf object *otg.Rocev2ConnectionType
	// and doesn't set defaults
	setMsg(*otg.Rocev2ConnectionType) Rocev2ConnectionType
	// provides marshal interface
	Marshal() marshalRocev2ConnectionType
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2ConnectionType
	// validate validates Rocev2ConnectionType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2ConnectionType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2ConnectionTypeChoiceEnum, set in Rocev2ConnectionType
	Choice() Rocev2ConnectionTypeChoiceEnum
	// setChoice assigns Rocev2ConnectionTypeChoiceEnum provided by user to Rocev2ConnectionType
	setChoice(value Rocev2ConnectionTypeChoiceEnum) Rocev2ConnectionType
	// HasChoice checks if Choice has been set in Rocev2ConnectionType
	HasChoice() bool
	// ReliableConnection returns Rocev2QPParameters, set in Rocev2ConnectionType.
	// Rocev2QPParameters is defines the parameters for configuring a RoCEv2 QP.
	ReliableConnection() Rocev2QPParameters
	// SetReliableConnection assigns Rocev2QPParameters provided by user to Rocev2ConnectionType.
	// Rocev2QPParameters is defines the parameters for configuring a RoCEv2 QP.
	SetReliableConnection(value Rocev2QPParameters) Rocev2ConnectionType
	// HasReliableConnection checks if ReliableConnection has been set in Rocev2ConnectionType
	HasReliableConnection() bool
	setNil()
}

type Rocev2ConnectionTypeChoiceEnum string

// Enum of Choice on Rocev2ConnectionType
var Rocev2ConnectionTypeChoice = struct {
	RELIABLE_CONNECTION Rocev2ConnectionTypeChoiceEnum
}{
	RELIABLE_CONNECTION: Rocev2ConnectionTypeChoiceEnum("reliable_connection"),
}

func (obj *rocev2ConnectionType) Choice() Rocev2ConnectionTypeChoiceEnum {
	return Rocev2ConnectionTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2ConnectionType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2ConnectionType) setChoice(value Rocev2ConnectionTypeChoiceEnum) Rocev2ConnectionType {
	intValue, ok := otg.Rocev2ConnectionType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2ConnectionTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2ConnectionType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.ReliableConnection = nil
	obj.reliableConnectionHolder = nil

	if value == Rocev2ConnectionTypeChoice.RELIABLE_CONNECTION {
		obj.obj.ReliableConnection = NewRocev2QPParameters().msg()
	}

	return obj
}

// description is TBD
// ReliableConnection returns a Rocev2QPParameters
func (obj *rocev2ConnectionType) ReliableConnection() Rocev2QPParameters {
	if obj.obj.ReliableConnection == nil {
		obj.setChoice(Rocev2ConnectionTypeChoice.RELIABLE_CONNECTION)
	}
	if obj.reliableConnectionHolder == nil {
		obj.reliableConnectionHolder = &rocev2QPParameters{obj: obj.obj.ReliableConnection}
	}
	return obj.reliableConnectionHolder
}

// description is TBD
// ReliableConnection returns a Rocev2QPParameters
func (obj *rocev2ConnectionType) HasReliableConnection() bool {
	return obj.obj.ReliableConnection != nil
}

// description is TBD
// SetReliableConnection sets the Rocev2QPParameters value in the Rocev2ConnectionType object
func (obj *rocev2ConnectionType) SetReliableConnection(value Rocev2QPParameters) Rocev2ConnectionType {
	obj.setChoice(Rocev2ConnectionTypeChoice.RELIABLE_CONNECTION)
	obj.reliableConnectionHolder = nil
	obj.obj.ReliableConnection = value.msg()

	return obj
}

func (obj *rocev2ConnectionType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ReliableConnection != nil {

		obj.ReliableConnection().validateObj(vObj, set_default)
	}

}

func (obj *rocev2ConnectionType) setDefault() {
	var choices_set int = 0
	var choice Rocev2ConnectionTypeChoiceEnum

	if obj.obj.ReliableConnection != nil {
		choices_set += 1
		choice = Rocev2ConnectionTypeChoice.RELIABLE_CONNECTION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2ConnectionTypeChoice.RELIABLE_CONNECTION)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2ConnectionType")
			}
		} else {
			intVal := otg.Rocev2ConnectionType_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2ConnectionType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
