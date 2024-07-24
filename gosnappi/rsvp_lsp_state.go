package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RsvpLspState *****
type rsvpLspState struct {
	validation
	obj          *otg.RsvpLspState
	marshaller   marshalRsvpLspState
	unMarshaller unMarshalRsvpLspState
}

func NewRsvpLspState() RsvpLspState {
	obj := rsvpLspState{obj: &otg.RsvpLspState{}}
	obj.setDefault()
	return &obj
}

func (obj *rsvpLspState) msg() *otg.RsvpLspState {
	return obj.obj
}

func (obj *rsvpLspState) setMsg(msg *otg.RsvpLspState) RsvpLspState {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalrsvpLspState struct {
	obj *rsvpLspState
}

type marshalRsvpLspState interface {
	// ToProto marshals RsvpLspState to protobuf object *otg.RsvpLspState
	ToProto() (*otg.RsvpLspState, error)
	// ToPbText marshals RsvpLspState to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RsvpLspState to YAML text
	ToYaml() (string, error)
	// ToJson marshals RsvpLspState to JSON text
	ToJson() (string, error)
}

type unMarshalrsvpLspState struct {
	obj *rsvpLspState
}

type unMarshalRsvpLspState interface {
	// FromProto unmarshals RsvpLspState from protobuf object *otg.RsvpLspState
	FromProto(msg *otg.RsvpLspState) (RsvpLspState, error)
	// FromPbText unmarshals RsvpLspState from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RsvpLspState from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RsvpLspState from JSON text
	FromJson(value string) error
}

func (obj *rsvpLspState) Marshal() marshalRsvpLspState {
	if obj.marshaller == nil {
		obj.marshaller = &marshalrsvpLspState{obj: obj}
	}
	return obj.marshaller
}

func (obj *rsvpLspState) Unmarshal() unMarshalRsvpLspState {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalrsvpLspState{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalrsvpLspState) ToProto() (*otg.RsvpLspState, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalrsvpLspState) FromProto(msg *otg.RsvpLspState) (RsvpLspState, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalrsvpLspState) ToPbText() (string, error) {
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

func (m *unMarshalrsvpLspState) FromPbText(value string) error {
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

func (m *marshalrsvpLspState) ToYaml() (string, error) {
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

func (m *unMarshalrsvpLspState) FromYaml(value string) error {
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

func (m *marshalrsvpLspState) ToJson() (string, error) {
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

func (m *unMarshalrsvpLspState) FromJson(value string) error {
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

func (obj *rsvpLspState) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *rsvpLspState) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *rsvpLspState) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *rsvpLspState) Clone() (RsvpLspState, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRsvpLspState()
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

// RsvpLspState is iPv4 RSVP-TE Discovered LSPs.
type RsvpLspState interface {
	Validation
	// msg marshals RsvpLspState to protobuf object *otg.RsvpLspState
	// and doesn't set defaults
	msg() *otg.RsvpLspState
	// setMsg unmarshals RsvpLspState from protobuf object *otg.RsvpLspState
	// and doesn't set defaults
	setMsg(*otg.RsvpLspState) RsvpLspState
	// provides marshal interface
	Marshal() marshalRsvpLspState
	// provides unmarshal interface
	Unmarshal() unMarshalRsvpLspState
	// validate validates RsvpLspState
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RsvpLspState, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// TunnelId returns uint32, set in RsvpLspState.
	TunnelId() uint32
	// SetTunnelId assigns uint32 provided by user to RsvpLspState
	SetTunnelId(value uint32) RsvpLspState
	// HasTunnelId checks if TunnelId has been set in RsvpLspState
	HasTunnelId() bool
	// LspId returns uint32, set in RsvpLspState.
	LspId() uint32
	// SetLspId assigns uint32 provided by user to RsvpLspState
	SetLspId(value uint32) RsvpLspState
	// HasLspId checks if LspId has been set in RsvpLspState
	HasLspId() bool
	// SessionName returns string, set in RsvpLspState.
	SessionName() string
	// SetSessionName assigns string provided by user to RsvpLspState
	SetSessionName(value string) RsvpLspState
	// HasSessionName checks if SessionName has been set in RsvpLspState
	HasSessionName() bool
	// LabelIn returns uint32, set in RsvpLspState.
	LabelIn() uint32
	// SetLabelIn assigns uint32 provided by user to RsvpLspState
	SetLabelIn(value uint32) RsvpLspState
	// HasLabelIn checks if LabelIn has been set in RsvpLspState
	HasLabelIn() bool
	// LabelOut returns uint32, set in RsvpLspState.
	LabelOut() uint32
	// SetLabelOut assigns uint32 provided by user to RsvpLspState
	SetLabelOut(value uint32) RsvpLspState
	// HasLabelOut checks if LabelOut has been set in RsvpLspState
	HasLabelOut() bool
	// SessionStatus returns RsvpLspStateSessionStatusEnum, set in RsvpLspState
	SessionStatus() RsvpLspStateSessionStatusEnum
	// SetSessionStatus assigns RsvpLspStateSessionStatusEnum provided by user to RsvpLspState
	SetSessionStatus(value RsvpLspStateSessionStatusEnum) RsvpLspState
	// HasSessionStatus checks if SessionStatus has been set in RsvpLspState
	HasSessionStatus() bool
	// LastFlapReason returns RsvpLspStateLastFlapReasonEnum, set in RsvpLspState
	LastFlapReason() RsvpLspStateLastFlapReasonEnum
	// SetLastFlapReason assigns RsvpLspStateLastFlapReasonEnum provided by user to RsvpLspState
	SetLastFlapReason(value RsvpLspStateLastFlapReasonEnum) RsvpLspState
	// HasLastFlapReason checks if LastFlapReason has been set in RsvpLspState
	HasLastFlapReason() bool
	// UpTime returns uint64, set in RsvpLspState.
	UpTime() uint64
	// SetUpTime assigns uint64 provided by user to RsvpLspState
	SetUpTime(value uint64) RsvpLspState
	// HasUpTime checks if UpTime has been set in RsvpLspState
	HasUpTime() bool
}

// The tunnel id of RSVP session which acts as an identifier that remains constant over the life of the tunnel.
// TunnelId returns a uint32
func (obj *rsvpLspState) TunnelId() uint32 {

	return *obj.obj.TunnelId

}

// The tunnel id of RSVP session which acts as an identifier that remains constant over the life of the tunnel.
// TunnelId returns a uint32
func (obj *rsvpLspState) HasTunnelId() bool {
	return obj.obj.TunnelId != nil
}

// The tunnel id of RSVP session which acts as an identifier that remains constant over the life of the tunnel.
// SetTunnelId sets the uint32 value in the RsvpLspState object
func (obj *rsvpLspState) SetTunnelId(value uint32) RsvpLspState {

	obj.obj.TunnelId = &value
	return obj
}

// The lsp-id of RSVP session which acts as a differentiator for two lsps originating from the same headend, commonly used to distinguish RSVP sessions during make before break operations.
// LspId returns a uint32
func (obj *rsvpLspState) LspId() uint32 {

	return *obj.obj.LspId

}

// The lsp-id of RSVP session which acts as a differentiator for two lsps originating from the same headend, commonly used to distinguish RSVP sessions during make before break operations.
// LspId returns a uint32
func (obj *rsvpLspState) HasLspId() bool {
	return obj.obj.LspId != nil
}

// The lsp-id of RSVP session which acts as a differentiator for two lsps originating from the same headend, commonly used to distinguish RSVP sessions during make before break operations.
// SetLspId sets the uint32 value in the RsvpLspState object
func (obj *rsvpLspState) SetLspId(value uint32) RsvpLspState {

	obj.obj.LspId = &value
	return obj
}

// The value of RSVP-TE Session Name field of the Session Attribute object.
// SessionName returns a string
func (obj *rsvpLspState) SessionName() string {

	return *obj.obj.SessionName

}

// The value of RSVP-TE Session Name field of the Session Attribute object.
// SessionName returns a string
func (obj *rsvpLspState) HasSessionName() bool {
	return obj.obj.SessionName != nil
}

// The value of RSVP-TE Session Name field of the Session Attribute object.
// SetSessionName sets the string value in the RsvpLspState object
func (obj *rsvpLspState) SetSessionName(value string) RsvpLspState {

	obj.obj.SessionName = &value
	return obj
}

// The label received by RSVP-TE ingress.
// LabelIn returns a uint32
func (obj *rsvpLspState) LabelIn() uint32 {

	return *obj.obj.LabelIn

}

// The label received by RSVP-TE ingress.
// LabelIn returns a uint32
func (obj *rsvpLspState) HasLabelIn() bool {
	return obj.obj.LabelIn != nil
}

// The label received by RSVP-TE ingress.
// SetLabelIn sets the uint32 value in the RsvpLspState object
func (obj *rsvpLspState) SetLabelIn(value uint32) RsvpLspState {

	obj.obj.LabelIn = &value
	return obj
}

// The label assigned by RSVP-TE egress.
// LabelOut returns a uint32
func (obj *rsvpLspState) LabelOut() uint32 {

	return *obj.obj.LabelOut

}

// The label assigned by RSVP-TE egress.
// LabelOut returns a uint32
func (obj *rsvpLspState) HasLabelOut() bool {
	return obj.obj.LabelOut != nil
}

// The label assigned by RSVP-TE egress.
// SetLabelOut sets the uint32 value in the RsvpLspState object
func (obj *rsvpLspState) SetLabelOut(value uint32) RsvpLspState {

	obj.obj.LabelOut = &value
	return obj
}

type RsvpLspStateSessionStatusEnum string

// Enum of SessionStatus on RsvpLspState
var RsvpLspStateSessionStatus = struct {
	UP   RsvpLspStateSessionStatusEnum
	DOWN RsvpLspStateSessionStatusEnum
}{
	UP:   RsvpLspStateSessionStatusEnum("up"),
	DOWN: RsvpLspStateSessionStatusEnum("down"),
}

func (obj *rsvpLspState) SessionStatus() RsvpLspStateSessionStatusEnum {
	return RsvpLspStateSessionStatusEnum(obj.obj.SessionStatus.Enum().String())
}

// Operational state of the RSVP LSP.
// SessionStatus returns a string
func (obj *rsvpLspState) HasSessionStatus() bool {
	return obj.obj.SessionStatus != nil
}

func (obj *rsvpLspState) SetSessionStatus(value RsvpLspStateSessionStatusEnum) RsvpLspState {
	intValue, ok := otg.RsvpLspState_SessionStatus_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpLspStateSessionStatusEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpLspState_SessionStatus_Enum(intValue)
	obj.obj.SessionStatus = &enumValue

	return obj
}

type RsvpLspStateLastFlapReasonEnum string

// Enum of LastFlapReason on RsvpLspState
var RsvpLspStateLastFlapReason = struct {
	RESV_TEAR    RsvpLspStateLastFlapReasonEnum
	PATH_TEAR    RsvpLspStateLastFlapReasonEnum
	PATH_TIMEOUT RsvpLspStateLastFlapReasonEnum
}{
	RESV_TEAR:    RsvpLspStateLastFlapReasonEnum("resv_tear"),
	PATH_TEAR:    RsvpLspStateLastFlapReasonEnum("path_tear"),
	PATH_TIMEOUT: RsvpLspStateLastFlapReasonEnum("path_timeout"),
}

func (obj *rsvpLspState) LastFlapReason() RsvpLspStateLastFlapReasonEnum {
	return RsvpLspStateLastFlapReasonEnum(obj.obj.LastFlapReason.Enum().String())
}

// The reason for the last flap of this RSVP session.
// LastFlapReason returns a string
func (obj *rsvpLspState) HasLastFlapReason() bool {
	return obj.obj.LastFlapReason != nil
}

func (obj *rsvpLspState) SetLastFlapReason(value RsvpLspStateLastFlapReasonEnum) RsvpLspState {
	intValue, ok := otg.RsvpLspState_LastFlapReason_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on RsvpLspStateLastFlapReasonEnum", string(value)))
		return obj
	}
	enumValue := otg.RsvpLspState_LastFlapReason_Enum(intValue)
	obj.obj.LastFlapReason = &enumValue

	return obj
}

// The tunnel UP time in milli seconds. If the tunnel is DOWN the UP time will be zero.
// UpTime returns a uint64
func (obj *rsvpLspState) UpTime() uint64 {

	return *obj.obj.UpTime

}

// The tunnel UP time in milli seconds. If the tunnel is DOWN the UP time will be zero.
// UpTime returns a uint64
func (obj *rsvpLspState) HasUpTime() bool {
	return obj.obj.UpTime != nil
}

// The tunnel UP time in milli seconds. If the tunnel is DOWN the UP time will be zero.
// SetUpTime sets the uint64 value in the RsvpLspState object
func (obj *rsvpLspState) SetUpTime(value uint64) RsvpLspState {

	obj.obj.UpTime = &value
	return obj
}

func (obj *rsvpLspState) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

}

func (obj *rsvpLspState) setDefault() {

}
