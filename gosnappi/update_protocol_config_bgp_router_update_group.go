package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigBgpRouterUpdateGroup *****
type updateProtocolConfigBgpRouterUpdateGroup struct {
	validation
	obj            *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	marshaller     marshalUpdateProtocolConfigBgpRouterUpdateGroup
	unMarshaller   unMarshalUpdateProtocolConfigBgpRouterUpdateGroup
	advancedHolder UpdateProtocolConfigBgpRouterAdvancedAttribute
}

func NewUpdateProtocolConfigBgpRouterUpdateGroup() UpdateProtocolConfigBgpRouterUpdateGroup {
	obj := updateProtocolConfigBgpRouterUpdateGroup{obj: &otg.UpdateProtocolConfigBgpRouterUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) msg() *otg.UpdateProtocolConfigBgpRouterUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigBgpRouterUpdateGroup struct {
	obj *updateProtocolConfigBgpRouterUpdateGroup
}

type marshalUpdateProtocolConfigBgpRouterUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigBgpRouterUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigBgpRouterUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigBgpRouterUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigBgpRouterUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigBgpRouterUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigBgpRouterUpdateGroup struct {
	obj *updateProtocolConfigBgpRouterUpdateGroup
}

type unMarshalUpdateProtocolConfigBgpRouterUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigBgpRouterUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigBgpRouterUpdateGroup) (UpdateProtocolConfigBgpRouterUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigBgpRouterUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigBgpRouterUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigBgpRouterUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) Marshal() marshalUpdateProtocolConfigBgpRouterUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigBgpRouterUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigBgpRouterUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigBgpRouterUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigBgpRouterUpdateGroup) ToProto() (*otg.UpdateProtocolConfigBgpRouterUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigBgpRouterUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigBgpRouterUpdateGroup) (UpdateProtocolConfigBgpRouterUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigBgpRouterUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigBgpRouterUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigBgpRouterUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigBgpRouterUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigBgpRouterUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) Clone() (UpdateProtocolConfigBgpRouterUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigBgpRouterUpdateGroup()
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

func (obj *updateProtocolConfigBgpRouterUpdateGroup) setNil() {
	obj.advancedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigBgpRouterUpdateGroup is a BGP router entry for router-level attribute updates.
type UpdateProtocolConfigBgpRouterUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigBgpRouterUpdateGroup to protobuf object *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigBgpRouterUpdateGroup from protobuf object *otg.UpdateProtocolConfigBgpRouterUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigBgpRouterUpdateGroup) UpdateProtocolConfigBgpRouterUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigBgpRouterUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigBgpRouterUpdateGroup
	// validate validates UpdateProtocolConfigBgpRouterUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigBgpRouterUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	SetNames(value []string) UpdateProtocolConfigBgpRouterUpdateGroup
	// Choice returns UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum, set in UpdateProtocolConfigBgpRouterUpdateGroup
	Choice() UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
	// setChoice assigns UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	setChoice(value UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasChoice checks if Choice has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasChoice() bool
	// LocalAs returns uint32, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	LocalAs() uint32
	// SetLocalAs assigns uint32 provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	SetLocalAs(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasLocalAs checks if LocalAs has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasLocalAs() bool
	// RouterId returns string, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	RouterId() string
	// SetRouterId assigns string provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	SetRouterId(value string) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasRouterId checks if RouterId has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasRouterId() bool
	// HoldTime returns uint32, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	HoldTime() uint32
	// SetHoldTime assigns uint32 provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	SetHoldTime(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasHoldTime checks if HoldTime has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasHoldTime() bool
	// KeepaliveTime returns uint32, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	KeepaliveTime() uint32
	// SetKeepaliveTime assigns uint32 provided by user to UpdateProtocolConfigBgpRouterUpdateGroup
	SetKeepaliveTime(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasKeepaliveTime checks if KeepaliveTime has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasKeepaliveTime() bool
	// Advanced returns UpdateProtocolConfigBgpRouterAdvancedAttribute, set in UpdateProtocolConfigBgpRouterUpdateGroup.
	// UpdateProtocolConfigBgpRouterAdvancedAttribute is the advanced properties of a BGP router to be updated.
	Advanced() UpdateProtocolConfigBgpRouterAdvancedAttribute
	// SetAdvanced assigns UpdateProtocolConfigBgpRouterAdvancedAttribute provided by user to UpdateProtocolConfigBgpRouterUpdateGroup.
	// UpdateProtocolConfigBgpRouterAdvancedAttribute is the advanced properties of a BGP router to be updated.
	SetAdvanced(value UpdateProtocolConfigBgpRouterAdvancedAttribute) UpdateProtocolConfigBgpRouterUpdateGroup
	// HasAdvanced checks if Advanced has been set in UpdateProtocolConfigBgpRouterUpdateGroup
	HasAdvanced() bool
	setNil()
}

// Names of the BGP router to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// Names returns a []string
func (obj *updateProtocolConfigBgpRouterUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// Names of the BGP router to be updated.
//
// x-constraint:
// - /components/schemas/Device.BgpRouter/properties/name
//
// SetNames sets the []string value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetNames(value []string) UpdateProtocolConfigBgpRouterUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum string

// Enum of Choice on UpdateProtocolConfigBgpRouterUpdateGroup
var UpdateProtocolConfigBgpRouterUpdateGroupChoice = struct {
	LOCAL_AS       UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
	ROUTER_ID      UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
	HOLD_TIME      UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
	KEEPALIVE_TIME UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
	ADVANCED       UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum
}{
	LOCAL_AS:       UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum("local_as"),
	ROUTER_ID:      UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum("router_id"),
	HOLD_TIME:      UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum("hold_time"),
	KEEPALIVE_TIME: UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum("keepalive_time"),
	ADVANCED:       UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum("advanced"),
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) Choice() UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum {
	return UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum(obj.obj.Choice.Enum().String())
}

// The category of router attributes to be updated. The properties field contains the actual attribute updates.
// Choice returns a string
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) setChoice(value UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum) UpdateProtocolConfigBgpRouterUpdateGroup {
	intValue, ok := otg.UpdateProtocolConfigBgpRouterUpdateGroup_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigBgpRouterUpdateGroup_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Advanced = nil
	obj.advancedHolder = nil
	obj.obj.KeepaliveTime = nil
	obj.obj.HoldTime = nil
	obj.obj.RouterId = nil
	obj.obj.LocalAs = nil

	if value == UpdateProtocolConfigBgpRouterUpdateGroupChoice.LOCAL_AS {
		defaultValue := uint32(65000)
		obj.obj.LocalAs = &defaultValue
	}

	if value == UpdateProtocolConfigBgpRouterUpdateGroupChoice.HOLD_TIME {
		defaultValue := uint32(90)
		obj.obj.HoldTime = &defaultValue
	}

	if value == UpdateProtocolConfigBgpRouterUpdateGroupChoice.KEEPALIVE_TIME {
		defaultValue := uint32(30)
		obj.obj.KeepaliveTime = &defaultValue
	}

	if value == UpdateProtocolConfigBgpRouterUpdateGroupChoice.ADVANCED {
		obj.obj.Advanced = NewUpdateProtocolConfigBgpRouterAdvancedAttribute().msg()
	}

	return obj
}

// The local AS number to be updated on the BGP router.
// LocalAs returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) LocalAs() uint32 {

	if obj.obj.LocalAs == nil {
		obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.LOCAL_AS)
	}

	return *obj.obj.LocalAs

}

