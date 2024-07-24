package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatePortCapture *****
type statePortCapture struct {
	validation
	obj          *otg.StatePortCapture
	marshaller   marshalStatePortCapture
	unMarshaller unMarshalStatePortCapture
}

func NewStatePortCapture() StatePortCapture {
	obj := statePortCapture{obj: &otg.StatePortCapture{}}
	obj.setDefault()
	return &obj
}

func (obj *statePortCapture) msg() *otg.StatePortCapture {
	return obj.obj
}

func (obj *statePortCapture) setMsg(msg *otg.StatePortCapture) StatePortCapture {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatePortCapture struct {
	obj *statePortCapture
}

type marshalStatePortCapture interface {
	// ToProto marshals StatePortCapture to protobuf object *otg.StatePortCapture
	ToProto() (*otg.StatePortCapture, error)
	// ToPbText marshals StatePortCapture to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatePortCapture to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatePortCapture to JSON text
	ToJson() (string, error)
}

type unMarshalstatePortCapture struct {
	obj *statePortCapture
}

type unMarshalStatePortCapture interface {
	// FromProto unmarshals StatePortCapture from protobuf object *otg.StatePortCapture
	FromProto(msg *otg.StatePortCapture) (StatePortCapture, error)
	// FromPbText unmarshals StatePortCapture from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatePortCapture from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatePortCapture from JSON text
	FromJson(value string) error
}

func (obj *statePortCapture) Marshal() marshalStatePortCapture {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatePortCapture{obj: obj}
	}
	return obj.marshaller
}

func (obj *statePortCapture) Unmarshal() unMarshalStatePortCapture {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatePortCapture{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatePortCapture) ToProto() (*otg.StatePortCapture, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatePortCapture) FromProto(msg *otg.StatePortCapture) (StatePortCapture, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatePortCapture) ToPbText() (string, error) {
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

func (m *unMarshalstatePortCapture) FromPbText(value string) error {
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

func (m *marshalstatePortCapture) ToYaml() (string, error) {
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

func (m *unMarshalstatePortCapture) FromYaml(value string) error {
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

func (m *marshalstatePortCapture) ToJson() (string, error) {
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

func (m *unMarshalstatePortCapture) FromJson(value string) error {
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

func (obj *statePortCapture) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statePortCapture) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statePortCapture) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statePortCapture) Clone() (StatePortCapture, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatePortCapture()
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

// StatePortCapture is sets the capture state of configured ports
type StatePortCapture interface {
	Validation
	// msg marshals StatePortCapture to protobuf object *otg.StatePortCapture
	// and doesn't set defaults
	msg() *otg.StatePortCapture
	// setMsg unmarshals StatePortCapture from protobuf object *otg.StatePortCapture
	// and doesn't set defaults
	setMsg(*otg.StatePortCapture) StatePortCapture
	// provides marshal interface
	Marshal() marshalStatePortCapture
	// provides unmarshal interface
	Unmarshal() unMarshalStatePortCapture
	// validate validates StatePortCapture
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatePortCapture, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in StatePortCapture.
	PortNames() []string
	// SetPortNames assigns []string provided by user to StatePortCapture
	SetPortNames(value []string) StatePortCapture
	// State returns StatePortCaptureStateEnum, set in StatePortCapture
	State() StatePortCaptureStateEnum
	// SetState assigns StatePortCaptureStateEnum provided by user to StatePortCapture
	SetState(value StatePortCaptureStateEnum) StatePortCapture
}

// The names of ports to which the capture state will be applied to. If the list of port_names is empty or null the state will be applied to all configured ports.
// If the list is not empty any port that is not included in the list of port_names MUST be ignored and not included in the state change.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *statePortCapture) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// The names of ports to which the capture state will be applied to. If the list of port_names is empty or null the state will be applied to all configured ports.
// If the list is not empty any port that is not included in the list of port_names MUST be ignored and not included in the state change.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the StatePortCapture object
func (obj *statePortCapture) SetPortNames(value []string) StatePortCapture {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

type StatePortCaptureStateEnum string

// Enum of State on StatePortCapture
var StatePortCaptureState = struct {
	START StatePortCaptureStateEnum
	STOP  StatePortCaptureStateEnum
}{
	START: StatePortCaptureStateEnum("start"),
	STOP:  StatePortCaptureStateEnum("stop"),
}

func (obj *statePortCapture) State() StatePortCaptureStateEnum {
	return StatePortCaptureStateEnum(obj.obj.State.Enum().String())
}

func (obj *statePortCapture) SetState(value StatePortCaptureStateEnum) StatePortCapture {
	intValue, ok := otg.StatePortCapture_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatePortCaptureStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StatePortCapture_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *statePortCapture) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StatePortCapture")
	}
}

func (obj *statePortCapture) setDefault() {

}
