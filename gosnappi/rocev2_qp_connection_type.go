package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Rocev2QPConnectionType *****
type rocev2QPConnectionType struct {
	validation
	obj                      *otg.Rocev2QPConnectionType
	marshaller               marshalRocev2QPConnectionType
	unMarshaller             unMarshalRocev2QPConnectionType
	reliableConnectionHolder Rocev2AckAndNak
}

func NewRocev2QPConnectionType() Rocev2QPConnectionType {
	obj := rocev2QPConnectionType{obj: &otg.Rocev2QPConnectionType{}}
	obj.setDefault()
	return &obj
}

func (obj *rocev2QPConnectionType) msg() *otg.Rocev2QPConnectionType {
	return obj.obj
}

func (obj *rocev2QPConnectionType) setMsg(msg *otg.Rocev2QPConnectionType) Rocev2QPConnectionType {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrocev2QPConnectionType struct {
	obj *rocev2QPConnectionType
}

type marshalRocev2QPConnectionType interface {
	// ToProto marshals Rocev2QPConnectionType to protobuf object *otg.Rocev2QPConnectionType
	ToProto() (*otg.Rocev2QPConnectionType, error)
	// ToPbText marshals Rocev2QPConnectionType to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Rocev2QPConnectionType to YAML text
	ToYaml() (string, error)
	// ToJson marshals Rocev2QPConnectionType to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals Rocev2QPConnectionType to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalrocev2QPConnectionType struct {
	obj *rocev2QPConnectionType
}

type unMarshalRocev2QPConnectionType interface {
	// FromProto unmarshals Rocev2QPConnectionType from protobuf object *otg.Rocev2QPConnectionType
	FromProto(msg *otg.Rocev2QPConnectionType) (Rocev2QPConnectionType, error)
	// FromPbText unmarshals Rocev2QPConnectionType from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Rocev2QPConnectionType from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Rocev2QPConnectionType from JSON text
	FromJson(value string) error
}

func (obj *rocev2QPConnectionType) Marshal() marshalRocev2QPConnectionType {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrocev2QPConnectionType{obj: obj}
	}
	return obj.marshaller
}

