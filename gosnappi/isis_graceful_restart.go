package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisGracefulRestart *****
type isisGracefulRestart struct {
	validation
	obj          *otg.IsisGracefulRestart
	marshaller   marshalIsisGracefulRestart
	unMarshaller unMarshalIsisGracefulRestart
}

func NewIsisGracefulRestart() IsisGracefulRestart {
	obj := isisGracefulRestart{obj: &otg.IsisGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *isisGracefulRestart) msg() *otg.IsisGracefulRestart {
	return obj.obj
}

func (obj *isisGracefulRestart) setMsg(msg *otg.IsisGracefulRestart) IsisGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisGracefulRestart struct {
	obj *isisGracefulRestart
}

type marshalIsisGracefulRestart interface {
	// ToProto marshals IsisGracefulRestart to protobuf object *otg.IsisGracefulRestart
	ToProto() (*otg.IsisGracefulRestart, error)
	// ToPbText marshals IsisGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalisisGracefulRestart struct {
	obj *isisGracefulRestart
}

type unMarshalIsisGracefulRestart interface {
	// FromProto unmarshals IsisGracefulRestart from protobuf object *otg.IsisGracefulRestart
	FromProto(msg *otg.IsisGracefulRestart) (IsisGracefulRestart, error)
	// FromPbText unmarshals IsisGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *isisGracefulRestart) Marshal() marshalIsisGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisGracefulRestart) Unmarshal() unMarshalIsisGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisGracefulRestart) ToProto() (*otg.IsisGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisGracefulRestart) FromProto(msg *otg.IsisGracefulRestart) (IsisGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromPbText(value string) error {
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

func (m *marshalisisGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromYaml(value string) error {
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

func (m *marshalisisGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalisisGracefulRestart) FromJson(value string) error {
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

func (obj *isisGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisGracefulRestart) Clone() (IsisGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisGracefulRestart()
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

// IsisGracefulRestart is contains IS-IS Graceful configuration parameters. https://datatracker.ietf.org/doc/html/rfc5306
type IsisGracefulRestart interface {
	Validation
	// msg marshals IsisGracefulRestart to protobuf object *otg.IsisGracefulRestart
	// and doesn't set defaults
	msg() *otg.IsisGracefulRestart
	// setMsg unmarshals IsisGracefulRestart from protobuf object *otg.IsisGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.IsisGracefulRestart) IsisGracefulRestart
	// provides marshal interface
	Marshal() marshalIsisGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalIsisGracefulRestart
	// validate validates IsisGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns IsisGracefulRestartChoiceEnum, set in IsisGracefulRestart
	Choice() IsisGracefulRestartChoiceEnum
	// setChoice assigns IsisGracefulRestartChoiceEnum provided by user to IsisGracefulRestart
	setChoice(value IsisGracefulRestartChoiceEnum) IsisGracefulRestart
	// HasChoice checks if Choice has been set in IsisGracefulRestart
	HasChoice() bool
	// getter for GracefulRestartMode to set choice.
	GracefulRestartMode()
	// getter for NormalMode to set choice.
	NormalMode()
	// getter for StartingMode to set choice.
	StartingMode()
	// getter for HelperMode to set choice.
	HelperMode()
	// RestartTime returns uint32, set in IsisGracefulRestart.
	RestartTime() uint32
	// SetRestartTime assigns uint32 provided by user to IsisGracefulRestart
	SetRestartTime(value uint32) IsisGracefulRestart
	// HasRestartTime checks if RestartTime has been set in IsisGracefulRestart
	HasRestartTime() bool
}

type IsisGracefulRestartChoiceEnum string

// Enum of Choice on IsisGracefulRestart
var IsisGracefulRestartChoice = struct {
	NORMAL_MODE           IsisGracefulRestartChoiceEnum
	STARTING_MODE         IsisGracefulRestartChoiceEnum
	GRACEFUL_RESTART_MODE IsisGracefulRestartChoiceEnum
	HELPER_MODE           IsisGracefulRestartChoiceEnum
}{
	NORMAL_MODE:           IsisGracefulRestartChoiceEnum("normal_mode"),
	STARTING_MODE:         IsisGracefulRestartChoiceEnum("starting_mode"),
	GRACEFUL_RESTART_MODE: IsisGracefulRestartChoiceEnum("graceful_restart_mode"),
	HELPER_MODE:           IsisGracefulRestartChoiceEnum("helper_mode"),
}

func (obj *isisGracefulRestart) Choice() IsisGracefulRestartChoiceEnum {
	return IsisGracefulRestartChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for GracefulRestartMode to set choice
func (obj *isisGracefulRestart) GracefulRestartMode() {
	obj.setChoice(IsisGracefulRestartChoice.GRACEFUL_RESTART_MODE)
}

// getter for NormalMode to set choice
func (obj *isisGracefulRestart) NormalMode() {
	obj.setChoice(IsisGracefulRestartChoice.NORMAL_MODE)
}

// getter for StartingMode to set choice
func (obj *isisGracefulRestart) StartingMode() {
	obj.setChoice(IsisGracefulRestartChoice.STARTING_MODE)
}

// getter for HelperMode to set choice
func (obj *isisGracefulRestart) HelperMode() {
	obj.setChoice(IsisGracefulRestartChoice.HELPER_MODE)
}

// Choice of Graceful Restart mode.
// - normal_mode: The emulated IS-IS router sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit unset.
// - starting_mode: The emulated IS-IS router is in the "Starting" mode.
// It sends to the neighbor router (DUT) an IIH containing a Restart TLV with the SA (Suppress Adjacency Advertisement) bit set.
// - graceful_restart: The emulated IS-IS router is in the unplanned "Restarting" mode.
// It sends to the neighbor router (DUT) an IIH containing a Restart TLV with the RR (Restart Request) bit set.
// - helper_mode: The emulated IS-IS router is acting as the "Helper" for the DUT that is restarting.
// It acknowledges the Restart TLV sent by the DUT by sending an IIH containing a Restart TLV with the RA (Restart Acknowledgment) bit set.
// Choice returns a string
func (obj *isisGracefulRestart) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *isisGracefulRestart) setChoice(value IsisGracefulRestartChoiceEnum) IsisGracefulRestart {
	intValue, ok := otg.IsisGracefulRestart_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on IsisGracefulRestartChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.IsisGracefulRestart_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue

	return obj
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart.
// RestartTime returns a uint32
func (obj *isisGracefulRestart) RestartTime() uint32 {

	return *obj.obj.RestartTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart.
// RestartTime returns a uint32
func (obj *isisGracefulRestart) HasRestartTime() bool {
	return obj.obj.RestartTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a start/restart.
// SetRestartTime sets the uint32 value in the IsisGracefulRestart object
func (obj *isisGracefulRestart) SetRestartTime(value uint32) IsisGracefulRestart {

	obj.obj.RestartTime = &value
	return obj
}

func (obj *isisGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTime != nil {

		if *obj.obj.RestartTime > 1 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= IsisGracefulRestart.RestartTime <= 1 but Got %d", *obj.obj.RestartTime))
		}

	}

}

func (obj *isisGracefulRestart) setDefault() {
	var choices_set int = 0
	var choice IsisGracefulRestartChoiceEnum
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(IsisGracefulRestartChoice.NORMAL_MODE)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in IsisGracefulRestart")
			}
		} else {
			intVal := otg.IsisGracefulRestart_Choice_Enum_value[string(choice)]
			enumValue := otg.IsisGracefulRestart_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.RestartTime == nil {
		obj.SetRestartTime(30)
	}

}
