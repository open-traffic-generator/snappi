package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateUltraEthernetLlr *****
type stateUltraEthernetLlr struct {
	validation
	obj          *otg.StateUltraEthernetLlr
	marshaller   marshalStateUltraEthernetLlr
	unMarshaller unMarshalStateUltraEthernetLlr
}

func NewStateUltraEthernetLlr() StateUltraEthernetLlr {
	obj := stateUltraEthernetLlr{obj: &otg.StateUltraEthernetLlr{}}
	obj.setDefault()
	return &obj
}

func (obj *stateUltraEthernetLlr) msg() *otg.StateUltraEthernetLlr {
	return obj.obj
}

func (obj *stateUltraEthernetLlr) setMsg(msg *otg.StateUltraEthernetLlr) StateUltraEthernetLlr {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateUltraEthernetLlr struct {
	obj *stateUltraEthernetLlr
}

type marshalStateUltraEthernetLlr interface {
	// ToProto marshals StateUltraEthernetLlr to protobuf object *otg.StateUltraEthernetLlr
	ToProto() (*otg.StateUltraEthernetLlr, error)
	// ToPbText marshals StateUltraEthernetLlr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateUltraEthernetLlr to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateUltraEthernetLlr to JSON text
	ToJson() (string, error)
}

type unMarshalstateUltraEthernetLlr struct {
	obj *stateUltraEthernetLlr
}

type unMarshalStateUltraEthernetLlr interface {
	// FromProto unmarshals StateUltraEthernetLlr from protobuf object *otg.StateUltraEthernetLlr
	FromProto(msg *otg.StateUltraEthernetLlr) (StateUltraEthernetLlr, error)
	// FromPbText unmarshals StateUltraEthernetLlr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateUltraEthernetLlr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateUltraEthernetLlr from JSON text
	FromJson(value string) error
}

func (obj *stateUltraEthernetLlr) Marshal() marshalStateUltraEthernetLlr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateUltraEthernetLlr{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateUltraEthernetLlr) Unmarshal() unMarshalStateUltraEthernetLlr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateUltraEthernetLlr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateUltraEthernetLlr) ToProto() (*otg.StateUltraEthernetLlr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateUltraEthernetLlr) FromProto(msg *otg.StateUltraEthernetLlr) (StateUltraEthernetLlr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateUltraEthernetLlr) ToPbText() (string, error) {
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

func (m *unMarshalstateUltraEthernetLlr) FromPbText(value string) error {
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

func (m *marshalstateUltraEthernetLlr) ToYaml() (string, error) {
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

func (m *unMarshalstateUltraEthernetLlr) FromYaml(value string) error {
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

func (m *marshalstateUltraEthernetLlr) ToJson() (string, error) {
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

func (m *unMarshalstateUltraEthernetLlr) FromJson(value string) error {
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

func (obj *stateUltraEthernetLlr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateUltraEthernetLlr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateUltraEthernetLlr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateUltraEthernetLlr) Clone() (StateUltraEthernetLlr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateUltraEthernetLlr()
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

// StateUltraEthernetLlr is sets the Link Layer Retry (LLR) mode state of configured Ultra Ethernet instances.
type StateUltraEthernetLlr interface {
	Validation
	// msg marshals StateUltraEthernetLlr to protobuf object *otg.StateUltraEthernetLlr
	// and doesn't set defaults
	msg() *otg.StateUltraEthernetLlr
	// setMsg unmarshals StateUltraEthernetLlr from protobuf object *otg.StateUltraEthernetLlr
	// and doesn't set defaults
	setMsg(*otg.StateUltraEthernetLlr) StateUltraEthernetLlr
	// provides marshal interface
	Marshal() marshalStateUltraEthernetLlr
	// provides unmarshal interface
	Unmarshal() unMarshalStateUltraEthernetLlr
	// validate validates StateUltraEthernetLlr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateUltraEthernetLlr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// UltraEthernetNames returns []string, set in StateUltraEthernetLlr.
	UltraEthernetNames() []string
	// SetUltraEthernetNames assigns []string provided by user to StateUltraEthernetLlr
	SetUltraEthernetNames(value []string) StateUltraEthernetLlr
	// State returns StateUltraEthernetLlrStateEnum, set in StateUltraEthernetLlr
	State() StateUltraEthernetLlrStateEnum
	// SetState assigns StateUltraEthernetLlrStateEnum provided by user to StateUltraEthernetLlr
	SetState(value StateUltraEthernetLlrStateEnum) StateUltraEthernetLlr
}

// The names of target Ultra Ethernet instances. An empty or null list will target all Ultra Ethernet instances.
//
// x-constraint:
// - /components/schemas/UltraEthernet/properties/name
//
// UltraEthernetNames returns a []string
func (obj *stateUltraEthernetLlr) UltraEthernetNames() []string {
	if obj.obj.UltraEthernetNames == nil {
		obj.obj.UltraEthernetNames = make([]string, 0)
	}
	return obj.obj.UltraEthernetNames
}

// The names of target Ultra Ethernet instances. An empty or null list will target all Ultra Ethernet instances.
//
// x-constraint:
// - /components/schemas/UltraEthernet/properties/name
//
// SetUltraEthernetNames sets the []string value in the StateUltraEthernetLlr object
func (obj *stateUltraEthernetLlr) SetUltraEthernetNames(value []string) StateUltraEthernetLlr {

	if obj.obj.UltraEthernetNames == nil {
		obj.obj.UltraEthernetNames = make([]string, 0)
	}
	obj.obj.UltraEthernetNames = value

	return obj
}

type StateUltraEthernetLlrStateEnum string

// Enum of State on StateUltraEthernetLlr
var StateUltraEthernetLlrState = struct {
	LOCAL_ENABLE   StateUltraEthernetLlrStateEnum
	LOCAL_DISABLE  StateUltraEthernetLlrStateEnum
	REMOTE_ENABLE  StateUltraEthernetLlrStateEnum
	REMOTE_DISABLE StateUltraEthernetLlrStateEnum
	ENABLE_ALL     StateUltraEthernetLlrStateEnum
	DISABLE_ALL    StateUltraEthernetLlrStateEnum
}{
	LOCAL_ENABLE:   StateUltraEthernetLlrStateEnum("local_enable"),
	LOCAL_DISABLE:  StateUltraEthernetLlrStateEnum("local_disable"),
	REMOTE_ENABLE:  StateUltraEthernetLlrStateEnum("remote_enable"),
	REMOTE_DISABLE: StateUltraEthernetLlrStateEnum("remote_disable"),
	ENABLE_ALL:     StateUltraEthernetLlrStateEnum("enable_all"),
	DISABLE_ALL:    StateUltraEthernetLlrStateEnum("disable_all"),
}

func (obj *stateUltraEthernetLlr) State() StateUltraEthernetLlrStateEnum {
	return StateUltraEthernetLlrStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateUltraEthernetLlr) SetState(value StateUltraEthernetLlrStateEnum) StateUltraEthernetLlr {
	intValue, ok := otg.StateUltraEthernetLlr_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateUltraEthernetLlrStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateUltraEthernetLlr_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateUltraEthernetLlr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateUltraEthernetLlr")
	}
}

func (obj *stateUltraEthernetLlr) setDefault() {

}
