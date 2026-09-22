package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetLlr *****
type ultraEthernetLlr struct {
	validation
	obj          *otg.UltraEthernetLlr
	marshaller   marshalUltraEthernetLlr
	unMarshaller unMarshalUltraEthernetLlr
}

func NewUltraEthernetLlr() UltraEthernetLlr {
	obj := ultraEthernetLlr{obj: &otg.UltraEthernetLlr{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetLlr) msg() *otg.UltraEthernetLlr {
	return obj.obj
}

func (obj *ultraEthernetLlr) setMsg(msg *otg.UltraEthernetLlr) UltraEthernetLlr {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetLlr struct {
	obj *ultraEthernetLlr
}

type marshalUltraEthernetLlr interface {
	// ToProto marshals UltraEthernetLlr to protobuf object *otg.UltraEthernetLlr
	ToProto() (*otg.UltraEthernetLlr, error)
	// ToPbText marshals UltraEthernetLlr to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetLlr to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetLlr to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetLlr struct {
	obj *ultraEthernetLlr
}

type unMarshalUltraEthernetLlr interface {
	// FromProto unmarshals UltraEthernetLlr from protobuf object *otg.UltraEthernetLlr
	FromProto(msg *otg.UltraEthernetLlr) (UltraEthernetLlr, error)
	// FromPbText unmarshals UltraEthernetLlr from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetLlr from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetLlr from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetLlr) Marshal() marshalUltraEthernetLlr {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetLlr{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetLlr) Unmarshal() unMarshalUltraEthernetLlr {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetLlr{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetLlr) ToProto() (*otg.UltraEthernetLlr, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetLlr) FromProto(msg *otg.UltraEthernetLlr) (UltraEthernetLlr, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetLlr) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetLlr) FromPbText(value string) error {
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

func (m *marshalultraEthernetLlr) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetLlr) FromYaml(value string) error {
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

func (m *marshalultraEthernetLlr) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetLlr) FromJson(value string) error {
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

func (obj *ultraEthernetLlr) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetLlr) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetLlr) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetLlr) Clone() (UltraEthernetLlr, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetLlr()
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

// UltraEthernetLlr is ultra Ethernet Link Layer Retry (LLR) settings. LLR provides lossless link
// operation by assigning a sequence number to LLR-eligible frames, holding them
// in a replay buffer, and retransmitting them when the link partner reports a
// loss via LLR_NACK or when a replay timer expires.
//
// Reference: UE-Specification-1.0.3 Section 5.1; configuration registers
// Table 5-9, counters Table 5-13.
type UltraEthernetLlr interface {
	Validation
	// msg marshals UltraEthernetLlr to protobuf object *otg.UltraEthernetLlr
	// and doesn't set defaults
	msg() *otg.UltraEthernetLlr
	// setMsg unmarshals UltraEthernetLlr from protobuf object *otg.UltraEthernetLlr
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetLlr) UltraEthernetLlr
	// provides marshal interface
	Marshal() marshalUltraEthernetLlr
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetLlr
	// validate validates UltraEthernetLlr
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetLlr, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// LocalEnable returns bool, set in UltraEthernetLlr.
	LocalEnable() bool
	// SetLocalEnable assigns bool provided by user to UltraEthernetLlr
	SetLocalEnable(value bool) UltraEthernetLlr
	// HasLocalEnable checks if LocalEnable has been set in UltraEthernetLlr
	HasLocalEnable() bool
	// RemoteEnable returns bool, set in UltraEthernetLlr.
	RemoteEnable() bool
	// SetRemoteEnable assigns bool provided by user to UltraEthernetLlr
	SetRemoteEnable(value bool) UltraEthernetLlr
	// HasRemoteEnable checks if RemoteEnable has been set in UltraEthernetLlr
	HasRemoteEnable() bool
	// OutstandingSeqMax returns uint32, set in UltraEthernetLlr.
	OutstandingSeqMax() uint32
	// SetOutstandingSeqMax assigns uint32 provided by user to UltraEthernetLlr
	SetOutstandingSeqMax(value uint32) UltraEthernetLlr
	// HasOutstandingSeqMax checks if OutstandingSeqMax has been set in UltraEthernetLlr
	HasOutstandingSeqMax() bool
	// OutstandingDataMax returns uint32, set in UltraEthernetLlr.
	OutstandingDataMax() uint32
	// SetOutstandingDataMax assigns uint32 provided by user to UltraEthernetLlr
	SetOutstandingDataMax(value uint32) UltraEthernetLlr
	// HasOutstandingDataMax checks if OutstandingDataMax has been set in UltraEthernetLlr
	HasOutstandingDataMax() bool
	// ReplayTimerMax returns uint32, set in UltraEthernetLlr.
	ReplayTimerMax() uint32
	// SetReplayTimerMax assigns uint32 provided by user to UltraEthernetLlr
	SetReplayTimerMax(value uint32) UltraEthernetLlr
	// HasReplayTimerMax checks if ReplayTimerMax has been set in UltraEthernetLlr
	HasReplayTimerMax() bool
	// ReplayCountMax returns uint32, set in UltraEthernetLlr.
	ReplayCountMax() uint32
	// SetReplayCountMax assigns uint32 provided by user to UltraEthernetLlr
	SetReplayCountMax(value uint32) UltraEthernetLlr
	// HasReplayCountMax checks if ReplayCountMax has been set in UltraEthernetLlr
	HasReplayCountMax() bool
	// PcsLostStatusTimerMax returns uint64, set in UltraEthernetLlr.
	PcsLostStatusTimerMax() uint64
	// SetPcsLostStatusTimerMax assigns uint64 provided by user to UltraEthernetLlr
	SetPcsLostStatusTimerMax(value uint64) UltraEthernetLlr
	// HasPcsLostStatusTimerMax checks if PcsLostStatusTimerMax has been set in UltraEthernetLlr
	HasPcsLostStatusTimerMax() bool
	// DataAgeTimerMax returns uint64, set in UltraEthernetLlr.
	DataAgeTimerMax() uint64
	// SetDataAgeTimerMax assigns uint64 provided by user to UltraEthernetLlr
	SetDataAgeTimerMax(value uint64) UltraEthernetLlr
	// HasDataAgeTimerMax checks if DataAgeTimerMax has been set in UltraEthernetLlr
	HasDataAgeTimerMax() bool
	// InitBehavior returns UltraEthernetLlrInitBehaviorEnum, set in UltraEthernetLlr
	InitBehavior() UltraEthernetLlrInitBehaviorEnum
	// SetInitBehavior assigns UltraEthernetLlrInitBehaviorEnum provided by user to UltraEthernetLlr
	SetInitBehavior(value UltraEthernetLlrInitBehaviorEnum) UltraEthernetLlr
	// HasInitBehavior checks if InitBehavior has been set in UltraEthernetLlr
	HasInitBehavior() bool
	// FlushBehavior returns UltraEthernetLlrFlushBehaviorEnum, set in UltraEthernetLlr
	FlushBehavior() UltraEthernetLlrFlushBehaviorEnum
	// SetFlushBehavior assigns UltraEthernetLlrFlushBehaviorEnum provided by user to UltraEthernetLlr
	SetFlushBehavior(value UltraEthernetLlrFlushBehaviorEnum) UltraEthernetLlr
	// HasFlushBehavior checks if FlushBehavior has been set in UltraEthernetLlr
	HasFlushBehavior() bool
	// ReInitOnDiscard returns bool, set in UltraEthernetLlr.
	ReInitOnDiscard() bool
	// SetReInitOnDiscard assigns bool provided by user to UltraEthernetLlr
	SetReInitOnDiscard(value bool) UltraEthernetLlr
	// HasReInitOnDiscard checks if ReInitOnDiscard has been set in UltraEthernetLlr
	HasReInitOnDiscard() bool
}

// Enables LLR reception on the port. When enabled, the port is allowed to
// receive LLR_INIT and operate as an LLR receiver (llr_mode_local).
// LocalEnable returns a bool
func (obj *ultraEthernetLlr) LocalEnable() bool {

	return *obj.obj.LocalEnable

}

// Enables LLR reception on the port. When enabled, the port is allowed to
// receive LLR_INIT and operate as an LLR receiver (llr_mode_local).
// LocalEnable returns a bool
func (obj *ultraEthernetLlr) HasLocalEnable() bool {
	return obj.obj.LocalEnable != nil
}

// Enables LLR reception on the port. When enabled, the port is allowed to
// receive LLR_INIT and operate as an LLR receiver (llr_mode_local).
// SetLocalEnable sets the bool value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetLocalEnable(value bool) UltraEthernetLlr {

	obj.obj.LocalEnable = &value
	return obj
}

