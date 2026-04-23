package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UpdateProtocolConfigIsisRouterUpdateGroup *****
type updateProtocolConfigIsisRouterUpdateGroup struct {
	validation
	obj            *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	marshaller     marshalUpdateProtocolConfigIsisRouterUpdateGroup
	unMarshaller   unMarshalUpdateProtocolConfigIsisRouterUpdateGroup
	basicHolder    UpdateProtocolConfigIsisRouterBasicAttribute
	advancedHolder UpdateProtocolConfigIsisRouterAdvancedAttribute
}

func NewUpdateProtocolConfigIsisRouterUpdateGroup() UpdateProtocolConfigIsisRouterUpdateGroup {
	obj := updateProtocolConfigIsisRouterUpdateGroup{obj: &otg.UpdateProtocolConfigIsisRouterUpdateGroup{}}
	obj.setDefault()
	return &obj
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) msg() *otg.UpdateProtocolConfigIsisRouterUpdateGroup {
	return obj.obj
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) setMsg(msg *otg.UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisRouterUpdateGroup {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalupdateProtocolConfigIsisRouterUpdateGroup struct {
	obj *updateProtocolConfigIsisRouterUpdateGroup
}

type marshalUpdateProtocolConfigIsisRouterUpdateGroup interface {
	// ToProto marshals UpdateProtocolConfigIsisRouterUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	ToProto() (*otg.UpdateProtocolConfigIsisRouterUpdateGroup, error)
	// ToPbText marshals UpdateProtocolConfigIsisRouterUpdateGroup to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UpdateProtocolConfigIsisRouterUpdateGroup to YAML text
	ToYaml() (string, error)
	// ToJson marshals UpdateProtocolConfigIsisRouterUpdateGroup to JSON text
	ToJson() (string, error)
}

type unMarshalupdateProtocolConfigIsisRouterUpdateGroup struct {
	obj *updateProtocolConfigIsisRouterUpdateGroup
}

type unMarshalUpdateProtocolConfigIsisRouterUpdateGroup interface {
	// FromProto unmarshals UpdateProtocolConfigIsisRouterUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	FromProto(msg *otg.UpdateProtocolConfigIsisRouterUpdateGroup) (UpdateProtocolConfigIsisRouterUpdateGroup, error)
	// FromPbText unmarshals UpdateProtocolConfigIsisRouterUpdateGroup from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UpdateProtocolConfigIsisRouterUpdateGroup from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UpdateProtocolConfigIsisRouterUpdateGroup from JSON text
	FromJson(value string) error
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) Marshal() marshalUpdateProtocolConfigIsisRouterUpdateGroup {
	if obj.marshaller == nil {
		obj.marshaller = &marshalupdateProtocolConfigIsisRouterUpdateGroup{obj: obj}
	}
	return obj.marshaller
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) Unmarshal() unMarshalUpdateProtocolConfigIsisRouterUpdateGroup {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalupdateProtocolConfigIsisRouterUpdateGroup{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalupdateProtocolConfigIsisRouterUpdateGroup) ToProto() (*otg.UpdateProtocolConfigIsisRouterUpdateGroup, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalupdateProtocolConfigIsisRouterUpdateGroup) FromProto(msg *otg.UpdateProtocolConfigIsisRouterUpdateGroup) (UpdateProtocolConfigIsisRouterUpdateGroup, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalupdateProtocolConfigIsisRouterUpdateGroup) ToPbText() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterUpdateGroup) FromPbText(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterUpdateGroup) ToYaml() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterUpdateGroup) FromYaml(value string) error {
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

func (m *marshalupdateProtocolConfigIsisRouterUpdateGroup) ToJson() (string, error) {
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

func (m *unMarshalupdateProtocolConfigIsisRouterUpdateGroup) FromJson(value string) error {
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

func (obj *updateProtocolConfigIsisRouterUpdateGroup) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) Clone() (UpdateProtocolConfigIsisRouterUpdateGroup, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUpdateProtocolConfigIsisRouterUpdateGroup()
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

func (obj *updateProtocolConfigIsisRouterUpdateGroup) setNil() {
	obj.basicHolder = nil
	obj.advancedHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UpdateProtocolConfigIsisRouterUpdateGroup is an update group that applies a common set of property updates to one or more IS-IS routers. All routers listed in names receive the same property updates defined in properties. Use multiple update groups in a single request to apply different property values to different subsets of routers (asymmetric updates).
type UpdateProtocolConfigIsisRouterUpdateGroup interface {
	Validation
	// msg marshals UpdateProtocolConfigIsisRouterUpdateGroup to protobuf object *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	// and doesn't set defaults
	msg() *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	// setMsg unmarshals UpdateProtocolConfigIsisRouterUpdateGroup from protobuf object *otg.UpdateProtocolConfigIsisRouterUpdateGroup
	// and doesn't set defaults
	setMsg(*otg.UpdateProtocolConfigIsisRouterUpdateGroup) UpdateProtocolConfigIsisRouterUpdateGroup
	// provides marshal interface
	Marshal() marshalUpdateProtocolConfigIsisRouterUpdateGroup
	// provides unmarshal interface
	Unmarshal() unMarshalUpdateProtocolConfigIsisRouterUpdateGroup
	// validate validates UpdateProtocolConfigIsisRouterUpdateGroup
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UpdateProtocolConfigIsisRouterUpdateGroup, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Names returns []string, set in UpdateProtocolConfigIsisRouterUpdateGroup.
	Names() []string
	// SetNames assigns []string provided by user to UpdateProtocolConfigIsisRouterUpdateGroup
	SetNames(value []string) UpdateProtocolConfigIsisRouterUpdateGroup
	// Choice returns UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum, set in UpdateProtocolConfigIsisRouterUpdateGroup
	Choice() UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum
	// setChoice assigns UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum provided by user to UpdateProtocolConfigIsisRouterUpdateGroup
	setChoice(value UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum) UpdateProtocolConfigIsisRouterUpdateGroup
	// HasChoice checks if Choice has been set in UpdateProtocolConfigIsisRouterUpdateGroup
	HasChoice() bool
	// getter for Instance to set choice.
	Instance()
	// Basic returns UpdateProtocolConfigIsisRouterBasicAttribute, set in UpdateProtocolConfigIsisRouterUpdateGroup.
	// UpdateProtocolConfigIsisRouterBasicAttribute is the basic properties of an ISIS Router to be updated.
	Basic() UpdateProtocolConfigIsisRouterBasicAttribute
	// SetBasic assigns UpdateProtocolConfigIsisRouterBasicAttribute provided by user to UpdateProtocolConfigIsisRouterUpdateGroup.
	// UpdateProtocolConfigIsisRouterBasicAttribute is the basic properties of an ISIS Router to be updated.
	SetBasic(value UpdateProtocolConfigIsisRouterBasicAttribute) UpdateProtocolConfigIsisRouterUpdateGroup
	// HasBasic checks if Basic has been set in UpdateProtocolConfigIsisRouterUpdateGroup
	HasBasic() bool
	// Advanced returns UpdateProtocolConfigIsisRouterAdvancedAttribute, set in UpdateProtocolConfigIsisRouterUpdateGroup.
	// UpdateProtocolConfigIsisRouterAdvancedAttribute is the advanced properties of an ISIS Router to be updated.
	Advanced() UpdateProtocolConfigIsisRouterAdvancedAttribute
	// SetAdvanced assigns UpdateProtocolConfigIsisRouterAdvancedAttribute provided by user to UpdateProtocolConfigIsisRouterUpdateGroup.
	// UpdateProtocolConfigIsisRouterAdvancedAttribute is the advanced properties of an ISIS Router to be updated.
	SetAdvanced(value UpdateProtocolConfigIsisRouterAdvancedAttribute) UpdateProtocolConfigIsisRouterUpdateGroup
	// HasAdvanced checks if Advanced has been set in UpdateProtocolConfigIsisRouterUpdateGroup
	HasAdvanced() bool
	// InstanceId returns uint32, set in UpdateProtocolConfigIsisRouterUpdateGroup.
	InstanceId() uint32
	// SetInstanceId assigns uint32 provided by user to UpdateProtocolConfigIsisRouterUpdateGroup
	SetInstanceId(value uint32) UpdateProtocolConfigIsisRouterUpdateGroup
	// HasInstanceId checks if InstanceId has been set in UpdateProtocolConfigIsisRouterUpdateGroup
	HasInstanceId() bool
	setNil()
}

// The names of the IS-IS routers to which all property updates in this group will be applied.
// Names returns a []string
func (obj *updateProtocolConfigIsisRouterUpdateGroup) Names() []string {
	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	return obj.obj.Names
}

// The names of the IS-IS routers to which all property updates in this group will be applied.
// SetNames sets the []string value in the UpdateProtocolConfigIsisRouterUpdateGroup object
func (obj *updateProtocolConfigIsisRouterUpdateGroup) SetNames(value []string) UpdateProtocolConfigIsisRouterUpdateGroup {

	if obj.obj.Names == nil {
		obj.obj.Names = make([]string, 0)
	}
	obj.obj.Names = value

	return obj
}

type UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum string

// Enum of Choice on UpdateProtocolConfigIsisRouterUpdateGroup
var UpdateProtocolConfigIsisRouterUpdateGroupChoice = struct {
	BASIC    UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum
	ADVANCED UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum
	INSTANCE UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum
}{
	BASIC:    UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum("basic"),
	ADVANCED: UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum("advanced"),
	INSTANCE: UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum("instance"),
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) Choice() UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum {
	return UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum(obj.obj.Choice.Enum().String())
}

// getter for Instance to set choice
func (obj *updateProtocolConfigIsisRouterUpdateGroup) Instance() {
	obj.setChoice(UpdateProtocolConfigIsisRouterUpdateGroupChoice.INSTANCE)
}

// The category of router attributes to be updated. The properties field contains the actual attribute updates.
// Choice returns a string
func (obj *updateProtocolConfigIsisRouterUpdateGroup) HasChoice() bool {
	return obj.obj.Choice != nil
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) setChoice(value UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum) UpdateProtocolConfigIsisRouterUpdateGroup {
	intValue, ok := otg.UpdateProtocolConfigIsisRouterUpdateGroup_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.UpdateProtocolConfigIsisRouterUpdateGroup_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Advanced = nil
	obj.advancedHolder = nil
	obj.obj.Basic = nil
	obj.basicHolder = nil

	if value == UpdateProtocolConfigIsisRouterUpdateGroupChoice.BASIC {
		obj.obj.Basic = NewUpdateProtocolConfigIsisRouterBasicAttribute().msg()
	}

	if value == UpdateProtocolConfigIsisRouterUpdateGroupChoice.ADVANCED {
		obj.obj.Advanced = NewUpdateProtocolConfigIsisRouterAdvancedAttribute().msg()
	}

	return obj
}

// The rate at which LSPs are re-sent in seconds.
// Basic returns a UpdateProtocolConfigIsisRouterBasicAttribute
func (obj *updateProtocolConfigIsisRouterUpdateGroup) Basic() UpdateProtocolConfigIsisRouterBasicAttribute {
	if obj.obj.Basic == nil {
		obj.setChoice(UpdateProtocolConfigIsisRouterUpdateGroupChoice.BASIC)
	}
	if obj.basicHolder == nil {
		obj.basicHolder = &updateProtocolConfigIsisRouterBasicAttribute{obj: obj.obj.Basic}
	}
	return obj.basicHolder
}

// The rate at which LSPs are re-sent in seconds.
// Basic returns a UpdateProtocolConfigIsisRouterBasicAttribute
func (obj *updateProtocolConfigIsisRouterUpdateGroup) HasBasic() bool {
	return obj.obj.Basic != nil
}

// The rate at which LSPs are re-sent in seconds.
// SetBasic sets the UpdateProtocolConfigIsisRouterBasicAttribute value in the UpdateProtocolConfigIsisRouterUpdateGroup object
func (obj *updateProtocolConfigIsisRouterUpdateGroup) SetBasic(value UpdateProtocolConfigIsisRouterBasicAttribute) UpdateProtocolConfigIsisRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigIsisRouterUpdateGroupChoice.BASIC)
	obj.basicHolder = nil
	obj.obj.Basic = value.msg()

	return obj
}

// The advanced properties of an ISIS Router to be updated.
// Advanced returns a UpdateProtocolConfigIsisRouterAdvancedAttribute
func (obj *updateProtocolConfigIsisRouterUpdateGroup) Advanced() UpdateProtocolConfigIsisRouterAdvancedAttribute {
	if obj.obj.Advanced == nil {
		obj.setChoice(UpdateProtocolConfigIsisRouterUpdateGroupChoice.ADVANCED)
	}
	if obj.advancedHolder == nil {
		obj.advancedHolder = &updateProtocolConfigIsisRouterAdvancedAttribute{obj: obj.obj.Advanced}
	}
	return obj.advancedHolder
}

// The advanced properties of an ISIS Router to be updated.
// Advanced returns a UpdateProtocolConfigIsisRouterAdvancedAttribute
func (obj *updateProtocolConfigIsisRouterUpdateGroup) HasAdvanced() bool {
	return obj.obj.Advanced != nil
}

// The advanced properties of an ISIS Router to be updated.
// SetAdvanced sets the UpdateProtocolConfigIsisRouterAdvancedAttribute value in the UpdateProtocolConfigIsisRouterUpdateGroup object
func (obj *updateProtocolConfigIsisRouterUpdateGroup) SetAdvanced(value UpdateProtocolConfigIsisRouterAdvancedAttribute) UpdateProtocolConfigIsisRouterUpdateGroup {
	obj.setChoice(UpdateProtocolConfigIsisRouterUpdateGroupChoice.ADVANCED)
	obj.advancedHolder = nil
	obj.obj.Advanced = value.msg()

	return obj
}

// The instance properties for this emulated ISIS router.
// InstanceId returns a uint32
func (obj *updateProtocolConfigIsisRouterUpdateGroup) InstanceId() uint32 {

	return *obj.obj.InstanceId

}

// The instance properties for this emulated ISIS router.
// InstanceId returns a uint32
func (obj *updateProtocolConfigIsisRouterUpdateGroup) HasInstanceId() bool {
	return obj.obj.InstanceId != nil
}

// The instance properties for this emulated ISIS router.
// SetInstanceId sets the uint32 value in the UpdateProtocolConfigIsisRouterUpdateGroup object
func (obj *updateProtocolConfigIsisRouterUpdateGroup) SetInstanceId(value uint32) UpdateProtocolConfigIsisRouterUpdateGroup {

	obj.obj.InstanceId = &value
	return obj
}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.Basic != nil {

		obj.Basic().validateObj(vObj, set_default)
	}

	if obj.obj.Advanced != nil {

		obj.Advanced().validateObj(vObj, set_default)
	}

}

func (obj *updateProtocolConfigIsisRouterUpdateGroup) setDefault() {
	var choices_set int = 0
	var choice UpdateProtocolConfigIsisRouterUpdateGroupChoiceEnum

	if obj.obj.Basic != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisRouterUpdateGroupChoice.BASIC
	}

	if obj.obj.Advanced != nil {
		choices_set += 1
		choice = UpdateProtocolConfigIsisRouterUpdateGroupChoice.ADVANCED
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in UpdateProtocolConfigIsisRouterUpdateGroup")
			}
		} else {
			intVal := otg.UpdateProtocolConfigIsisRouterUpdateGroup_Choice_Enum_value[string(choice)]
			enumValue := otg.UpdateProtocolConfigIsisRouterUpdateGroup_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

	if obj.obj.InstanceId == nil {
		obj.SetInstanceId(1)
	}

}
