package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** VxlanV4TunnelDestinationIPMode *****
type vxlanV4TunnelDestinationIPMode struct {
	validation
	obj             *otg.VxlanV4TunnelDestinationIPMode
	marshaller      marshalVxlanV4TunnelDestinationIPMode
	unMarshaller    unMarshalVxlanV4TunnelDestinationIPMode
	unicastHolder   VxlanV4TunnelDestinationIPModeUnicast
	multicastHolder VxlanV4TunnelDestinationIPModeMulticast
}

func NewVxlanV4TunnelDestinationIPMode() VxlanV4TunnelDestinationIPMode {
	obj := vxlanV4TunnelDestinationIPMode{obj: &otg.VxlanV4TunnelDestinationIPMode{}}
	obj.setDefault()
	return &obj
}

func (obj *vxlanV4TunnelDestinationIPMode) msg() *otg.VxlanV4TunnelDestinationIPMode {
	return obj.obj
}

func (obj *vxlanV4TunnelDestinationIPMode) setMsg(msg *otg.VxlanV4TunnelDestinationIPMode) VxlanV4TunnelDestinationIPMode {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalvxlanV4TunnelDestinationIPMode struct {
	obj *vxlanV4TunnelDestinationIPMode
}

type marshalVxlanV4TunnelDestinationIPMode interface {
	// ToProto marshals VxlanV4TunnelDestinationIPMode to protobuf object *otg.VxlanV4TunnelDestinationIPMode
	ToProto() (*otg.VxlanV4TunnelDestinationIPMode, error)
	// ToPbText marshals VxlanV4TunnelDestinationIPMode to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals VxlanV4TunnelDestinationIPMode to YAML text
	ToYaml() (string, error)
	// ToJson marshals VxlanV4TunnelDestinationIPMode to JSON text
	ToJson() (string, error)
	// ToJsonRaw marshals VxlanV4TunnelDestinationIPMode to raw JSON text
	ToJsonRaw() (string, error)
}

type unMarshalvxlanV4TunnelDestinationIPMode struct {
	obj *vxlanV4TunnelDestinationIPMode
}

type unMarshalVxlanV4TunnelDestinationIPMode interface {
	// FromProto unmarshals VxlanV4TunnelDestinationIPMode from protobuf object *otg.VxlanV4TunnelDestinationIPMode
	FromProto(msg *otg.VxlanV4TunnelDestinationIPMode) (VxlanV4TunnelDestinationIPMode, error)
	// FromPbText unmarshals VxlanV4TunnelDestinationIPMode from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals VxlanV4TunnelDestinationIPMode from YAML text
	FromYaml(value string) error
	// FromJson unmarshals VxlanV4TunnelDestinationIPMode from JSON text
	FromJson(value string) error
}

func (obj *vxlanV4TunnelDestinationIPMode) Marshal() marshalVxlanV4TunnelDestinationIPMode {
	if obj.marshaller == nil {
		obj.marshaller = &marshalvxlanV4TunnelDestinationIPMode{obj: obj}
	}
	return obj.marshaller
}

func (obj *vxlanV4TunnelDestinationIPMode) Unmarshal() unMarshalVxlanV4TunnelDestinationIPMode {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalvxlanV4TunnelDestinationIPMode{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalvxlanV4TunnelDestinationIPMode) ToProto() (*otg.VxlanV4TunnelDestinationIPMode, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalvxlanV4TunnelDestinationIPMode) FromProto(msg *otg.VxlanV4TunnelDestinationIPMode) (VxlanV4TunnelDestinationIPMode, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalvxlanV4TunnelDestinationIPMode) ToPbText() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPMode) FromPbText(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPMode) ToYaml() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPMode) FromYaml(value string) error {
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

func (m *marshalvxlanV4TunnelDestinationIPMode) ToJsonRaw() (string, error) {
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
	return string(data), nil
}

func (m *marshalvxlanV4TunnelDestinationIPMode) ToJson() (string, error) {
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

func (m *unMarshalvxlanV4TunnelDestinationIPMode) FromJson(value string) error {
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

func (obj *vxlanV4TunnelDestinationIPMode) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPMode) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *vxlanV4TunnelDestinationIPMode) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *vxlanV4TunnelDestinationIPMode) Clone() (VxlanV4TunnelDestinationIPMode, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewVxlanV4TunnelDestinationIPMode()
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

func (obj *vxlanV4TunnelDestinationIPMode) setNil() {
	obj.unicastHolder = nil
	obj.multicastHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// VxlanV4TunnelDestinationIPMode is communication mode between the VTEPs, either unicast or multicast.
type VxlanV4TunnelDestinationIPMode interface {
	Validation
	// msg marshals VxlanV4TunnelDestinationIPMode to protobuf object *otg.VxlanV4TunnelDestinationIPMode
	// and doesn't set defaults
	msg() *otg.VxlanV4TunnelDestinationIPMode
	// setMsg unmarshals VxlanV4TunnelDestinationIPMode from protobuf object *otg.VxlanV4TunnelDestinationIPMode
	// and doesn't set defaults
	setMsg(*otg.VxlanV4TunnelDestinationIPMode) VxlanV4TunnelDestinationIPMode
	// provides marshal interface
	Marshal() marshalVxlanV4TunnelDestinationIPMode
	// provides unmarshal interface
	Unmarshal() unMarshalVxlanV4TunnelDestinationIPMode
	// validate validates VxlanV4TunnelDestinationIPMode
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (VxlanV4TunnelDestinationIPMode, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns VxlanV4TunnelDestinationIPModeChoiceEnum, set in VxlanV4TunnelDestinationIPMode
	Choice() VxlanV4TunnelDestinationIPModeChoiceEnum
	// setChoice assigns VxlanV4TunnelDestinationIPModeChoiceEnum provided by user to VxlanV4TunnelDestinationIPMode
	setChoice(value VxlanV4TunnelDestinationIPModeChoiceEnum) VxlanV4TunnelDestinationIPMode
	// HasChoice checks if Choice has been set in VxlanV4TunnelDestinationIPMode
	HasChoice() bool
	// Unicast returns VxlanV4TunnelDestinationIPModeUnicast, set in VxlanV4TunnelDestinationIPMode.
	// VxlanV4TunnelDestinationIPModeUnicast is description is TBD
	Unicast() VxlanV4TunnelDestinationIPModeUnicast
	// SetUnicast assigns VxlanV4TunnelDestinationIPModeUnicast provided by user to VxlanV4TunnelDestinationIPMode.
	// VxlanV4TunnelDestinationIPModeUnicast is description is TBD
	SetUnicast(value VxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPMode
	// HasUnicast checks if Unicast has been set in VxlanV4TunnelDestinationIPMode
	HasUnicast() bool
	// Multicast returns VxlanV4TunnelDestinationIPModeMulticast, set in VxlanV4TunnelDestinationIPMode.
	// VxlanV4TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
	Multicast() VxlanV4TunnelDestinationIPModeMulticast
	// SetMulticast assigns VxlanV4TunnelDestinationIPModeMulticast provided by user to VxlanV4TunnelDestinationIPMode.
	// VxlanV4TunnelDestinationIPModeMulticast is multicast Group address for member VNI(VXLAN Network Identifier)
	SetMulticast(value VxlanV4TunnelDestinationIPModeMulticast) VxlanV4TunnelDestinationIPMode
	// HasMulticast checks if Multicast has been set in VxlanV4TunnelDestinationIPMode
	HasMulticast() bool
	setNil()
}

type VxlanV4TunnelDestinationIPModeChoiceEnum string

// Enum of Choice on VxlanV4TunnelDestinationIPMode
var VxlanV4TunnelDestinationIPModeChoice = struct {
	UNICAST   VxlanV4TunnelDestinationIPModeChoiceEnum
	MULTICAST VxlanV4TunnelDestinationIPModeChoiceEnum
}{
	UNICAST:   VxlanV4TunnelDestinationIPModeChoiceEnum("unicast"),
	MULTICAST: VxlanV4TunnelDestinationIPModeChoiceEnum("multicast"),
}

func (obj *vxlanV4TunnelDestinationIPMode) Choice() VxlanV4TunnelDestinationIPModeChoiceEnum {
	return VxlanV4TunnelDestinationIPModeChoiceEnum(obj.obj.Choice.Enum().String())
}

// unicast or multicast
// Choice returns a string
func (obj *vxlanV4TunnelDestinationIPMode) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *vxlanV4TunnelDestinationIPMode) setChoice(value VxlanV4TunnelDestinationIPModeChoiceEnum) VxlanV4TunnelDestinationIPMode {
	intValue, ok := otg.VxlanV4TunnelDestinationIPMode_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on VxlanV4TunnelDestinationIPModeChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.VxlanV4TunnelDestinationIPMode_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Multicast = nil
	obj.multicastHolder = nil
	obj.obj.Unicast = nil
	obj.unicastHolder = nil

	if value == VxlanV4TunnelDestinationIPModeChoice.UNICAST {
		obj.obj.Unicast = NewVxlanV4TunnelDestinationIPModeUnicast().msg()
	}

	if value == VxlanV4TunnelDestinationIPModeChoice.MULTICAST {
		obj.obj.Multicast = NewVxlanV4TunnelDestinationIPModeMulticast().msg()
	}

	return obj
}

// description is TBD
// Unicast returns a VxlanV4TunnelDestinationIPModeUnicast
func (obj *vxlanV4TunnelDestinationIPMode) Unicast() VxlanV4TunnelDestinationIPModeUnicast {
	if obj.obj.Unicast == nil {
		obj.setChoice(VxlanV4TunnelDestinationIPModeChoice.UNICAST)
	}
	if obj.unicastHolder == nil {
		obj.unicastHolder = &vxlanV4TunnelDestinationIPModeUnicast{obj: obj.obj.Unicast}
	}
	return obj.unicastHolder
}

// description is TBD
// Unicast returns a VxlanV4TunnelDestinationIPModeUnicast
func (obj *vxlanV4TunnelDestinationIPMode) HasUnicast() bool {
	return obj.obj.Unicast != nil
}

// description is TBD
// SetUnicast sets the VxlanV4TunnelDestinationIPModeUnicast value in the VxlanV4TunnelDestinationIPMode object
func (obj *vxlanV4TunnelDestinationIPMode) SetUnicast(value VxlanV4TunnelDestinationIPModeUnicast) VxlanV4TunnelDestinationIPMode {
	obj.setChoice(VxlanV4TunnelDestinationIPModeChoice.UNICAST)
	obj.unicastHolder = nil
	obj.obj.Unicast = value.msg()

	return obj
}

// description is TBD
// Multicast returns a VxlanV4TunnelDestinationIPModeMulticast
func (obj *vxlanV4TunnelDestinationIPMode) Multicast() VxlanV4TunnelDestinationIPModeMulticast {
	if obj.obj.Multicast == nil {
		obj.setChoice(VxlanV4TunnelDestinationIPModeChoice.MULTICAST)
	}
	if obj.multicastHolder == nil {
		obj.multicastHolder = &vxlanV4TunnelDestinationIPModeMulticast{obj: obj.obj.Multicast}
	}
	return obj.multicastHolder
}

// description is TBD
// Multicast returns a VxlanV4TunnelDestinationIPModeMulticast
func (obj *vxlanV4TunnelDestinationIPMode) HasMulticast() bool {
	return obj.obj.Multicast != nil
}

// description is TBD
// SetMulticast sets the VxlanV4TunnelDestinationIPModeMulticast value in the VxlanV4TunnelDestinationIPMode object
func (obj *vxlanV4TunnelDestinationIPMode) SetMulticast(value VxlanV4TunnelDestinationIPModeMulticast) VxlanV4TunnelDestinationIPMode {
	obj.setChoice(VxlanV4TunnelDestinationIPModeChoice.MULTICAST)
	obj.multicastHolder = nil
	obj.obj.Multicast = value.msg()

	return obj
}

func (obj *vxlanV4TunnelDestinationIPMode) validateObj(vObj *validation, set_default bool) {
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

func (obj *vxlanV4TunnelDestinationIPMode) setDefault() {
	var choices_set int = 0
	var choice VxlanV4TunnelDestinationIPModeChoiceEnum

	if obj.obj.Unicast != nil {
		choices_set += 1
		choice = VxlanV4TunnelDestinationIPModeChoice.UNICAST
	}

	if obj.obj.Multicast != nil {
		choices_set += 1
		choice = VxlanV4TunnelDestinationIPModeChoice.MULTICAST
	}
	if choices_set == 0 {
		if obj.obj.Choice == nil {
			obj.setChoice(VxlanV4TunnelDestinationIPModeChoice.MULTICAST)

		}

	} else if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in VxlanV4TunnelDestinationIPMode")
			}
		} else {
			intVal := otg.VxlanV4TunnelDestinationIPMode_Choice_Enum_value[string(choice)]
			enumValue := otg.VxlanV4TunnelDestinationIPMode_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