// Enables LLR transmission on the port. When enabled, the port may send
// LLR_INIT and operate as an LLR transmitter (llr_mode_remote).
// RemoteEnable returns a bool
func (obj *ultraEthernetLlr) RemoteEnable() bool {

	return *obj.obj.RemoteEnable

}

// Enables LLR transmission on the port. When enabled, the port may send
// LLR_INIT and operate as an LLR transmitter (llr_mode_remote).
// RemoteEnable returns a bool
func (obj *ultraEthernetLlr) HasRemoteEnable() bool {
	return obj.obj.RemoteEnable != nil
}

// Enables LLR transmission on the port. When enabled, the port may send
// LLR_INIT and operate as an LLR transmitter (llr_mode_remote).
// SetRemoteEnable sets the bool value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetRemoteEnable(value bool) UltraEthernetLlr {

	obj.obj.RemoteEnable = &value
	return obj
}

// The maximum number of unacknowledged (in flight) LLR frames permitted.
// The specification permits an absolute maximum of 524288, however an
// implementation may support a lower maximum depending on the link speed
// and minimum frame size.
// OutstandingSeqMax returns a uint32
func (obj *ultraEthernetLlr) OutstandingSeqMax() uint32 {

	return *obj.obj.OutstandingSeqMax

}

// The maximum number of unacknowledged (in flight) LLR frames permitted.
// The specification permits an absolute maximum of 524288, however an
// implementation may support a lower maximum depending on the link speed
// and minimum frame size.
// OutstandingSeqMax returns a uint32
func (obj *ultraEthernetLlr) HasOutstandingSeqMax() bool {
	return obj.obj.OutstandingSeqMax != nil
}