// The local AS number to be updated on the BGP router.
// LocalAs returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasLocalAs() bool {
	return obj.obj.LocalAs != nil
}

// The local AS number to be updated on the BGP router.
// SetLocalAs sets the uint32 value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetLocalAs(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.LOCAL_AS)
	obj.obj.LocalAs = &value
	return obj
}

// The router ID to be updated on the BGP router.
// RouterId returns a string
func (obj *updateProtocolConfigBgpRouterUpdateGroup) RouterId() string {

	if obj.obj.RouterId == nil {
		obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.ROUTER_ID)
	}

	return *obj.obj.RouterId

}

// The router ID to be updated on the BGP router.
// RouterId returns a string
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasRouterId() bool {
	return obj.obj.RouterId != nil
}

// The router ID to be updated on the BGP router.
// SetRouterId sets the string value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetRouterId(value string) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.ROUTER_ID)
	obj.obj.RouterId = &value
	return obj
}

// The hold time in seconds to be updated on the BGP router.
// HoldTime returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HoldTime() uint32 {

	if obj.obj.HoldTime == nil {
		obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.HOLD_TIME)
	}

	return *obj.obj.HoldTime

}

// The hold time in seconds to be updated on the BGP router.
// HoldTime returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasHoldTime() bool {
	return obj.obj.HoldTime != nil
}

// The hold time in seconds to be updated on the BGP router.
// SetHoldTime sets the uint32 value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetHoldTime(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.HOLD_TIME)
	obj.obj.HoldTime = &value
	return obj
}

