package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV6TunnelDestinationIPMode *****
type vxlanV6TunnelDestinationIPMode struct {
	validation
	obj             *otg.VxlanV6TunnelDestinationIPMode
	marshaller      marshalVxlanV6TunnelDestinationIPMode
	unMarshaller    unMarshalVxlanV6TunnelDestinationIPMode
	unicastHolder   VxlanV6TunnelDestinationIPModeUnicast
	multicastHolder VxlanV6TunnelDestinationIPModeMulticast
}

func NewVxlanV6TunnelDestinationIPMode() VxlanV6TunnelDestinationIPMode {
	obj := vxlanV6TunnelDestinationIPMode{obj: &otg.VxlanV6TunnelDestinationIPMode{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV6TunnelDestinationIPMode) msg() *otg.VxlanV6TunnelDestinationIPMode {
	return obj.obj
}

func (obj *vxlanV6TunnelDestinationIPMode) setMsg(msg *otg.VxlanV6TunnelDestinationIPMode) VxlanV6TunnelDestinationIPMode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV6TunnelDestinationIPMode struct {
	obj *vxlanV6TunnelDestinationIPMode
}

type marshalVxlanV6TunnelDestinationIPMode interface {
	// ToProto marshals VxlanV6TunnelDestinationIPMode to protobuf object *otg.VxlanV6TunnelDestinationIPMode
	ToProto() (*otg.VxlanV6TunnelDestinationIPMode, error)
	// ToPbText marshals VxlanV6TunnelDestinationIPMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV6TunnelDestinationIPMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV6TunnelDestinationIPMode to JSON text
	ToJson() (string, error)
}

type unMarshalvxlanV6TunnelDestinationIPMode struct {
	obj *vxlanV6TunnelDestinationIPMode
}

type unMarshalVxlanV6TunnelDestinationIPMode interface {
	// FromProto unmarshals VxlanV6TunnelDestinationIPMode from protobuf object *otg.VxlanV6TunnelDestinationIPMode
	FromProto(msg *otg.VxlanV6TunnelDestinationIPMode) (VxlanV6TunnelDestinationIPMode, error)
	// FromPbText unmarshals VxlanV6TunnelDestinationIPMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV6TunnelDestinationIPMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV6TunnelDestinationIPMode from JSON text
	FromJson(value string) error
}

func (obj *vxlanV6TunnelDestinationIPMode) Marshal() marshalVxlanV6TunnelDestinationIPMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV6TunnelDestinationIPMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV6TunnelDestinationIPMode) Unmarshal() unMarshalVxlanV6TunnelDestinationIPMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV6TunnelDestinationIPMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV6TunnelDestinationIPMode) ToProto() (*otg.VxlanV6TunnelDestinationIPMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV6TunnelDestinationIPMode) FromProto(msg *otg.VxlanV6TunnelDestinationIPMode) (VxlanV6TunnelDestinationIPMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV6TunnelDestinationIPMode) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPMode) FromPbText(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPMode) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPMode) FromYaml(value string) error {
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

func (m *marshalvxlanV6TunnelDestinationIPMode) ToJson() (string, error) {
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

func (m *unMarshalvxlanV6TunnelDestinationIPMode) FromJson(value string) error {
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

func (obj *vxlanV6TunnelDestinationIPMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV6TunnelDestinationIPMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV6TunnelDestinationIPMode) Clone() (VxlanV6TunnelDestinationIPMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV6TunnelDestinationIPMode()
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

func (obj *vxlanV6TunnelDestinationIPMode) setNil() {
	obj.unicastHolder = nil
	obj.multicastHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV6TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
type VxlanV6TunnelDestinationIPMode interface {
	Validation
	// msg marshals VxlanV6TunnelDestinationIPMode to protobuf object *otg.VxlanV6TunnelDestinationIPMode
	// and doesn't set defaults
	msg() *otg.VxlanV6TunnelDestinationIPMode
	// setMsg unmarshals VxlanV6TunnelDestinationIPMode from protobuf object *otg.VxlanV6TunnelDestinationIPMode
	// and doesn't set defaults
	setMsg(*otg.VxlanV6TunnelDestinationIPMode) VxlanV6TunnelDestinationIPMode
	// provides marshal interface
	Marshal() marshalVxlanV6TunnelDestinationIPMode
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV6TunnelDestinationIPMode
	// validate validates VxlanV6TunnelDestinationIPMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV6TunnelDestinationIPMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns VxlanV6TunnelDestinationIPModeChoiceEnum, set in VxlanV6TunnelDestinationIPMode
	Choice() VxlanV6TunnelDestinationIPModeChoiceEnum
	// setChoice assigns VxlanV6TunnelDestinationIPModeChoiceEnum provided by user to VxlanV6TunnelDestinationIPMode
	setChoice(value VxlanV6TunnelDestinationIPModeChoiceEnum) VxlanV6TunnelDestinationIPMode
	// HasChoice checks if Choice has been set in VxlanV6TunnelDestinationIPMode
	HasChoice() bool
	// Unicast returns VxlanV6TunnelDestinationIPModeUnicast, set in VxlanV6TunnelDestinationIPMode.
	// VxlanV6TunnelDestinationIPModeUnicast is description is TBD
	Unicast() VxlanV6TunnelDestinationIPModeUnicast
	// SetUnicast assigns VxlanV6TunnelDestinationIPModeUnicast provided by user to VxlanV6TunnelDestinationIPMode.
	// VxlanV6TunnelDestinationIPModeUnicast is description is TBD
	SetUnicast(value VxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPMode
	// HasUnicast checks if Unicast has been set in VxlanV6TunnelDestinationIPMode
	HasUnicast() bool
	// Multicast returns VxlanV6TunnelDestinationIPModeMulticast, set in VxlanV6TunnelDestinationIPMode.
	// VxlanV6TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
	Multicast() VxlanV6TunnelDestinationIPModeMulticast
	// SetMulticast assigns VxlanV6TunnelDestinationIPModeMulticast provided by user to VxlanV6TunnelDestinationIPMode.
	// VxlanV6TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
	SetMulticast(value VxlanV6TunnelDestinationIPModeMulticast) VxlanV6TunnelDestinationIPMode
	// HasMulticast checks if Multicast has been set in VxlanV6TunnelDestinationIPMode
	HasMulticast() bool
	setNil()
}

type VxlanV6TunnelDestinationIPModeChoiceEnum string

// Enum of Choice on VxlanV6TunnelDestinationIPMode
var VxlanV6TunnelDestinationIPModeChoice = struct {
	UNICAST   VxlanV6TunnelDestinationIPModeChoiceEnum
	MULTICAST VxlanV6TunnelDestinationIPModeChoiceEnum
}{
	UNICAST:   VxlanV6TunnelDestinationIPModeChoiceEnum("unicast"),
	MULTICAST: VxlanV6TunnelDestinationIPModeChoiceEnum("multicast"),
}

func (obj *vxlanV6TunnelDestinationIPMode) Choice() VxlanV6TunnelDestinationIPModeChoiceEnum {
	return VxlanV6TunnelDestinationIPModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// unicast or multicast
// Choice returns a string
func (obj *vxlanV6TunnelDestinationIPMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *vxlanV6TunnelDestinationIPMode) setChoice(value VxlanV6TunnelDestinationIPModeChoiceEnum) VxlanV6TunnelDestinationIPMode {
	intValue, ok := otg.VxlanV6TunnelDestinationIPMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on VxlanV6TunnelDestinationIPModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.VxlanV6TunnelDestinationIPMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Multicast = nil
	obj.multicastHolder = nil
	obj.obj.Unicast = nil
	obj.unicastHolder = nil

	if value == VxlanV6TunnelDestinationIPModeChoice.UNICAST {
		obj.obj.Unicast = NewVxlanV6TunnelDestinationIPModeUnicast().msg()
	}

	if value == VxlanV6TunnelDestinationIPModeChoice.MULTICAST {
		obj.obj.Multicast = NewVxlanV6TunnelDestinationIPModeMulticast().msg()
	}

	return obj
}

// description is TBD
// Unicast returns a VxlanV6TunnelDestinationIPModeUnicast
func (obj *vxlanV6TunnelDestinationIPMode) Unicast() VxlanV6TunnelDestinationIPModeUnicast {
	if obj.obj.Unicast == nil {
		obj.setChoice(VxlanV6TunnelDestinationIPModeChoice.UNICAST)
	}
	if obj.unicastHolder == nil {
		obj.unicastHolder = &vxlanV6TunnelDestinationIPModeUnicast{obj: obj.obj.Unicast}
	}
	return obj.unicastHolder
}

// description is TBD
// Unicast returns a VxlanV6TunnelDestinationIPModeUnicast
func (obj *vxlanV6TunnelDestinationIPMode) HasUnicast() bool {
	return obj.obj.Unicast != nil
}

// description is TBD
// SetUnicast sets the VxlanV6TunnelDestinationIPModeUnicast value in the VxlanV6TunnelDestinationIPMode object
func (obj *vxlanV6TunnelDestinationIPMode) SetUnicast(value VxlanV6TunnelDestinationIPModeUnicast) VxlanV6TunnelDestinationIPMode {
	obj.setChoice(VxlanV6TunnelDestinationIPModeChoice.UNICAST)
	obj.unicastHolder = nil
	obj.obj.Unicast = value.msg()

	return obj
}

// description is TBD
// Multicast returns a VxlanV6TunnelDestinationIPModeMulticast
func (obj *vxlanV6TunnelDestinationIPMode) Multicast() VxlanV6TunnelDestinationIPModeMulticast {
	if obj.obj.Multicast == nil {
		obj.setChoice(VxlanV6TunnelDestinationIPModeChoice.MULTICAST)
	}
	if obj.multicastHolder == nil {
		obj.multicastHolder = &vxlanV6TunnelDestinationIPModeMulticast{obj: obj.obj.Multicast}
	}
	return obj.multicastHolder
}

// description is TBD
// Multicast returns a VxlanV6TunnelDestinationIPModeMulticast
func (obj *vxlanV6TunnelDestinationIPMode) HasMulticast() bool {
	return obj.obj.Multicast != nil
}

// description is TBD
// SetMulticast sets the VxlanV6TunnelDestinationIPModeMulticast value in the VxlanV6TunnelDestinationIPMode object
func (obj *vxlanV6TunnelDestinationIPMode) SetMulticast(value VxlanV6TunnelDestinationIPModeMulticast) VxlanV6TunnelDestinationIPMode {
	obj.setChoice(VxlanV6TunnelDestinationIPModeChoice.MULTICAST)
	obj.multicastHolder = nil
	obj.obj.Multicast = value.msg()

	return obj
}

func (obj *vxlanV6TunnelDestinationIPMode) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Unicast != nil {

		obj.Unicast().validateObj(vObj, set_default)
	}

	if obj.obj.Multicast != nil {

		obj.Multicast().validateObj(vObj, set_default)
	}

}

func (obj *vxlanV6TunnelDestinationIPMode) setDefault() {
	var choices_set int = 0
	var choice VxlanV6TunnelDestinationIPModeChoiceEnum

	if obj.obj.Unicast != nil {
		choices_set += 1
		choice = VxlanV6TunnelDestinationIPModeChoice.UNICAST
	}

	if obj.obj.Multicast != nil {
		choices_set += 1
		choice = VxlanV6TunnelDestinationIPModeChoice.MULTICAST
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(VxlanV6TunnelDestinationIPModeChoice.MULTICAST)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in VxlanV6TunnelDestinationIPMode")
			}
		} else {
			intVal := otg.VxlanV6TunnelDestinationIPMode_Choice_Enum_value[string(choice)]
			enumValue := otg.VxlanV6TunnelDestinationIPMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