// The maximum number of unacknowledged (in flight) LLR frames permitted.
// The specification permits an absolute maximum of 524288, however an
// implementation may support a lower maximum depending on the link speed
// and minimum frame size.
// SetOutstandingSeqMax sets the uint32 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetOutstandingSeqMax(value uint32) UltraEthernetLlr {

	obj.obj.OutstandingSeqMax = &value
	return obj
}

// The maximum number of unacknowledged (in flight) bytes permitted. This
// should be set to the link's bandwidth-delay product to ensure correct
// operation of the pause and PFC mechanisms.
// OutstandingDataMax returns a uint32
func (obj *ultraEthernetLlr) OutstandingDataMax() uint32 {

	return *obj.obj.OutstandingDataMax

}

// The maximum number of unacknowledged (in flight) bytes permitted. This
// should be set to the link's bandwidth-delay product to ensure correct
// operation of the pause and PFC mechanisms.
// OutstandingDataMax returns a uint32
func (obj *ultraEthernetLlr) HasOutstandingDataMax() bool {
	return obj.obj.OutstandingDataMax != nil
}

// The maximum number of unacknowledged (in flight) bytes permitted. This
// should be set to the link's bandwidth-delay product to ensure correct
// operation of the pause and PFC mechanisms.
// SetOutstandingDataMax sets the uint32 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetOutstandingDataMax(value uint32) UltraEthernetLlr {

	obj.obj.OutstandingDataMax = &value
	return obj
}