// The keepalive time in seconds to be updated on the BGP router.
// KeepaliveTime returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) KeepaliveTime() uint32 {

	if obj.obj.KeepaliveTime == nil {
		obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.KEEPALIVE_TIME)
	}

	return *obj.obj.KeepaliveTime

}

// The keepalive time in seconds to be updated on the BGP router.
// KeepaliveTime returns a uint32
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasKeepaliveTime() bool {
	return obj.obj.KeepaliveTime != nil
}

// The keepalive time in seconds to be updated on the BGP router.
// SetKeepaliveTime sets the uint32 value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetKeepaliveTime(value uint32) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.KEEPALIVE_TIME)
	obj.obj.KeepaliveTime = &value
	return obj
}

// The advanced properties of a BGP Router to be updated.
// Advanced returns a UpdateProtocolConfigBgpRouterAdvancedAttribute
func (obj *updateProtocolConfigBgpRouterUpdateGroup) Advanced() UpdateProtocolConfigBgpRouterAdvancedAttribute {
	if obj.obj.Advanced == nil {
		obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.ADVANCED)
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &updateProtocolConfigBgpRouterAdvancedAttribute{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// The advanced properties of a BGP Router to be updated.
// Advanced returns a UpdateProtocolConfigBgpRouterAdvancedAttribute
func (obj *updateProtocolConfigBgpRouterUpdateGroup) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// The advanced properties of a BGP Router to be updated.
// SetAdvanced sets the UpdateProtocolConfigBgpRouterAdvancedAttribute value in the UpdateProtocolConfigBgpRouterUpdateGroup object
func (obj *updateProtocolConfigBgpRouterUpdateGroup) SetAdvanced(value UpdateProtocolConfigBgpRouterAdvancedAttribute) UpdateProtocolConfigBgpRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigBgpRouterUpdateGroupChoice.ADVANCED)
	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.LocalAs != nil {

		if *obj.obj.LocalAs > 4294967295 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigBgpRouterUpdateGroup.LocalAs <= 4294967295 but Got %d", *obj.obj.LocalAs))
		}

	}

	if obj.obj.RouterId != nil {

		err := obj.validateIpv4(obj.RouterId())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on UpdateProtocolConfigBgpRouterUpdateGroup.RouterId"))
		}

	}

	if obj.obj.HoldTime != nil {

		if *obj.obj.HoldTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigBgpRouterUpdateGroup.HoldTime <= 65535 but Got %d", *obj.obj.HoldTime))
		}

	}

	if obj.obj.KeepaliveTime != nil {

		if *obj.obj.KeepaliveTime > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UpdateProtocolConfigBgpRouterUpdateGroup.KeepaliveTime <= 65535 but Got %d", *obj.obj.KeepaliveTime))
		}

	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

}

func (obj *updateProtocolConfigBgpRouterUpdateGroup) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigBgpRouterUpdateGroupChoiceEnum

	if obj.obj.LocalAs != nil {
		choices_set += 1
		choice = UpdateProtocolConfigBgpRouterUpdateGroupChoice.LOCAL_AS
	}

	if obj.obj.RouterId != nil {
		choices_set += 1
		choice = UpdateProtocolConfigBgpRouterUpdateGroupChoice.ROUTER_ID
	}

	if obj.obj.HoldTime != nil {
		choices_set += 1
		choice = UpdateProtocolConfigBgpRouterUpdateGroupChoice.HOLD_TIME
	}

	if obj.obj.KeepaliveTime != nil {
		choices_set += 1
		choice = UpdateProtocolConfigBgpRouterUpdateGroupChoice.KEEPALIVE_TIME
	}

	if obj.obj.Advanced != nil {
		choices_set += 1
		choice = UpdateProtocolConfigBgpRouterUpdateGroupChoice.ADVANCED
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigBgpRouterUpdateGroup")
			}
		} else {
			intVal := otg.UpdateProtocolConfigBgpRouterUpdateGroup_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigBgpRouterUpdateGroup_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.LocalAs == nil && choice == UpdateProtocolConfigBgpRouterUpdateGroupChoice.LOCAL_AS {
		obj.SetLocalAs(65000)
	}
	if obj.obj.HoldTime == nil && choice == UpdateProtocolConfigBgpRouterUpdateGroupChoice.HOLD_TIME {
		obj.SetHoldTime(90)
	}
	if obj.obj.KeepaliveTime == nil && choice == UpdateProtocolConfigBgpRouterUpdateGroupChoice.KEEPALIVE_TIME {
		obj.SetKeepaliveTime(30)
	}

}
