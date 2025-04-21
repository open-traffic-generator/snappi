package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisInitiateGracefulRestart *****
type actionProtocolIsisInitiateGracefulRestart struct {
	validation
	obj          *otg.ActionProtocolIsisInitiateGracefulRestart
	marshaller   marshalActionProtocolIsisInitiateGracefulRestart
	unMarshaller unMarshalActionProtocolIsisInitiateGracefulRestart
}

func NewActionProtocolIsisInitiateGracefulRestart() ActionProtocolIsisInitiateGracefulRestart {
	obj := actionProtocolIsisInitiateGracefulRestart{obj: &otg.ActionProtocolIsisInitiateGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisInitiateGracefulRestart) msg() *otg.ActionProtocolIsisInitiateGracefulRestart {
	return obj.obj
}

func (obj *actionProtocolIsisInitiateGracefulRestart) setMsg(msg *otg.ActionProtocolIsisInitiateGracefulRestart) ActionProtocolIsisInitiateGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisInitiateGracefulRestart struct {
	obj *actionProtocolIsisInitiateGracefulRestart
}

type marshalActionProtocolIsisInitiateGracefulRestart interface {
	// ToProto marshals ActionProtocolIsisInitiateGracefulRestart to protobuf object *otg.ActionProtocolIsisInitiateGracefulRestart
	ToProto() (*otg.ActionProtocolIsisInitiateGracefulRestart, error)
	// ToPbText marshals ActionProtocolIsisInitiateGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisInitiateGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisInitiateGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisInitiateGracefulRestart struct {
	obj *actionProtocolIsisInitiateGracefulRestart
}

type unMarshalActionProtocolIsisInitiateGracefulRestart interface {
	// FromProto unmarshals ActionProtocolIsisInitiateGracefulRestart from protobuf object *otg.ActionProtocolIsisInitiateGracefulRestart
	FromProto(msg *otg.ActionProtocolIsisInitiateGracefulRestart) (ActionProtocolIsisInitiateGracefulRestart, error)
	// FromPbText unmarshals ActionProtocolIsisInitiateGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisInitiateGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisInitiateGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisInitiateGracefulRestart) Marshal() marshalActionProtocolIsisInitiateGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisInitiateGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisInitiateGracefulRestart) Unmarshal() unMarshalActionProtocolIsisInitiateGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisInitiateGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisInitiateGracefulRestart) ToProto() (*otg.ActionProtocolIsisInitiateGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisInitiateGracefulRestart) FromProto(msg *otg.ActionProtocolIsisInitiateGracefulRestart) (ActionProtocolIsisInitiateGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisInitiateGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulRestart) FromPbText(value string) error {
	retObj := proto.Unmarshal([]byte(value), m.obj.msg())
	if retObj != nil {
		return retObj
	}

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return retObj
}

func (m *marshalactionProtocolIsisInitiateGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulRestart) FromYaml(value string) error {
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

	vErr := m.obj.validateToAndFrom()
	if vErr != nil {
		return vErr
	}
	return nil
}

func (m *marshalactionProtocolIsisInitiateGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulRestart) FromJson(value string) error {
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

	err := m.obj.validateToAndFrom()
	if err != nil {
		return err
	}
	return nil
}

func (obj *actionProtocolIsisInitiateGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisInitiateGracefulRestart) Clone() (ActionProtocolIsisInitiateGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisInitiateGracefulRestart()
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

// ActionProtocolIsisInitiateGracefulRestart is initiates IS-IS Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. To emulate scenarios where a router sends Graceful Start TlV this will result in  Graceful Restart scenario to be triggered as perReference: https://datatracker.ietf.org/doc/html/rfc5306.
type ActionProtocolIsisInitiateGracefulRestart interface {
	Validation
	// msg marshals ActionProtocolIsisInitiateGracefulRestart to protobuf object *otg.ActionProtocolIsisInitiateGracefulRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisInitiateGracefulRestart
	// setMsg unmarshals ActionProtocolIsisInitiateGracefulRestart from protobuf object *otg.ActionProtocolIsisInitiateGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisInitiateGracefulRestart) ActionProtocolIsisInitiateGracefulRestart
	// provides marshal interface
	Marshal() marshalActionProtocolIsisInitiateGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisInitiateGracefulRestart
	// validate validates ActionProtocolIsisInitiateGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisInitiateGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolIsisInitiateGracefulRestart.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolIsisInitiateGracefulRestart
	SetRouterNames(value []string) ActionProtocolIsisInitiateGracefulRestart
}

// The names of device rfc5306 objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolIsisInitiateGracefulRestart) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of device rfc5306 objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// SetRouterNames sets the []string value in the ActionProtocolIsisInitiateGracefulRestart object
func (obj *actionProtocolIsisInitiateGracefulRestart) SetRouterNames(value []string) ActionProtocolIsisInitiateGracefulRestart {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

func (obj *actionProtocolIsisInitiateGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *actionProtocolIsisInitiateGracefulRestart) setDefault() {

}