// The time, in nanoseconds, before a missing acknowledgement triggers a
// replay of the unacknowledged frames held in the replay buffer.
// ReplayTimerMax returns a uint32
func (obj *ultraEthernetLlr) ReplayTimerMax() uint32 {

	return *obj.obj.ReplayTimerMax

}

// The time, in nanoseconds, before a missing acknowledgement triggers a
// replay of the unacknowledged frames held in the replay buffer.
// ReplayTimerMax returns a uint32
func (obj *ultraEthernetLlr) HasReplayTimerMax() bool {
	return obj.obj.ReplayTimerMax != nil
}

// The time, in nanoseconds, before a missing acknowledgement triggers a
// replay of the unacknowledged frames held in the replay buffer.
// SetReplayTimerMax sets the uint32 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetReplayTimerMax(value uint32) UltraEthernetLlr {

	obj.obj.ReplayTimerMax = &value
	return obj
}

// The maximum number of replay attempts before the LLR gives up and enters
// the FLUSH state. A value of 255 indicates there is no maximum (unlimited
// retries).
// ReplayCountMax returns a uint32
func (obj *ultraEthernetLlr) ReplayCountMax() uint32 {

	return *obj.obj.ReplayCountMax

}

// The maximum number of replay attempts before the LLR gives up and enters
// the FLUSH state. A value of 255 indicates there is no maximum (unlimited
// retries).
// ReplayCountMax returns a uint32
func (obj *ultraEthernetLlr) HasReplayCountMax() bool {
	return obj.obj.ReplayCountMax != nil
}

// The maximum number of replay attempts before the LLR gives up and enters
// the FLUSH state. A value of 255 indicates there is no maximum (unlimited
// retries).
// SetReplayCountMax sets the uint32 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetReplayCountMax(value uint32) UltraEthernetLlr {

	obj.obj.ReplayCountMax = &value
	return obj
}

// The time, in nanoseconds, that the PCS link may remain down before the LLR
// transmit state machine is forced to transition to the FLUSH state. A value
// of 0 causes immediate expiration.
// PcsLostStatusTimerMax returns a uint64
func (obj *ultraEthernetLlr) PcsLostStatusTimerMax() uint64 {

	return *obj.obj.PcsLostStatusTimerMax

}

// The time, in nanoseconds, that the PCS link may remain down before the LLR
// transmit state machine is forced to transition to the FLUSH state. A value
// of 0 causes immediate expiration.
// PcsLostStatusTimerMax returns a uint64
func (obj *ultraEthernetLlr) HasPcsLostStatusTimerMax() bool {
	return obj.obj.PcsLostStatusTimerMax != nil
}

// The time, in nanoseconds, that the PCS link may remain down before the LLR
// transmit state machine is forced to transition to the FLUSH state. A value
// of 0 causes immediate expiration.
// SetPcsLostStatusTimerMax sets the uint64 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetPcsLostStatusTimerMax(value uint64) UltraEthernetLlr {

	obj.obj.PcsLostStatusTimerMax = &value
	return obj
}

// The maximum time, in nanoseconds, that data may sit unacknowledged in the
// replay buffer before it is discarded as too old.
// DataAgeTimerMax returns a uint64
func (obj *ultraEthernetLlr) DataAgeTimerMax() uint64 {

	return *obj.obj.DataAgeTimerMax

}

// The maximum time, in nanoseconds, that data may sit unacknowledged in the
// replay buffer before it is discarded as too old.
// DataAgeTimerMax returns a uint64
func (obj *ultraEthernetLlr) HasDataAgeTimerMax() bool {
	return obj.obj.DataAgeTimerMax != nil
}

