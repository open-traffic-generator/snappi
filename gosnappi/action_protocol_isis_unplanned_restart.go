package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** ActionProtocolIsisUnplannedRestart *****
type actionProtocolIsisUnplannedRestart struct {
	validation
	obj          *otg.ActionProtocolIsisUnplannedRestart
	marshaller   marshalActionProtocolIsisUnplannedRestart
	unMarshaller unMarshalActionProtocolIsisUnplannedRestart
}

func NewActionProtocolIsisUnplannedRestart() ActionProtocolIsisUnplannedRestart {
	obj := actionProtocolIsisUnplannedRestart{obj: &otg.ActionProtocolIsisUnplannedRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *actionProtocolIsisUnplannedRestart) msg() *otg.ActionProtocolIsisUnplannedRestart {
	return obj.obj
}

func (obj *actionProtocolIsisUnplannedRestart) setMsg(msg *otg.ActionProtocolIsisUnplannedRestart) ActionProtocolIsisUnplannedRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalactionProtocolIsisUnplannedRestart struct {
	obj *actionProtocolIsisUnplannedRestart
}

type marshalActionProtocolIsisUnplannedRestart interface {
	// ToProto marshals ActionProtocolIsisUnplannedRestart to protobuf object *otg.ActionProtocolIsisUnplannedRestart
	ToProto() (*otg.ActionProtocolIsisUnplannedRestart, error)
	// ToPbText marshals ActionProtocolIsisUnplannedRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals ActionProtocolIsisUnplannedRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals ActionProtocolIsisUnplannedRestart to JSON text
	ToJson() (string, error)
}

type unMarshalactionProtocolIsisUnplannedRestart struct {
	obj *actionProtocolIsisUnplannedRestart
}

type unMarshalActionProtocolIsisUnplannedRestart interface {
	// FromProto unmarshals ActionProtocolIsisUnplannedRestart from protobuf object *otg.ActionProtocolIsisUnplannedRestart
	FromProto(msg *otg.ActionProtocolIsisUnplannedRestart) (ActionProtocolIsisUnplannedRestart, error)
	// FromPbText unmarshals ActionProtocolIsisUnplannedRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals ActionProtocolIsisUnplannedRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals ActionProtocolIsisUnplannedRestart from JSON text
	FromJson(value string) error
}

func (obj *actionProtocolIsisUnplannedRestart) Marshal() marshalActionProtocolIsisUnplannedRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalactionProtocolIsisUnplannedRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *actionProtocolIsisUnplannedRestart) Unmarshal() unMarshalActionProtocolIsisUnplannedRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalactionProtocolIsisUnplannedRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalactionProtocolIsisUnplannedRestart) ToProto() (*otg.ActionProtocolIsisUnplannedRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalactionProtocolIsisUnplannedRestart) FromProto(msg *otg.ActionProtocolIsisUnplannedRestart) (ActionProtocolIsisUnplannedRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalactionProtocolIsisUnplannedRestart) ToPbText() (string, error) {
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

func (m *unMarshalactionProtocolIsisUnplannedRestart) FromPbText(value string) error {
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

func (m *marshalactionProtocolIsisUnplannedRestart) ToYaml() (string, error) {
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

func (m *unMarshalactionProtocolIsisUnplannedRestart) FromYaml(value string) error {
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

func (m *marshalactionProtocolIsisUnplannedRestart) ToJson() (string, error) {
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

func (m *unMarshalactionProtocolIsisUnplannedRestart) FromJson(value string) error {
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

func (obj *actionProtocolIsisUnplannedRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *actionProtocolIsisUnplannedRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *actionProtocolIsisUnplannedRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *actionProtocolIsisUnplannedRestart) Clone() (ActionProtocolIsisUnplannedRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewActionProtocolIsisUnplannedRestart()
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

// ActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the RR (Restart Request) bit set and holding_time updated to as specified by user to indicate the maximum time within which this router  or routers will complete the graceful restart.  It waits for RA (Restart Acknowledge) in an IIH PDU from Neigbhor(s).  The timer T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated.
type ActionProtocolIsisUnplannedRestart interface {
	Validation
	// msg marshals ActionProtocolIsisUnplannedRestart to protobuf object *otg.ActionProtocolIsisUnplannedRestart
	// and doesn't set defaults
	msg() *otg.ActionProtocolIsisUnplannedRestart
	// setMsg unmarshals ActionProtocolIsisUnplannedRestart from protobuf object *otg.ActionProtocolIsisUnplannedRestart
	// and doesn't set defaults
	setMsg(*otg.ActionProtocolIsisUnplannedRestart) ActionProtocolIsisUnplannedRestart
	// provides marshal interface
	Marshal() marshalActionProtocolIsisUnplannedRestart
	// provides unmarshal interface
	Unmarshal() unMarshalActionProtocolIsisUnplannedRestart
	// validate validates ActionProtocolIsisUnplannedRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (ActionProtocolIsisUnplannedRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HoldingTime returns uint32, set in ActionProtocolIsisUnplannedRestart.
	HoldingTime() uint32
	// SetHoldingTime assigns uint32 provided by user to ActionProtocolIsisUnplannedRestart
	SetHoldingTime(value uint32) ActionProtocolIsisUnplannedRestart
	// HasHoldingTime checks if HoldingTime has been set in ActionProtocolIsisUnplannedRestart
	HasHoldingTime() bool
	// RestartAfter returns uint32, set in ActionProtocolIsisUnplannedRestart.
	RestartAfter() uint32
	// SetRestartAfter assigns uint32 provided by user to ActionProtocolIsisUnplannedRestart
	SetRestartAfter(value uint32) ActionProtocolIsisUnplannedRestart
	// HasRestartAfter checks if RestartAfter has been set in ActionProtocolIsisUnplannedRestart
	HasRestartAfter() bool
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// HoldingTime returns a uint32
func (obj *actionProtocolIsisUnplannedRestart) HoldingTime() uint32 {

	return *obj.obj.HoldingTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// HoldingTime returns a uint32
func (obj *actionProtocolIsisUnplannedRestart) HasHoldingTime() bool {
	return obj.obj.HoldingTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// SetHoldingTime sets the uint32 value in the ActionProtocolIsisUnplannedRestart object
func (obj *actionProtocolIsisUnplannedRestart) SetHoldingTime(value uint32) ActionProtocolIsisUnplannedRestart {

	obj.obj.HoldingTime = &value
	return obj
}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU, having Restart Tlv with RR bit unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed,  not started i.e. RR bit is cleared and hold_timer is reset to normal.
// RestartAfter returns a uint32
func (obj *actionProtocolIsisUnplannedRestart) RestartAfter() uint32 {

	return *obj.obj.RestartAfter

}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU, having Restart Tlv with RR bit unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed,  not started i.e. RR bit is cleared and hold_timer is reset to normal.
// RestartAfter returns a uint32
func (obj *actionProtocolIsisUnplannedRestart) HasRestartAfter() bool {
	return obj.obj.RestartAfter != nil
}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU, having Restart Tlv with RR bit unset, will be sent. This should result in IIH to be transmitted indicating restart is Completed,  not started i.e. RR bit is cleared and hold_timer is reset to normal.
// SetRestartAfter sets the uint32 value in the ActionProtocolIsisUnplannedRestart object
func (obj *actionProtocolIsisUnplannedRestart) SetRestartAfter(value uint32) ActionProtocolIsisUnplannedRestart {

	obj.obj.RestartAfter = &value
	return obj
}

func (obj *actionProtocolIsisUnplannedRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HoldingTime != nil {

		if *obj.obj.HoldingTime < 1 || *obj.obj.HoldingTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= ActionProtocolIsisUnplannedRestart.HoldingTime <= 4096 but Got %d", *obj.obj.HoldingTime))
		}

	}

	if obj.obj.RestartAfter != nil {

		if *obj.obj.RestartAfter < 1 || *obj.obj.RestartAfter > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("1 <= ActionProtocolIsisUnplannedRestart.RestartAfter <= 4096 but Got %d", *obj.obj.RestartAfter))
		}

	}

}

func (obj *actionProtocolIsisUnplannedRestart) setDefault() {
	if obj.obj.HoldingTime == nil {
		obj.SetHoldingTime(30)
	}
	if obj.obj.RestartAfter == nil {
		obj.SetRestartAfter(10)
	}

}
