package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmTlvs *****
type flowCfmTlvs struct {
	validation
	obj                   *otg.FlowCfmTlvs
	marshaller            marshalFlowCfmTlvs
	unMarshaller          unMarshalFlowCfmTlvs
	lengthHolder          PatternFlowCfmTlvsLength
	senderIdHolder        FlowCfmSenderIdTlv
	portStatusHolder      FlowCfmPortStatusTlv
	interfaceStatusHolder FlowCfmInterfaceStatusTlv
	dataHolder            FlowCfmDataTlv
}

func NewFlowCfmTlvs() FlowCfmTlvs {
	obj := flowCfmTlvs{obj: &otg.FlowCfmTlvs{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmTlvs) msg() *otg.FlowCfmTlvs {
	return obj.obj
}

func (obj *flowCfmTlvs) setMsg(msg *otg.FlowCfmTlvs) FlowCfmTlvs {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmTlvs struct {
	obj *flowCfmTlvs
}

type marshalFlowCfmTlvs interface {
	// ToProto marshals FlowCfmTlvs to protobuf object *otg.FlowCfmTlvs
	ToProto() (*otg.FlowCfmTlvs, error)
	// ToPbText marshals FlowCfmTlvs to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmTlvs to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmTlvs to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmTlvs struct {
	obj *flowCfmTlvs
}

type unMarshalFlowCfmTlvs interface {
	// FromProto unmarshals FlowCfmTlvs from protobuf object *otg.FlowCfmTlvs
	FromProto(msg *otg.FlowCfmTlvs) (FlowCfmTlvs, error)
	// FromPbText unmarshals FlowCfmTlvs from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmTlvs from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmTlvs from JSON text
	FromJson(value string) error
}

func (obj *flowCfmTlvs) Marshal() marshalFlowCfmTlvs {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmTlvs{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmTlvs) Unmarshal() unMarshalFlowCfmTlvs {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmTlvs{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmTlvs) ToProto() (*otg.FlowCfmTlvs, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmTlvs) FromProto(msg *otg.FlowCfmTlvs) (FlowCfmTlvs, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmTlvs) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmTlvs) FromPbText(value string) error {
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

func (m *marshalflowCfmTlvs) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmTlvs) FromYaml(value string) error {
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

func (m *marshalflowCfmTlvs) ToJson() (string, error) {
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

func (m *unMarshalflowCfmTlvs) FromJson(value string) error {
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

func (obj *flowCfmTlvs) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmTlvs) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmTlvs) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmTlvs) Clone() (FlowCfmTlvs, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmTlvs()
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

func (obj *flowCfmTlvs) setNil() {
	obj.lengthHolder = nil
	obj.senderIdHolder = nil
	obj.portStatusHolder = nil
	obj.interfaceStatusHolder = nil
	obj.dataHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmTlvs is the CFM TLV format is defined in 8021Q. Common fields are Type (1 octet) defined by choice below, Length (2 octet) and Value (variable) defined for each tlv type.
type FlowCfmTlvs interface {
	Validation
	// msg marshals FlowCfmTlvs to protobuf object *otg.FlowCfmTlvs
	// and doesn't set defaults
	msg() *otg.FlowCfmTlvs
	// setMsg unmarshals FlowCfmTlvs from protobuf object *otg.FlowCfmTlvs
	// and doesn't set defaults
	setMsg(*otg.FlowCfmTlvs) FlowCfmTlvs
	// provides marshal interface
	Marshal() marshalFlowCfmTlvs
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmTlvs
	// validate validates FlowCfmTlvs
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmTlvs, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns FlowCfmTlvsChoiceEnum, set in FlowCfmTlvs
	Choice() FlowCfmTlvsChoiceEnum
	// setChoice assigns FlowCfmTlvsChoiceEnum provided by user to FlowCfmTlvs
	setChoice(value FlowCfmTlvsChoiceEnum) FlowCfmTlvs
	// HasChoice checks if Choice has been set in FlowCfmTlvs
	HasChoice() bool
	// getter for End to set choice.
	End()
	// Length returns PatternFlowCfmTlvsLength, set in FlowCfmTlvs.
	// PatternFlowCfmTlvsLength is specifies the length of the tlv data in octets.
	Length() PatternFlowCfmTlvsLength
	// SetLength assigns PatternFlowCfmTlvsLength provided by user to FlowCfmTlvs.
	// PatternFlowCfmTlvsLength is specifies the length of the tlv data in octets.
	SetLength(value PatternFlowCfmTlvsLength) FlowCfmTlvs
	// HasLength checks if Length has been set in FlowCfmTlvs
	HasLength() bool
	// SenderId returns FlowCfmSenderIdTlv, set in FlowCfmTlvs.
	// FlowCfmSenderIdTlv is sender id
	SenderId() FlowCfmSenderIdTlv
	// SetSenderId assigns FlowCfmSenderIdTlv provided by user to FlowCfmTlvs.
	// FlowCfmSenderIdTlv is sender id
	SetSenderId(value FlowCfmSenderIdTlv) FlowCfmTlvs
	// HasSenderId checks if SenderId has been set in FlowCfmTlvs
	HasSenderId() bool
	// PortStatus returns FlowCfmPortStatusTlv, set in FlowCfmTlvs.
	// FlowCfmPortStatusTlv is port status
	PortStatus() FlowCfmPortStatusTlv
	// SetPortStatus assigns FlowCfmPortStatusTlv provided by user to FlowCfmTlvs.
	// FlowCfmPortStatusTlv is port status
	SetPortStatus(value FlowCfmPortStatusTlv) FlowCfmTlvs
	// HasPortStatus checks if PortStatus has been set in FlowCfmTlvs
	HasPortStatus() bool
	// InterfaceStatus returns FlowCfmInterfaceStatusTlv, set in FlowCfmTlvs.
	// FlowCfmInterfaceStatusTlv is interface status
	InterfaceStatus() FlowCfmInterfaceStatusTlv
	// SetInterfaceStatus assigns FlowCfmInterfaceStatusTlv provided by user to FlowCfmTlvs.
	// FlowCfmInterfaceStatusTlv is interface status
	SetInterfaceStatus(value FlowCfmInterfaceStatusTlv) FlowCfmTlvs
	// HasInterfaceStatus checks if InterfaceStatus has been set in FlowCfmTlvs
	HasInterfaceStatus() bool
	// Data returns FlowCfmDataTlv, set in FlowCfmTlvs.
	// FlowCfmDataTlv is data tlv
	Data() FlowCfmDataTlv
	// SetData assigns FlowCfmDataTlv provided by user to FlowCfmTlvs.
	// FlowCfmDataTlv is data tlv
	SetData(value FlowCfmDataTlv) FlowCfmTlvs
	// HasData checks if Data has been set in FlowCfmTlvs
	HasData() bool
	setNil()
}

type FlowCfmTlvsChoiceEnum string

// Enum of Choice on FlowCfmTlvs
var FlowCfmTlvsChoice = struct {
	END              FlowCfmTlvsChoiceEnum
	SENDER_ID        FlowCfmTlvsChoiceEnum
	PORT_STATUS      FlowCfmTlvsChoiceEnum
	INTERFACE_STATUS FlowCfmTlvsChoiceEnum
	DATA             FlowCfmTlvsChoiceEnum
}{
	END:              FlowCfmTlvsChoiceEnum("end"),
	SENDER_ID:        FlowCfmTlvsChoiceEnum("sender_id"),
	PORT_STATUS:      FlowCfmTlvsChoiceEnum("port_status"),
	INTERFACE_STATUS: FlowCfmTlvsChoiceEnum("interface_status"),
	DATA:             FlowCfmTlvsChoiceEnum("data"),
}

func (obj *flowCfmTlvs) Choice() FlowCfmTlvsChoiceEnum {
	return FlowCfmTlvsChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for End to set choice
func (obj *flowCfmTlvs) End() {
	obj.setChoice(FlowCfmTlvsChoice.END)
}

// description is TBD
// Choice returns a string
func (obj *flowCfmTlvs) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *flowCfmTlvs) setChoice(value FlowCfmTlvsChoiceEnum) FlowCfmTlvs {
	intValue, ok := otg.FlowCfmTlvs_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmTlvsChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmTlvs_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Data = nil
	obj.dataHolder = nil
	obj.obj.InterfaceStatus = nil
	obj.interfaceStatusHolder = nil
	obj.obj.PortStatus = nil
	obj.portStatusHolder = nil
	obj.obj.SenderId = nil
	obj.senderIdHolder = nil

	if value == FlowCfmTlvsChoice.SENDER_ID {
		obj.obj.SenderId = NewFlowCfmSenderIdTlv().msg()
	}

	if value == FlowCfmTlvsChoice.PORT_STATUS {
		obj.obj.PortStatus = NewFlowCfmPortStatusTlv().msg()
	}

	if value == FlowCfmTlvsChoice.INTERFACE_STATUS {
		obj.obj.InterfaceStatus = NewFlowCfmInterfaceStatusTlv().msg()
	}

	if value == FlowCfmTlvsChoice.DATA {
		obj.obj.Data = NewFlowCfmDataTlv().msg()
	}

	return obj
}

// description is TBD
// Length returns a PatternFlowCfmTlvsLength
func (obj *flowCfmTlvs) Length() PatternFlowCfmTlvsLength {
	if obj.obj.Length == nil {
		obj.obj.Length = NewPatternFlowCfmTlvsLength().msg()
	}
	if obj.lengthHolder == nil {
		obj.lengthHolder = &patternFlowCfmTlvsLength{obj: obj.obj.Length}
	}
	return obj.lengthHolder
}

// description is TBD
// Length returns a PatternFlowCfmTlvsLength
func (obj *flowCfmTlvs) HasLength() bool {
	return obj.obj.Length != nil
}

// description is TBD
// SetLength sets the PatternFlowCfmTlvsLength value in the FlowCfmTlvs object
func (obj *flowCfmTlvs) SetLength(value PatternFlowCfmTlvsLength) FlowCfmTlvs {

	obj.lengthHolder = nil
	obj.obj.Length = value.msg()

	return obj
}

// description is TBD
// SenderId returns a FlowCfmSenderIdTlv
func (obj *flowCfmTlvs) SenderId() FlowCfmSenderIdTlv {
	if obj.obj.SenderId == nil {
		obj.setChoice(FlowCfmTlvsChoice.SENDER_ID)
	}
	if obj.senderIdHolder == nil {
		obj.senderIdHolder = &flowCfmSenderIdTlv{obj: obj.obj.SenderId}
	}
	return obj.senderIdHolder
}

// description is TBD
// SenderId returns a FlowCfmSenderIdTlv
func (obj *flowCfmTlvs) HasSenderId() bool {
	return obj.obj.SenderId != nil
}

// description is TBD
// SetSenderId sets the FlowCfmSenderIdTlv value in the FlowCfmTlvs object
func (obj *flowCfmTlvs) SetSenderId(value FlowCfmSenderIdTlv) FlowCfmTlvs {
	obj.setChoice(FlowCfmTlvsChoice.SENDER_ID)
	obj.senderIdHolder = nil
	obj.obj.SenderId = value.msg()

	return obj
}

// description is TBD
// PortStatus returns a FlowCfmPortStatusTlv
func (obj *flowCfmTlvs) PortStatus() FlowCfmPortStatusTlv {
	if obj.obj.PortStatus == nil {
		obj.setChoice(FlowCfmTlvsChoice.PORT_STATUS)
	}
	if obj.portStatusHolder == nil {
		obj.portStatusHolder = &flowCfmPortStatusTlv{obj: obj.obj.PortStatus}
	}
	return obj.portStatusHolder
}

// description is TBD
// PortStatus returns a FlowCfmPortStatusTlv
func (obj *flowCfmTlvs) HasPortStatus() bool {
	return obj.obj.PortStatus != nil
}

// description is TBD
// SetPortStatus sets the FlowCfmPortStatusTlv value in the FlowCfmTlvs object
func (obj *flowCfmTlvs) SetPortStatus(value FlowCfmPortStatusTlv) FlowCfmTlvs {
	obj.setChoice(FlowCfmTlvsChoice.PORT_STATUS)
	obj.portStatusHolder = nil
	obj.obj.PortStatus = value.msg()

	return obj
}

// description is TBD
// InterfaceStatus returns a FlowCfmInterfaceStatusTlv
func (obj *flowCfmTlvs) InterfaceStatus() FlowCfmInterfaceStatusTlv {
	if obj.obj.InterfaceStatus == nil {
		obj.setChoice(FlowCfmTlvsChoice.INTERFACE_STATUS)
	}
	if obj.interfaceStatusHolder == nil {
		obj.interfaceStatusHolder = &flowCfmInterfaceStatusTlv{obj: obj.obj.InterfaceStatus}
	}
	return obj.interfaceStatusHolder
}

// description is TBD
// InterfaceStatus returns a FlowCfmInterfaceStatusTlv
func (obj *flowCfmTlvs) HasInterfaceStatus() bool {
	return obj.obj.InterfaceStatus != nil
}

// description is TBD
// SetInterfaceStatus sets the FlowCfmInterfaceStatusTlv value in the FlowCfmTlvs object
func (obj *flowCfmTlvs) SetInterfaceStatus(value FlowCfmInterfaceStatusTlv) FlowCfmTlvs {
	obj.setChoice(FlowCfmTlvsChoice.INTERFACE_STATUS)
	obj.interfaceStatusHolder = nil
	obj.obj.InterfaceStatus = value.msg()

	return obj
}

// description is TBD
// Data returns a FlowCfmDataTlv
func (obj *flowCfmTlvs) Data() FlowCfmDataTlv {
	if obj.obj.Data == nil {
		obj.setChoice(FlowCfmTlvsChoice.DATA)
	}
	if obj.dataHolder == nil {
		obj.dataHolder = &flowCfmDataTlv{obj: obj.obj.Data}
	}
	return obj.dataHolder
}

// description is TBD
// Data returns a FlowCfmDataTlv
func (obj *flowCfmTlvs) HasData() bool {
	return obj.obj.Data != nil
}

// description is TBD
// SetData sets the FlowCfmDataTlv value in the FlowCfmTlvs object
func (obj *flowCfmTlvs) SetData(value FlowCfmDataTlv) FlowCfmTlvs {
	obj.setChoice(FlowCfmTlvsChoice.DATA)
	obj.dataHolder = nil
	obj.obj.Data = value.msg()

	return obj
}

func (obj *flowCfmTlvs) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Length != nil {

		obj.Length().validateObj(vObj, set_default)
	}

	if obj.obj.SenderId != nil {

		obj.SenderId().validateObj(vObj, set_default)
	}

	if obj.obj.PortStatus != nil {

		obj.PortStatus().validateObj(vObj, set_default)
	}

	if obj.obj.InterfaceStatus != nil {

		obj.InterfaceStatus().validateObj(vObj, set_default)
	}

	if obj.obj.Data != nil {

		obj.Data().validateObj(vObj, set_default)
	}

}

func (obj *flowCfmTlvs) setDefault() {
	var choices_set int = 0
	var choice FlowCfmTlvsChoiceEnum

	if obj.obj.SenderId != nil {
		choices_set += 1
		choice = FlowCfmTlvsChoice.SENDER_ID
	}

	if obj.obj.PortStatus != nil {
		choices_set += 1
		choice = FlowCfmTlvsChoice.PORT_STATUS
	}

	if obj.obj.InterfaceStatus != nil {
		choices_set += 1
		choice = FlowCfmTlvsChoice.INTERFACE_STATUS
	}

	if obj.obj.Data != nil {
		choices_set += 1
		choice = FlowCfmTlvsChoice.DATA
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(FlowCfmTlvsChoice.END)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in FlowCfmTlvs")
			}
		} else {
			intVal := otg.FlowCfmTlvs_Choice_Enum_value[string(choice)]
			enumValue := otg.FlowCfmTlvs_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