// The maximum time, in nanoseconds, that data may sit unacknowledged in the
// replay buffer before it is discarded as too old.
// SetDataAgeTimerMax sets the uint64 value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetDataAgeTimerMax(value uint64) UltraEthernetLlr {

	obj.obj.DataAgeTimerMax = &value
	return obj
}

type UltraEthernetLlrInitBehaviorEnum string

// Enum of InitBehavior on UltraEthernetLlr
var UltraEthernetLlrInitBehavior = struct {
	DISCARD     UltraEthernetLlrInitBehaviorEnum
	BLOCK       UltraEthernetLlrInitBehaviorEnum
	BEST_EFFORT UltraEthernetLlrInitBehaviorEnum
}{
	DISCARD:     UltraEthernetLlrInitBehaviorEnum("discard"),
	BLOCK:       UltraEthernetLlrInitBehaviorEnum("block"),
	BEST_EFFORT: UltraEthernetLlrInitBehaviorEnum("best_effort"),
}

func (obj *ultraEthernetLlr) InitBehavior() UltraEthernetLlrInitBehaviorEnum {
	return UltraEthernetLlrInitBehaviorEnum(obj.obj.InitBehavior.Enum().String())
}

// Controls how the transmit LLR handles outgoing LLR-desired frames while in
// the INIT state.
//
// - discard: frames are discarded without being transmitted.
// - block: frames are not accepted from the MAC client, blocking transmission.
// - best_effort: frames are transmitted as LLR-ineligible (no replay protection).
// InitBehavior returns a string
func (obj *ultraEthernetLlr) HasInitBehavior() bool {
	return obj.obj.InitBehavior != nil
}

func (obj *ultraEthernetLlr) SetInitBehavior(value UltraEthernetLlrInitBehaviorEnum) UltraEthernetLlr {
	intValue, ok := otg.UltraEthernetLlr_InitBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetLlrInitBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetLlr_InitBehavior_Enum(intValue)
	obj.obj.InitBehavior = &enumValue

	return obj
}

type UltraEthernetLlrFlushBehaviorEnum string

// Enum of FlushBehavior on UltraEthernetLlr
var UltraEthernetLlrFlushBehavior = struct {
	DISCARD     UltraEthernetLlrFlushBehaviorEnum
	BLOCK       UltraEthernetLlrFlushBehaviorEnum
	BEST_EFFORT UltraEthernetLlrFlushBehaviorEnum
}{
	DISCARD:     UltraEthernetLlrFlushBehaviorEnum("discard"),
	BLOCK:       UltraEthernetLlrFlushBehaviorEnum("block"),
	BEST_EFFORT: UltraEthernetLlrFlushBehaviorEnum("best_effort"),
}

func (obj *ultraEthernetLlr) FlushBehavior() UltraEthernetLlrFlushBehaviorEnum {
	return UltraEthernetLlrFlushBehaviorEnum(obj.obj.FlushBehavior.Enum().String())
}

// Controls how the transmit LLR handles outgoing LLR-desired frames while in
// the FLUSH state.
//
// - discard: frames are discarded without being transmitted.
// - block: frames are not accepted from the MAC client, blocking transmission.
// - best_effort: frames are transmitted as LLR-ineligible (no replay protection).
// FlushBehavior returns a string
func (obj *ultraEthernetLlr) HasFlushBehavior() bool {
	return obj.obj.FlushBehavior != nil
}

func (obj *ultraEthernetLlr) SetFlushBehavior(value UltraEthernetLlrFlushBehaviorEnum) UltraEthernetLlr {
	intValue, ok := otg.UltraEthernetLlr_FlushBehavior_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetLlrFlushBehaviorEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetLlr_FlushBehavior_Enum(intValue)
	obj.obj.FlushBehavior = &enumValue

	return obj
}

