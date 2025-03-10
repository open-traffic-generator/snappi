package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateProtocolAll *****
type stateProtocolAll struct {
	validation
	obj          *otg.StateProtocolAll
	marshaller   marshalStateProtocolAll
	unMarshaller unMarshalStateProtocolAll
}

func NewStateProtocolAll() StateProtocolAll {
	obj := stateProtocolAll{obj: &otg.StateProtocolAll{}}
	obj.setDefault()
	return &obj
}

func (obj *stateProtocolAll) msg() *otg.StateProtocolAll {
	return obj.obj
}

func (obj *stateProtocolAll) setMsg(msg *otg.StateProtocolAll) StateProtocolAll {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateProtocolAll struct {
	obj *stateProtocolAll
}

type marshalStateProtocolAll interface {
	// ToProto marshals StateProtocolAll to protobuf object *otg.StateProtocolAll
	ToProto() (*otg.StateProtocolAll, error)
	// ToPbText marshals StateProtocolAll to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateProtocolAll to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateProtocolAll to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals StateProtocolAll to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalstateProtocolAll struct {
	obj *stateProtocolAll
}

type unMarshalStateProtocolAll interface {
	// FromProto unmarshals StateProtocolAll from protobuf object *otg.StateProtocolAll
	FromProto(msg *otg.StateProtocolAll) (StateProtocolAll, error)
	// FromPbText unmarshals StateProtocolAll from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateProtocolAll from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateProtocolAll from JSON text
	FromJson(value string) error
}

func (obj *stateProtocolAll) Marshal() marshalStateProtocolAll {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateProtocolAll{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateProtocolAll) Unmarshal() unMarshalStateProtocolAll {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateProtocolAll{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateProtocolAll) ToProto() (*otg.StateProtocolAll, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateProtocolAll) FromProto(msg *otg.StateProtocolAll) (StateProtocolAll, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateProtocolAll) ToPbText() (string, error) {
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

func (m *unMarshalstateProtocolAll) FromPbText(value string) error {
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

func (m *marshalstateProtocolAll) ToYaml() (string, error) {
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

func (m *unMarshalstateProtocolAll) FromYaml(value string) error {
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

func (m *marshalstateProtocolAll) ToJsonRaw() (string, error) {
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

func (m *marshalstateProtocolAll) ToJson() (string, error) {
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

func (m *unMarshalstateProtocolAll) FromJson(value string) error {
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

func (obj *stateProtocolAll) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateProtocolAll) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateProtocolAll) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateProtocolAll) Clone() (StateProtocolAll, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateProtocolAll()
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

// StateProtocolAll is sets all configured protocols to `start` or `stop` state.
// Setting protocol state to `start` shall be a no-op if preceding `set_config` API call was made with `config.options.protocol_options.auto_start_all` set to `true` or if all the configured protocols are already started.
type StateProtocolAll interface {
	Validation
	// msg marshals StateProtocolAll to protobuf object *otg.StateProtocolAll
	// and doesn't set defaults
	msg() *otg.StateProtocolAll
	// setMsg unmarshals StateProtocolAll from protobuf object *otg.StateProtocolAll
	// and doesn't set defaults
	setMsg(*otg.StateProtocolAll) StateProtocolAll
	// provides marshal interface
	Marshal() marshalStateProtocolAll
	// provides unmarshal interface
	Unmarshal() unMarshalStateProtocolAll
	// validate validates StateProtocolAll
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateProtocolAll, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// State returns StateProtocolAllStateEnum, set in StateProtocolAll
	State() StateProtocolAllStateEnum
	// SetState assigns StateProtocolAllStateEnum provided by user to StateProtocolAll
	SetState(value StateProtocolAllStateEnum) StateProtocolAll
}

type StateProtocolAllStateEnum string

// Enum of State on StateProtocolAll
var StateProtocolAllState = struct {
	START StateProtocolAllStateEnum
	STOP  StateProtocolAllStateEnum
}{
	START: StateProtocolAllStateEnum("start"),
	STOP:  StateProtocolAllStateEnum("stop"),
}

func (obj *stateProtocolAll) State() StateProtocolAllStateEnum {
	return StateProtocolAllStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateProtocolAll) SetState(value StateProtocolAllStateEnum) StateProtocolAll {
	intValue, ok := otg.StateProtocolAll_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateProtocolAllStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateProtocolAll_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateProtocolAll) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateProtocolAll")
	}
}

func (obj *stateProtocolAll) setDefault() {

}
