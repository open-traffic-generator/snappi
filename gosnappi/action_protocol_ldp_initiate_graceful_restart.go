package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolLdpInitiateGracefulRestart *****
type actionProtocolLdpInitiateGracefulRestart struct {
	validation
	obj          *otg.ActionProtocolLdpInitiateGracefulRestart
	marshaller   marshalActionProtocolLdpInitiateGracefulRestart
	unMarshaller unMarshalActionProtocolLdpInitiateGracefulRestart
}

func NewActionProtocolLdpInitiateGracefulRestart() ActionProtocolLdpInitiateGracefulRestart {
	obj := actionProtocolLdpInitiateGracefulRestart{obj: &otg.ActionProtocolLdpInitiateGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolLdpInitiateGracefulRestart) msg() *otg.ActionProtocolLdpInitiateGracefulRestart {
	return obj.obj
}

func (obj *actionProtocolLdpInitiateGracefulRestart) setMsg(msg *otg.ActionProtocolLdpInitiateGracefulRestart) ActionProtocolLdpInitiateGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolLdpInitiateGracefulRestart struct {
	obj *actionProtocolLdpInitiateGracefulRestart
}

type marshalActionProtocolLdpInitiateGracefulRestart interface {
	// ToProto marshals ActionProtocolLdpInitiateGracefulRestart to protobuf object *otg.ActionProtocolLdpInitiateGracefulRestart
	ToProto() (*otg.ActionProtocolLdpInitiateGracefulRestart, error)
	// ToPbText marshals ActionProtocolLdpInitiateGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolLdpInitiateGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolLdpInitiateGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolLdpInitiateGracefulRestart struct {
	obj *actionProtocolLdpInitiateGracefulRestart
}

type unMarshalActionProtocolLdpInitiateGracefulRestart interface {
	// FromProto unmarshals ActionProtocolLdpInitiateGracefulRestart from protobuf object *otg.ActionProtocolLdpInitiateGracefulRestart
	FromProto(msg *otg.ActionProtocolLdpInitiateGracefulRestart) (ActionProtocolLdpInitiateGracefulRestart, error)
	// FromPbText unmarshals ActionProtocolLdpInitiateGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolLdpInitiateGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolLdpInitiateGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolLdpInitiateGracefulRestart) Marshal() marshalActionProtocolLdpInitiateGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolLdpInitiateGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolLdpInitiateGracefulRestart) Unmarshal() unMarshalActionProtocolLdpInitiateGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolLdpInitiateGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolLdpInitiateGracefulRestart) ToProto() (*otg.ActionProtocolLdpInitiateGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolLdpInitiateGracefulRestart) FromProto(msg *otg.ActionProtocolLdpInitiateGracefulRestart) (ActionProtocolLdpInitiateGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolLdpInitiateGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolLdpInitiateGracefulRestart) FromPbText(value string) error {
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

func (m *marshalactionProtocolLdpInitiateGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolLdpInitiateGracefulRestart) FromYaml(value string) error {
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

func (m *marshalactionProtocolLdpInitiateGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolLdpInitiateGracefulRestart) FromJson(value string) error {
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

func (obj *actionProtocolLdpInitiateGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolLdpInitiateGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolLdpInitiateGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolLdpInitiateGracefulRestart) Clone() (ActionProtocolLdpInitiateGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolLdpInitiateGracefulRestart()
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

// ActionProtocolLdpInitiateGracefulRestart is initiates an LDP graceful restart (RFC 3478) on the selected LDP routers. Each selected router closes its LDP sessions without withdrawing its labels, and re-establishes them after restart_delay seconds. The Initialization messages sent after the restart carry the FT Session TLV with the configured reconnect_time and recovery_time (RFC 3478 Section 2). A selected router must have graceful_restart configured; otherwise the implementation returns an error. If no name is specified, all configured LDP routers are restarted.
type ActionProtocolLdpInitiateGracefulRestart interface {
	Validation
	// msg marshals ActionProtocolLdpInitiateGracefulRestart to protobuf object *otg.ActionProtocolLdpInitiateGracefulRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolLdpInitiateGracefulRestart
	// setMsg unmarshals ActionProtocolLdpInitiateGracefulRestart from protobuf object *otg.ActionProtocolLdpInitiateGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolLdpInitiateGracefulRestart) ActionProtocolLdpInitiateGracefulRestart
	// provides marshal interface
	Marshal() marshalActionProtocolLdpInitiateGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolLdpInitiateGracefulRestart
	// validate validates ActionProtocolLdpInitiateGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolLdpInitiateGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RouterNames returns []string, set in ActionProtocolLdpInitiateGracefulRestart.
	RouterNames() []string
	// SetRouterNames assigns []string provided by user to ActionProtocolLdpInitiateGracefulRestart
	SetRouterNames(value []string) ActionProtocolLdpInitiateGracefulRestart
	// RestartDelay returns uint32, set in ActionProtocolLdpInitiateGracefulRestart.
	RestartDelay() uint32
	// SetRestartDelay assigns uint32 provided by user to ActionProtocolLdpInitiateGracefulRestart
	SetRestartDelay(value uint32) ActionProtocolLdpInitiateGracefulRestart
	// HasRestartDelay checks if RestartDelay has been set in ActionProtocolLdpInitiateGracefulRestart
	HasRestartDelay() bool
}

// The names of LDP routers to restart.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// RouterNames returns a []string
func (obj *actionProtocolLdpInitiateGracefulRestart) RouterNames() []string {
	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	return obj.obj.RouterNames
}

// The names of LDP routers to restart.
//
// x-constraint:
// - /components/schemas/Device.LdpRouter/properties/name
//
// SetRouterNames sets the []string value in the ActionProtocolLdpInitiateGracefulRestart object
func (obj *actionProtocolLdpInitiateGracefulRestart) SetRouterNames(value []string) ActionProtocolLdpInitiateGracefulRestart {

	if obj.obj.RouterNames == nil {
		obj.obj.RouterNames = make([]string, 0)
	}
	obj.obj.RouterNames = value

	return obj
}

// The time in seconds after which the selected LDP routers re-establish their sessions.
// RestartDelay returns a uint32
func (obj *actionProtocolLdpInitiateGracefulRestart) RestartDelay() uint32 {

	return *obj.obj.RestartDelay

}

// The time in seconds after which the selected LDP routers re-establish their sessions.
// RestartDelay returns a uint32
func (obj *actionProtocolLdpInitiateGracefulRestart) HasRestartDelay() bool {
	return obj.obj.RestartDelay != nil
}

// The time in seconds after which the selected LDP routers re-establish their sessions.
// SetRestartDelay sets the uint32 value in the ActionProtocolLdpInitiateGracefulRestart object
func (obj *actionProtocolLdpInitiateGracefulRestart) SetRestartDelay(value uint32) ActionProtocolLdpInitiateGracefulRestart {

	obj.obj.RestartDelay = &value
	return obj
}

func (obj *actionProtocolLdpInitiateGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartDelay != nil {

		if *obj.obj.RestartDelay > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ActionProtocolLdpInitiateGracefulRestart.RestartDelay <= 3600 but Got %d", *obj.obj.RestartDelay))
		}

	}

}

func (obj *actionProtocolLdpInitiateGracefulRestart) setDefault() {
	if obj.obj.RestartDelay == nil {
		obj.SetRestartDelay(30)
	}

}
