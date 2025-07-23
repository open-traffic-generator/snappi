package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisPlannedRestart *****
type actionProtocolIsisPlannedRestart struct {
	validation
	obj          *otg.ActionProtocolIsisPlannedRestart
	marshaller   marshalActionProtocolIsisPlannedRestart
	unMarshaller unMarshalActionProtocolIsisPlannedRestart
}

func NewActionProtocolIsisPlannedRestart() ActionProtocolIsisPlannedRestart {
	obj := actionProtocolIsisPlannedRestart{obj: &otg.ActionProtocolIsisPlannedRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisPlannedRestart) msg() *otg.ActionProtocolIsisPlannedRestart {
	return obj.obj
}

func (obj *actionProtocolIsisPlannedRestart) setMsg(msg *otg.ActionProtocolIsisPlannedRestart) ActionProtocolIsisPlannedRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisPlannedRestart struct {
	obj *actionProtocolIsisPlannedRestart
}

type marshalActionProtocolIsisPlannedRestart interface {
	// ToProto marshals ActionProtocolIsisPlannedRestart to protobuf object *otg.ActionProtocolIsisPlannedRestart
	ToProto() (*otg.ActionProtocolIsisPlannedRestart, error)
	// ToPbText marshals ActionProtocolIsisPlannedRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisPlannedRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisPlannedRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisPlannedRestart struct {
	obj *actionProtocolIsisPlannedRestart
}

type unMarshalActionProtocolIsisPlannedRestart interface {
	// FromProto unmarshals ActionProtocolIsisPlannedRestart from protobuf object *otg.ActionProtocolIsisPlannedRestart
	FromProto(msg *otg.ActionProtocolIsisPlannedRestart) (ActionProtocolIsisPlannedRestart, error)
	// FromPbText unmarshals ActionProtocolIsisPlannedRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisPlannedRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisPlannedRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisPlannedRestart) Marshal() marshalActionProtocolIsisPlannedRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisPlannedRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisPlannedRestart) Unmarshal() unMarshalActionProtocolIsisPlannedRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisPlannedRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisPlannedRestart) ToProto() (*otg.ActionProtocolIsisPlannedRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisPlannedRestart) FromProto(msg *otg.ActionProtocolIsisPlannedRestart) (ActionProtocolIsisPlannedRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisPlannedRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisPlannedRestart) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisPlannedRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisPlannedRestart) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisPlannedRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisPlannedRestart) FromJson(value string) error {
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

func (obj *actionProtocolIsisPlannedRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisPlannedRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisPlannedRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisPlannedRestart) Clone() (ActionProtocolIsisPlannedRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisPlannedRestart()
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

// ActionProtocolIsisPlannedRestart is initiates IS-IS Planned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the planned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the PR (Planned Restart Request) bit set and set the Remaining Time with the restart_time greater than the expected control-plane restart time  that is the maximum time within which this router or routers will complete the graceful restart. It waits for PA (Planned Restart Acknowledge) in an IIH PDU from Neigbhor(s).  The use of the PR bit provides a means to safely support restart periods that are significantly longer than standard Holding Times. The PR bit SHOULD remain set in IIHs until the restart is initiated. Refrence: https://datatracker.ietf.org/doc/html/rfc8706#section-3.2.3. Once the Restaring Router receives the Restart Tlv with PA bit is set, it intiates the Restart Request with the RR bit is set.  The holding_time is set as the Remaining Time as received in Restart Tlv or by remaining time of restart_time that was sent  in the Planned Restart Request Tlv. This is left by the implementation.
type ActionProtocolIsisPlannedRestart interface {
	Validation
	// msg marshals ActionProtocolIsisPlannedRestart to protobuf object *otg.ActionProtocolIsisPlannedRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisPlannedRestart
	// setMsg unmarshals ActionProtocolIsisPlannedRestart from protobuf object *otg.ActionProtocolIsisPlannedRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisPlannedRestart) ActionProtocolIsisPlannedRestart
	// provides marshal interface
	Marshal() marshalActionProtocolIsisPlannedRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisPlannedRestart
	// validate validates ActionProtocolIsisPlannedRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisPlannedRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// RestartTime returns uint32, set in ActionProtocolIsisPlannedRestart.
	RestartTime() uint32
	// SetRestartTime assigns uint32 provided by user to ActionProtocolIsisPlannedRestart
	SetRestartTime(value uint32) ActionProtocolIsisPlannedRestart
	// HasRestartTime checks if RestartTime has been set in ActionProtocolIsisPlannedRestart
	HasRestartTime() bool
	// RestartAfter returns uint32, set in ActionProtocolIsisPlannedRestart.
	RestartAfter() uint32
	// SetRestartAfter assigns uint32 provided by user to ActionProtocolIsisPlannedRestart
	SetRestartAfter(value uint32) ActionProtocolIsisPlannedRestart
	// HasRestartAfter checks if RestartAfter has been set in ActionProtocolIsisPlannedRestart
	HasRestartAfter() bool
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a Planned Restart Request. The Remaining Time is set with the restart_time in the Restarting Tlv.
// RestartTime returns a uint32
func (obj *actionProtocolIsisPlannedRestart) RestartTime() uint32 {

	return *obj.obj.RestartTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a Planned Restart Request. The Remaining Time is set with the restart_time in the Restarting Tlv.
// RestartTime returns a uint32
func (obj *actionProtocolIsisPlannedRestart) HasRestartTime() bool {
	return obj.obj.RestartTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a Planned Restart Request. The Remaining Time is set with the restart_time in the Restarting Tlv.
// SetRestartTime sets the uint32 value in the ActionProtocolIsisPlannedRestart object
func (obj *actionProtocolIsisPlannedRestart) SetRestartTime(value uint32) ActionProtocolIsisPlannedRestart {

	obj.obj.RestartTime = &value
	return obj
}

// Once it receives Restarting TLV having PA bit set in a IIH PDU from a neighbor, the time (in seconds), after which IIH PDU, having Restart Tlv with RA & RR bits unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed, and hold_timer is reset to normal.
// RestartAfter returns a uint32
func (obj *actionProtocolIsisPlannedRestart) RestartAfter() uint32 {

	return *obj.obj.RestartAfter

}

// Once it receives Restarting TLV having PA bit set in a IIH PDU from a neighbor, the time (in seconds), after which IIH PDU, having Restart Tlv with RA & RR bits unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed, and hold_timer is reset to normal.
// RestartAfter returns a uint32
func (obj *actionProtocolIsisPlannedRestart) HasRestartAfter() bool {
	return obj.obj.RestartAfter != nil
}

// Once it receives Restarting TLV having PA bit set in a IIH PDU from a neighbor, the time (in seconds), after which IIH PDU, having Restart Tlv with RA & RR bits unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed, and hold_timer is reset to normal.
// SetRestartAfter sets the uint32 value in the ActionProtocolIsisPlannedRestart object
func (obj *actionProtocolIsisPlannedRestart) SetRestartAfter(value uint32) ActionProtocolIsisPlannedRestart {

	obj.obj.RestartAfter = &value
	return obj
}

func (obj *actionProtocolIsisPlannedRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.RestartTime != nil {

		if *obj.obj.RestartTime < 1 || *obj.obj.RestartTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= ActionProtocolIsisPlannedRestart.RestartTime <= 4096 but Got %d", *obj.obj.RestartTime))
		}

	}

	if obj.obj.RestartAfter != nil {

		if *obj.obj.RestartAfter > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= ActionProtocolIsisPlannedRestart.RestartAfter <= 4096 but Got %d", *obj.obj.RestartAfter))
		}

	}

}

func (obj *actionProtocolIsisPlannedRestart) setDefault() {
	if obj.obj.RestartTime == nil {
		obj.SetRestartTime(30)
	}
	if obj.obj.RestartAfter == nil {
		obj.SetRestartAfter(10)
	}

}
