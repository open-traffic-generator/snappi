package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowTxRx *****
type flowTxRx struct {
	validation
	obj          *otg.FlowTxRx
	marshaller   marshalFlowTxRx
	unMarshaller unMarshalFlowTxRx
	portHolder   FlowPort
	deviceHolder FlowRouter
}

func NewFlowTxRx() FlowTxRx {
	obj := flowTxRx{obj: &otg.FlowTxRx{}}
	obj.setDefault()
	return &obj
}

func (obj *flowTxRx) msg() *otg.FlowTxRx {
	return obj.obj
}

func (obj *flowTxRx) setMsg(msg *otg.FlowTxRx) FlowTxRx {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowTxRx struct {
	obj *flowTxRx
}

type marshalFlowTxRx interface {
	// ToProto marshals FlowTxRx to protobuf object *otg.FlowTxRx
	ToProto() (*otg.FlowTxRx, error)
	// ToPbText marshals FlowTxRx to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowTxRx to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowTxRx to JSON text
	ToJson() (string, error)
}

type unMarshalflowTxRx struct {
	obj *flowTxRx
}

type unMarshalFlowTxRx interface {
	// FromProto unmarshals FlowTxRx from protobuf object *otg.FlowTxRx
	FromProto(msg *otg.FlowTxRx) (FlowTxRx, error)
	// FromPbText unmarshals FlowTxRx from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowTxRx from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowTxRx from JSON text
	FromJson(value string) error
}

func (obj *flowTxRx) Marshal() marshalFlowTxRx {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowTxRx{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowTxRx) Unmarshal() unMarshalFlowTxRx {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowTxRx{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowTxRx) ToProto() (*otg.FlowTxRx, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowTxRx) FromProto(msg *otg.FlowTxRx) (FlowTxRx, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowTxRx) ToPbText() (string, error) {
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

func (m *unMarshalflowTxRx) FromPbText(value string) error {
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

func (m *marshalflowTxRx) ToYaml() (string, error) {
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

func (m *unMarshalflowTxRx) FromYaml(value string) error {
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

func (m *marshalflowTxRx) ToJson() (string, error) {
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

func (m *unMarshalflowTxRx) FromJson(value string) error {
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

func (obj *flowTxRx) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowTxRx) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowTxRx) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowTxRx) Clone() (FlowTxRx, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowTxRx()
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

func (obj *flowTxRx) setNil() {
	obj.portHolder = nil
	obj.deviceHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowTxRx is a container for different types of transmit and receive
// endpoint containers.
type FlowTxRx interface {
	Validation
	// msg marshals FlowTxRx to protobuf object *otg.FlowTxRx
	// and doesn't set defaults
	msg() *otg.FlowTxRx
	// setMsg unmarshals FlowTxRx from protobuf object *otg.FlowTxRx
	// and doesn't set defaults
	setMsg(*otg.FlowTxRx) FlowTxRx
	// provides marshal interface
	Marshal() marshalFlowTxRx
	// provides unmarshal interface
	Unmarshal() unMarshalFlowTxRx
	// validate validates FlowTxRx
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowTxRx, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowTxRxChoiceEnum, set in FlowTxRx
	Choice() FlowTxRxChoiceEnum
	// setChoice assigns FlowTxRxChoiceEnum provided by user to FlowTxRx
	setChoice(value FlowTxRxChoiceEnum) FlowTxRx
	// HasChoice checks if Choice has been set in FlowTxRx
	HasChoice() bool
	// Port returns FlowPort, set in FlowTxRx.
	// FlowPort is a container for a transmit port and 0..n intended receive ports.
	// When assigning this container to a flow the flows's
	// packet headers will not be populated with any address resolution
	// information such as source and/or destination addresses.
	// For example Flow.Ethernet dst mac address values will be defaulted to 0.
	// For full control over the Flow.properties.packet header contents use this
	// container.
	Port() FlowPort
	// SetPort assigns FlowPort provided by user to FlowTxRx.
	// FlowPort is a container for a transmit port and 0..n intended receive ports.
	// When assigning this container to a flow the flows's
	// packet headers will not be populated with any address resolution
	// information such as source and/or destination addresses.
	// For example Flow.Ethernet dst mac address values will be defaulted to 0.
	// For full control over the Flow.properties.packet header contents use this
	// container.
	SetPort(value FlowPort) FlowTxRx
	// HasPort checks if Port has been set in FlowTxRx
	HasPort() bool
	// Device returns FlowRouter, set in FlowTxRx.
	// FlowRouter is a container for declaring a map of 1..n transmit devices to 1..n receive devices. This allows for a single flow to have  different tx to rx device flows such as a single one to one map or a  many to many map.
	Device() FlowRouter
	// SetDevice assigns FlowRouter provided by user to FlowTxRx.
	// FlowRouter is a container for declaring a map of 1..n transmit devices to 1..n receive devices. This allows for a single flow to have  different tx to rx device flows such as a single one to one map or a  many to many map.
	SetDevice(value FlowRouter) FlowTxRx
	// HasDevice checks if Device has been set in FlowTxRx
	HasDevice() bool
	setNil()
}

type FlowTxRxChoiceEnum string

// Enum of Choice on FlowTxRx
var FlowTxRxChoice = struct {
	PORT   FlowTxRxChoiceEnum
	DEVICE FlowTxRxChoiceEnum
}{
	PORT:   FlowTxRxChoiceEnum("port"),
	DEVICE: FlowTxRxChoiceEnum("device"),
}

func (obj *flowTxRx) Choice() FlowTxRxChoiceEnum {
	return FlowTxRxChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of transmit and receive container used by the flow.
// Choice returns a string
func (obj *flowTxRx) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowTxRx) setChoice(value FlowTxRxChoiceEnum) FlowTxRx {
	intValue, ok := otg.FlowTxRx_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowTxRxChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowTxRx_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Device = nil
	obj.deviceHolder = nil
	obj.obj.Port = nil
	obj.portHolder = nil

	if value == FlowTxRxChoice.PORT {
		obj.obj.Port = NewFlowPort().msg()
	}

	if value == FlowTxRxChoice.DEVICE {
		obj.obj.Device = NewFlowRouter().msg()
	}

	return obj
}

// description is TBD
// Port returns a FlowPort
func (obj *flowTxRx) Port() FlowPort {
	if obj.obj.Port == nil {
		obj.setChoice(FlowTxRxChoice.PORT)
	}
	if obj.portHolder == nil {
		obj.portHolder = &flowPort{obj: obj.obj.Port}
	}
	return obj.portHolder
}

// description is TBD
// Port returns a FlowPort
func (obj *flowTxRx) HasPort() bool {
	return obj.obj.Port != nil
}

// description is TBD
// SetPort sets the FlowPort value in the FlowTxRx object
func (obj *flowTxRx) SetPort(value FlowPort) FlowTxRx {
	obj.setChoice(FlowTxRxChoice.PORT)
	obj.portHolder = nil
	obj.obj.Port = value.msg()

	return obj
}

// description is TBD
// Device returns a FlowRouter
func (obj *flowTxRx) Device() FlowRouter {
	if obj.obj.Device == nil {
		obj.setChoice(FlowTxRxChoice.DEVICE)
	}
	if obj.deviceHolder == nil {
		obj.deviceHolder = &flowRouter{obj: obj.obj.Device}
	}
	return obj.deviceHolder
}

// description is TBD
// Device returns a FlowRouter
func (obj *flowTxRx) HasDevice() bool {
	return obj.obj.Device != nil
}

// description is TBD
// SetDevice sets the FlowRouter value in the FlowTxRx object
func (obj *flowTxRx) SetDevice(value FlowRouter) FlowTxRx {
	obj.setChoice(FlowTxRxChoice.DEVICE)
	obj.deviceHolder = nil
	obj.obj.Device = value.msg()

	return obj
}

func (obj *flowTxRx) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Port != nil {

		obj.Port().validateObj(vObj, set_default)
	}

	if obj.obj.Device != nil {

		obj.Device().validateObj(vObj, set_default)
	}

}

func (obj *flowTxRx) setDefault() {
	var choices_set int = 0
	var choice FlowTxRxChoiceEnum

	if obj.obj.Port != nil {
		choices_set += 1
		choice = FlowTxRxChoice.PORT
	}

	if obj.obj.Device != nil {
		choices_set += 1
		choice = FlowTxRxChoice.DEVICE
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowTxRxChoice.PORT)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowTxRx")
			}
		} else {
			intVal := otg.FlowTxRx_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowTxRx_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
