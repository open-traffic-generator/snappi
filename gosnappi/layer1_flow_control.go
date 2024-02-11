package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** Layer1FlowControl *****
type layer1FlowControl struct {
	validation
	obj                 *otg.Layer1FlowControl
	marshaller          marshalLayer1FlowControl
	unMarshaller        unMarshalLayer1FlowControl
	ieee_802_1QbbHolder Layer1Ieee8021Qbb
	ieee_802_3XHolder   Layer1Ieee8023X
}

func NewLayer1FlowControl() Layer1FlowControl {
	obj := layer1FlowControl{obj: &otg.Layer1FlowControl{}}
	obj.setDefault()
	return &obj
}

func (obj *layer1FlowControl) msg() *otg.Layer1FlowControl {
	return obj.obj
}

func (obj *layer1FlowControl) setMsg(msg *otg.Layer1FlowControl) Layer1FlowControl {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshallayer1FlowControl struct {
	obj *layer1FlowControl
}

type marshalLayer1FlowControl interface {
	// ToProto marshals Layer1FlowControl to protobuf object *otg.Layer1FlowControl
	ToProto() (*otg.Layer1FlowControl, error)
	// ToPbText marshals Layer1FlowControl to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals Layer1FlowControl to YAML text
	ToYaml() (string, error)
	// ToJson marshals Layer1FlowControl to JSON text
	ToJson() (string, error)
}

type unMarshallayer1FlowControl struct {
	obj *layer1FlowControl
}

type unMarshalLayer1FlowControl interface {
	// FromProto unmarshals Layer1FlowControl from protobuf object *otg.Layer1FlowControl
	FromProto(msg *otg.Layer1FlowControl) (Layer1FlowControl, error)
	// FromPbText unmarshals Layer1FlowControl from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals Layer1FlowControl from YAML text
	FromYaml(value string) error
	// FromJson unmarshals Layer1FlowControl from JSON text
	FromJson(value string) error
}

func (obj *layer1FlowControl) Marshal() marshalLayer1FlowControl {
	if obj.marshaller == nil {
		obj.marshaller = &marshallayer1FlowControl{obj: obj}
	}
	return obj.marshaller
}

func (obj *layer1FlowControl) Unmarshal() unMarshalLayer1FlowControl {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshallayer1FlowControl{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshallayer1FlowControl) ToProto() (*otg.Layer1FlowControl, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshallayer1FlowControl) FromProto(msg *otg.Layer1FlowControl) (Layer1FlowControl, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshallayer1FlowControl) ToPbText() (string, error) {
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

func (m *unMarshallayer1FlowControl) FromPbText(value string) error {
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

func (m *marshallayer1FlowControl) ToYaml() (string, error) {
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

func (m *unMarshallayer1FlowControl) FromYaml(value string) error {
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

func (m *marshallayer1FlowControl) ToJson() (string, error) {
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

func (m *unMarshallayer1FlowControl) FromJson(value string) error {
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

func (obj *layer1FlowControl) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *layer1FlowControl) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *layer1FlowControl) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *layer1FlowControl) Clone() (Layer1FlowControl, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewLayer1FlowControl()
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

func (obj *layer1FlowControl) setNil() {
	obj.ieee_802_1QbbHolder = nil
	obj.ieee_802_3XHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// Layer1FlowControl is a container for layer1 receive flow control settings.
// To enable flow control settings on ports this object must be a valid
// object not a null value.
type Layer1FlowControl interface {
	Validation
	// msg marshals Layer1FlowControl to protobuf object *otg.Layer1FlowControl
	// and doesn't set defaults
	msg() *otg.Layer1FlowControl
	// setMsg unmarshals Layer1FlowControl from protobuf object *otg.Layer1FlowControl
	// and doesn't set defaults
	setMsg(*otg.Layer1FlowControl) Layer1FlowControl
	// provides marshal interface
	Marshal() marshalLayer1FlowControl
	// provides unmarshal interface
	Unmarshal() unMarshalLayer1FlowControl
	// validate validates Layer1FlowControl
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (Layer1FlowControl, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// DirectedAddress returns string, set in Layer1FlowControl.
	DirectedAddress() string
	// SetDirectedAddress assigns string provided by user to Layer1FlowControl
	SetDirectedAddress(value string) Layer1FlowControl
	// HasDirectedAddress checks if DirectedAddress has been set in Layer1FlowControl
	HasDirectedAddress() bool
	// Choice returns Layer1FlowControlChoiceEnum, set in Layer1FlowControl
	Choice() Layer1FlowControlChoiceEnum
	// setChoice assigns Layer1FlowControlChoiceEnum provided by user to Layer1FlowControl
	setChoice(value Layer1FlowControlChoiceEnum) Layer1FlowControl
	// HasChoice checks if Choice has been set in Layer1FlowControl
	HasChoice() bool
	// Ieee8021Qbb returns Layer1Ieee8021Qbb, set in Layer1FlowControl.
	Ieee8021Qbb() Layer1Ieee8021Qbb
	// SetIeee8021Qbb assigns Layer1Ieee8021Qbb provided by user to Layer1FlowControl.
	SetIeee8021Qbb(value Layer1Ieee8021Qbb) Layer1FlowControl
	// HasIeee8021Qbb checks if Ieee8021Qbb has been set in Layer1FlowControl
	HasIeee8021Qbb() bool
	// Ieee8023X returns Layer1Ieee8023X, set in Layer1FlowControl.
	Ieee8023X() Layer1Ieee8023X
	// SetIeee8023X assigns Layer1Ieee8023X provided by user to Layer1FlowControl.
	SetIeee8023X(value Layer1Ieee8023X) Layer1FlowControl
	// HasIeee8023X checks if Ieee8023X has been set in Layer1FlowControl
	HasIeee8023X() bool
	setNil()
}

// The 48bit mac address that the layer1 port names will listen on
// for a directed pause.
// DirectedAddress returns a string
func (obj *layer1FlowControl) DirectedAddress() string {

	return *obj.obj.DirectedAddress

}

// The 48bit mac address that the layer1 port names will listen on
// for a directed pause.
// DirectedAddress returns a string
func (obj *layer1FlowControl) HasDirectedAddress() bool {
	return obj.obj.DirectedAddress != nil
}

// The 48bit mac address that the layer1 port names will listen on
// for a directed pause.
// SetDirectedAddress sets the string value in the Layer1FlowControl object
func (obj *layer1FlowControl) SetDirectedAddress(value string) Layer1FlowControl {

	obj.obj.DirectedAddress = &value
	return obj
}

type Layer1FlowControlChoiceEnum string

// Enum of Choice on Layer1FlowControl
var Layer1FlowControlChoice = struct {
	IEEE_802_1QBB Layer1FlowControlChoiceEnum
	IEEE_802_3X   Layer1FlowControlChoiceEnum
}{
	IEEE_802_1QBB: Layer1FlowControlChoiceEnum("ieee_802_1qbb"),
	IEEE_802_3X:   Layer1FlowControlChoiceEnum("ieee_802_3x"),
}

func (obj *layer1FlowControl) Choice() Layer1FlowControlChoiceEnum {
	return Layer1FlowControlChoiceEnum(obj.obj.Choice.Enum().String())
}

// The type of priority flow control.
// Choice returns a string
func (obj *layer1FlowControl) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *layer1FlowControl) setChoice(value Layer1FlowControlChoiceEnum) Layer1FlowControl {
	intValue, ok := otg.Layer1FlowControl_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on Layer1FlowControlChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.Layer1FlowControl_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Ieee_802_3X = nil
	obj.ieee_802_3XHolder = nil
	obj.obj.Ieee_802_1Qbb = nil
	obj.ieee_802_1QbbHolder = nil

	if value == Layer1FlowControlChoice.IEEE_802_1QBB {
		obj.obj.Ieee_802_1Qbb = NewLayer1Ieee8021Qbb().msg()
	}

	if value == Layer1FlowControlChoice.IEEE_802_3X {
		obj.obj.Ieee_802_3X = NewLayer1Ieee8023X().msg()
	}

	return obj
}

// description is TBD
// Ieee8021Qbb returns a Layer1Ieee8021Qbb
func (obj *layer1FlowControl) Ieee8021Qbb() Layer1Ieee8021Qbb {
	if obj.obj.Ieee_802_1Qbb == nil {
		obj.setChoice(Layer1FlowControlChoice.IEEE_802_1QBB)
	}
	if obj.ieee_802_1QbbHolder == nil {
		obj.ieee_802_1QbbHolder = &layer1Ieee8021Qbb{obj: obj.obj.Ieee_802_1Qbb}
	}
	return obj.ieee_802_1QbbHolder
}

// description is TBD
// Ieee8021Qbb returns a Layer1Ieee8021Qbb
func (obj *layer1FlowControl) HasIeee8021Qbb() bool {
	return obj.obj.Ieee_802_1Qbb != nil
}

// description is TBD
// SetIeee8021Qbb sets the Layer1Ieee8021Qbb value in the Layer1FlowControl object
func (obj *layer1FlowControl) SetIeee8021Qbb(value Layer1Ieee8021Qbb) Layer1FlowControl {
	obj.setChoice(Layer1FlowControlChoice.IEEE_802_1QBB)
	obj.ieee_802_1QbbHolder = nil
	obj.obj.Ieee_802_1Qbb = value.msg()

	return obj
}

// description is TBD
// Ieee8023X returns a Layer1Ieee8023X
func (obj *layer1FlowControl) Ieee8023X() Layer1Ieee8023X {
	if obj.obj.Ieee_802_3X == nil {
		obj.setChoice(Layer1FlowControlChoice.IEEE_802_3X)
	}
	if obj.ieee_802_3XHolder == nil {
		obj.ieee_802_3XHolder = &layer1Ieee8023X{obj: obj.obj.Ieee_802_3X}
	}
	return obj.ieee_802_3XHolder
}

// description is TBD
// Ieee8023X returns a Layer1Ieee8023X
func (obj *layer1FlowControl) HasIeee8023X() bool {
	return obj.obj.Ieee_802_3X != nil
}

// description is TBD
// SetIeee8023X sets the Layer1Ieee8023X value in the Layer1FlowControl object
func (obj *layer1FlowControl) SetIeee8023X(value Layer1Ieee8023X) Layer1FlowControl {
	obj.setChoice(Layer1FlowControlChoice.IEEE_802_3X)
	obj.ieee_802_3XHolder = nil
	obj.obj.Ieee_802_3X = value.msg()

	return obj
}

func (obj *layer1FlowControl) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.DirectedAddress != nil {

		err := obj.validateMac(obj.DirectedAddress())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on Layer1FlowControl.DirectedAddress"))
		}

	}

	if obj.obj.Ieee_802_1Qbb != nil {

		obj.Ieee8021Qbb().validateObj(vObj, set_default)
	}

	if obj.obj.Ieee_802_3X != nil {

		obj.Ieee8023X().validateObj(vObj, set_default)
	}

}

func (obj *layer1FlowControl) setDefault() {
	if obj.obj.Choice == nil {
		obj.setChoice(Layer1FlowControlChoice.IEEE_802_1QBB)

	}
	if obj.obj.DirectedAddress == nil {
		obj.SetDirectedAddress("01:80:C2:00:00:01")
	}

}
