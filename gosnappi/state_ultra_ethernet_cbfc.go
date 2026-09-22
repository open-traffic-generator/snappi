package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateUltraEthernetCbfc *****
type stateUltraEthernetCbfc struct {
	validation
	obj          *otg.StateUltraEthernetCbfc
	marshaller   marshalStateUltraEthernetCbfc
	unMarshaller unMarshalStateUltraEthernetCbfc
}

func NewStateUltraEthernetCbfc() StateUltraEthernetCbfc {
	obj := stateUltraEthernetCbfc{obj: &otg.StateUltraEthernetCbfc{}}
	obj.setDefault()
	return &obj
}

func (obj *stateUltraEthernetCbfc) msg() *otg.StateUltraEthernetCbfc {
	return obj.obj
}

func (obj *stateUltraEthernetCbfc) setMsg(msg *otg.StateUltraEthernetCbfc) StateUltraEthernetCbfc {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateUltraEthernetCbfc struct {
	obj *stateUltraEthernetCbfc
}

type marshalStateUltraEthernetCbfc interface {
	// ToProto marshals StateUltraEthernetCbfc to protobuf object *otg.StateUltraEthernetCbfc
	ToProto() (*otg.StateUltraEthernetCbfc, error)
	// ToPbText marshals StateUltraEthernetCbfc to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateUltraEthernetCbfc to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateUltraEthernetCbfc to JSON text
	ToJson() (string, error)
}

type unMarshalstateUltraEthernetCbfc struct {
	obj *stateUltraEthernetCbfc
}

type unMarshalStateUltraEthernetCbfc interface {
	// FromProto unmarshals StateUltraEthernetCbfc from protobuf object *otg.StateUltraEthernetCbfc
	FromProto(msg *otg.StateUltraEthernetCbfc) (StateUltraEthernetCbfc, error)
	// FromPbText unmarshals StateUltraEthernetCbfc from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateUltraEthernetCbfc from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateUltraEthernetCbfc from JSON text
	FromJson(value string) error
}

func (obj *stateUltraEthernetCbfc) Marshal() marshalStateUltraEthernetCbfc {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateUltraEthernetCbfc{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateUltraEthernetCbfc) Unmarshal() unMarshalStateUltraEthernetCbfc {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateUltraEthernetCbfc{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateUltraEthernetCbfc) ToProto() (*otg.StateUltraEthernetCbfc, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateUltraEthernetCbfc) FromProto(msg *otg.StateUltraEthernetCbfc) (StateUltraEthernetCbfc, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateUltraEthernetCbfc) ToPbText() (string, error) {
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

func (m *unMarshalstateUltraEthernetCbfc) FromPbText(value string) error {
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

func (m *marshalstateUltraEthernetCbfc) ToYaml() (string, error) {
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

func (m *unMarshalstateUltraEthernetCbfc) FromYaml(value string) error {
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

func (m *marshalstateUltraEthernetCbfc) ToJson() (string, error) {
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

func (m *unMarshalstateUltraEthernetCbfc) FromJson(value string) error {
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

func (obj *stateUltraEthernetCbfc) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateUltraEthernetCbfc) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateUltraEthernetCbfc) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateUltraEthernetCbfc) Clone() (StateUltraEthernetCbfc, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateUltraEthernetCbfc()
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

// StateUltraEthernetCbfc is sets the Credit-based Flow Control (CBFC) state of configured Ultra Ethernet instances.
type StateUltraEthernetCbfc interface {
	Validation
	// msg marshals StateUltraEthernetCbfc to protobuf object *otg.StateUltraEthernetCbfc
	// and doesn't set defaults
	msg() *otg.StateUltraEthernetCbfc
	// setMsg unmarshals StateUltraEthernetCbfc from protobuf object *otg.StateUltraEthernetCbfc
	// and doesn't set defaults
	setMsg(*otg.StateUltraEthernetCbfc) StateUltraEthernetCbfc
	// provides marshal interface
	Marshal() marshalStateUltraEthernetCbfc
	// provides unmarshal interface
	Unmarshal() unMarshalStateUltraEthernetCbfc
	// validate validates StateUltraEthernetCbfc
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateUltraEthernetCbfc, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// UltraEthernetNames returns []string, set in StateUltraEthernetCbfc.
	UltraEthernetNames() []string
	// SetUltraEthernetNames assigns []string provided by user to StateUltraEthernetCbfc
	SetUltraEthernetNames(value []string) StateUltraEthernetCbfc
	// State returns StateUltraEthernetCbfcStateEnum, set in StateUltraEthernetCbfc
	State() StateUltraEthernetCbfcStateEnum
	// SetState assigns StateUltraEthernetCbfcStateEnum provided by user to StateUltraEthernetCbfc
	SetState(value StateUltraEthernetCbfcStateEnum) StateUltraEthernetCbfc
}

// The names of target Ultra Ethernet instances. An empty or null list will target all Ultra Ethernet instances.
//
// x-constraint:
// - /components/schemas/UltraEthernet/properties/name
//
// UltraEthernetNames returns a []string
func (obj *stateUltraEthernetCbfc) UltraEthernetNames() []string {
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
// SetUltraEthernetNames sets the []string value in the StateUltraEthernetCbfc object
func (obj *stateUltraEthernetCbfc) SetUltraEthernetNames(value []string) StateUltraEthernetCbfc {

	if obj.obj.UltraEthernetNames == nil {
		obj.obj.UltraEthernetNames = make([]string, 0)
	}
	obj.obj.UltraEthernetNames = value

	return obj
}

type StateUltraEthernetCbfcStateEnum string

// Enum of State on StateUltraEthernetCbfc
var StateUltraEthernetCbfcState = struct {
	SENDER_ENABLE    StateUltraEthernetCbfcStateEnum
	SENDER_DISABLE   StateUltraEthernetCbfcStateEnum
	RECEIVER_ENABLE  StateUltraEthernetCbfcStateEnum
	RECEIVER_DISABLE StateUltraEthernetCbfcStateEnum
	ENABLE_ALL       StateUltraEthernetCbfcStateEnum
	DISABLE_ALL      StateUltraEthernetCbfcStateEnum
}{
	SENDER_ENABLE:    StateUltraEthernetCbfcStateEnum("sender_enable"),
	SENDER_DISABLE:   StateUltraEthernetCbfcStateEnum("sender_disable"),
	RECEIVER_ENABLE:  StateUltraEthernetCbfcStateEnum("receiver_enable"),
	RECEIVER_DISABLE: StateUltraEthernetCbfcStateEnum("receiver_disable"),
	ENABLE_ALL:       StateUltraEthernetCbfcStateEnum("enable_all"),
	DISABLE_ALL:      StateUltraEthernetCbfcStateEnum("disable_all"),
}

func (obj *stateUltraEthernetCbfc) State() StateUltraEthernetCbfcStateEnum {
	return StateUltraEthernetCbfcStateEnum(obj.obj.State.Enum().String())
}

func (obj *stateUltraEthernetCbfc) SetState(value StateUltraEthernetCbfcStateEnum) StateUltraEthernetCbfc {
	intValue, ok := otg.StateUltraEthernetCbfc_State_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateUltraEthernetCbfcStateEnum", string(value)))
		return obj
	}
	enumValue := otg.StateUltraEthernetCbfc_State_Enum(intValue)
	obj.obj.State = &enumValue

	return obj
}

func (obj *stateUltraEthernetCbfc) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// State is required
	if obj.obj.State == nil {
		vObj.validationErrors = append(vObj.validationErrors, "State is required field on interface StateUltraEthernetCbfc")
	}
}

func (obj *stateUltraEthernetCbfc) setDefault() {

}