func (obj *rocev2QPConnectionType) Unmarshal() unMarshalRocev2QPConnectionType {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrocev2QPConnectionType{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrocev2QPConnectionType) ToProto() (*otg.Rocev2QPConnectionType, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrocev2QPConnectionType) FromProto(msg *otg.Rocev2QPConnectionType) (Rocev2QPConnectionType, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrocev2QPConnectionType) ToPbText() (string, error) {
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

func (m *unMarshalrocev2QPConnectionType) FromPbText(value string) error {
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

func (m *marshalrocev2QPConnectionType) ToYaml() (string, error) {
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

func (m *unMarshalrocev2QPConnectionType) FromYaml(value string) error {
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

func (m *marshalrocev2QPConnectionType) ToJsonRaw() (string, error) {
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

func (m *marshalrocev2QPConnectionType) ToJson() (string, error) {
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

func (m *unMarshalrocev2QPConnectionType) FromJson(value string) error {
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

func (obj *rocev2QPConnectionType) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rocev2QPConnectionType) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rocev2QPConnectionType) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rocev2QPConnectionType) Clone() (Rocev2QPConnectionType, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRocev2QPConnectionType()
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

func (obj *rocev2QPConnectionType) setNil() {
	obj.reliableConnectionHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Rocev2QPConnectionType is specifies the connection type for the QP, determining what and how the QP transfers data.
type Rocev2QPConnectionType interface {
	Validation
	// msg marshals Rocev2QPConnectionType to protobuf object *otg.Rocev2QPConnectionType
	// and doesn't set defaults
	msg() *otg.Rocev2QPConnectionType
	// setMsg unmarshals Rocev2QPConnectionType from protobuf object *otg.Rocev2QPConnectionType
	// and doesn't set defaults
	setMsg(*otg.Rocev2QPConnectionType) Rocev2QPConnectionType
	// provides marshal interface
	Marshal() marshalRocev2QPConnectionType
	// provides unmarshal interface
	Unmarshal() unMarshalRocev2QPConnectionType
	// validate validates Rocev2QPConnectionType
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Rocev2QPConnectionType, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns Rocev2QPConnectionTypeChoiceEnum, set in Rocev2QPConnectionType
	Choice() Rocev2QPConnectionTypeChoiceEnum
	// setChoice assigns Rocev2QPConnectionTypeChoiceEnum provided by user to Rocev2QPConnectionType
	setChoice(value Rocev2QPConnectionTypeChoiceEnum) Rocev2QPConnectionType
	// HasChoice checks if Choice has been set in Rocev2QPConnectionType
	HasChoice() bool
	// ReliableConnection returns Rocev2AckAndNak, set in Rocev2QPConnectionType.
	// Rocev2AckAndNak is defines the ACK and NAK settings for RoCEv2. This configuration ensures reliable data delivery by controlling how the system responds to successful and failed packet transmissions.
	ReliableConnection() Rocev2AckAndNak
	// SetReliableConnection assigns Rocev2AckAndNak provided by user to Rocev2QPConnectionType.
	// Rocev2AckAndNak is defines the ACK and NAK settings for RoCEv2. This configuration ensures reliable data delivery by controlling how the system responds to successful and failed packet transmissions.
	SetReliableConnection(value Rocev2AckAndNak) Rocev2QPConnectionType
	// HasReliableConnection checks if ReliableConnection has been set in Rocev2QPConnectionType
	HasReliableConnection() bool
	setNil()
}

type Rocev2QPConnectionTypeChoiceEnum string

// Enum of Choice on Rocev2QPConnectionType
var Rocev2QPConnectionTypeChoice = struct {
	RELIABLE_CONNECTION Rocev2QPConnectionTypeChoiceEnum
}{
	RELIABLE_CONNECTION: Rocev2QPConnectionTypeChoiceEnum("reliable_connection"),
}

func (obj *rocev2QPConnectionType) Choice() Rocev2QPConnectionTypeChoiceEnum {
	return Rocev2QPConnectionTypeChoiceEnum(obj.obj.Choice.Enum().String())
}

// description is TBD
// Choice returns a string
func (obj *rocev2QPConnectionType) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *rocev2QPConnectionType) setChoice(value Rocev2QPConnectionTypeChoiceEnum) Rocev2QPConnectionType {
	intValue, ok := otg.Rocev2QPConnectionType_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Rocev2QPConnectionTypeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Rocev2QPConnectionType_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.ReliableConnection = nil
	obj.reliableConnectionHolder = nil

	if value == Rocev2QPConnectionTypeChoice.RELIABLE_CONNECTION {
		obj.obj.ReliableConnection = NewRocev2AckAndNak().msg()
	}

	return obj
}

// description is TBD
// ReliableConnection returns a Rocev2AckAndNak
func (obj *rocev2QPConnectionType) ReliableConnection() Rocev2AckAndNak {
	if obj.obj.ReliableConnection == nil {
		obj.setChoice(Rocev2QPConnectionTypeChoice.RELIABLE_CONNECTION)
	}
	if obj.reliableConnectionHolder == nil {
		obj.reliableConnectionHolder = &rocev2AckAndNak{obj: obj.obj.ReliableConnection}
	}
	return obj.reliableConnectionHolder
}

// description is TBD
// ReliableConnection returns a Rocev2AckAndNak
func (obj *rocev2QPConnectionType) HasReliableConnection() bool {
	return obj.obj.ReliableConnection != nil
}

// description is TBD
// SetReliableConnection sets the Rocev2AckAndNak value in the Rocev2QPConnectionType object
func (obj *rocev2QPConnectionType) SetReliableConnection(value Rocev2AckAndNak) Rocev2QPConnectionType {
	obj.setChoice(Rocev2QPConnectionTypeChoice.RELIABLE_CONNECTION)
	obj.reliableConnectionHolder = nil
	obj.obj.ReliableConnection = value.msg()

	return obj
}

func (obj *rocev2QPConnectionType) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ReliableConnection != nil {

		obj.ReliableConnection().validateObj(vObj, set_default)
	}

}

func (obj *rocev2QPConnectionType) setDefault() {
	var choices_set int = 0
	var choice Rocev2QPConnectionTypeChoiceEnum

	if obj.obj.ReliableConnection != nil {
		choices_set += 1
		choice = Rocev2QPConnectionTypeChoice.RELIABLE_CONNECTION
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(Rocev2QPConnectionTypeChoice.RELIABLE_CONNECTION)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in Rocev2QPConnectionType")
			}
		} else {
			intVal := otg.Rocev2QPConnectionType_Choice_Enum_value[string(choice)]
			enumValue := otg.Rocev2QPConnectionType_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
