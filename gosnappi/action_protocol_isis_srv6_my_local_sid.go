package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6MyLocalSid *****
type actionProtocolIsisSrv6MyLocalSid struct {
	validation
	obj          *otg.ActionProtocolIsisSrv6MyLocalSid
	marshaller   marshalActionProtocolIsisSrv6MyLocalSid
	unMarshaller unMarshalActionProtocolIsisSrv6MyLocalSid
	addHolder    ActionProtocolIsisSrv6MyLocalSidAdd
	modifyHolder ActionProtocolIsisSrv6MyLocalSidModify
	deleteHolder ActionProtocolIsisSrv6MyLocalSidDelete
}

func NewActionProtocolIsisSrv6MyLocalSid() ActionProtocolIsisSrv6MyLocalSid {
	obj := actionProtocolIsisSrv6MyLocalSid{obj: &otg.ActionProtocolIsisSrv6MyLocalSid{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6MyLocalSid) msg() *otg.ActionProtocolIsisSrv6MyLocalSid {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6MyLocalSid) setMsg(msg *otg.ActionProtocolIsisSrv6MyLocalSid) ActionProtocolIsisSrv6MyLocalSid {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6MyLocalSid struct {
	obj *actionProtocolIsisSrv6MyLocalSid
}

type marshalActionProtocolIsisSrv6MyLocalSid interface {
	// ToProto marshals ActionProtocolIsisSrv6MyLocalSid to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSid
	ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSid, error)
	// ToPbText marshals ActionProtocolIsisSrv6MyLocalSid to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6MyLocalSid to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6MyLocalSid to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6MyLocalSid struct {
	obj *actionProtocolIsisSrv6MyLocalSid
}

type unMarshalActionProtocolIsisSrv6MyLocalSid interface {
	// FromProto unmarshals ActionProtocolIsisSrv6MyLocalSid from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSid
	FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSid) (ActionProtocolIsisSrv6MyLocalSid, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6MyLocalSid from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6MyLocalSid from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6MyLocalSid from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6MyLocalSid) Marshal() marshalActionProtocolIsisSrv6MyLocalSid {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6MyLocalSid{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6MyLocalSid) Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSid {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6MyLocalSid{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6MyLocalSid) ToProto() (*otg.ActionProtocolIsisSrv6MyLocalSid, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6MyLocalSid) FromProto(msg *otg.ActionProtocolIsisSrv6MyLocalSid) (ActionProtocolIsisSrv6MyLocalSid, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6MyLocalSid) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSid) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSid) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSid) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6MyLocalSid) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6MyLocalSid) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6MyLocalSid) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSid) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6MyLocalSid) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6MyLocalSid) Clone() (ActionProtocolIsisSrv6MyLocalSid, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6MyLocalSid()
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

func (obj *actionProtocolIsisSrv6MyLocalSid) setNil() {
	obj.addHolder = nil
	obj.modifyHolder = nil
	obj.deleteHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisSrv6MyLocalSid is trigger a My Local SID lifecycle operation (add, modify, or delete) on the specified IS-IS routers (RFC 8986 Section 4 Instantiation, Re-instantiation, and Un-instantiation). In IxNetwork, this operation manipulates isisSRv6EndSIDList entries under the IS-IS router and triggers IS-IS LSP re-advertisement.
type ActionProtocolIsisSrv6MyLocalSid interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6MyLocalSid to protobuf object *otg.ActionProtocolIsisSrv6MyLocalSid
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6MyLocalSid
	// setMsg unmarshals ActionProtocolIsisSrv6MyLocalSid from protobuf object *otg.ActionProtocolIsisSrv6MyLocalSid
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6MyLocalSid) ActionProtocolIsisSrv6MyLocalSid
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6MyLocalSid
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6MyLocalSid
	// validate validates ActionProtocolIsisSrv6MyLocalSid
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6MyLocalSid, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolIsisSrv6MyLocalSid.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolIsisSrv6MyLocalSid
	SetRouterNames(value []string) ActionProtocolIsisSrv6MyLocalSid
	// Choice returns ActionProtocolIsisSrv6MyLocalSidChoiceEnum, set in ActionProtocolIsisSrv6MyLocalSid
	Choice() ActionProtocolIsisSrv6MyLocalSidChoiceEnum
	// setChoice assigns ActionProtocolIsisSrv6MyLocalSidChoiceEnum provided by user to ActionProtocolIsisSrv6MyLocalSid
	setChoice(value ActionProtocolIsisSrv6MyLocalSidChoiceEnum) ActionProtocolIsisSrv6MyLocalSid
	// Add returns ActionProtocolIsisSrv6MyLocalSidAdd, set in ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidAdd is program new SRv6 My Local SID entries on the routers specified by router_names. After this operation the router matches packets whose DA falls within each SID prefix and applies the configured behavior. Session impact: implementation-specific. In IxNetwork the IS-IS router re-advertises its Locator TLV after the add.
	Add() ActionProtocolIsisSrv6MyLocalSidAdd
	// SetAdd assigns ActionProtocolIsisSrv6MyLocalSidAdd provided by user to ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidAdd is program new SRv6 My Local SID entries on the routers specified by router_names. After this operation the router matches packets whose DA falls within each SID prefix and applies the configured behavior. Session impact: implementation-specific. In IxNetwork the IS-IS router re-advertises its Locator TLV after the add.
	SetAdd(value ActionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSid
	// HasAdd checks if Add has been set in ActionProtocolIsisSrv6MyLocalSid
	HasAdd() bool
	// Modify returns ActionProtocolIsisSrv6MyLocalSidModify, set in ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidModify is modify existing SRv6 My Local SID entries on the routers specified by router_names. Each entry is matched by sid_prefix + prefix_length; the remaining fields replace the existing values without a delete-and-re-add cycle. Session impact: implementation-specific. In IxNetwork IS-IS re-sends its LSP with the updated SID set.
	Modify() ActionProtocolIsisSrv6MyLocalSidModify
	// SetModify assigns ActionProtocolIsisSrv6MyLocalSidModify provided by user to ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidModify is modify existing SRv6 My Local SID entries on the routers specified by router_names. Each entry is matched by sid_prefix + prefix_length; the remaining fields replace the existing values without a delete-and-re-add cycle. Session impact: implementation-specific. In IxNetwork IS-IS re-sends its LSP with the updated SID set.
	SetModify(value ActionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSid
	// HasModify checks if Modify has been set in ActionProtocolIsisSrv6MyLocalSid
	HasModify() bool
	// Delete returns ActionProtocolIsisSrv6MyLocalSidDelete, set in ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidDelete is remove SRv6 My Local SID entries from the routers specified by router_names. After deletion, packets whose DA matches these SID prefixes are no longer forwarded (RFC 8986 Section 4 Un-instantiation). Session impact: implementation-specific. In IxNetwork IS-IS withdraws the affected SIDs from its Locator TLV.
	Delete() ActionProtocolIsisSrv6MyLocalSidDelete
	// SetDelete assigns ActionProtocolIsisSrv6MyLocalSidDelete provided by user to ActionProtocolIsisSrv6MyLocalSid.
	// ActionProtocolIsisSrv6MyLocalSidDelete is remove SRv6 My Local SID entries from the routers specified by router_names. After deletion, packets whose DA matches these SID prefixes are no longer forwarded (RFC 8986 Section 4 Un-instantiation). Session impact: implementation-specific. In IxNetwork IS-IS withdraws the affected SIDs from its Locator TLV.
	SetDelete(value ActionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSid
	// HasDelete checks if Delete has been set in ActionProtocolIsisSrv6MyLocalSid
	HasDelete() bool
	setNil()
}

// Names of the IS-IS routers on which to perform the My Local SID operation. An empty or null list applies the operation to all configured IS-IS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolIsisSrv6MyLocalSid) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// Names of the IS-IS routers on which to perform the My Local SID operation. An empty or null list applies the operation to all configured IS-IS routers.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the ActionProtocolIsisSrv6MyLocalSid object
func (obj *actionProtocolIsisSrv6MyLocalSid) SetRouterNames(value []string) ActionProtocolIsisSrv6MyLocalSid {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

type ActionProtocolIsisSrv6MyLocalSidChoiceEnum string

// Enum of Choice on ActionProtocolIsisSrv6MyLocalSid
var ActionProtocolIsisSrv6MyLocalSidChoice = struct {
	ADD    ActionProtocolIsisSrv6MyLocalSidChoiceEnum
	MODIFY ActionProtocolIsisSrv6MyLocalSidChoiceEnum
	DELETE ActionProtocolIsisSrv6MyLocalSidChoiceEnum
}{
	ADD:    ActionProtocolIsisSrv6MyLocalSidChoiceEnum("add"),
	MODIFY: ActionProtocolIsisSrv6MyLocalSidChoiceEnum("modify"),
	DELETE: ActionProtocolIsisSrv6MyLocalSidChoiceEnum("delete"),
}

func (obj *actionProtocolIsisSrv6MyLocalSid) Choice() ActionProtocolIsisSrv6MyLocalSidChoiceEnum {
	return ActionProtocolIsisSrv6MyLocalSidChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolIsisSrv6MyLocalSid) setChoice(value ActionProtocolIsisSrv6MyLocalSidChoiceEnum) ActionProtocolIsisSrv6MyLocalSid {
	intValue, ok := otg.ActionProtocolIsisSrv6MyLocalSid_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisSrv6MyLocalSidChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsisSrv6MyLocalSid_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.Delete = nil
	obj.deleteHolder = nil
	obj.obj.Modify = nil
	obj.modifyHolder = nil
	obj.obj.Add = nil
	obj.addHolder = nil

	if value == ActionProtocolIsisSrv6MyLocalSidChoice.ADD {
		obj.obj.Add = NewActionProtocolIsisSrv6MyLocalSidAdd().msg()
	}

	if value == ActionProtocolIsisSrv6MyLocalSidChoice.MODIFY {
		obj.obj.Modify = NewActionProtocolIsisSrv6MyLocalSidModify().msg()
	}

	if value == ActionProtocolIsisSrv6MyLocalSidChoice.DELETE {
		obj.obj.Delete = NewActionProtocolIsisSrv6MyLocalSidDelete().msg()
	}

	return obj
}

// description is TBD
// Add returns a ActionProtocolIsisSrv6MyLocalSidAdd
func (obj *actionProtocolIsisSrv6MyLocalSid) Add() ActionProtocolIsisSrv6MyLocalSidAdd {
	if obj.obj.Add == nil {
		obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.ADD)
	}
	if obj.addHolder == nil {
		obj.addHolder = &actionProtocolIsisSrv6MyLocalSidAdd{obj: obj.obj.Add}
	}
	return obj.addHolder
}

// description is TBD
// Add returns a ActionProtocolIsisSrv6MyLocalSidAdd
func (obj *actionProtocolIsisSrv6MyLocalSid) HasAdd() bool {
	return obj.obj.Add != nil
}

// description is TBD
// SetAdd sets the ActionProtocolIsisSrv6MyLocalSidAdd value in the ActionProtocolIsisSrv6MyLocalSid object
func (obj *actionProtocolIsisSrv6MyLocalSid) SetAdd(value ActionProtocolIsisSrv6MyLocalSidAdd) ActionProtocolIsisSrv6MyLocalSid {
	obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.ADD)
	obj.addHolder = nil
	obj.obj.Add = value.msg()

	return obj
}

// description is TBD
// Modify returns a ActionProtocolIsisSrv6MyLocalSidModify
func (obj *actionProtocolIsisSrv6MyLocalSid) Modify() ActionProtocolIsisSrv6MyLocalSidModify {
	if obj.obj.Modify == nil {
		obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.MODIFY)
	}
	if obj.modifyHolder == nil {
		obj.modifyHolder = &actionProtocolIsisSrv6MyLocalSidModify{obj: obj.obj.Modify}
	}
	return obj.modifyHolder
}

// description is TBD
// Modify returns a ActionProtocolIsisSrv6MyLocalSidModify
func (obj *actionProtocolIsisSrv6MyLocalSid) HasModify() bool {
	return obj.obj.Modify != nil
}

// description is TBD
// SetModify sets the ActionProtocolIsisSrv6MyLocalSidModify value in the ActionProtocolIsisSrv6MyLocalSid object
func (obj *actionProtocolIsisSrv6MyLocalSid) SetModify(value ActionProtocolIsisSrv6MyLocalSidModify) ActionProtocolIsisSrv6MyLocalSid {
	obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.MODIFY)
	obj.modifyHolder = nil
	obj.obj.Modify = value.msg()

	return obj
}

