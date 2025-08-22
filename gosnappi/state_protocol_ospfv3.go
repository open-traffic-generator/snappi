package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolOspfv3 *****
type stateProtocolOspfv3 struct {
	validation
	obj           *otg.StateProtocolOspfv3
	marshaller    marshalStateProtocolOspfv3
	unMarshaller  unMarshalStateProtocolOspfv3
	routersHolder StateProtocolOspfv3Routers
}

func NewStateProtocolOspfv3() StateProtocolOspfv3 {
	obj := stateProtocolOspfv3{obj: &otg.StateProtocolOspfv3{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolOspfv3) msg() *otg.StateProtocolOspfv3 {
	return obj.obj
}

func (obj *stateProtocolOspfv3) setMsg(msg *otg.StateProtocolOspfv3) StateProtocolOspfv3 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolOspfv3 struct {
	obj *stateProtocolOspfv3
}

type marshalStateProtocolOspfv3 interface {
	// ToProto marshals StateProtocolOspfv3 to protobuf object *otg.StateProtocolOspfv3
	ToProto() (*otg.StateProtocolOspfv3, error)
	// ToPbText marshals StateProtocolOspfv3 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolOspfv3 to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolOspfv3 to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolOspfv3 struct {
	obj *stateProtocolOspfv3
}

type unMarshalStateProtocolOspfv3 interface {
	// FromProto unmarshals StateProtocolOspfv3 from protobuf object *otg.StateProtocolOspfv3
	FromProto(msg *otg.StateProtocolOspfv3) (StateProtocolOspfv3, error)
	// FromPbText unmarshals StateProtocolOspfv3 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolOspfv3 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolOspfv3 from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolOspfv3) Marshal() marshalStateProtocolOspfv3 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolOspfv3{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolOspfv3) Unmarshal() unMarshalStateProtocolOspfv3 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolOspfv3{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolOspfv3) ToProto() (*otg.StateProtocolOspfv3, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolOspfv3) FromProto(msg *otg.StateProtocolOspfv3) (StateProtocolOspfv3, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolOspfv3) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3) FromPbText(value string) error {
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

func (m *marshalstateProtocolOspfv3) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3) FromYaml(value string) error {
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

func (m *marshalstateProtocolOspfv3) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolOspfv3) FromJson(value string) error {
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

func (obj *stateProtocolOspfv3) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv3) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolOspfv3) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolOspfv3) Clone() (StateProtocolOspfv3, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolOspfv3()
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

func (obj *stateProtocolOspfv3) setNil() {
	obj.routersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolOspfv3 is sets state of configured OSPFv3 routers.
type StateProtocolOspfv3 interface {
	Validation
	// msg marshals StateProtocolOspfv3 to protobuf object *otg.StateProtocolOspfv3
	// and doesn't set defaults
	msg() *otg.StateProtocolOspfv3
	// setMsg unmarshals StateProtocolOspfv3 from protobuf object *otg.StateProtocolOspfv3
	// and doesn't set defaults
	setMsg(*otg.StateProtocolOspfv3) StateProtocolOspfv3
	// provides marshal interface
	Marshal() marshalStateProtocolOspfv3
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolOspfv3
	// validate validates StateProtocolOspfv3
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolOspfv3, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolOspfv3ChoiceEnum, set in StateProtocolOspfv3
	Choice() StateProtocolOspfv3ChoiceEnum
	// setChoice assigns StateProtocolOspfv3ChoiceEnum provided by user to StateProtocolOspfv3
	setChoice(value StateProtocolOspfv3ChoiceEnum) StateProtocolOspfv3
	// Routers returns StateProtocolOspfv3Routers, set in StateProtocolOspfv3.
	// StateProtocolOspfv3Routers is sets state of configured OSPFv3 routers.
	Routers() StateProtocolOspfv3Routers
	// SetRouters assigns StateProtocolOspfv3Routers provided by user to StateProtocolOspfv3.
	// StateProtocolOspfv3Routers is sets state of configured OSPFv3 routers.
	SetRouters(value StateProtocolOspfv3Routers) StateProtocolOspfv3
	// HasRouters checks if Routers has been set in StateProtocolOspfv3
	HasRouters() bool
	setNil()
}

type StateProtocolOspfv3ChoiceEnum string

// Enum of Choice on StateProtocolOspfv3
var StateProtocolOspfv3Choice = struct {
	ROUTERS StateProtocolOspfv3ChoiceEnum
}{
	ROUTERS: StateProtocolOspfv3ChoiceEnum("routers"),
}

func (obj *stateProtocolOspfv3) Choice() StateProtocolOspfv3ChoiceEnum {
	return StateProtocolOspfv3ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolOspfv3) setChoice(value StateProtocolOspfv3ChoiceEnum) StateProtocolOspfv3 {
	intValue, ok := otg.StateProtocolOspfv3_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolOspfv3ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolOspfv3_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Routers = nil
	obj.routersHolder = nil

	if value == StateProtocolOspfv3Choice.ROUTERS {
		obj.obj.Routers = NewStateProtocolOspfv3Routers().msg()
	}

	return obj
}

// description is TBD
// Routers returns a StateProtocolOspfv3Routers
func (obj *stateProtocolOspfv3) Routers() StateProtocolOspfv3Routers {
	if obj.obj.Routers == nil {
		obj.setChoice(StateProtocolOspfv3Choice.ROUTERS)
	}
	if obj.routersHolder == nil {
		obj.routersHolder = &stateProtocolOspfv3Routers{obj: obj.obj.Routers}
	}
	return obj.routersHolder
}

// description is TBD
// Routers returns a StateProtocolOspfv3Routers
func (obj *stateProtocolOspfv3) HasRouters() bool {
	return obj.obj.Routers != nil
}

// description is TBD
// SetRouters sets the StateProtocolOspfv3Routers value in the StateProtocolOspfv3 object
func (obj *stateProtocolOspfv3) SetRouters(value StateProtocolOspfv3Routers) StateProtocolOspfv3 {
	obj.setChoice(StateProtocolOspfv3Choice.ROUTERS)
	obj.routersHolder = nil
	obj.obj.Routers = value.msg()

	return obj
}

func (obj *stateProtocolOspfv3) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolOspfv3")
	}

	if obj.obj.Routers != nil {

		obj.Routers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolOspfv3) setDefault() {
	var choices_set int = 0
	var choice StateProtocolOspfv3ChoiceEnum

	if obj.obj.Routers != nil {
		choices_set += 1
		choice = StateProtocolOspfv3Choice.ROUTERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolOspfv3")
			}
		} else {
			intVal := otg.StateProtocolOspfv3_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolOspfv3_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
