package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** StateUltraEthernet *****
type stateUltraEthernet struct {
	validation
	obj          *otg.StateUltraEthernet
	marshaller   marshalStateUltraEthernet
	unMarshaller unMarshalStateUltraEthernet
	llrHolder    StateUltraEthernetLlr
	cbfcHolder   StateUltraEthernetCbfc
}

func NewStateUltraEthernet() StateUltraEthernet {
	obj := stateUltraEthernet{obj: &otg.StateUltraEthernet{}}
	obj.setDefault()
	return &obj
}

func (obj *stateUltraEthernet) msg() *otg.StateUltraEthernet {
	return obj.obj
}

func (obj *stateUltraEthernet) setMsg(msg *otg.StateUltraEthernet) StateUltraEthernet {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalstateUltraEthernet struct {
	obj *stateUltraEthernet
}

type marshalStateUltraEthernet interface {
	// ToProto marshals StateUltraEthernet to protobuf object *otg.StateUltraEthernet
	ToProto() (*otg.StateUltraEthernet, error)
	// ToPbText marshals StateUltraEthernet to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals StateUltraEthernet to YAML text
	ToYaml() (string, error)
	// ToJson marshals StateUltraEthernet to JSON text
	ToJson() (string, error)
}

type unMarshalstateUltraEthernet struct {
	obj *stateUltraEthernet
}

type unMarshalStateUltraEthernet interface {
	// FromProto unmarshals StateUltraEthernet from protobuf object *otg.StateUltraEthernet
	FromProto(msg *otg.StateUltraEthernet) (StateUltraEthernet, error)
	// FromPbText unmarshals StateUltraEthernet from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals StateUltraEthernet from YAML text
	FromYaml(value string) error
	// FromJson unmarshals StateUltraEthernet from JSON text
	FromJson(value string) error
}

func (obj *stateUltraEthernet) Marshal() marshalStateUltraEthernet {
	if obj.marshaller == nil {
		obj.marshaller = &marshalstateUltraEthernet{obj: obj}
	}
	return obj.marshaller
}

func (obj *stateUltraEthernet) Unmarshal() unMarshalStateUltraEthernet {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalstateUltraEthernet{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalstateUltraEthernet) ToProto() (*otg.StateUltraEthernet, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalstateUltraEthernet) FromProto(msg *otg.StateUltraEthernet) (StateUltraEthernet, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalstateUltraEthernet) ToPbText() (string, error) {
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

func (m *unMarshalstateUltraEthernet) FromPbText(value string) error {
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

func (m *marshalstateUltraEthernet) ToYaml() (string, error) {
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

func (m *unMarshalstateUltraEthernet) FromYaml(value string) error {
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

func (m *marshalstateUltraEthernet) ToJson() (string, error) {
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

func (m *unMarshalstateUltraEthernet) FromJson(value string) error {
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

func (obj *stateUltraEthernet) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *stateUltraEthernet) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *stateUltraEthernet) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *stateUltraEthernet) Clone() (StateUltraEthernet, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewStateUltraEthernet()
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

func (obj *stateUltraEthernet) setNil() {
	obj.llrHolder = nil
	obj.cbfcHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// StateUltraEthernet is states associated with Ultra Ethernet link layer features on configured resources.
type StateUltraEthernet interface {
	Validation
	// msg marshals StateUltraEthernet to protobuf object *otg.StateUltraEthernet
	// and doesn't set defaults
	msg() *otg.StateUltraEthernet
	// setMsg unmarshals StateUltraEthernet from protobuf object *otg.StateUltraEthernet
	// and doesn't set defaults
	setMsg(*otg.StateUltraEthernet) StateUltraEthernet
	// provides marshal interface
	Marshal() marshalStateUltraEthernet
	// provides unmarshal interface
	Unmarshal() unMarshalStateUltraEthernet
	// validate validates StateUltraEthernet
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (StateUltraEthernet, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns StateUltraEthernetChoiceEnum, set in StateUltraEthernet
	Choice() StateUltraEthernetChoiceEnum
	// setChoice assigns StateUltraEthernetChoiceEnum provided by user to StateUltraEthernet
	setChoice(value StateUltraEthernetChoiceEnum) StateUltraEthernet
	// Llr returns StateUltraEthernetLlr, set in StateUltraEthernet.
	// StateUltraEthernetLlr is sets the Link Layer Retry (LLR) mode state of configured Ultra Ethernet instances.
	Llr() StateUltraEthernetLlr
	// SetLlr assigns StateUltraEthernetLlr provided by user to StateUltraEthernet.
	// StateUltraEthernetLlr is sets the Link Layer Retry (LLR) mode state of configured Ultra Ethernet instances.
	SetLlr(value StateUltraEthernetLlr) StateUltraEthernet
	// HasLlr checks if Llr has been set in StateUltraEthernet
	HasLlr() bool
	// Cbfc returns StateUltraEthernetCbfc, set in StateUltraEthernet.
	// StateUltraEthernetCbfc is sets the Credit-based Flow Control (CBFC) state of configured Ultra Ethernet instances.
	Cbfc() StateUltraEthernetCbfc
	// SetCbfc assigns StateUltraEthernetCbfc provided by user to StateUltraEthernet.
	// StateUltraEthernetCbfc is sets the Credit-based Flow Control (CBFC) state of configured Ultra Ethernet instances.
	SetCbfc(value StateUltraEthernetCbfc) StateUltraEthernet
	// HasCbfc checks if Cbfc has been set in StateUltraEthernet
	HasCbfc() bool
	setNil()
}

type StateUltraEthernetChoiceEnum string

// Enum of Choice on StateUltraEthernet
var StateUltraEthernetChoice = struct {
	LLR  StateUltraEthernetChoiceEnum
	CBFC StateUltraEthernetChoiceEnum
}{
	LLR:  StateUltraEthernetChoiceEnum("llr"),
	CBFC: StateUltraEthernetChoiceEnum("cbfc"),
}

func (obj *stateUltraEthernet) Choice() StateUltraEthernetChoiceEnum {
	return StateUltraEthernetChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *stateUltraEthernet) setChoice(value StateUltraEthernetChoiceEnum) StateUltraEthernet {
	intValue, ok := otg.StateUltraEthernet_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on StateUltraEthernetChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.StateUltraEthernet_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Cbfc = nil
	obj.cbfcHolder = nil
	obj.obj.Llr = nil
	obj.llrHolder = nil

	if value == StateUltraEthernetChoice.LLR {
		obj.obj.Llr = NewStateUltraEthernetLlr().msg()
	}

	if value == StateUltraEthernetChoice.CBFC {
		obj.obj.Cbfc = NewStateUltraEthernetCbfc().msg()
	}

	return obj
}

// description is TBD
// Llr returns a StateUltraEthernetLlr
func (obj *stateUltraEthernet) Llr() StateUltraEthernetLlr {
	if obj.obj.Llr == nil {
		obj.setChoice(StateUltraEthernetChoice.LLR)
	}
	if obj.llrHolder == nil {
		obj.llrHolder = &stateUltraEthernetLlr{obj: obj.obj.Llr}
	}
	return obj.llrHolder
}

// description is TBD
// Llr returns a StateUltraEthernetLlr
func (obj *stateUltraEthernet) HasLlr() bool {
	return obj.obj.Llr != nil
}

// description is TBD
// SetLlr sets the StateUltraEthernetLlr value in the StateUltraEthernet object
func (obj *stateUltraEthernet) SetLlr(value StateUltraEthernetLlr) StateUltraEthernet {
	obj.setChoice(StateUltraEthernetChoice.LLR)
	obj.llrHolder = nil
	obj.obj.Llr = value.msg()

	return obj
}

// description is TBD
// Cbfc returns a StateUltraEthernetCbfc
func (obj *stateUltraEthernet) Cbfc() StateUltraEthernetCbfc {
	if obj.obj.Cbfc == nil {
		obj.setChoice(StateUltraEthernetChoice.CBFC)
	}
	if obj.cbfcHolder == nil {
		obj.cbfcHolder = &stateUltraEthernetCbfc{obj: obj.obj.Cbfc}
	}
	return obj.cbfcHolder
}

// description is TBD
// Cbfc returns a StateUltraEthernetCbfc
func (obj *stateUltraEthernet) HasCbfc() bool {
	return obj.obj.Cbfc != nil
}

// description is TBD
// SetCbfc sets the StateUltraEthernetCbfc value in the StateUltraEthernet object
func (obj *stateUltraEthernet) SetCbfc(value StateUltraEthernetCbfc) StateUltraEthernet {
	obj.setChoice(StateUltraEthernetChoice.CBFC)
	obj.cbfcHolder = nil
	obj.obj.Cbfc = value.msg()

	return obj
}

func (obj *stateUltraEthernet) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface StateUltraEthernet")
	}

	if obj.obj.Llr != nil {

		obj.Llr().validateObj(vObj, set_default)
	}

	if obj.obj.Cbfc != nil {

		obj.Cbfc().validateObj(vObj, set_default)
	}

}

func (obj *stateUltraEthernet) setDefault() {
	var choices_set int = 0
	var choice StateUltraEthernetChoiceEnum

	if obj.obj.Llr != nil {
		choices_set += 1
		choice = StateUltraEthernetChoice.LLR
	}

	if obj.obj.Cbfc != nil {
		choices_set += 1
		choice = StateUltraEthernetChoice.CBFC
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in StateUltraEthernet")
			}
		} else {
			intVal := otg.StateUltraEthernet_Choice_Enum_value[string(choice)]
			enumValue := otg.StateUltraEthernet_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
