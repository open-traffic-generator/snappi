package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** FlowCfmSenderIdTlv *****
type flowCfmSenderIdTlv struct {
	validation
	obj                           *otg.FlowCfmSenderIdTlv
	marshaller                    marshalFlowCfmSenderIdTlv
	unMarshaller                  unMarshalFlowCfmSenderIdTlv
	chassisIdLengthHolder         PatternFlowCfmSenderIdTlvChassisIdLength
	managementAddressLengthHolder PatternFlowCfmSenderIdTlvManagementAddressLength
}

func NewFlowCfmSenderIdTlv() FlowCfmSenderIdTlv {
	obj := flowCfmSenderIdTlv{obj: &otg.FlowCfmSenderIdTlv{}}
	obj.setDefault()
	return &obj
}

func (obj *flowCfmSenderIdTlv) msg() *otg.FlowCfmSenderIdTlv {
	return obj.obj
}

func (obj *flowCfmSenderIdTlv) setMsg(msg *otg.FlowCfmSenderIdTlv) FlowCfmSenderIdTlv {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalflowCfmSenderIdTlv struct {
	obj *flowCfmSenderIdTlv
}

type marshalFlowCfmSenderIdTlv interface {
	// ToProto marshals FlowCfmSenderIdTlv to protobuf object *otg.FlowCfmSenderIdTlv
	ToProto() (*otg.FlowCfmSenderIdTlv, error)
	// ToPbText marshals FlowCfmSenderIdTlv to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals FlowCfmSenderIdTlv to YAML text
	ToYaml() (string, error)
	// ToJson marshals FlowCfmSenderIdTlv to JSON text
	ToJson() (string, error)
}

type unMarshalflowCfmSenderIdTlv struct {
	obj *flowCfmSenderIdTlv
}

type unMarshalFlowCfmSenderIdTlv interface {
	// FromProto unmarshals FlowCfmSenderIdTlv from protobuf object *otg.FlowCfmSenderIdTlv
	FromProto(msg *otg.FlowCfmSenderIdTlv) (FlowCfmSenderIdTlv, error)
	// FromPbText unmarshals FlowCfmSenderIdTlv from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals FlowCfmSenderIdTlv from YAML text
	FromYaml(value string) error
	// FromJson unmarshals FlowCfmSenderIdTlv from JSON text
	FromJson(value string) error
}

func (obj *flowCfmSenderIdTlv) Marshal() marshalFlowCfmSenderIdTlv {
	if obj.marshaller == nil {
		obj.marshaller = &marshalflowCfmSenderIdTlv{obj: obj}
	}
	return obj.marshaller
}

func (obj *flowCfmSenderIdTlv) Unmarshal() unMarshalFlowCfmSenderIdTlv {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalflowCfmSenderIdTlv{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalflowCfmSenderIdTlv) ToProto() (*otg.FlowCfmSenderIdTlv, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalflowCfmSenderIdTlv) FromProto(msg *otg.FlowCfmSenderIdTlv) (FlowCfmSenderIdTlv, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalflowCfmSenderIdTlv) ToPbText() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromPbText(value string) error {
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

func (m *marshalflowCfmSenderIdTlv) ToYaml() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromYaml(value string) error {
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

func (m *marshalflowCfmSenderIdTlv) ToJson() (string, error) {
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

func (m *unMarshalflowCfmSenderIdTlv) FromJson(value string) error {
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

func (obj *flowCfmSenderIdTlv) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlv) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *flowCfmSenderIdTlv) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *flowCfmSenderIdTlv) Clone() (FlowCfmSenderIdTlv, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewFlowCfmSenderIdTlv()
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

func (obj *flowCfmSenderIdTlv) setNil() {
	obj.chassisIdLengthHolder = nil
	obj.managementAddressLengthHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// FlowCfmSenderIdTlv is sender id
type FlowCfmSenderIdTlv interface {
	Validation
	// msg marshals FlowCfmSenderIdTlv to protobuf object *otg.FlowCfmSenderIdTlv
	// and doesn't set defaults
	msg() *otg.FlowCfmSenderIdTlv
	// setMsg unmarshals FlowCfmSenderIdTlv from protobuf object *otg.FlowCfmSenderIdTlv
	// and doesn't set defaults
	setMsg(*otg.FlowCfmSenderIdTlv) FlowCfmSenderIdTlv
	// provides marshal interface
	Marshal() marshalFlowCfmSenderIdTlv
	// provides unmarshal interface
	Unmarshal() unMarshalFlowCfmSenderIdTlv
	// validate validates FlowCfmSenderIdTlv
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (FlowCfmSenderIdTlv, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// ChassisIdLength returns PatternFlowCfmSenderIdTlvChassisIdLength, set in FlowCfmSenderIdTlv.
	// PatternFlowCfmSenderIdTlvChassisIdLength is specifies the length of the chassis ID field.
	ChassisIdLength() PatternFlowCfmSenderIdTlvChassisIdLength
	// SetChassisIdLength assigns PatternFlowCfmSenderIdTlvChassisIdLength provided by user to FlowCfmSenderIdTlv.
	// PatternFlowCfmSenderIdTlvChassisIdLength is specifies the length of the chassis ID field.
	SetChassisIdLength(value PatternFlowCfmSenderIdTlvChassisIdLength) FlowCfmSenderIdTlv
	// HasChassisIdLength checks if ChassisIdLength has been set in FlowCfmSenderIdTlv
	HasChassisIdLength() bool
	// ChassisIdSubtype returns FlowCfmSenderIdTlvChassisIdSubtypeEnum, set in FlowCfmSenderIdTlv
	ChassisIdSubtype() FlowCfmSenderIdTlvChassisIdSubtypeEnum
	// SetChassisIdSubtype assigns FlowCfmSenderIdTlvChassisIdSubtypeEnum provided by user to FlowCfmSenderIdTlv
	SetChassisIdSubtype(value FlowCfmSenderIdTlvChassisIdSubtypeEnum) FlowCfmSenderIdTlv
	// HasChassisIdSubtype checks if ChassisIdSubtype has been set in FlowCfmSenderIdTlv
	HasChassisIdSubtype() bool
	// ChassisId returns string, set in FlowCfmSenderIdTlv.
	ChassisId() string
	// SetChassisId assigns string provided by user to FlowCfmSenderIdTlv
	SetChassisId(value string) FlowCfmSenderIdTlv
	// HasChassisId checks if ChassisId has been set in FlowCfmSenderIdTlv
	HasChassisId() bool
	// ManagementAddressLength returns PatternFlowCfmSenderIdTlvManagementAddressLength, set in FlowCfmSenderIdTlv.
	// PatternFlowCfmSenderIdTlvManagementAddressLength is specifies the length of the management address.
	ManagementAddressLength() PatternFlowCfmSenderIdTlvManagementAddressLength
	// SetManagementAddressLength assigns PatternFlowCfmSenderIdTlvManagementAddressLength provided by user to FlowCfmSenderIdTlv.
	// PatternFlowCfmSenderIdTlvManagementAddressLength is specifies the length of the management address.
	SetManagementAddressLength(value PatternFlowCfmSenderIdTlvManagementAddressLength) FlowCfmSenderIdTlv
	// HasManagementAddressLength checks if ManagementAddressLength has been set in FlowCfmSenderIdTlv
	HasManagementAddressLength() bool
	// ManagementAddress returns string, set in FlowCfmSenderIdTlv.
	ManagementAddress() string
	// SetManagementAddress assigns string provided by user to FlowCfmSenderIdTlv
	SetManagementAddress(value string) FlowCfmSenderIdTlv
	// HasManagementAddress checks if ManagementAddress has been set in FlowCfmSenderIdTlv
	HasManagementAddress() bool
	setNil()
}

// description is TBD
// ChassisIdLength returns a PatternFlowCfmSenderIdTlvChassisIdLength
func (obj *flowCfmSenderIdTlv) ChassisIdLength() PatternFlowCfmSenderIdTlvChassisIdLength {
	if obj.obj.ChassisIdLength == nil {
		obj.obj.ChassisIdLength = NewPatternFlowCfmSenderIdTlvChassisIdLength().msg()
	}
	if obj.chassisIdLengthHolder == nil {
		obj.chassisIdLengthHolder = &patternFlowCfmSenderIdTlvChassisIdLength{obj: obj.obj.ChassisIdLength}
	}
	return obj.chassisIdLengthHolder
}

// description is TBD
// ChassisIdLength returns a PatternFlowCfmSenderIdTlvChassisIdLength
func (obj *flowCfmSenderIdTlv) HasChassisIdLength() bool {
	return obj.obj.ChassisIdLength != nil
}

// description is TBD
// SetChassisIdLength sets the PatternFlowCfmSenderIdTlvChassisIdLength value in the FlowCfmSenderIdTlv object
func (obj *flowCfmSenderIdTlv) SetChassisIdLength(value PatternFlowCfmSenderIdTlvChassisIdLength) FlowCfmSenderIdTlv {

	obj.chassisIdLengthHolder = nil
	obj.obj.ChassisIdLength = value.msg()

	return obj
}

type FlowCfmSenderIdTlvChassisIdSubtypeEnum string

// Enum of ChassisIdSubtype on FlowCfmSenderIdTlv
var FlowCfmSenderIdTlvChassisIdSubtype = struct {
	CHASSIS_COMPONENT FlowCfmSenderIdTlvChassisIdSubtypeEnum
	INTERFACE_ALIAS   FlowCfmSenderIdTlvChassisIdSubtypeEnum
	PORT_COMPONENT    FlowCfmSenderIdTlvChassisIdSubtypeEnum
	MAC_ADDRESS       FlowCfmSenderIdTlvChassisIdSubtypeEnum
	NETWORK_ADDRESS   FlowCfmSenderIdTlvChassisIdSubtypeEnum
	INTERFACE_NAME    FlowCfmSenderIdTlvChassisIdSubtypeEnum
	LOCALLY_ASSIGNED  FlowCfmSenderIdTlvChassisIdSubtypeEnum
}{
	CHASSIS_COMPONENT: FlowCfmSenderIdTlvChassisIdSubtypeEnum("chassis_component"),
	INTERFACE_ALIAS:   FlowCfmSenderIdTlvChassisIdSubtypeEnum("interface_alias"),
	PORT_COMPONENT:    FlowCfmSenderIdTlvChassisIdSubtypeEnum("port_component"),
	MAC_ADDRESS:       FlowCfmSenderIdTlvChassisIdSubtypeEnum("mac_address"),
	NETWORK_ADDRESS:   FlowCfmSenderIdTlvChassisIdSubtypeEnum("network_address"),
	INTERFACE_NAME:    FlowCfmSenderIdTlvChassisIdSubtypeEnum("interface_name"),
	LOCALLY_ASSIGNED:  FlowCfmSenderIdTlvChassisIdSubtypeEnum("locally_assigned"),
}

func (obj *flowCfmSenderIdTlv) ChassisIdSubtype() FlowCfmSenderIdTlvChassisIdSubtypeEnum {
	return FlowCfmSenderIdTlvChassisIdSubtypeEnum(obj.obj.ChassisIdSubtype.Enum().String())
}

// Specifies the Type of chassis ID.
// ChassisIdSubtype returns a string
func (obj *flowCfmSenderIdTlv) HasChassisIdSubtype() bool {
	return obj.obj.ChassisIdSubtype != nil
}

func (obj *flowCfmSenderIdTlv) SetChassisIdSubtype(value FlowCfmSenderIdTlvChassisIdSubtypeEnum) FlowCfmSenderIdTlv {
	intValue, ok := otg.FlowCfmSenderIdTlv_ChassisIdSubtype_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on FlowCfmSenderIdTlvChassisIdSubtypeEnum", string(value)))
		return obj
	}
	enumValue := otg.FlowCfmSenderIdTlv_ChassisIdSubtype_Enum(intValue)
	obj.obj.ChassisIdSubtype = &enumValue

	return obj
}

// Chassis ID
// ChassisId returns a string
func (obj *flowCfmSenderIdTlv) ChassisId() string {

	return *obj.obj.ChassisId

}

// Chassis ID
// ChassisId returns a string
func (obj *flowCfmSenderIdTlv) HasChassisId() bool {
	return obj.obj.ChassisId != nil
}

// Chassis ID
// SetChassisId sets the string value in the FlowCfmSenderIdTlv object
func (obj *flowCfmSenderIdTlv) SetChassisId(value string) FlowCfmSenderIdTlv {

	obj.obj.ChassisId = &value
	return obj
}

// description is TBD
// ManagementAddressLength returns a PatternFlowCfmSenderIdTlvManagementAddressLength
func (obj *flowCfmSenderIdTlv) ManagementAddressLength() PatternFlowCfmSenderIdTlvManagementAddressLength {
	if obj.obj.ManagementAddressLength == nil {
		obj.obj.ManagementAddressLength = NewPatternFlowCfmSenderIdTlvManagementAddressLength().msg()
	}
	if obj.managementAddressLengthHolder == nil {
		obj.managementAddressLengthHolder = &patternFlowCfmSenderIdTlvManagementAddressLength{obj: obj.obj.ManagementAddressLength}
	}
	return obj.managementAddressLengthHolder
}

// description is TBD
// ManagementAddressLength returns a PatternFlowCfmSenderIdTlvManagementAddressLength
func (obj *flowCfmSenderIdTlv) HasManagementAddressLength() bool {
	return obj.obj.ManagementAddressLength != nil
}

// description is TBD
// SetManagementAddressLength sets the PatternFlowCfmSenderIdTlvManagementAddressLength value in the FlowCfmSenderIdTlv object
func (obj *flowCfmSenderIdTlv) SetManagementAddressLength(value PatternFlowCfmSenderIdTlvManagementAddressLength) FlowCfmSenderIdTlv {

	obj.managementAddressLengthHolder = nil
	obj.obj.ManagementAddressLength = value.msg()

	return obj
}

// Management address
// ManagementAddress returns a string
func (obj *flowCfmSenderIdTlv) ManagementAddress() string {

	return *obj.obj.ManagementAddress

}

// Management address
// ManagementAddress returns a string
func (obj *flowCfmSenderIdTlv) HasManagementAddress() bool {
	return obj.obj.ManagementAddress != nil
}

// Management address
// SetManagementAddress sets the string value in the FlowCfmSenderIdTlv object
func (obj *flowCfmSenderIdTlv) SetManagementAddress(value string) FlowCfmSenderIdTlv {

	obj.obj.ManagementAddress = &value
	return obj
}

func (obj *flowCfmSenderIdTlv) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.ChassisIdLength != nil {

		obj.ChassisIdLength().validateObj(vObj, set_default)
	}

	if obj.obj.ChassisId != nil {

		err := obj.validateHex(obj.ChassisId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmSenderIdTlv.ChassisId"))
		}

	}

	if obj.obj.ManagementAddressLength != nil {

		obj.ManagementAddressLength().validateObj(vObj, set_default)
	}

	if obj.obj.ManagementAddress != nil {

		err := obj.validateHex(obj.ManagementAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on FlowCfmSenderIdTlv.ManagementAddress"))
		}

	}

}

func (obj *flowCfmSenderIdTlv) setDefault() {
	if obj.obj.ChassisIdSubtype == nil {
		obj.SetChassisIdSubtype(FlowCfmSenderIdTlvChassisIdSubtype.MAC_ADDRESS)

	}
	if obj.obj.ChassisId == nil {
		obj.SetChassisId("00")
	}
	if obj.obj.ManagementAddress == nil {
		obj.SetManagementAddress("00")
	}

}
