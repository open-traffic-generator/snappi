package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisInitiateGracefulStart *****
type actionProtocolIsisInitiateGracefulStart struct {
	validation
	obj          *otg.ActionProtocolIsisInitiateGracefulStart
	marshaller   marshalActionProtocolIsisInitiateGracefulStart
	unMarshaller unMarshalActionProtocolIsisInitiateGracefulStart
}

func NewActionProtocolIsisInitiateGracefulStart() ActionProtocolIsisInitiateGracefulStart {
	obj := actionProtocolIsisInitiateGracefulStart{obj: &otg.ActionProtocolIsisInitiateGracefulStart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisInitiateGracefulStart) msg() *otg.ActionProtocolIsisInitiateGracefulStart {
	return obj.obj
}

func (obj *actionProtocolIsisInitiateGracefulStart) setMsg(msg *otg.ActionProtocolIsisInitiateGracefulStart) ActionProtocolIsisInitiateGracefulStart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisInitiateGracefulStart struct {
	obj *actionProtocolIsisInitiateGracefulStart
}

type marshalActionProtocolIsisInitiateGracefulStart interface {
	// ToProto marshals ActionProtocolIsisInitiateGracefulStart to protobuf object *otg.ActionProtocolIsisInitiateGracefulStart
	ToProto() (*otg.ActionProtocolIsisInitiateGracefulStart, error)
	// ToPbText marshals ActionProtocolIsisInitiateGracefulStart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisInitiateGracefulStart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisInitiateGracefulStart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisInitiateGracefulStart struct {
	obj *actionProtocolIsisInitiateGracefulStart
}

type unMarshalActionProtocolIsisInitiateGracefulStart interface {
	// FromProto unmarshals ActionProtocolIsisInitiateGracefulStart from protobuf object *otg.ActionProtocolIsisInitiateGracefulStart
	FromProto(msg *otg.ActionProtocolIsisInitiateGracefulStart) (ActionProtocolIsisInitiateGracefulStart, error)
	// FromPbText unmarshals ActionProtocolIsisInitiateGracefulStart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisInitiateGracefulStart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisInitiateGracefulStart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisInitiateGracefulStart) Marshal() marshalActionProtocolIsisInitiateGracefulStart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisInitiateGracefulStart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisInitiateGracefulStart) Unmarshal() unMarshalActionProtocolIsisInitiateGracefulStart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisInitiateGracefulStart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisInitiateGracefulStart) ToProto() (*otg.ActionProtocolIsisInitiateGracefulStart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisInitiateGracefulStart) FromProto(msg *otg.ActionProtocolIsisInitiateGracefulStart) (ActionProtocolIsisInitiateGracefulStart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisInitiateGracefulStart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulStart) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisInitiateGracefulStart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulStart) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisInitiateGracefulStart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisInitiateGracefulStart) FromJson(value string) error {
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

func (obj *actionProtocolIsisInitiateGracefulStart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateGracefulStart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisInitiateGracefulStart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisInitiateGracefulStart) Clone() (ActionProtocolIsisInitiateGracefulStart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisInitiateGracefulStart()
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

// ActionProtocolIsisInitiateGracefulStart is initiates IS-IS Graceful Start process for the selected IS-IS routers. If no name is specified then Graceful Start will be sent to all configured IS-IS routers. To emulate scenarios where a router sends Graceful Restart TlV this will result in  Graceful tart scenario to be triggered as per Reference: https://datatracker.ietf.org/doc/html/rfc5306.
type ActionProtocolIsisInitiateGracefulStart interface {
	Validation
	// msg marshals ActionProtocolIsisInitiateGracefulStart to protobuf object *otg.ActionProtocolIsisInitiateGracefulStart
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisInitiateGracefulStart
	// setMsg unmarshals ActionProtocolIsisInitiateGracefulStart from protobuf object *otg.ActionProtocolIsisInitiateGracefulStart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisInitiateGracefulStart) ActionProtocolIsisInitiateGracefulStart
	// provides marshal interface
	Marshal() marshalActionProtocolIsisInitiateGracefulStart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisInitiateGracefulStart
	// validate validates ActionProtocolIsisInitiateGracefulStart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisInitiateGracefulStart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolIsisInitiateGracefulStart.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolIsisInitiateGracefulStart
	SetRouterNames(value []string) ActionProtocolIsisInitiateGracefulStart
}

// The names of device rfc5306 objects to control.
//
// x-constraint:
// - /components/schemas/Device.IsisRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolIsisInitiateGracefulStart) RouterNames() []string {
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
// SetRouterNames sets the []string value in the ActionProtocolIsisInitiateGracefulStart object
func (obj *actionProtocolIsisInitiateGracefulStart) SetRouterNames(value []string) ActionProtocolIsisInitiateGracefulStart {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

func (obj *actionProtocolIsisInitiateGracefulStart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *actionProtocolIsisInitiateGracefulStart) setDefault() {

}
