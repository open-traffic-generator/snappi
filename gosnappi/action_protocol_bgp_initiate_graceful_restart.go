package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolBgpInitiateGracefulRestart *****
type actionProtocolBgpInitiateGracefulRestart struct {
	validation
	obj          *otg.ActionProtocolBgpInitiateGracefulRestart
	marshaller   marshalActionProtocolBgpInitiateGracefulRestart
	unMarshaller unMarshalActionProtocolBgpInitiateGracefulRestart
}

func NewActionProtocolBgpInitiateGracefulRestart() ActionProtocolBgpInitiateGracefulRestart {
	obj := actionProtocolBgpInitiateGracefulRestart{obj: &otg.ActionProtocolBgpInitiateGracefulRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolBgpInitiateGracefulRestart) msg() *otg.ActionProtocolBgpInitiateGracefulRestart {
	return obj.obj
}

func (obj *actionProtocolBgpInitiateGracefulRestart) setMsg(msg *otg.ActionProtocolBgpInitiateGracefulRestart) ActionProtocolBgpInitiateGracefulRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolBgpInitiateGracefulRestart struct {
	obj *actionProtocolBgpInitiateGracefulRestart
}

type marshalActionProtocolBgpInitiateGracefulRestart interface {
	// ToProto marshals ActionProtocolBgpInitiateGracefulRestart to protobuf object *otg.ActionProtocolBgpInitiateGracefulRestart
	ToProto() (*otg.ActionProtocolBgpInitiateGracefulRestart, error)
	// ToPbText marshals ActionProtocolBgpInitiateGracefulRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolBgpInitiateGracefulRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolBgpInitiateGracefulRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolBgpInitiateGracefulRestart struct {
	obj *actionProtocolBgpInitiateGracefulRestart
}

type unMarshalActionProtocolBgpInitiateGracefulRestart interface {
	// FromProto unmarshals ActionProtocolBgpInitiateGracefulRestart from protobuf object *otg.ActionProtocolBgpInitiateGracefulRestart
	FromProto(msg *otg.ActionProtocolBgpInitiateGracefulRestart) (ActionProtocolBgpInitiateGracefulRestart, error)
	// FromPbText unmarshals ActionProtocolBgpInitiateGracefulRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolBgpInitiateGracefulRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolBgpInitiateGracefulRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolBgpInitiateGracefulRestart) Marshal() marshalActionProtocolBgpInitiateGracefulRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolBgpInitiateGracefulRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolBgpInitiateGracefulRestart) Unmarshal() unMarshalActionProtocolBgpInitiateGracefulRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolBgpInitiateGracefulRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolBgpInitiateGracefulRestart) ToProto() (*otg.ActionProtocolBgpInitiateGracefulRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolBgpInitiateGracefulRestart) FromProto(msg *otg.ActionProtocolBgpInitiateGracefulRestart) (ActionProtocolBgpInitiateGracefulRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolBgpInitiateGracefulRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolBgpInitiateGracefulRestart) FromPbText(value string) error {
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

func (m *marshalactionProtocolBgpInitiateGracefulRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolBgpInitiateGracefulRestart) FromYaml(value string) error {
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

func (m *marshalactionProtocolBgpInitiateGracefulRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolBgpInitiateGracefulRestart) FromJson(value string) error {
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

func (obj *actionProtocolBgpInitiateGracefulRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolBgpInitiateGracefulRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolBgpInitiateGracefulRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolBgpInitiateGracefulRestart) Clone() (ActionProtocolBgpInitiateGracefulRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolBgpInitiateGracefulRestart()
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

// ActionProtocolBgpInitiateGracefulRestart is initiates BGP Graceful Restart process for the selected BGP peers. If no name is specified then Graceful Restart will be sent to all configured BGP peers.
type ActionProtocolBgpInitiateGracefulRestart interface {
	Validation
	// msg marshals ActionProtocolBgpInitiateGracefulRestart to protobuf object *otg.ActionProtocolBgpInitiateGracefulRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolBgpInitiateGracefulRestart
	// setMsg unmarshals ActionProtocolBgpInitiateGracefulRestart from protobuf object *otg.ActionProtocolBgpInitiateGracefulRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolBgpInitiateGracefulRestart) ActionProtocolBgpInitiateGracefulRestart
	// provides marshal interface
	Marshal() marshalActionProtocolBgpInitiateGracefulRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolBgpInitiateGracefulRestart
	// validate validates ActionProtocolBgpInitiateGracefulRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolBgpInitiateGracefulRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// PeerNames returns []string, set in ActionProtocolBgpInitiateGracefulRestart.
	PeerNames() []string
	// SetPeerNames assigns []string provided by user to ActionProtocolBgpInitiateGracefulRestart
	SetPeerNames(value []string) ActionProtocolBgpInitiateGracefulRestart
	// RestartDelay returns uint32, set in ActionProtocolBgpInitiateGracefulRestart.
	RestartDelay() uint32
	// SetRestartDelay assigns uint32 provided by user to ActionProtocolBgpInitiateGracefulRestart
	SetRestartDelay(value uint32) ActionProtocolBgpInitiateGracefulRestart
	// HasRestartDelay checks if RestartDelay has been set in ActionProtocolBgpInitiateGracefulRestart
	HasRestartDelay() bool
}

// The names of device BGP peers objects to control.
//
// x-constraint:
// - /components/schemas/Device.Bgp/properties/name
//
// PeerNames returns a []string
func (obj *actionProtocolBgpInitiateGracefulRestart) PeerNames() []string {
	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	return obj.obj.PeerNames
}

// The names of device BGP peers objects to control.
//
// x-constraint:
// - /components/schemas/Device.Bgp/properties/name
//
// SetPeerNames sets the []string value in the ActionProtocolBgpInitiateGracefulRestart object
func (obj *actionProtocolBgpInitiateGracefulRestart) SetPeerNames(value []string) ActionProtocolBgpInitiateGracefulRestart {

	if obj.obj.PeerNames == nil {
		obj.obj.PeerNames = make([]string, 0)
	}
	obj.obj.PeerNames = value

	return obj
}

// Duration (in seconds) after which selected BGP peers will initiate
// Graceful restart by sending the Open Message with Restart State bit set in the Graceful Restart capability.
// RestartDelay returns a uint32
func (obj *actionProtocolBgpInitiateGracefulRestart) RestartDelay() uint32 {

	return *obj.obj.RestartDelay

}

// Duration (in seconds) after which selected BGP peers will initiate
// Graceful restart by sending the Open Message with Restart State bit set in the Graceful Restart capability.
// RestartDelay returns a uint32
func (obj *actionProtocolBgpInitiateGracefulRestart) HasRestartDelay() bool {
	return obj.obj.RestartDelay != nil
}

// Duration (in seconds) after which selected BGP peers will initiate
// Graceful restart by sending the Open Message with Restart State bit set in the Graceful Restart capability.
// SetRestartDelay sets the uint32 value in the ActionProtocolBgpInitiateGracefulRestart object
func (obj *actionProtocolBgpInitiateGracefulRestart) SetRestartDelay(value uint32) ActionProtocolBgpInitiateGracefulRestart {

	obj.obj.RestartDelay = &value
	return obj
}

func (obj *actionProtocolBgpInitiateGracefulRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartDelay != nil {

		if *obj.obj.RestartDelay > 3600 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ActionProtocolBgpInitiateGracefulRestart.RestartDelay <= 3600 but Got %d", *obj.obj.RestartDelay))
		}

	}

}

func (obj *actionProtocolBgpInitiateGracefulRestart) setDefault() {
	if obj.obj.RestartDelay == nil {
		obj.SetRestartDelay(30)
	}

}
