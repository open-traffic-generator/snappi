package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** IsisActionProtocolIsisUnplannedRestart *****
type isisActionProtocolIsisUnplannedRestart struct {
	validation
	obj          *otg.IsisActionProtocolIsisUnplannedRestart
	marshaller   marshalIsisActionProtocolIsisUnplannedRestart
	unMarshaller unMarshalIsisActionProtocolIsisUnplannedRestart
}

func NewIsisActionProtocolIsisUnplannedRestart() IsisActionProtocolIsisUnplannedRestart {
	obj := isisActionProtocolIsisUnplannedRestart{obj: &otg.IsisActionProtocolIsisUnplannedRestart{}}
	obj.setDefault()
	return &obj
}

func (obj *isisActionProtocolIsisUnplannedRestart) msg() *otg.IsisActionProtocolIsisUnplannedRestart {
	return obj.obj
}

func (obj *isisActionProtocolIsisUnplannedRestart) setMsg(msg *otg.IsisActionProtocolIsisUnplannedRestart) IsisActionProtocolIsisUnplannedRestart {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalisisActionProtocolIsisUnplannedRestart struct {
	obj *isisActionProtocolIsisUnplannedRestart
}

type marshalIsisActionProtocolIsisUnplannedRestart interface {
	// ToProto marshals IsisActionProtocolIsisUnplannedRestart to protobuf object *otg.IsisActionProtocolIsisUnplannedRestart
	ToProto() (*otg.IsisActionProtocolIsisUnplannedRestart, error)
	// ToPbText marshals IsisActionProtocolIsisUnplannedRestart to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals IsisActionProtocolIsisUnplannedRestart to YAML text
	ToYaml() (string, error)
	// ToJson marshals IsisActionProtocolIsisUnplannedRestart to JSON text
	ToJson() (string, error)
}

type unMarshalisisActionProtocolIsisUnplannedRestart struct {
	obj *isisActionProtocolIsisUnplannedRestart
}

type unMarshalIsisActionProtocolIsisUnplannedRestart interface {
	// FromProto unmarshals IsisActionProtocolIsisUnplannedRestart from protobuf object *otg.IsisActionProtocolIsisUnplannedRestart
	FromProto(msg *otg.IsisActionProtocolIsisUnplannedRestart) (IsisActionProtocolIsisUnplannedRestart, error)
	// FromPbText unmarshals IsisActionProtocolIsisUnplannedRestart from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals IsisActionProtocolIsisUnplannedRestart from YAML text
	FromYaml(value string) error
	// FromJson unmarshals IsisActionProtocolIsisUnplannedRestart from JSON text
	FromJson(value string) error
}

func (obj *isisActionProtocolIsisUnplannedRestart) Marshal() marshalIsisActionProtocolIsisUnplannedRestart {
	if obj.marshaller == nil {
		obj.marshaller = &marshalisisActionProtocolIsisUnplannedRestart{obj: obj}
	}
	return obj.marshaller
}

func (obj *isisActionProtocolIsisUnplannedRestart) Unmarshal() unMarshalIsisActionProtocolIsisUnplannedRestart {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalisisActionProtocolIsisUnplannedRestart{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalisisActionProtocolIsisUnplannedRestart) ToProto() (*otg.IsisActionProtocolIsisUnplannedRestart, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalisisActionProtocolIsisUnplannedRestart) FromProto(msg *otg.IsisActionProtocolIsisUnplannedRestart) (IsisActionProtocolIsisUnplannedRestart, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalisisActionProtocolIsisUnplannedRestart) ToPbText() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisUnplannedRestart) FromPbText(value string) error {
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

func (m *marshalisisActionProtocolIsisUnplannedRestart) ToYaml() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisUnplannedRestart) FromYaml(value string) error {
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

func (m *marshalisisActionProtocolIsisUnplannedRestart) ToJson() (string, error) {
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

func (m *unMarshalisisActionProtocolIsisUnplannedRestart) FromJson(value string) error {
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

func (obj *isisActionProtocolIsisUnplannedRestart) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisUnplannedRestart) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *isisActionProtocolIsisUnplannedRestart) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *isisActionProtocolIsisUnplannedRestart) Clone() (IsisActionProtocolIsisUnplannedRestart, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewIsisActionProtocolIsisUnplannedRestart()
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

// IsisActionProtocolIsisUnplannedRestart is initiates IS-IS Unplanned Graceful Restart process for the selected IS-IS routers. If no name is specified then Graceful Restart will be sent to all configured IS-IS routers. When an emulated IS-IS router is in the unplanned "Restarting" mode, it sends an IIH PDU containing a Restart TLV with the RR (Restart Request) bit set and  waits for RA (Restart Acknowledge) in an IIH PDU from Neigbhor(s).  The time T1 is maintained per interface and indicates the time after which an unacknowledged (re)start attempt will be repeated.
type IsisActionProtocolIsisUnplannedRestart interface {
	Validation
	// msg marshals IsisActionProtocolIsisUnplannedRestart to protobuf object *otg.IsisActionProtocolIsisUnplannedRestart
	// and doesn't set defaults
	msg() *otg.IsisActionProtocolIsisUnplannedRestart
	// setMsg unmarshals IsisActionProtocolIsisUnplannedRestart from protobuf object *otg.IsisActionProtocolIsisUnplannedRestart
	// and doesn't set defaults
	setMsg(*otg.IsisActionProtocolIsisUnplannedRestart) IsisActionProtocolIsisUnplannedRestart
	// provides marshal interface
	Marshal() marshalIsisActionProtocolIsisUnplannedRestart
	// provides unmarshal interface
	Unmarshal() unMarshalIsisActionProtocolIsisUnplannedRestart
	// validate validates IsisActionProtocolIsisUnplannedRestart
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (IsisActionProtocolIsisUnplannedRestart, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// HoldingTime returns uint32, set in IsisActionProtocolIsisUnplannedRestart.
	HoldingTime() uint32
	// SetHoldingTime assigns uint32 provided by user to IsisActionProtocolIsisUnplannedRestart
	SetHoldingTime(value uint32) IsisActionProtocolIsisUnplannedRestart
	// HasHoldingTime checks if HoldingTime has been set in IsisActionProtocolIsisUnplannedRestart
	HasHoldingTime() bool
	// RestartAfter returns uint32, set in IsisActionProtocolIsisUnplannedRestart.
	RestartAfter() uint32
	// SetRestartAfter assigns uint32 provided by user to IsisActionProtocolIsisUnplannedRestart
	SetRestartAfter(value uint32) IsisActionProtocolIsisUnplannedRestart
	// HasRestartAfter checks if RestartAfter has been set in IsisActionProtocolIsisUnplannedRestart
	HasRestartAfter() bool
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// HoldingTime returns a uint32
func (obj *isisActionProtocolIsisUnplannedRestart) HoldingTime() uint32 {

	return *obj.obj.HoldingTime

}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// HoldingTime returns a uint32
func (obj *isisActionProtocolIsisUnplannedRestart) HasHoldingTime() bool {
	return obj.obj.HoldingTime != nil
}

// This is the estimated duration (in seconds) it will take for the IS-IS session to be re-established after a restart. The hold-timer in the IIH PDU is updated with this time.
// SetHoldingTime sets the uint32 value in the IsisActionProtocolIsisUnplannedRestart object
func (obj *isisActionProtocolIsisUnplannedRestart) SetHoldingTime(value uint32) IsisActionProtocolIsisUnplannedRestart {

	obj.obj.HoldingTime = &value
	return obj
}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU will be sent having Restart Tlv with RR bit unset.
// RestartAfter returns a uint32
func (obj *isisActionProtocolIsisUnplannedRestart) RestartAfter() uint32 {

	return *obj.obj.RestartAfter

}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU will be sent having Restart Tlv with RR bit unset.
// RestartAfter returns a uint32
func (obj *isisActionProtocolIsisUnplannedRestart) HasRestartAfter() bool {
	return obj.obj.RestartAfter != nil
}

// Once it receives Restarting TLV having RA bit set in a IIH PDU,  time (in seconds), after which IIH PDU will be sent having Restart Tlv with RR bit unset.
// SetRestartAfter sets the uint32 value in the IsisActionProtocolIsisUnplannedRestart object
func (obj *isisActionProtocolIsisUnplannedRestart) SetRestartAfter(value uint32) IsisActionProtocolIsisUnplannedRestart {

	obj.obj.RestartAfter = &value
	return obj
}

func (obj *isisActionProtocolIsisUnplannedRestart) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.HoldingTime != nil {

		if *obj.obj.HoldingTime < 10 || *obj.obj.HoldingTime > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisUnplannedRestart.HoldingTime <= 4096 but Got %d", *obj.obj.HoldingTime))
		}

	}

	if obj.obj.RestartAfter != nil {

		if *obj.obj.RestartAfter < 10 || *obj.obj.RestartAfter > 4096 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("10 <= IsisActionProtocolIsisUnplannedRestart.RestartAfter <= 4096 but Got %d", *obj.obj.RestartAfter))
		}

	}

}

func (obj *isisActionProtocolIsisUnplannedRestart) setDefault() {
	if obj.obj.HoldingTime == nil {
		obj.SetHoldingTime(30)
	}
	if obj.obj.RestartAfter == nil {
		obj.SetRestartAfter(10)
	}

}
