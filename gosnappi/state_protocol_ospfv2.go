package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolOspfv2 *****
type stateProtocolOspfv2 struct {
	validation
	obj           *otg.StateProtocolOspfv2
	marshaller    marshalStateProtocolOspfv2
	unMarshaller  unMarshalStateProtocolOspfv2
	routersHolder StateProtocolOspfv2Routers
}

func NewStateProtocolOspfv2() StateProtocolOspfv2 {
	obj := stateProtocolOspfv2{obj: &otg.StateProtocolOspfv2{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolOspfv2) msg() *otg.StateProtocolOspfv2 {
	return obj.obj
}

func (obj *stateProtocolOspfv2) setMsg(msg *otg.StateProtocolOspfv2) StateProtocolOspfv2 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolOspfv2 struct {
	obj *stateProtocolOspfv2
}

type marshalStateProtocolOspfv2 interface {
	// ToProto marshals StateProtocolOspfv2 to protobuf object *otg.StateProtocolOspfv2
	ToProto() (*otg.StateProtocolOspfv2, error)
	// ToPbText marshals StateProtocolOspfv2 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolOspfv2 to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolOspfv2 to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolOspfv2 struct {
	obj *stateProtocolOspfv2
}

type unMarshalStateProtocolOspfv2 interface {
	// FromProto unmarshals StateProtocolOspfv2 from protobuf object *otg.StateProtocolOspfv2
	FromProto(msg *otg.StateProtocolOspfv2) (StateProtocolOspfv2, error)
	// FromPbText unmarshals StateProtocolOspfv2 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolOspfv2 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolOspfv2 from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolOspfv2) Marshal() marshalStateProtocolOspfv2 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolOspfv2{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolOspfv2) Unmarshal() unMarshalStateProtocolOspfv2 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolOspfv2{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolOspfv2) ToProto() (*otg.StateProtocolOspfv2, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolOspfv2) FromProto(msg *otg.StateProtocolOspfv2) (StateProtocolOspfv2, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolOspfv2) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2) FromPbText(value string) error {
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

func (m *marshalstateProtocolOspfv2) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2) FromYaml(value string) error {
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

func (m *marshalstateProtocolOspfv2) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolOspfv2) FromJson(value string) error {
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

func (obj *stateProtocolOspfv2) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv2) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolOspfv2) Clone() (StateProtocolOspfv2, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolOspfv2()
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

func (obj *stateProtocolOspfv2) setNil() {
	obj.routersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolOspfv2 is sets state of configured OSPFv2 routers.
type StateProtocolOspfv2 interface {
	Validation
	// msg marshals StateProtocolOspfv2 to protobuf object *otg.StateProtocolOspfv2
	// and doesn't set defaults
	msg() *otg.StateProtocolOspfv2
	// setMsg unmarshals StateProtocolOspfv2 from protobuf object *otg.StateProtocolOspfv2
	// and doesn't set defaults
	setMsg(*otg.StateProtocolOspfv2) StateProtocolOspfv2
	// provides marshal interface
	Marshal() marshalStateProtocolOspfv2
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolOspfv2
	// validate validates StateProtocolOspfv2
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolOspfv2, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolOspfv2ChoiceEnum, set in StateProtocolOspfv2
	Choice() StateProtocolOspfv2ChoiceEnum
	// setChoice assigns StateProtocolOspfv2ChoiceEnum provided by user to StateProtocolOspfv2
	setChoice(value StateProtocolOspfv2ChoiceEnum) StateProtocolOspfv2
	// Routers returns StateProtocolOspfv2Routers, set in StateProtocolOspfv2.
	// StateProtocolOspfv2Routers is sets state of configured OSPFv2 routers.
	Routers() StateProtocolOspfv2Routers
	// SetRouters assigns StateProtocolOspfv2Routers provided by user to StateProtocolOspfv2.
	// StateProtocolOspfv2Routers is sets state of configured OSPFv2 routers.
	SetRouters(value StateProtocolOspfv2Routers) StateProtocolOspfv2
	// HasRouters checks if Routers has been set in StateProtocolOspfv2
	HasRouters() bool
	setNil()
}

type StateProtocolOspfv2ChoiceEnum string

// Enum of Choice on StateProtocolOspfv2
var StateProtocolOspfv2Choice = struct {
	ROUTERS StateProtocolOspfv2ChoiceEnum
}{
	ROUTERS: StateProtocolOspfv2ChoiceEnum("routers"),
}

func (obj *stateProtocolOspfv2) Choice() StateProtocolOspfv2ChoiceEnum {
	return StateProtocolOspfv2ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolOspfv2) setChoice(value StateProtocolOspfv2ChoiceEnum) StateProtocolOspfv2 {
	intValue, ok := otg.StateProtocolOspfv2_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolOspfv2ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolOspfv2_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Routers = nil
	obj.routersHolder = nil

	if value == StateProtocolOspfv2Choice.ROUTERS {
		obj.obj.Routers = NewStateProtocolOspfv2Routers().msg()
	}

	return obj
}

// description is TBD
// Routers returns a StateProtocolOspfv2Routers
func (obj *stateProtocolOspfv2) Routers() StateProtocolOspfv2Routers {
	if obj.obj.Routers == nil {
		obj.setChoice(StateProtocolOspfv2Choice.ROUTERS)
	}
	if obj.routersHolder == nil {
		obj.routersHolder = &stateProtocolOspfv2Routers{obj: obj.obj.Routers}
	}
	return obj.routersHolder
}

// description is TBD
// Routers returns a StateProtocolOspfv2Routers
func (obj *stateProtocolOspfv2) HasRouters() bool {
	return obj.obj.Routers != nil
}

// description is TBD
// SetRouters sets the StateProtocolOspfv2Routers value in the StateProtocolOspfv2 object
func (obj *stateProtocolOspfv2) SetRouters(value StateProtocolOspfv2Routers) StateProtocolOspfv2 {
	obj.setChoice(StateProtocolOspfv2Choice.ROUTERS)
	obj.routersHolder = nil
	obj.obj.Routers = value.msg()

	return obj
}

func (obj *stateProtocolOspfv2) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolOspfv2")
	}

	if obj.obj.Routers != nil {

		obj.Routers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolOspfv2) setDefault() {
	var choices_set int = 0
	var choice StateProtocolOspfv2ChoiceEnum

	if obj.obj.Routers != nil {
		choices_set += 1
		choice = StateProtocolOspfv2Choice.ROUTERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolOspfv2")
			}
		} else {
			intVal := otg.StateProtocolOspfv2_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolOspfv2_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
