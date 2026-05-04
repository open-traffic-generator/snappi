package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisSrv6 *****
type actionProtocolIsisSrv6 struct {
	validation
	obj              *otg.ActionProtocolIsisSrv6
	marshaller       marshalActionProtocolIsisSrv6
	unMarshaller     unMarshalActionProtocolIsisSrv6
	myLocalSidHolder ActionProtocolIsisSrv6MyLocalSid
}

func NewActionProtocolIsisSrv6() ActionProtocolIsisSrv6 {
	obj := actionProtocolIsisSrv6{obj: &otg.ActionProtocolIsisSrv6{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisSrv6) msg() *otg.ActionProtocolIsisSrv6 {
	return obj.obj
}

func (obj *actionProtocolIsisSrv6) setMsg(msg *otg.ActionProtocolIsisSrv6) ActionProtocolIsisSrv6 {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisSrv6 struct {
	obj *actionProtocolIsisSrv6
}

type marshalActionProtocolIsisSrv6 interface {
	// ToProto marshals ActionProtocolIsisSrv6 to protobuf object *otg.ActionProtocolIsisSrv6
	ToProto() (*otg.ActionProtocolIsisSrv6, error)
	// ToPbText marshals ActionProtocolIsisSrv6 to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisSrv6 to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisSrv6 to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisSrv6 struct {
	obj *actionProtocolIsisSrv6
}

type unMarshalActionProtocolIsisSrv6 interface {
	// FromProto unmarshals ActionProtocolIsisSrv6 from protobuf object *otg.ActionProtocolIsisSrv6
	FromProto(msg *otg.ActionProtocolIsisSrv6) (ActionProtocolIsisSrv6, error)
	// FromPbText unmarshals ActionProtocolIsisSrv6 from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisSrv6 from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisSrv6 from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisSrv6) Marshal() marshalActionProtocolIsisSrv6 {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisSrv6{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisSrv6) Unmarshal() unMarshalActionProtocolIsisSrv6 {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisSrv6{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisSrv6) ToProto() (*otg.ActionProtocolIsisSrv6, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisSrv6) FromProto(msg *otg.ActionProtocolIsisSrv6) (ActionProtocolIsisSrv6, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisSrv6) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisSrv6) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisSrv6) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisSrv6) FromJson(value string) error {
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

func (obj *actionProtocolIsisSrv6) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisSrv6) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisSrv6) Clone() (ActionProtocolIsisSrv6, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisSrv6()
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

func (obj *actionProtocolIsisSrv6) setNil() {
	obj.myLocalSidHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// ActionProtocolIsisSrv6 is sRv6 control actions for IS-IS emulated routers.
type ActionProtocolIsisSrv6 interface {
	Validation
	// msg marshals ActionProtocolIsisSrv6 to protobuf object *otg.ActionProtocolIsisSrv6
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisSrv6
	// setMsg unmarshals ActionProtocolIsisSrv6 from protobuf object *otg.ActionProtocolIsisSrv6
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisSrv6) ActionProtocolIsisSrv6
	// provides marshal interface
	Marshal() marshalActionProtocolIsisSrv6
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisSrv6
	// validate validates ActionProtocolIsisSrv6
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisSrv6, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Choice returns ActionProtocolIsisSrv6ChoiceEnum, set in ActionProtocolIsisSrv6
	Choice() ActionProtocolIsisSrv6ChoiceEnum
	// setChoice assigns ActionProtocolIsisSrv6ChoiceEnum provided by user to ActionProtocolIsisSrv6
	setChoice(value ActionProtocolIsisSrv6ChoiceEnum) ActionProtocolIsisSrv6
	// MyLocalSid returns ActionProtocolIsisSrv6MyLocalSid, set in ActionProtocolIsisSrv6.
	// ActionProtocolIsisSrv6MyLocalSid is trigger a My Local SID lifecycle operation (add, modify, or delete) on the specified IS-IS routers (RFC 8986 Section 4 Instantiation, Re-instantiation, and Un-instantiation). In IxNetwork, this operation manipulates isisSRv6EndSIDList entries under the IS-IS router and triggers IS-IS LSP re-advertisement.
	MyLocalSid() ActionProtocolIsisSrv6MyLocalSid
	// SetMyLocalSid assigns ActionProtocolIsisSrv6MyLocalSid provided by user to ActionProtocolIsisSrv6.
	// ActionProtocolIsisSrv6MyLocalSid is trigger a My Local SID lifecycle operation (add, modify, or delete) on the specified IS-IS routers (RFC 8986 Section 4 Instantiation, Re-instantiation, and Un-instantiation). In IxNetwork, this operation manipulates isisSRv6EndSIDList entries under the IS-IS router and triggers IS-IS LSP re-advertisement.
	SetMyLocalSid(value ActionProtocolIsisSrv6MyLocalSid) ActionProtocolIsisSrv6
	// HasMyLocalSid checks if MyLocalSid has been set in ActionProtocolIsisSrv6
	HasMyLocalSid() bool
	setNil()
}

type ActionProtocolIsisSrv6ChoiceEnum string

// Enum of Choice on ActionProtocolIsisSrv6
var ActionProtocolIsisSrv6Choice = struct {
	MY_LOCAL_SID ActionProtocolIsisSrv6ChoiceEnum
}{
	MY_LOCAL_SID: ActionProtocolIsisSrv6ChoiceEnum("my_local_sid"),
}

func (obj *actionProtocolIsisSrv6) Choice() ActionProtocolIsisSrv6ChoiceEnum {
	return ActionProtocolIsisSrv6ChoiceEnum(obj.obj.Choice.Enum().String())
}

func (obj *actionProtocolIsisSrv6) setChoice(value ActionProtocolIsisSrv6ChoiceEnum) ActionProtocolIsisSrv6 {
	intValue, ok := otg.ActionProtocolIsisSrv6_Choice_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on ActionProtocolIsisSrv6ChoiceEnum", string(value)))
		return obj
	}
	enumValue := otg.ActionProtocolIsisSrv6_Choice_Enum(intValue)
	obj.obj.Choice = &enumValue
	obj.obj.MyLocalSid = nil
	obj.myLocalSidHolder = nil

	if value == ActionProtocolIsisSrv6Choice.MY_LOCAL_SID {
		obj.obj.MyLocalSid = NewActionProtocolIsisSrv6MyLocalSid().msg()
	}

	return obj
}

// description is TBD
// MyLocalSid returns a ActionProtocolIsisSrv6MyLocalSid
func (obj *actionProtocolIsisSrv6) MyLocalSid() ActionProtocolIsisSrv6MyLocalSid {
	if obj.obj.MyLocalSid == nil {
		obj.setChoice(ActionProtocolIsisSrv6Choice.MY_LOCAL_SID)
	}
	if obj.myLocalSidHolder == nil {
		obj.myLocalSidHolder = &actionProtocolIsisSrv6MyLocalSid{obj: obj.obj.MyLocalSid}
	}
	return obj.myLocalSidHolder
}

// description is TBD
// MyLocalSid returns a ActionProtocolIsisSrv6MyLocalSid
func (obj *actionProtocolIsisSrv6) HasMyLocalSid() bool {
	return obj.obj.MyLocalSid != nil
}

// description is TBD
// SetMyLocalSid sets the ActionProtocolIsisSrv6MyLocalSid value in the ActionProtocolIsisSrv6 object
func (obj *actionProtocolIsisSrv6) SetMyLocalSid(value ActionProtocolIsisSrv6MyLocalSid) ActionProtocolIsisSrv6 {
	obj.setChoice(ActionProtocolIsisSrv6Choice.MY_LOCAL_SID)
	obj.myLocalSidHolder = nil
	obj.obj.MyLocalSid = value.msg()

	return obj
}

func (obj *actionProtocolIsisSrv6) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	// Choice is required
	if obj.obj.Choice == nil {
		vObj.validationErrors = append(vObj.validationErrors, "Choice is required field on interface ActionProtocolIsisSrv6")
	}

	if obj.obj.MyLocalSid != nil {

		obj.MyLocalSid().validateObj(vObj, set_default)
	}

}

func (obj *actionProtocolIsisSrv6) setDefault() {
	var choices_set int = 0
	var choice ActionProtocolIsisSrv6ChoiceEnum

	if obj.obj.MyLocalSid != nil {
		choices_set += 1
		choice = ActionProtocolIsisSrv6Choice.MY_LOCAL_SID
	}
	if choices_set == 1 && choice != "" {
		if obj.obj.Choice != nil {
			if obj.Choice() != choice {
				obj.validationErrors = append(obj.validationErrors, "choice not matching with property in ActionProtocolIsisSrv6")
			}
		} else {
			intVal := otg.ActionProtocolIsisSrv6_Choice_Enum_value[string(choice)]
			enumValue := otg.ActionProtocolIsisSrv6_Choice_Enum(intVal)
			obj.obj.Choice = &enumValue
		}
	}

}
