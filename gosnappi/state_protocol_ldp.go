package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolLdp *****
type stateProtocolLdp struct {
	validation
	obj           *otg.StateProtocolLdp
	marshaller    marshalStateProtocolLdp
	unMarshaller  unMarshalStateProtocolLdp
	routersHolder StateProtocolLdpRouters
}

func NewStateProtocolLdp() StateProtocolLdp {
	obj := stateProtocolLdp{obj: &otg.StateProtocolLdp{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolLdp) msg() *otg.StateProtocolLdp {
	return obj.obj
}

func (obj *stateProtocolLdp) setMsg(msg *otg.StateProtocolLdp) StateProtocolLdp {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolLdp struct {
	obj *stateProtocolLdp
}

type marshalStateProtocolLdp interface {
	// ToProto marshals StateProtocolLdp to protobuf object *otg.StateProtocolLdp
	ToProto() (*otg.StateProtocolLdp, error)
	// ToPbText marshals StateProtocolLdp to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolLdp to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolLdp to JSON text
	ToJson() (string, error)
}

type unMarshalstateProtocolLdp struct {
	obj *stateProtocolLdp
}

type unMarshalStateProtocolLdp interface {
	// FromProto unmarshals StateProtocolLdp from protobuf object *otg.StateProtocolLdp
	FromProto(msg *otg.StateProtocolLdp) (StateProtocolLdp, error)
	// FromPbText unmarshals StateProtocolLdp from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolLdp from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolLdp from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolLdp) Marshal() marshalStateProtocolLdp {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolLdp{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolLdp) Unmarshal() unMarshalStateProtocolLdp {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolLdp{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolLdp) ToProto() (*otg.StateProtocolLdp, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolLdp) FromProto(msg *otg.StateProtocolLdp) (StateProtocolLdp, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolLdp) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolLdp) FromPbText(value string) error {
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

func (m *marshalstateProtocolLdp) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolLdp) FromYaml(value string) error {
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

func (m *marshalstateProtocolLdp) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolLdp) FromJson(value string) error {
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

func (obj *stateProtocolLdp) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolLdp) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolLdp) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolLdp) Clone() (StateProtocolLdp, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolLdp()
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

func (obj *stateProtocolLdp) setNil() {
	obj.routersHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateProtocolLdp is sets state of configured LDP routers.
type StateProtocolLdp interface {
	Validation
	// msg marshals StateProtocolLdp to protobuf object *otg.StateProtocolLdp
	// and doesn't set defaults
	msg() *otg.StateProtocolLdp
	// setMsg unmarshals StateProtocolLdp from protobuf object *otg.StateProtocolLdp
	// and doesn't set defaults
	setMsg(*otg.StateProtocolLdp) StateProtocolLdp
	// provides marshal interface
	Marshal() marshalStateProtocolLdp
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolLdp
	// validate validates StateProtocolLdp
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolLdp, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateProtocolLdpChoiceEnum, set in StateProtocolLdp
	Choice() StateProtocolLdpChoiceEnum
	// setChoice assigns StateProtocolLdpChoiceEnum provided by user to StateProtocolLdp
	setChoice(value StateProtocolLdpChoiceEnum) StateProtocolLdp
	// Routers returns StateProtocolLdpRouters, set in StateProtocolLdp.
	// StateProtocolLdpRouters is sets state of configured LDP routers.
	Routers() StateProtocolLdpRouters
	// SetRouters assigns StateProtocolLdpRouters provided by user to StateProtocolLdp.
	// StateProtocolLdpRouters is sets state of configured LDP routers.
	SetRouters(value StateProtocolLdpRouters) StateProtocolLdp
	// HasRouters checks if Routers has been set in StateProtocolLdp
	HasRouters() bool
	setNil()
}

type StateProtocolLdpChoiceEnum string

// Enum of Choice on StateProtocolLdp
var StateProtocolLdpChoice = struct {
	ROUTERS StateProtocolLdpChoiceEnum
}{
	ROUTERS: StateProtocolLdpChoiceEnum("routers"),
}

func (obj *stateProtocolLdp) Choice() StateProtocolLdpChoiceEnum {
	return StateProtocolLdpChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateProtocolLdp) setChoice(value StateProtocolLdpChoiceEnum) StateProtocolLdp {
	intValue, ok := otg.StateProtocolLdp_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolLdpChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolLdp_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Routers = nil
	obj.routersHolder = nil

	if value == StateProtocolLdpChoice.ROUTERS {
		obj.obj.Routers = NewStateProtocolLdpRouters().msg()
	}

	return obj
}

// Sets the state of LDP routers.
// Routers returns a StateProtocolLdpRouters
func (obj *stateProtocolLdp) Routers() StateProtocolLdpRouters {
	if obj.obj.Routers == nil {
		obj.setChoice(StateProtocolLdpChoice.ROUTERS)
	}
	if obj.routersHolder == nil {
		obj.routersHolder = &stateProtocolLdpRouters{obj: obj.obj.Routers}
	}
	return obj.routersHolder
}

// Sets the state of LDP routers.
// Routers returns a StateProtocolLdpRouters
func (obj *stateProtocolLdp) HasRouters() bool {
	return obj.obj.Routers != nil
}

// Sets the state of LDP routers.
// SetRouters sets the StateProtocolLdpRouters value in the StateProtocolLdp object
func (obj *stateProtocolLdp) SetRouters(value StateProtocolLdpRouters) StateProtocolLdp {
	obj.setChoice(StateProtocolLdpChoice.ROUTERS)
	obj.routersHolder = nil
	obj.obj.Routers = value.msg()

	return obj
}

func (obj *stateProtocolLdp) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateProtocolLdp")
	}

	if obj.obj.Routers != nil {

		obj.Routers().validateObj(vObj, set_default)
	}

}

func (obj *stateProtocolLdp) setDefault() {
	var choices_set int = 0
	var choice StateProtocolLdpChoiceEnum

	if obj.obj.Routers != nil {
		choices_set += 1
		choice = StateProtocolLdpChoice.ROUTERS
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateProtocolLdp")
			}
		} else {
			intVal := otg.StateProtocolLdp_Choice_Enum_value[string(choice)]
			enumValue := otg.StateProtocolLdp_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
