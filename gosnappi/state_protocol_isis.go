package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolIsis *****
type stateProtocolIsis struct {
	validation
	obj           *otg.StateProtocolIsis
	marshaller    marshalStateProtocolIsis
	unMarshaller  unMarshalStateProtocolIsis
	routersHolder StateProtocolIsisRouters
}

func NewStateProtocolIsis() StateProtocolIsis {
	obj := stateProtocolIsis{obj: &otg.StateProtocolIsis{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolIsis) msg() *otg.StateProtocolIsis {
	return obj.obj
}

func (obj *stateProtocolIsis) setMsg(msg *otg.StateProtocolIsis) StateProtocolIsis {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolIsis struct {
	obj *stateProtocolIsis
}

type marshalStateProtocolIsis interface {
	// ToProto marshals StateProtocolIsis to protobuf object *otg.StateProtocolIsis
	ToProto() (*otg.StateProtocolIsis, error)
	// ToPbText marshals StateProtocolIsis to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolIsis to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolIsis to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolIsis struct {
	obj *stateProtocolIsis
}

type unMarshalStateProtocolIsis interface {
	// FromProto unmarshals StateProtocolIsis from protobuf object *otg.StateProtocolIsis
	FromProto(msg *otg.StateProtocolIsis) (StateProtocolIsis, error)
	// FromPbText unmarshals StateProtocolIsis from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolIsis from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolIsis from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolIsis) Marshal() marshalStateProtocolIsis {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolIsis{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolIsis) Unmarshal() unMarshalStateProtocolIsis {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolIsis{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolIsis) ToProto() (*otg.StateProtocolIsis, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolIsis) FromProto(msg *otg.StateProtocolIsis) (StateProtocolIsis, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolIsis) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolIsis) FromPbText(value string) error {
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

func (m *marshalstateProtocolIsis) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolIsis) FromYaml(value string) error {
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

func (m *marshalstateProtocolIsis) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolIsis) FromJson(value string) error {
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

func (obj *stateProtocolIsis) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolIsis) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolIsis) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolIsis) Clone() (StateProtocolIsis, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolIsis()
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

func (obj *stateProtocolIsis) setNil() {
	obj.routersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolIsis is sets state of configured ISIS routers.
type StateProtocolIsis interface {
	Validation
	// msg marshals StateProtocolIsis to protobuf object *otg.StateProtocolIsis
	// and doesn't set defaults
	msg() *otg.StateProtocolIsis
	// setMsg unmarshals StateProtocolIsis from protobuf object *otg.StateProtocolIsis
	// and doesn't set defaults
	setMsg(*otg.StateProtocolIsis) StateProtocolIsis
	// provides marshal interface
	Marshal() marshalStateProtocolIsis
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolIsis
	// validate validates StateProtocolIsis
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolIsis, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolIsisChoiceEnum, set in StateProtocolIsis
	Choice() StateProtocolIsisChoiceEnum
	// setChoice assigns StateProtocolIsisChoiceEnum provided by user to StateProtocolIsis
	setChoice(value StateProtocolIsisChoiceEnum) StateProtocolIsis
	// Routers returns StateProtocolIsisRouters, set in StateProtocolIsis.
	// StateProtocolIsisRouters is sets state of configured ISIS routers.
	Routers() StateProtocolIsisRouters
	// SetRouters assigns StateProtocolIsisRouters provided by user to StateProtocolIsis.
	// StateProtocolIsisRouters is sets state of configured ISIS routers.
	SetRouters(value StateProtocolIsisRouters) StateProtocolIsis
	// HasRouters checks if Routers has been set in StateProtocolIsis
	HasRouters() bool
	setNil()
}

type StateProtocolIsisChoiceEnum string

// Enum of Choice on StateProtocolIsis
var StateProtocolIsisChoice = struct {
	ROUTERS StateProtocolIsisChoiceEnum
}{
	ROUTERS: StateProtocolIsisChoiceEnum("routers"),
}

func (obj *stateProtocolIsis) Choice() StateProtocolIsisChoiceEnum {
	return StateProtocolIsisChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolIsis) setChoice(value StateProtocolIsisChoiceEnum) StateProtocolIsis {
	intValue, ok := otg.StateProtocolIsis_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolIsisChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolIsis_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Routers = nil
	obj.routersHolder = nil

	if value == StateProtocolIsisChoice.ROUTERS {
		obj.obj.Routers = NewStateProtocolIsisRouters().msg()
	}

	return obj
}

// description is TBD
// Routers returns a StateProtocolIsisRouters
func (obj *stateProtocolIsis) Routers() StateProtocolIsisRouters {
	if obj.obj.Routers == nil {
		obj.setChoice(StateProtocolIsisChoice.ROUTERS)
	}
	if obj.routersHolder == nil {
		obj.routersHolder = &stateProtocolIsisRouters{obj: obj.obj.Routers}
	}
	return obj.routersHolder
}

// description is TBD
// Routers returns a StateProtocolIsisRouters
func (obj *stateProtocolIsis) HasRouters() bool {
	return obj.obj.Routers != nil
}

// description is TBD
// SetRouters sets the StateProtocolIsisRouters value in the StateProtocolIsis object
func (obj *stateProtocolIsis) SetRouters(value StateProtocolIsisRouters) StateProtocolIsis {
	obj.setChoice(StateProtocolIsisChoice.ROUTERS)
	obj.routersHolder = nil
	obj.obj.Routers = value.msg()

	return obj
}

func (obj *stateProtocolIsis) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolIsis")
	}

	if obj.obj.Routers != nil {

		obj.Routers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolIsis) setDefault() {
	var choices_set int = 0
	var choice StateProtocolIsisChoiceEnum

	if obj.obj.Routers != nil {
		choices_set += 1
		choice = StateProtocolIsisChoice.ROUTERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolIsis")
			}
		} else {
			intVal := otg.StateProtocolIsis_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolIsis_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