// description is TBD
// Delete returns a ActionProtocolIsisSrv6MyLocalSidDelete
func (obj *actionProtocolIsisSrv6MyLocalSid) Delete() ActionProtocolIsisSrv6MyLocalSidDelete {
	if obj.obj.Delete == nil {
		obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.DELETE)
	}
	if obj.deleteHolder == nil {
		obj.deleteHolder = &actionProtocolIsisSrv6MyLocalSidDelete{obj: obj.obj.Delete}
	}
	return obj.deleteHolder
}

// description is TBD
// Delete returns a ActionProtocolIsisSrv6MyLocalSidDelete
func (obj *actionProtocolIsisSrv6MyLocalSid) HasDelete() bool {
	return obj.obj.Delete != nil
}

// description is TBD
// SetDelete sets the ActionProtocolIsisSrv6MyLocalSidDelete value in the ActionProtocolIsisSrv6MyLocalSid object
func (obj *actionProtocolIsisSrv6MyLocalSid) SetDelete(value ActionProtocolIsisSrv6MyLocalSidDelete) ActionProtocolIsisSrv6MyLocalSid {
	obj.setChoice(ActionProtocolIsisSrv6MyLocalSidChoice.DELETE)
	obj.deleteHolder = nil
	obj.obj.Delete = value.msg()

	return obj
}

func (obj *actionProtocolIsisSrv6MyLocalSid) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolIsisSrv6MyLocalSid")
	}

	if obj.obj.Add != nil {

		obj.Add().validateObj(vObj, set_default)
	}

	if obj.obj.Modify != nil {

		obj.Modify().validateObj(vObj, set_default)
	}

	if obj.obj.Delete != nil {

		obj.Delete().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIsisSrv6MyLocalSid) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisSrv6MyLocalSidChoiceEnum

	if obj.obj.Add != nil {
		choices_set += 1
		choice = ActionProtocolIsisSrv6MyLocalSidChoice.ADD
	}

	if obj.obj.Modify != nil {
		choices_set += 1
		choice = ActionProtocolIsisSrv6MyLocalSidChoice.MODIFY
	}

	if obj.obj.Delete != nil {
		choices_set += 1
		choice = ActionProtocolIsisSrv6MyLocalSidChoice.DELETE
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIsisSrv6MyLocalSid")
			}
		} else {
			intVal := otg.ActionProtocolIsisSrv6MyLocalSid_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIsisSrv6MyLocalSid_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
