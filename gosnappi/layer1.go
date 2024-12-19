package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Layer1 *****
type layer1 struct {
	validation
	obj                   *otg.Layer1
	marshaller            marshalLayer1
	unMarshaller          unMarshalLayer1
	autoNegotiationHolder Layer1AutoNegotiation
	flowControlHolder     Layer1FlowControl
}

func NewLayer1() Layer1 {
	obj := layer1{obj: &otg.Layer1{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1) msg() *otg.Layer1 {
	return obj.obj
}

func (obj *layer1) setMsg(msg *otg.Layer1) Layer1 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallayer1 struct {
	obj *layer1
}

type marshalLayer1 interface {
	// ToProto marshals Layer1 to protobuf object *otg.Layer1
	ToProto() (*otg.Layer1, error)
	// ToPbText marshals Layer1 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1 to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1 to JSON text
	ToJson() (string, error)
}

type unMarshallayer1 struct {
	obj *layer1
}

type unMarshalLayer1 interface {
	// FromProto unmarshals Layer1 from protobuf object *otg.Layer1
	FromProto(msg *otg.Layer1) (Layer1, error)
	// FromPbText unmarshals Layer1 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1 from JSON text
	FromJson(value string) error
}

func (obj *layer1) Marshal() marshalLayer1 {
	if obj.marshaller == nil {
		obj.marshaller = &marshallayer1{obj: obj}
	}
	return obj.marshaller
}

func (obj *layer1) Unmarshal() unMarshalLayer1 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallayer1{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallayer1) ToProto() (*otg.Layer1, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallayer1) FromProto(msg *otg.Layer1) (Layer1, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallayer1) ToPbText() (string, error) {
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

func (m *unMarshallayer1) FromPbText(value string) error {
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

func (m *marshallayer1) ToYaml() (string, error) {
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

func (m *unMarshallayer1) FromYaml(value string) error {
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

func (m *marshallayer1) ToJson() (string, error) {
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

func (m *unMarshallayer1) FromJson(value string) error {
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

func (obj *layer1) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1) Clone() (Layer1, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1()
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

func (obj *layer1) setNil() {
	obj.autoNegotiationHolder = nil
	obj.flowControlHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Layer1 is a container for layer1 settings.
type Layer1 interface {
	Validation
	// msg marshals Layer1 to protobuf object *otg.Layer1
	// and doesn't set defaults
	msg() *otg.Layer1
	// setMsg unmarshals Layer1 from protobuf object *otg.Layer1
	// and doesn't set defaults
	setMsg(*otg.Layer1) Layer1
	// provides marshal interface
	Marshal() marshalLayer1
	// provides unmarshal interface
	Unmarshal() unMarshalLayer1
	// validate validates Layer1
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PortNames returns []string, set in Layer1.
	PortNames() []string
	// SetPortNames assigns []string provided by user to Layer1
	SetPortNames(value []string) Layer1
	// Speed returns Layer1SpeedEnum, set in Layer1
	Speed() Layer1SpeedEnum
	// SetSpeed assigns Layer1SpeedEnum provided by user to Layer1
	SetSpeed(value Layer1SpeedEnum) Layer1
	// HasSpeed checks if Speed has been set in Layer1
	HasSpeed() bool
	// CustomSpeed returns string, set in Layer1.
	CustomSpeed() string
	// SetCustomSpeed assigns string provided by user to Layer1
	SetCustomSpeed(value string) Layer1
	// HasCustomSpeed checks if CustomSpeed has been set in Layer1
	HasCustomSpeed() bool
	// Media returns Layer1MediaEnum, set in Layer1
	Media() Layer1MediaEnum
	// SetMedia assigns Layer1MediaEnum provided by user to Layer1
	SetMedia(value Layer1MediaEnum) Layer1
	// HasMedia checks if Media has been set in Layer1
	HasMedia() bool
	// Promiscuous returns bool, set in Layer1.
	Promiscuous() bool
	// SetPromiscuous assigns bool provided by user to Layer1
	SetPromiscuous(value bool) Layer1
	// HasPromiscuous checks if Promiscuous has been set in Layer1
	HasPromiscuous() bool
	// Mtu returns uint32, set in Layer1.
	Mtu() uint32
	// SetMtu assigns uint32 provided by user to Layer1
	SetMtu(value uint32) Layer1
	// HasMtu checks if Mtu has been set in Layer1
	HasMtu() bool
	// IeeeMediaDefaults returns bool, set in Layer1.
	IeeeMediaDefaults() bool
	// SetIeeeMediaDefaults assigns bool provided by user to Layer1
	SetIeeeMediaDefaults(value bool) Layer1
	// HasIeeeMediaDefaults checks if IeeeMediaDefaults has been set in Layer1
	HasIeeeMediaDefaults() bool
	// AutoNegotiate returns bool, set in Layer1.
	AutoNegotiate() bool
	// SetAutoNegotiate assigns bool provided by user to Layer1
	SetAutoNegotiate(value bool) Layer1
	// HasAutoNegotiate checks if AutoNegotiate has been set in Layer1
	HasAutoNegotiate() bool
	// AutoNegotiation returns Layer1AutoNegotiation, set in Layer1.
	// Layer1AutoNegotiation is configuration for auto negotiation settings
	AutoNegotiation() Layer1AutoNegotiation
	// SetAutoNegotiation assigns Layer1AutoNegotiation provided by user to Layer1.
	// Layer1AutoNegotiation is configuration for auto negotiation settings
	SetAutoNegotiation(value Layer1AutoNegotiation) Layer1
	// HasAutoNegotiation checks if AutoNegotiation has been set in Layer1
	HasAutoNegotiation() bool
	// FlowControl returns Layer1FlowControl, set in Layer1.
	// Layer1FlowControl is a container for layer1 receive flow control settings.
	// To enable flow control settings on ports this object must be a valid
	// object not a null value.
	FlowControl() Layer1FlowControl
	// SetFlowControl assigns Layer1FlowControl provided by user to Layer1.
	// Layer1FlowControl is a container for layer1 receive flow control settings.
	// To enable flow control settings on ports this object must be a valid
	// object not a null value.
	SetFlowControl(value Layer1FlowControl) Layer1
	// HasFlowControl checks if FlowControl has been set in Layer1
	HasFlowControl() bool
	// Name returns string, set in Layer1.
	Name() string
	// SetName assigns string provided by user to Layer1
	SetName(value string) Layer1
	setNil()
}

// A list of unique names of port objects that will share the
// choice settings.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// PortNames returns a []string
func (obj *layer1) PortNames() []string {
	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	return obj.obj.PortNames
}

// A list of unique names of port objects that will share the
// choice settings.
//
// x-constraint:
// - /components/schemas/Port/properties/name
//
// SetPortNames sets the []string value in the Layer1 object
func (obj *layer1) SetPortNames(value []string) Layer1 {

	if obj.obj.PortNames == nil {
		obj.obj.PortNames = make([]string, 0)
	}
	obj.obj.PortNames = value

	return obj
}

type Layer1SpeedEnum string

// Enum of Speed on Layer1
var Layer1Speed = struct {
	SPEED_10_FD_MBPS  Layer1SpeedEnum
	SPEED_10_HD_MBPS  Layer1SpeedEnum
	SPEED_100_FD_MBPS Layer1SpeedEnum
	SPEED_100_HD_MBPS Layer1SpeedEnum
	SPEED_1_GBPS      Layer1SpeedEnum
	SPEED_10_GBPS     Layer1SpeedEnum
	SPEED_25_GBPS     Layer1SpeedEnum
	SPEED_40_GBPS     Layer1SpeedEnum
	SPEED_50_GBPS     Layer1SpeedEnum
	SPEED_100_GBPS    Layer1SpeedEnum
	SPEED_200_GBPS    Layer1SpeedEnum
	SPEED_400_GBPS    Layer1SpeedEnum
	SPEED_800_GBPS    Layer1SpeedEnum
	CUSTOM_SPEED      Layer1SpeedEnum
}{
	SPEED_10_FD_MBPS:  Layer1SpeedEnum("speed_10_fd_mbps"),
	SPEED_10_HD_MBPS:  Layer1SpeedEnum("speed_10_hd_mbps"),
	SPEED_100_FD_MBPS: Layer1SpeedEnum("speed_100_fd_mbps"),
	SPEED_100_HD_MBPS: Layer1SpeedEnum("speed_100_hd_mbps"),
	SPEED_1_GBPS:      Layer1SpeedEnum("speed_1_gbps"),
	SPEED_10_GBPS:     Layer1SpeedEnum("speed_10_gbps"),
	SPEED_25_GBPS:     Layer1SpeedEnum("speed_25_gbps"),
	SPEED_40_GBPS:     Layer1SpeedEnum("speed_40_gbps"),
	SPEED_50_GBPS:     Layer1SpeedEnum("speed_50_gbps"),
	SPEED_100_GBPS:    Layer1SpeedEnum("speed_100_gbps"),
	SPEED_200_GBPS:    Layer1SpeedEnum("speed_200_gbps"),
	SPEED_400_GBPS:    Layer1SpeedEnum("speed_400_gbps"),
	SPEED_800_GBPS:    Layer1SpeedEnum("speed_800_gbps"),
	CUSTOM_SPEED:      Layer1SpeedEnum("custom_speed"),
}

func (obj *layer1) Speed() Layer1SpeedEnum {
	return Layer1SpeedEnum(obj.obj.Speed.Enum().String())
}

// Set the speed if supported. When no speed is explicitly set, the current
// speed of underlying test interface shall be assumed.
// Speed returns a string
func (obj *layer1) HasSpeed() bool {
	return obj.obj.Speed != nil
}

func (obj *layer1) SetSpeed(value Layer1SpeedEnum) Layer1 {
	intValue, ok := otg.Layer1_Speed_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Layer1SpeedEnum", string(value)))
		return obj
	}
	enumValue := otg.Layer1_Speed_Enum(intValue)
	obj.obj.Speed = &enumValue

	return obj
}

// Vendor specific custom speed.
// CustomSpeed returns a string
func (obj *layer1) CustomSpeed() string {

	return *obj.obj.CustomSpeed

}

// Vendor specific custom speed.
// CustomSpeed returns a string
func (obj *layer1) HasCustomSpeed() bool {
	return obj.obj.CustomSpeed != nil
}

// Vendor specific custom speed.
// SetCustomSpeed sets the string value in the Layer1 object
func (obj *layer1) SetCustomSpeed(value string) Layer1 {

	obj.obj.CustomSpeed = &value
	return obj
}

type Layer1MediaEnum string

// Enum of Media on Layer1
var Layer1Media = struct {
	COPPER Layer1MediaEnum
	FIBER  Layer1MediaEnum
	SGMII  Layer1MediaEnum
}{
	COPPER: Layer1MediaEnum("copper"),
	FIBER:  Layer1MediaEnum("fiber"),
	SGMII:  Layer1MediaEnum("sgmii"),
}

func (obj *layer1) Media() Layer1MediaEnum {
	return Layer1MediaEnum(obj.obj.Media.Enum().String())
}

// Set the type of media for test interface if supported. When no media
// type is explicitly set, the current media type of underlying test
// interface shall be assumed.
// Media returns a string
func (obj *layer1) HasMedia() bool {
	return obj.obj.Media != nil
}

func (obj *layer1) SetMedia(value Layer1MediaEnum) Layer1 {
	intValue, ok := otg.Layer1_Media_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Layer1MediaEnum", string(value)))
		return obj
	}
	enumValue := otg.Layer1_Media_Enum(intValue)
	obj.obj.Media = &enumValue

	return obj
}

// Enable promiscuous mode on test interface. A warning shall be raised if
// this field is set to `true`, even when it's not supported, ignoring
// the setting altogether.
// Promiscuous returns a bool
func (obj *layer1) Promiscuous() bool {

	return *obj.obj.Promiscuous

}

// Enable promiscuous mode on test interface. A warning shall be raised if
// this field is set to `true`, even when it's not supported, ignoring
// the setting altogether.
// Promiscuous returns a bool
func (obj *layer1) HasPromiscuous() bool {
	return obj.obj.Promiscuous != nil
}

// Enable promiscuous mode on test interface. A warning shall be raised if
// this field is set to `true`, even when it's not supported, ignoring
// the setting altogether.
// SetPromiscuous sets the bool value in the Layer1 object
func (obj *layer1) SetPromiscuous(value bool) Layer1 {

	obj.obj.Promiscuous = &value
	return obj
}

// Set the maximum transmission unit size. A warning shall be raised if
// the specified value is valid but not supported, ignoring the setting altogether.
// Mtu returns a uint32
func (obj *layer1) Mtu() uint32 {

	return *obj.obj.Mtu

}

// Set the maximum transmission unit size. A warning shall be raised if
// the specified value is valid but not supported, ignoring the setting altogether.
// Mtu returns a uint32
func (obj *layer1) HasMtu() bool {
	return obj.obj.Mtu != nil
}

// Set the maximum transmission unit size. A warning shall be raised if
// the specified value is valid but not supported, ignoring the setting altogether.
// SetMtu sets the uint32 value in the Layer1 object
func (obj *layer1) SetMtu(value uint32) Layer1 {

	obj.obj.Mtu = &value
	return obj
}

// Under Review: This field is currently under review for pending exploration on use cases
//
// Set to true to override the auto_negotiate, link_training
// and rs_fec settings for gigabit ethernet interfaces.
// IeeeMediaDefaults returns a bool
func (obj *layer1) IeeeMediaDefaults() bool {

	return *obj.obj.IeeeMediaDefaults

}

// Under Review: This field is currently under review for pending exploration on use cases
//
// Set to true to override the auto_negotiate, link_training
// and rs_fec settings for gigabit ethernet interfaces.
// IeeeMediaDefaults returns a bool
func (obj *layer1) HasIeeeMediaDefaults() bool {
	return obj.obj.IeeeMediaDefaults != nil
}

// Under Review: This field is currently under review for pending exploration on use cases
//
// Set to true to override the auto_negotiate, link_training
// and rs_fec settings for gigabit ethernet interfaces.
// SetIeeeMediaDefaults sets the bool value in the Layer1 object
func (obj *layer1) SetIeeeMediaDefaults(value bool) Layer1 {

	obj.obj.IeeeMediaDefaults = &value
	return obj
}

// Under Review: This field is currently under review for pending exploration on use cases, given that a separate configuration called `AutoNegotiation` already exists.
//
// Enable/disable auto negotiation.
// AutoNegotiate returns a bool
func (obj *layer1) AutoNegotiate() bool {

	return *obj.obj.AutoNegotiate

}

// Under Review: This field is currently under review for pending exploration on use cases, given that a separate configuration called `AutoNegotiation` already exists.
//
// Enable/disable auto negotiation.
// AutoNegotiate returns a bool
func (obj *layer1) HasAutoNegotiate() bool {
	return obj.obj.AutoNegotiate != nil
}

// Under Review: This field is currently under review for pending exploration on use cases, given that a separate configuration called `AutoNegotiation` already exists.
//
// Enable/disable auto negotiation.
// SetAutoNegotiate sets the bool value in the Layer1 object
func (obj *layer1) SetAutoNegotiate(value bool) Layer1 {

	obj.obj.AutoNegotiate = &value
	return obj
}

// description is TBD
// AutoNegotiation returns a Layer1AutoNegotiation
func (obj *layer1) AutoNegotiation() Layer1AutoNegotiation {
	if obj.obj.AutoNegotiation == nil {
		obj.obj.AutoNegotiation = NewLayer1AutoNegotiation().msg()
	}
	if obj.autoNegotiationHolder == nil {
		obj.autoNegotiationHolder = &layer1AutoNegotiation{obj: obj.obj.AutoNegotiation}
	}
	return obj.autoNegotiationHolder
}

// description is TBD
// AutoNegotiation returns a Layer1AutoNegotiation
func (obj *layer1) HasAutoNegotiation() bool {
	return obj.obj.AutoNegotiation != nil
}

// description is TBD
// SetAutoNegotiation sets the Layer1AutoNegotiation value in the Layer1 object
func (obj *layer1) SetAutoNegotiation(value Layer1AutoNegotiation) Layer1 {

	obj.autoNegotiationHolder = nil
	obj.obj.AutoNegotiation = value.msg()

	return obj
}

// description is TBD
// FlowControl returns a Layer1FlowControl
func (obj *layer1) FlowControl() Layer1FlowControl {
	if obj.obj.FlowControl == nil {
		obj.obj.FlowControl = NewLayer1FlowControl().msg()
	}
	if obj.flowControlHolder == nil {
		obj.flowControlHolder = &layer1FlowControl{obj: obj.obj.FlowControl}
	}
	return obj.flowControlHolder
}

// description is TBD
// FlowControl returns a Layer1FlowControl
func (obj *layer1) HasFlowControl() bool {
	return obj.obj.FlowControl != nil
}

// description is TBD
// SetFlowControl sets the Layer1FlowControl value in the Layer1 object
func (obj *layer1) SetFlowControl(value Layer1FlowControl) Layer1 {

	obj.flowControlHolder = nil
	obj.obj.FlowControl = value.msg()

	return obj
}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// Name returns a string
func (obj *layer1) Name() string {

	return *obj.obj.Name

}

// Globally unique name of an object. It also serves as the primary key for arrays of objects.
// SetName sets the string value in the Layer1 object
func (obj *layer1) SetName(value string) Layer1 {

	obj.obj.Name = &value
	return obj
}

func (obj *layer1) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Mtu != nil {

		if *obj.obj.Mtu < 64 || *obj.obj.Mtu > 9000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("64 <= Layer1.Mtu <= 9000 but Got %d", *obj.obj.Mtu))
		}

	}

	if obj.obj.AutoNegotiation != nil {

		obj.AutoNegotiation().validateObj(vObj, set_default)
	}

	if obj.obj.FlowControl != nil {

		obj.FlowControl().validateObj(vObj, set_default)
	}

	// Name is required
	if obj.obj.Name == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Name is required field on interface Layer1")
	}
}

func (obj *layer1) setDefault() {
	if obj.obj.Promiscuous == nil {
		obj.SetPromiscuous(true)
	}
	if obj.obj.Mtu == nil {
		obj.SetMtu(1500)
	}

}