// When true, the LLR automatically re-initializes after a replay failure.
// When false, it waits for management intervention.
// ReInitOnDiscard returns a bool
func (obj *ultraEthernetLlr) ReInitOnDiscard() bool {

	return *obj.obj.ReInitOnDiscard

}

// When true, the LLR automatically re-initializes after a replay failure.
// When false, it waits for management intervention.
// ReInitOnDiscard returns a bool
func (obj *ultraEthernetLlr) HasReInitOnDiscard() bool {
	return obj.obj.ReInitOnDiscard != nil
}

// When true, the LLR automatically re-initializes after a replay failure.
// When false, it waits for management intervention.
// SetReInitOnDiscard sets the bool value in the UltraEthernetLlr object
func (obj *ultraEthernetLlr) SetReInitOnDiscard(value bool) UltraEthernetLlr {

	obj.obj.ReInitOnDiscard = &value
	return obj
}

func (obj *ultraEthernetLlr) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.OutstandingSeqMax != nil {

		if *obj.obj.OutstandingSeqMax > 524288 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.OutstandingSeqMax <= 524288 but Got %d", *obj.obj.OutstandingSeqMax))
		}

	}

	if obj.obj.OutstandingDataMax != nil {

		if *obj.obj.OutstandingDataMax > 2147483647 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.OutstandingDataMax <= 2147483647 but Got %d", *obj.obj.OutstandingDataMax))
		}

	}

	if obj.obj.ReplayTimerMax != nil {

		if *obj.obj.ReplayTimerMax > 65535 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.ReplayTimerMax <= 65535 but Got %d", *obj.obj.ReplayTimerMax))
		}

	}

	if obj.obj.ReplayCountMax != nil {

		if *obj.obj.ReplayCountMax > 255 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.ReplayCountMax <= 255 but Got %d", *obj.obj.ReplayCountMax))
		}

	}

	if obj.obj.PcsLostStatusTimerMax != nil {

		if *obj.obj.PcsLostStatusTimerMax > 4290000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.PcsLostStatusTimerMax <= 4290000000 but Got %d", *obj.obj.PcsLostStatusTimerMax))
		}

	}

	if obj.obj.DataAgeTimerMax != nil {

		if *obj.obj.DataAgeTimerMax > 4290000000 {
			vObj.validationErrors = append(
				vObj.validationErrors,
				fmt.Sprintf("0 <= UltraEthernetLlr.DataAgeTimerMax <= 4290000000 but Got %d", *obj.obj.DataAgeTimerMax))
		}

	}

}

func (obj *ultraEthernetLlr) setDefault() {
	if obj.obj.LocalEnable == nil {
		obj.SetLocalEnable(false)
	}
	if obj.obj.RemoteEnable == nil {
		obj.SetRemoteEnable(false)
	}
	if obj.obj.OutstandingSeqMax == nil {
		obj.SetOutstandingSeqMax(2047)
	}
	if obj.obj.OutstandingDataMax == nil {
		obj.SetOutstandingDataMax(64000000)
	}
	if obj.obj.ReplayTimerMax == nil {
		obj.SetReplayTimerMax(65535)
	}
	if obj.obj.ReplayCountMax == nil {
		obj.SetReplayCountMax(255)
	}
	if obj.obj.PcsLostStatusTimerMax == nil {
		obj.SetPcsLostStatusTimerMax(4200000000)
	}
	if obj.obj.DataAgeTimerMax == nil {
		obj.SetDataAgeTimerMax(4200000000)
	}
	if obj.obj.InitBehavior == nil {
		obj.SetInitBehavior(UltraEthernetLlrInitBehavior.BEST_EFFORT)

	}
	if obj.obj.FlushBehavior == nil {
		obj.SetFlushBehavior(UltraEthernetLlrFlushBehavior.BEST_EFFORT)

	}
	if obj.obj.ReInitOnDiscard == nil {
		obj.SetReInitOnDiscard(false)
	}

}
