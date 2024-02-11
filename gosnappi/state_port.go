package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StatePort *****
type statePort struct {
	validation
	obj           *otg.StatePort
	marshaller    marshalStatePort
	unMarshaller  unMarshalStatePort
	linkHolder    StatePortLink
	captureHolder StatePortCapture
}

func NewStatePort() StatePort {
	obj := statePort{obj: &otg.StatePort{}}
	obj.setDefault()
	return &obj
}

func (obj *statePort) msg() *otg.StatePort {
	return obj.obj
}

func (obj *statePort) setMsg(msg *otg.StatePort) StatePort {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstatePort struct {
	obj *statePort
}

type marshalStatePort interface {
	// ToProto marshals StatePort to protobuf object *otg.StatePort
	ToProto() (*otg.StatePort, error)
	// ToPbText marshals StatePort to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StatePort to YAML text
	ToYaml() (string, error)
	// ToJson marshals StatePort to JSON text
	ToJson() (string, error)
}

type unMarshalstatePort struct {
	obj *statePort
}

type unMarshalStatePort interface {
	// FromProto unmarshals StatePort from protobuf object *otg.StatePort
	FromProto(msg *otg.StatePort) (StatePort, error)
	// FromPbText unmarshals StatePort from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StatePort from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StatePort from JSON text
	FromJson(value string) error
}

func (obj *statePort) Marshal() marshalStatePort {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstatePort{obj: obj}
	}
	return obj.marshaller
}

func (obj *statePort) Unmarshal() unMarshalStatePort {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstatePort{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstatePort) ToProto() (*otg.StatePort, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstatePort) FromProto(msg *otg.StatePort) (StatePort, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstatePort) ToPbText() (string, error) {
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

func (m *unMarshalstatePort) FromPbText(value string) error {
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

func (m *marshalstatePort) ToYaml() (string, error) {
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

func (m *unMarshalstatePort) FromYaml(value string) error {
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

func (m *marshalstatePort) ToJson() (string, error) {
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

func (m *unMarshalstatePort) FromJson(value string) error {
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

func (obj *statePort) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *statePort) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *statePort) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *statePort) Clone() (StatePort, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStatePort()
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

func (obj *statePort) setNil() {
	obj.linkHolder = nil
	obj.captureHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StatePort is states associated with configured ports.
type StatePort interface {
	Validation
	// msg marshals StatePort to protobuf object *otg.StatePort
	// and doesn't set defaults
	msg() *otg.StatePort
	// setMsg unmarshals StatePort from protobuf object *otg.StatePort
	// and doesn't set defaults
	setMsg(*otg.StatePort) StatePort
	// provides marshal interface
	Marshal() marshalStatePort
	// provides unmarshal interface
	Unmarshal() unMarshalStatePort
	// validate validates StatePort
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StatePort, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StatePortChoiceEnum, set in StatePort
	Choice() StatePortChoiceEnum
	// setChoice assigns StatePortChoiceEnum provided by user to StatePort
	setChoice(value StatePortChoiceEnum) StatePort
	// Link returns StatePortLink, set in StatePort.
	// StatePortLink is sets the link of configured ports.
	Link() StatePortLink
	// SetLink assigns StatePortLink provided by user to StatePort.
	// StatePortLink is sets the link of configured ports.
	SetLink(value StatePortLink) StatePort
	// HasLink checks if Link has been set in StatePort
	HasLink() bool
	// Capture returns StatePortCapture, set in StatePort.
	// StatePortCapture is sets the capture state of configured ports
	Capture() StatePortCapture
	// SetCapture assigns StatePortCapture provided by user to StatePort.
	// StatePortCapture is sets the capture state of configured ports
	SetCapture(value StatePortCapture) StatePort
	// HasCapture checks if Capture has been set in StatePort
	HasCapture() bool
	setNil()
}

type StatePortChoiceEnum string

// Enum of Choice on StatePort
var StatePortChoice = struct {
	LINK    StatePortChoiceEnum
	CAPTURE StatePortChoiceEnum
}{
	LINK:    StatePortChoiceEnum("link"),
	CAPTURE: StatePortChoiceEnum("capture"),
}

func (obj *statePort) Choice() StatePortChoiceEnum {
	return StatePortChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *statePort) setChoice(value StatePortChoiceEnum) StatePort {
	intValue, ok := otg.StatePort_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StatePortChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StatePort_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Capture = nil
	obj.captureHolder = nil
	obj.obj.Link = nil
	obj.linkHolder = nil

	if value == StatePortChoice.LINK {
		obj.obj.Link = NewStatePortLink().msg()
	}

	if value == StatePortChoice.CAPTURE {
		obj.obj.Capture = NewStatePortCapture().msg()
	}

	return obj
}

// description is TBD
// Link returns a StatePortLink
func (obj *statePort) Link() StatePortLink {
	if obj.obj.Link == nil {
		obj.setChoice(StatePortChoice.LINK)
	}
	if obj.linkHolder == nil {
		obj.linkHolder = &statePortLink{obj: obj.obj.Link}
	}
	return obj.linkHolder
}

// description is TBD
// Link returns a StatePortLink
func (obj *statePort) HasLink() bool {
	return obj.obj.Link != nil
}

// description is TBD
// SetLink sets the StatePortLink value in the StatePort object
func (obj *statePort) SetLink(value StatePortLink) StatePort {
	obj.setChoice(StatePortChoice.LINK)
	obj.linkHolder = nil
	obj.obj.Link = value.msg()

	return obj
}

// description is TBD
// Capture returns a StatePortCapture
func (obj *statePort) Capture() StatePortCapture {
	if obj.obj.Capture == nil {
		obj.setChoice(StatePortChoice.CAPTURE)
	}
	if obj.captureHolder == nil {
		obj.captureHolder = &statePortCapture{obj: obj.obj.Capture}
	}
	return obj.captureHolder
}

// description is TBD
// Capture returns a StatePortCapture
func (obj *statePort) HasCapture() bool {
	return obj.obj.Capture != nil
}

// description is TBD
// SetCapture sets the StatePortCapture value in the StatePort object
func (obj *statePort) SetCapture(value StatePortCapture) StatePort {
	obj.setChoice(StatePortChoice.CAPTURE)
	obj.captureHolder = nil
	obj.obj.Capture = value.msg()

	return obj
}

func (obj *statePort) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StatePort")
	}

	if obj.obj.Link != nil {

		obj.Link().validateObj(vObj, set_default)
	}

	if obj.obj.Capture != nil {

		obj.Capture().validateObj(vObj, set_default)
	}

}

func (obj *statePort) setDefault() {

}
