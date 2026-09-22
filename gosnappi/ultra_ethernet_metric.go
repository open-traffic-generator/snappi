package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** UltraEthernetMetric *****
type ultraEthernetMetric struct {
	validation
	obj                               *otg.UltraEthernetMetric
	marshaller                        marshalUltraEthernetMetric
	unMarshaller                      unMarshalUltraEthernetMetric
	cbfcSenderVirtualChannelsHolder   UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	cbfcReceiverVirtualChannelsHolder UltraEthernetMetricUltraEthernetCbfcVcMetricIter
}

func NewUltraEthernetMetric() UltraEthernetMetric {
	obj := ultraEthernetMetric{obj: &otg.UltraEthernetMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *ultraEthernetMetric) msg() *otg.UltraEthernetMetric {
	return obj.obj
}

func (obj *ultraEthernetMetric) setMsg(msg *otg.UltraEthernetMetric) UltraEthernetMetric {
	obj.setNil()
	proto.Merge(obj.obj, msg)
	return obj
}

type marshalultraEthernetMetric struct {
	obj *ultraEthernetMetric
}

type marshalUltraEthernetMetric interface {
	// ToProto marshals UltraEthernetMetric to protobuf object *otg.UltraEthernetMetric
	ToProto() (*otg.UltraEthernetMetric, error)
	// ToPbText marshals UltraEthernetMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals UltraEthernetMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals UltraEthernetMetric to JSON text
	ToJson() (string, error)
}

type unMarshalultraEthernetMetric struct {
	obj *ultraEthernetMetric
}

type unMarshalUltraEthernetMetric interface {
	// FromProto unmarshals UltraEthernetMetric from protobuf object *otg.UltraEthernetMetric
	FromProto(msg *otg.UltraEthernetMetric) (UltraEthernetMetric, error)
	// FromPbText unmarshals UltraEthernetMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals UltraEthernetMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals UltraEthernetMetric from JSON text
	FromJson(value string) error
}

func (obj *ultraEthernetMetric) Marshal() marshalUltraEthernetMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalultraEthernetMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *ultraEthernetMetric) Unmarshal() unMarshalUltraEthernetMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalultraEthernetMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalultraEthernetMetric) ToProto() (*otg.UltraEthernetMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalultraEthernetMetric) FromProto(msg *otg.UltraEthernetMetric) (UltraEthernetMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalultraEthernetMetric) ToPbText() (string, error) {
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

func (m *unMarshalultraEthernetMetric) FromPbText(value string) error {
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

func (m *marshalultraEthernetMetric) ToYaml() (string, error) {
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

func (m *unMarshalultraEthernetMetric) FromYaml(value string) error {
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

func (m *marshalultraEthernetMetric) ToJson() (string, error) {
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

func (m *unMarshalultraEthernetMetric) FromJson(value string) error {
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

func (obj *ultraEthernetMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *ultraEthernetMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *ultraEthernetMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *ultraEthernetMetric) Clone() (UltraEthernetMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewUltraEthernetMetric()
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

func (obj *ultraEthernetMetric) setNil() {
	obj.cbfcSenderVirtualChannelsHolder = nil
	obj.cbfcReceiverVirtualChannelsHolder = nil
	obj.validationErrors = nil
	obj.warnings = nil
	obj.constraints = make(map[string]map[string]Constraints)
}

// UltraEthernetMetric is ultra Ethernet per port statistics information covering LLR, PHY/CtlOS and the CBFC per virtual channel state. LLR/PHY counters per UE-Specification-1.0.3 Table 5-13.
type UltraEthernetMetric interface {
	Validation
	// msg marshals UltraEthernetMetric to protobuf object *otg.UltraEthernetMetric
	// and doesn't set defaults
	msg() *otg.UltraEthernetMetric
	// setMsg unmarshals UltraEthernetMetric from protobuf object *otg.UltraEthernetMetric
	// and doesn't set defaults
	setMsg(*otg.UltraEthernetMetric) UltraEthernetMetric
	// provides marshal interface
	Marshal() marshalUltraEthernetMetric
	// provides unmarshal interface
	Unmarshal() unMarshalUltraEthernetMetric
	// validate validates UltraEthernetMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (UltraEthernetMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// Name returns string, set in UltraEthernetMetric.
	Name() string
	// SetName assigns string provided by user to UltraEthernetMetric
	SetName(value string) UltraEthernetMetric
	// HasName checks if Name has been set in UltraEthernetMetric
	HasName() bool
	// LlrModeLocal returns UltraEthernetMetricLlrModeLocalEnum, set in UltraEthernetMetric
	LlrModeLocal() UltraEthernetMetricLlrModeLocalEnum
	// SetLlrModeLocal assigns UltraEthernetMetricLlrModeLocalEnum provided by user to UltraEthernetMetric
	SetLlrModeLocal(value UltraEthernetMetricLlrModeLocalEnum) UltraEthernetMetric
	// HasLlrModeLocal checks if LlrModeLocal has been set in UltraEthernetMetric
	HasLlrModeLocal() bool
	// LlrModeRemote returns UltraEthernetMetricLlrModeRemoteEnum, set in UltraEthernetMetric
	LlrModeRemote() UltraEthernetMetricLlrModeRemoteEnum
	// SetLlrModeRemote assigns UltraEthernetMetricLlrModeRemoteEnum provided by user to UltraEthernetMetric
	SetLlrModeRemote(value UltraEthernetMetricLlrModeRemoteEnum) UltraEthernetMetric
	// HasLlrModeRemote checks if LlrModeRemote has been set in UltraEthernetMetric
	HasLlrModeRemote() bool
	// LlrTransmitState returns UltraEthernetMetricLlrTransmitStateEnum, set in UltraEthernetMetric
	LlrTransmitState() UltraEthernetMetricLlrTransmitStateEnum
	// SetLlrTransmitState assigns UltraEthernetMetricLlrTransmitStateEnum provided by user to UltraEthernetMetric
	SetLlrTransmitState(value UltraEthernetMetricLlrTransmitStateEnum) UltraEthernetMetric
	// HasLlrTransmitState checks if LlrTransmitState has been set in UltraEthernetMetric
	HasLlrTransmitState() bool
	// LlrAckNackTransmitState returns UltraEthernetMetricLlrAckNackTransmitStateEnum, set in UltraEthernetMetric
	LlrAckNackTransmitState() UltraEthernetMetricLlrAckNackTransmitStateEnum
	// SetLlrAckNackTransmitState assigns UltraEthernetMetricLlrAckNackTransmitStateEnum provided by user to UltraEthernetMetric
	SetLlrAckNackTransmitState(value UltraEthernetMetricLlrAckNackTransmitStateEnum) UltraEthernetMetric
	// HasLlrAckNackTransmitState checks if LlrAckNackTransmitState has been set in UltraEthernetMetric
	HasLlrAckNackTransmitState() bool
	// RoundTripTimeNs returns uint64, set in UltraEthernetMetric.
	RoundTripTimeNs() uint64
	// SetRoundTripTimeNs assigns uint64 provided by user to UltraEthernetMetric
	SetRoundTripTimeNs(value uint64) UltraEthernetMetric
	// HasRoundTripTimeNs checks if RoundTripTimeNs has been set in UltraEthernetMetric
	HasRoundTripTimeNs() bool
	// RoundTripTimeValid returns bool, set in UltraEthernetMetric.
	RoundTripTimeValid() bool
	// SetRoundTripTimeValid assigns bool provided by user to UltraEthernetMetric
	SetRoundTripTimeValid(value bool) UltraEthernetMetric
	// HasRoundTripTimeValid checks if RoundTripTimeValid has been set in UltraEthernetMetric
	HasRoundTripTimeValid() bool
	// MeanTimeBetweenPhyErrors returns float64, set in UltraEthernetMetric.
	MeanTimeBetweenPhyErrors() float64
	// SetMeanTimeBetweenPhyErrors assigns float64 provided by user to UltraEthernetMetric
	SetMeanTimeBetweenPhyErrors(value float64) UltraEthernetMetric
	// HasMeanTimeBetweenPhyErrors checks if MeanTimeBetweenPhyErrors has been set in UltraEthernetMetric
	HasMeanTimeBetweenPhyErrors() bool
	// TxInitCtlOs returns uint64, set in UltraEthernetMetric.
	TxInitCtlOs() uint64
	// SetTxInitCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetTxInitCtlOs(value uint64) UltraEthernetMetric
	// HasTxInitCtlOs checks if TxInitCtlOs has been set in UltraEthernetMetric
	HasTxInitCtlOs() bool
	// TxInitEchoCtlOs returns uint64, set in UltraEthernetMetric.
	TxInitEchoCtlOs() uint64
	// SetTxInitEchoCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetTxInitEchoCtlOs(value uint64) UltraEthernetMetric
	// HasTxInitEchoCtlOs checks if TxInitEchoCtlOs has been set in UltraEthernetMetric
	HasTxInitEchoCtlOs() bool
	// TxAckCtlOs returns uint64, set in UltraEthernetMetric.
	TxAckCtlOs() uint64
	// SetTxAckCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetTxAckCtlOs(value uint64) UltraEthernetMetric
	// HasTxAckCtlOs checks if TxAckCtlOs has been set in UltraEthernetMetric
	HasTxAckCtlOs() bool
	// TxNackCtlOs returns uint64, set in UltraEthernetMetric.
	TxNackCtlOs() uint64
	// SetTxNackCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetTxNackCtlOs(value uint64) UltraEthernetMetric
	// HasTxNackCtlOs checks if TxNackCtlOs has been set in UltraEthernetMetric
	HasTxNackCtlOs() bool
	// TxDiscard returns uint64, set in UltraEthernetMetric.
	TxDiscard() uint64
	// SetTxDiscard assigns uint64 provided by user to UltraEthernetMetric
	SetTxDiscard(value uint64) UltraEthernetMetric
	// HasTxDiscard checks if TxDiscard has been set in UltraEthernetMetric
	HasTxDiscard() bool
	// TxOk returns uint64, set in UltraEthernetMetric.
	TxOk() uint64
	// SetTxOk assigns uint64 provided by user to UltraEthernetMetric
	SetTxOk(value uint64) UltraEthernetMetric
	// HasTxOk checks if TxOk has been set in UltraEthernetMetric
	HasTxOk() bool
	// TxPoisoned returns uint64, set in UltraEthernetMetric.
	TxPoisoned() uint64
	// SetTxPoisoned assigns uint64 provided by user to UltraEthernetMetric
	SetTxPoisoned(value uint64) UltraEthernetMetric
	// HasTxPoisoned checks if TxPoisoned has been set in UltraEthernetMetric
	HasTxPoisoned() bool
	// TxReplayEvent returns uint64, set in UltraEthernetMetric.
	TxReplayEvent() uint64
	// SetTxReplayEvent assigns uint64 provided by user to UltraEthernetMetric
	SetTxReplayEvent(value uint64) UltraEthernetMetric
	// HasTxReplayEvent checks if TxReplayEvent has been set in UltraEthernetMetric
	HasTxReplayEvent() bool
	// ReplayedPacket returns uint64, set in UltraEthernetMetric.
	ReplayedPacket() uint64
	// SetReplayedPacket assigns uint64 provided by user to UltraEthernetMetric
	SetReplayedPacket(value uint64) UltraEthernetMetric
	// HasReplayedPacket checks if ReplayedPacket has been set in UltraEthernetMetric
	HasReplayedPacket() bool
	// ReplayedByte returns uint64, set in UltraEthernetMetric.
	ReplayedByte() uint64
	// SetReplayedByte assigns uint64 provided by user to UltraEthernetMetric
	SetReplayedByte(value uint64) UltraEthernetMetric
	// HasReplayedByte checks if ReplayedByte has been set in UltraEthernetMetric
	HasReplayedByte() bool
	// TxSeq returns uint64, set in UltraEthernetMetric.
	TxSeq() uint64
	// SetTxSeq assigns uint64 provided by user to UltraEthernetMetric
	SetTxSeq(value uint64) UltraEthernetMetric
	// HasTxSeq checks if TxSeq has been set in UltraEthernetMetric
	HasTxSeq() bool
	// TxOutstandingSeq returns uint64, set in UltraEthernetMetric.
	TxOutstandingSeq() uint64
	// SetTxOutstandingSeq assigns uint64 provided by user to UltraEthernetMetric
	SetTxOutstandingSeq(value uint64) UltraEthernetMetric
	// HasTxOutstandingSeq checks if TxOutstandingSeq has been set in UltraEthernetMetric
	HasTxOutstandingSeq() bool
	// RxInitCtlOs returns uint64, set in UltraEthernetMetric.
	RxInitCtlOs() uint64
	// SetRxInitCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetRxInitCtlOs(value uint64) UltraEthernetMetric
	// HasRxInitCtlOs checks if RxInitCtlOs has been set in UltraEthernetMetric
	HasRxInitCtlOs() bool
	// RxInitEchoCtlOs returns uint64, set in UltraEthernetMetric.
	RxInitEchoCtlOs() uint64
	// SetRxInitEchoCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetRxInitEchoCtlOs(value uint64) UltraEthernetMetric
	// HasRxInitEchoCtlOs checks if RxInitEchoCtlOs has been set in UltraEthernetMetric
	HasRxInitEchoCtlOs() bool
	// RxAckCtlOs returns uint64, set in UltraEthernetMetric.
	RxAckCtlOs() uint64
	// SetRxAckCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetRxAckCtlOs(value uint64) UltraEthernetMetric
	// HasRxAckCtlOs checks if RxAckCtlOs has been set in UltraEthernetMetric
	HasRxAckCtlOs() bool
	// RxNackCtlOs returns uint64, set in UltraEthernetMetric.
	RxNackCtlOs() uint64
	// SetRxNackCtlOs assigns uint64 provided by user to UltraEthernetMetric
	SetRxNackCtlOs(value uint64) UltraEthernetMetric
	// HasRxNackCtlOs checks if RxNackCtlOs has been set in UltraEthernetMetric
	HasRxNackCtlOs() bool
	// RxAckNackSeqError returns uint64, set in UltraEthernetMetric.
	RxAckNackSeqError() uint64
	// SetRxAckNackSeqError assigns uint64 provided by user to UltraEthernetMetric
	SetRxAckNackSeqError(value uint64) UltraEthernetMetric
	// HasRxAckNackSeqError checks if RxAckNackSeqError has been set in UltraEthernetMetric
	HasRxAckNackSeqError() bool
	// RxOk returns uint64, set in UltraEthernetMetric.
	RxOk() uint64
	// SetRxOk assigns uint64 provided by user to UltraEthernetMetric
	SetRxOk(value uint64) UltraEthernetMetric
	// HasRxOk checks if RxOk has been set in UltraEthernetMetric
	HasRxOk() bool
	// RxPoisoned returns uint64, set in UltraEthernetMetric.
	RxPoisoned() uint64
	// SetRxPoisoned assigns uint64 provided by user to UltraEthernetMetric
	SetRxPoisoned(value uint64) UltraEthernetMetric
	// HasRxPoisoned checks if RxPoisoned has been set in UltraEthernetMetric
	HasRxPoisoned() bool
	// RxBad returns uint64, set in UltraEthernetMetric.
	RxBad() uint64
	// SetRxBad assigns uint64 provided by user to UltraEthernetMetric
	SetRxBad(value uint64) UltraEthernetMetric
	// HasRxBad checks if RxBad has been set in UltraEthernetMetric
	HasRxBad() bool
	// RxExpectedSeqGood returns uint64, set in UltraEthernetMetric.
	RxExpectedSeqGood() uint64
	// SetRxExpectedSeqGood assigns uint64 provided by user to UltraEthernetMetric
	SetRxExpectedSeqGood(value uint64) UltraEthernetMetric
	// HasRxExpectedSeqGood checks if RxExpectedSeqGood has been set in UltraEthernetMetric
	HasRxExpectedSeqGood() bool
	// RxExpectedSeqPoisoned returns uint64, set in UltraEthernetMetric.
	RxExpectedSeqPoisoned() uint64
	// SetRxExpectedSeqPoisoned assigns uint64 provided by user to UltraEthernetMetric
	SetRxExpectedSeqPoisoned(value uint64) UltraEthernetMetric
	// HasRxExpectedSeqPoisoned checks if RxExpectedSeqPoisoned has been set in UltraEthernetMetric
	HasRxExpectedSeqPoisoned() bool
	// RxExpectedSeqBad returns uint64, set in UltraEthernetMetric.
	RxExpectedSeqBad() uint64
	// SetRxExpectedSeqBad assigns uint64 provided by user to UltraEthernetMetric
	SetRxExpectedSeqBad(value uint64) UltraEthernetMetric
	// HasRxExpectedSeqBad checks if RxExpectedSeqBad has been set in UltraEthernetMetric
	HasRxExpectedSeqBad() bool
	// RxMissingSeq returns uint64, set in UltraEthernetMetric.
	RxMissingSeq() uint64
	// SetRxMissingSeq assigns uint64 provided by user to UltraEthernetMetric
	SetRxMissingSeq(value uint64) UltraEthernetMetric
	// HasRxMissingSeq checks if RxMissingSeq has been set in UltraEthernetMetric
	HasRxMissingSeq() bool
	// RxDuplicateSeq returns uint64, set in UltraEthernetMetric.
	RxDuplicateSeq() uint64
	// SetRxDuplicateSeq assigns uint64 provided by user to UltraEthernetMetric
	SetRxDuplicateSeq(value uint64) UltraEthernetMetric
	// HasRxDuplicateSeq checks if RxDuplicateSeq has been set in UltraEthernetMetric
	HasRxDuplicateSeq() bool
	// RxReplay returns uint64, set in UltraEthernetMetric.
	RxReplay() uint64
	// SetRxReplay assigns uint64 provided by user to UltraEthernetMetric
	SetRxReplay(value uint64) UltraEthernetMetric
	// HasRxReplay checks if RxReplay has been set in UltraEthernetMetric
	HasRxReplay() bool
	// RxNextSeq returns uint64, set in UltraEthernetMetric.
	RxNextSeq() uint64
	// SetRxNextSeq assigns uint64 provided by user to UltraEthernetMetric
	SetRxNextSeq(value uint64) UltraEthernetMetric
	// HasRxNextSeq checks if RxNextSeq has been set in UltraEthernetMetric
	HasRxNextSeq() bool
	// LlrInitCtlOsSpacingMin returns uint64, set in UltraEthernetMetric.
	LlrInitCtlOsSpacingMin() uint64
	// SetLlrInitCtlOsSpacingMin assigns uint64 provided by user to UltraEthernetMetric
	SetLlrInitCtlOsSpacingMin(value uint64) UltraEthernetMetric
	// HasLlrInitCtlOsSpacingMin checks if LlrInitCtlOsSpacingMin has been set in UltraEthernetMetric
	HasLlrInitCtlOsSpacingMin() bool
	// LlrInitCtlOsSpacingError returns uint64, set in UltraEthernetMetric.
	LlrInitCtlOsSpacingError() uint64
	// SetLlrInitCtlOsSpacingError assigns uint64 provided by user to UltraEthernetMetric
	SetLlrInitCtlOsSpacingError(value uint64) UltraEthernetMetric
	// HasLlrInitCtlOsSpacingError checks if LlrInitCtlOsSpacingError has been set in UltraEthernetMetric
	HasLlrInitCtlOsSpacingError() bool
	// LlrAckNackCtlOsSpacingMin returns uint64, set in UltraEthernetMetric.
	LlrAckNackCtlOsSpacingMin() uint64
	// SetLlrAckNackCtlOsSpacingMin assigns uint64 provided by user to UltraEthernetMetric
	SetLlrAckNackCtlOsSpacingMin(value uint64) UltraEthernetMetric
	// HasLlrAckNackCtlOsSpacingMin checks if LlrAckNackCtlOsSpacingMin has been set in UltraEthernetMetric
	HasLlrAckNackCtlOsSpacingMin() bool
	// LlrAckNackCtlOsSpacingError returns uint64, set in UltraEthernetMetric.
	LlrAckNackCtlOsSpacingError() uint64
	// SetLlrAckNackCtlOsSpacingError assigns uint64 provided by user to UltraEthernetMetric
	SetLlrAckNackCtlOsSpacingError(value uint64) UltraEthernetMetric
	// HasLlrAckNackCtlOsSpacingError checks if LlrAckNackCtlOsSpacingError has been set in UltraEthernetMetric
	HasLlrAckNackCtlOsSpacingError() bool
	// LlrInitEchoInitSeqMismatch returns uint64, set in UltraEthernetMetric.
	LlrInitEchoInitSeqMismatch() uint64
	// SetLlrInitEchoInitSeqMismatch assigns uint64 provided by user to UltraEthernetMetric
	SetLlrInitEchoInitSeqMismatch(value uint64) UltraEthernetMetric
	// HasLlrInitEchoInitSeqMismatch checks if LlrInitEchoInitSeqMismatch has been set in UltraEthernetMetric
	HasLlrInitEchoInitSeqMismatch() bool
	// RxAckCtlOsDropped returns uint64, set in UltraEthernetMetric.
	RxAckCtlOsDropped() uint64
	// SetRxAckCtlOsDropped assigns uint64 provided by user to UltraEthernetMetric
	SetRxAckCtlOsDropped(value uint64) UltraEthernetMetric
	// HasRxAckCtlOsDropped checks if RxAckCtlOsDropped has been set in UltraEthernetMetric
	HasRxAckCtlOsDropped() bool
	// RxNackCtlOsDropped returns uint64, set in UltraEthernetMetric.
	RxNackCtlOsDropped() uint64
	// SetRxNackCtlOsDropped assigns uint64 provided by user to UltraEthernetMetric
	SetRxNackCtlOsDropped(value uint64) UltraEthernetMetric
	// HasRxNackCtlOsDropped checks if RxNackCtlOsDropped has been set in UltraEthernetMetric
	HasRxNackCtlOsDropped() bool
	// RxInitCtlOsDropped returns uint64, set in UltraEthernetMetric.
	RxInitCtlOsDropped() uint64
	// SetRxInitCtlOsDropped assigns uint64 provided by user to UltraEthernetMetric
	SetRxInitCtlOsDropped(value uint64) UltraEthernetMetric
	// HasRxInitCtlOsDropped checks if RxInitCtlOsDropped has been set in UltraEthernetMetric
	HasRxInitCtlOsDropped() bool
	// RxInitEchoCtlOsDropped returns uint64, set in UltraEthernetMetric.
	RxInitEchoCtlOsDropped() uint64
	// SetRxInitEchoCtlOsDropped assigns uint64 provided by user to UltraEthernetMetric
	SetRxInitEchoCtlOsDropped(value uint64) UltraEthernetMetric
	// HasRxInitEchoCtlOsDropped checks if RxInitEchoCtlOsDropped has been set in UltraEthernetMetric
	HasRxInitEchoCtlOsDropped() bool
	// UeRxCtlOsFrameHeaderError returns uint64, set in UltraEthernetMetric.
	UeRxCtlOsFrameHeaderError() uint64
	// SetUeRxCtlOsFrameHeaderError assigns uint64 provided by user to UltraEthernetMetric
	SetUeRxCtlOsFrameHeaderError(value uint64) UltraEthernetMetric
	// HasUeRxCtlOsFrameHeaderError checks if UeRxCtlOsFrameHeaderError has been set in UltraEthernetMetric
	HasUeRxCtlOsFrameHeaderError() bool
	// UeRxCtlOsIntraFrameSpacingError returns uint64, set in UltraEthernetMetric.
	UeRxCtlOsIntraFrameSpacingError() uint64
	// SetUeRxCtlOsIntraFrameSpacingError assigns uint64 provided by user to UltraEthernetMetric
	SetUeRxCtlOsIntraFrameSpacingError(value uint64) UltraEthernetMetric
	// HasUeRxCtlOsIntraFrameSpacingError checks if UeRxCtlOsIntraFrameSpacingError has been set in UltraEthernetMetric
	HasUeRxCtlOsIntraFrameSpacingError() bool
	// UeCtlOsSpacingMin returns uint64, set in UltraEthernetMetric.
	UeCtlOsSpacingMin() uint64
	// SetUeCtlOsSpacingMin assigns uint64 provided by user to UltraEthernetMetric
	SetUeCtlOsSpacingMin(value uint64) UltraEthernetMetric
	// HasUeCtlOsSpacingMin checks if UeCtlOsSpacingMin has been set in UltraEthernetMetric
	HasUeCtlOsSpacingMin() bool
	// UeCtlOsSpacingError returns uint64, set in UltraEthernetMetric.
	UeCtlOsSpacingError() uint64
	// SetUeCtlOsSpacingError assigns uint64 provided by user to UltraEthernetMetric
	SetUeCtlOsSpacingError(value uint64) UltraEthernetMetric
	// HasUeCtlOsSpacingError checks if UeCtlOsSpacingError has been set in UltraEthernetMetric
	HasUeCtlOsSpacingError() bool
	// CbfcSenderVirtualChannels returns UltraEthernetMetricUltraEthernetCbfcVcMetricIterIter, set in UltraEthernetMetric
	CbfcSenderVirtualChannels() UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	// CbfcReceiverVirtualChannels returns UltraEthernetMetricUltraEthernetCbfcVcMetricIterIter, set in UltraEthernetMetric
	CbfcReceiverVirtualChannels() UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	setNil()
}

// The name of the configured Ultra Ethernet instance.
// Name returns a string
func (obj *ultraEthernetMetric) Name() string {

	return *obj.obj.Name

}

// The name of the configured Ultra Ethernet instance.
// Name returns a string
func (obj *ultraEthernetMetric) HasName() bool {
	return obj.obj.Name != nil
}

// The name of the configured Ultra Ethernet instance.
// SetName sets the string value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetName(value string) UltraEthernetMetric {

	obj.obj.Name = &value
	return obj
}

type UltraEthernetMetricLlrModeLocalEnum string

// Enum of LlrModeLocal on UltraEthernetMetric
var UltraEthernetMetricLlrModeLocal = struct {
	OFF UltraEthernetMetricLlrModeLocalEnum
	ON  UltraEthernetMetricLlrModeLocalEnum
}{
	OFF: UltraEthernetMetricLlrModeLocalEnum("off"),
	ON:  UltraEthernetMetricLlrModeLocalEnum("on"),
}

func (obj *ultraEthernetMetric) LlrModeLocal() UltraEthernetMetricLlrModeLocalEnum {
	return UltraEthernetMetricLlrModeLocalEnum(obj.obj.LlrModeLocal.Enum().String())
}

// The current LLR local mode (receive) reported by the device.
// LlrModeLocal returns a string
func (obj *ultraEthernetMetric) HasLlrModeLocal() bool {
	return obj.obj.LlrModeLocal != nil
}

func (obj *ultraEthernetMetric) SetLlrModeLocal(value UltraEthernetMetricLlrModeLocalEnum) UltraEthernetMetric {
	intValue, ok := otg.UltraEthernetMetric_LlrModeLocal_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetMetricLlrModeLocalEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetMetric_LlrModeLocal_Enum(intValue)
	obj.obj.LlrModeLocal = &enumValue

	return obj
}

type UltraEthernetMetricLlrModeRemoteEnum string

// Enum of LlrModeRemote on UltraEthernetMetric
var UltraEthernetMetricLlrModeRemote = struct {
	OFF UltraEthernetMetricLlrModeRemoteEnum
	ON  UltraEthernetMetricLlrModeRemoteEnum
}{
	OFF: UltraEthernetMetricLlrModeRemoteEnum("off"),
	ON:  UltraEthernetMetricLlrModeRemoteEnum("on"),
}

func (obj *ultraEthernetMetric) LlrModeRemote() UltraEthernetMetricLlrModeRemoteEnum {
	return UltraEthernetMetricLlrModeRemoteEnum(obj.obj.LlrModeRemote.Enum().String())
}

// The current LLR remote mode (transmit) reported by the device.
// LlrModeRemote returns a string
func (obj *ultraEthernetMetric) HasLlrModeRemote() bool {
	return obj.obj.LlrModeRemote != nil
}

func (obj *ultraEthernetMetric) SetLlrModeRemote(value UltraEthernetMetricLlrModeRemoteEnum) UltraEthernetMetric {
	intValue, ok := otg.UltraEthernetMetric_LlrModeRemote_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetMetricLlrModeRemoteEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetMetric_LlrModeRemote_Enum(intValue)
	obj.obj.LlrModeRemote = &enumValue

	return obj
}

type UltraEthernetMetricLlrTransmitStateEnum string

// Enum of LlrTransmitState on UltraEthernetMetric
var UltraEthernetMetricLlrTransmitState = struct {
	OFF     UltraEthernetMetricLlrTransmitStateEnum
	INIT    UltraEthernetMetricLlrTransmitStateEnum
	ADVANCE UltraEthernetMetricLlrTransmitStateEnum
	REPLAY  UltraEthernetMetricLlrTransmitStateEnum
	FLUSH   UltraEthernetMetricLlrTransmitStateEnum
}{
	OFF:     UltraEthernetMetricLlrTransmitStateEnum("off"),
	INIT:    UltraEthernetMetricLlrTransmitStateEnum("init"),
	ADVANCE: UltraEthernetMetricLlrTransmitStateEnum("advance"),
	REPLAY:  UltraEthernetMetricLlrTransmitStateEnum("replay"),
	FLUSH:   UltraEthernetMetricLlrTransmitStateEnum("flush"),
}

func (obj *ultraEthernetMetric) LlrTransmitState() UltraEthernetMetricLlrTransmitStateEnum {
	return UltraEthernetMetricLlrTransmitStateEnum(obj.obj.LlrTransmitState.Enum().String())
}

// The current LLR remote transmit state machine value.
// LlrTransmitState returns a string
func (obj *ultraEthernetMetric) HasLlrTransmitState() bool {
	return obj.obj.LlrTransmitState != nil
}

func (obj *ultraEthernetMetric) SetLlrTransmitState(value UltraEthernetMetricLlrTransmitStateEnum) UltraEthernetMetric {
	intValue, ok := otg.UltraEthernetMetric_LlrTransmitState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetMetricLlrTransmitStateEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetMetric_LlrTransmitState_Enum(intValue)
	obj.obj.LlrTransmitState = &enumValue

	return obj
}

type UltraEthernetMetricLlrAckNackTransmitStateEnum string

// Enum of LlrAckNackTransmitState on UltraEthernetMetric
var UltraEthernetMetricLlrAckNackTransmitState = struct {
	OFF       UltraEthernetMetricLlrAckNackTransmitStateEnum
	SEND_ACKS UltraEthernetMetricLlrAckNackTransmitStateEnum
	SEND_NACK UltraEthernetMetricLlrAckNackTransmitStateEnum
	NACK_SENT UltraEthernetMetricLlrAckNackTransmitStateEnum
}{
	OFF:       UltraEthernetMetricLlrAckNackTransmitStateEnum("off"),
	SEND_ACKS: UltraEthernetMetricLlrAckNackTransmitStateEnum("send_acks"),
	SEND_NACK: UltraEthernetMetricLlrAckNackTransmitStateEnum("send_nack"),
	NACK_SENT: UltraEthernetMetricLlrAckNackTransmitStateEnum("nack_sent"),
}

func (obj *ultraEthernetMetric) LlrAckNackTransmitState() UltraEthernetMetricLlrAckNackTransmitStateEnum {
	return UltraEthernetMetricLlrAckNackTransmitStateEnum(obj.obj.LlrAckNackTransmitState.Enum().String())
}

// The current LLR ACK/NACK transmit state machine value.
// LlrAckNackTransmitState returns a string
func (obj *ultraEthernetMetric) HasLlrAckNackTransmitState() bool {
	return obj.obj.LlrAckNackTransmitState != nil
}

func (obj *ultraEthernetMetric) SetLlrAckNackTransmitState(value UltraEthernetMetricLlrAckNackTransmitStateEnum) UltraEthernetMetric {
	intValue, ok := otg.UltraEthernetMetric_LlrAckNackTransmitState_Enum_value[string(value)]
	if !ok {
		obj.validationErrors = append(obj.validationErrors, fmt.Sprintf(
			"%s is not a valid choice on UltraEthernetMetricLlrAckNackTransmitStateEnum", string(value)))
		return obj
	}
	enumValue := otg.UltraEthernetMetric_LlrAckNackTransmitState_Enum(intValue)
	obj.obj.LlrAckNackTransmitState = &enumValue

	return obj
}

// The measured round trip time, in nanoseconds, of the LLR_INIT control ordered set handshake.
// RoundTripTimeNs returns a uint64
func (obj *ultraEthernetMetric) RoundTripTimeNs() uint64 {

	return *obj.obj.RoundTripTimeNs

}

// The measured round trip time, in nanoseconds, of the LLR_INIT control ordered set handshake.
// RoundTripTimeNs returns a uint64
func (obj *ultraEthernetMetric) HasRoundTripTimeNs() bool {
	return obj.obj.RoundTripTimeNs != nil
}

// The measured round trip time, in nanoseconds, of the LLR_INIT control ordered set handshake.
// SetRoundTripTimeNs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRoundTripTimeNs(value uint64) UltraEthernetMetric {

	obj.obj.RoundTripTimeNs = &value
	return obj
}

// Indicates whether the round_trip_time_ns measurement is currently valid.
// RoundTripTimeValid returns a bool
func (obj *ultraEthernetMetric) RoundTripTimeValid() bool {

	return *obj.obj.RoundTripTimeValid

}

// Indicates whether the round_trip_time_ns measurement is currently valid.
// RoundTripTimeValid returns a bool
func (obj *ultraEthernetMetric) HasRoundTripTimeValid() bool {
	return obj.obj.RoundTripTimeValid != nil
}

// Indicates whether the round_trip_time_ns measurement is currently valid.
// SetRoundTripTimeValid sets the bool value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRoundTripTimeValid(value bool) UltraEthernetMetric {

	obj.obj.RoundTripTimeValid = &value
	return obj
}

// The mean time, in seconds, between PHY errors.
// MeanTimeBetweenPhyErrors returns a float64
func (obj *ultraEthernetMetric) MeanTimeBetweenPhyErrors() float64 {

	return *obj.obj.MeanTimeBetweenPhyErrors

}

// The mean time, in seconds, between PHY errors.
// MeanTimeBetweenPhyErrors returns a float64
func (obj *ultraEthernetMetric) HasMeanTimeBetweenPhyErrors() bool {
	return obj.obj.MeanTimeBetweenPhyErrors != nil
}

// The mean time, in seconds, between PHY errors.
// SetMeanTimeBetweenPhyErrors sets the float64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetMeanTimeBetweenPhyErrors(value float64) UltraEthernetMetric {

	obj.obj.MeanTimeBetweenPhyErrors = &value
	return obj
}

// Number of LLR_INIT control ordered sets transmitted.
// TxInitCtlOs returns a uint64
func (obj *ultraEthernetMetric) TxInitCtlOs() uint64 {

	return *obj.obj.TxInitCtlOs

}

// Number of LLR_INIT control ordered sets transmitted.
// TxInitCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasTxInitCtlOs() bool {
	return obj.obj.TxInitCtlOs != nil
}

// Number of LLR_INIT control ordered sets transmitted.
// SetTxInitCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxInitCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.TxInitCtlOs = &value
	return obj
}

// Number of LLR_INIT_ECHO control ordered sets transmitted.
// TxInitEchoCtlOs returns a uint64
func (obj *ultraEthernetMetric) TxInitEchoCtlOs() uint64 {

	return *obj.obj.TxInitEchoCtlOs

}

// Number of LLR_INIT_ECHO control ordered sets transmitted.
// TxInitEchoCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasTxInitEchoCtlOs() bool {
	return obj.obj.TxInitEchoCtlOs != nil
}

// Number of LLR_INIT_ECHO control ordered sets transmitted.
// SetTxInitEchoCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxInitEchoCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.TxInitEchoCtlOs = &value
	return obj
}

// Number of LLR_ACK control ordered sets transmitted.
// TxAckCtlOs returns a uint64
func (obj *ultraEthernetMetric) TxAckCtlOs() uint64 {

	return *obj.obj.TxAckCtlOs

}

// Number of LLR_ACK control ordered sets transmitted.
// TxAckCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasTxAckCtlOs() bool {
	return obj.obj.TxAckCtlOs != nil
}

// Number of LLR_ACK control ordered sets transmitted.
// SetTxAckCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxAckCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.TxAckCtlOs = &value
	return obj
}

// Number of LLR_NACK control ordered sets transmitted.
// TxNackCtlOs returns a uint64
func (obj *ultraEthernetMetric) TxNackCtlOs() uint64 {

	return *obj.obj.TxNackCtlOs

}

// Number of LLR_NACK control ordered sets transmitted.
// TxNackCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasTxNackCtlOs() bool {
	return obj.obj.TxNackCtlOs != nil
}

// Number of LLR_NACK control ordered sets transmitted.
// SetTxNackCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxNackCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.TxNackCtlOs = &value
	return obj
}

// Number of LLR-eligible frames discarded by the LLR transmitter while in the INIT or FLUSH state with a discard behavior.
// TxDiscard returns a uint64
func (obj *ultraEthernetMetric) TxDiscard() uint64 {

	return *obj.obj.TxDiscard

}

// Number of LLR-eligible frames discarded by the LLR transmitter while in the INIT or FLUSH state with a discard behavior.
// TxDiscard returns a uint64
func (obj *ultraEthernetMetric) HasTxDiscard() bool {
	return obj.obj.TxDiscard != nil
}

// Number of LLR-eligible frames discarded by the LLR transmitter while in the INIT or FLUSH state with a discard behavior.
// SetTxDiscard sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxDiscard(value uint64) UltraEthernetMetric {

	obj.obj.TxDiscard = &value
	return obj
}

// Number of LLR-eligible frames transmitted with a good FCS.
// TxOk returns a uint64
func (obj *ultraEthernetMetric) TxOk() uint64 {

	return *obj.obj.TxOk

}

// Number of LLR-eligible frames transmitted with a good FCS.
// TxOk returns a uint64
func (obj *ultraEthernetMetric) HasTxOk() bool {
	return obj.obj.TxOk != nil
}

// Number of LLR-eligible frames transmitted with a good FCS.
// SetTxOk sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxOk(value uint64) UltraEthernetMetric {

	obj.obj.TxOk = &value
	return obj
}

// Number of LLR-eligible frames transmitted with a poisoned FCS.
// TxPoisoned returns a uint64
func (obj *ultraEthernetMetric) TxPoisoned() uint64 {

	return *obj.obj.TxPoisoned

}

// Number of LLR-eligible frames transmitted with a poisoned FCS.
// TxPoisoned returns a uint64
func (obj *ultraEthernetMetric) HasTxPoisoned() bool {
	return obj.obj.TxPoisoned != nil
}

// Number of LLR-eligible frames transmitted with a poisoned FCS.
// SetTxPoisoned sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxPoisoned(value uint64) UltraEthernetMetric {

	obj.obj.TxPoisoned = &value
	return obj
}

// Number of times the transmitter completed a replay operation (exited the REPLAY state).
// TxReplayEvent returns a uint64
func (obj *ultraEthernetMetric) TxReplayEvent() uint64 {

	return *obj.obj.TxReplayEvent

}

// Number of times the transmitter completed a replay operation (exited the REPLAY state).
// TxReplayEvent returns a uint64
func (obj *ultraEthernetMetric) HasTxReplayEvent() bool {
	return obj.obj.TxReplayEvent != nil
}

// Number of times the transmitter completed a replay operation (exited the REPLAY state).
// SetTxReplayEvent sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxReplayEvent(value uint64) UltraEthernetMetric {

	obj.obj.TxReplayEvent = &value
	return obj
}

// Number of LLR-eligible frames replayed.
// ReplayedPacket returns a uint64
func (obj *ultraEthernetMetric) ReplayedPacket() uint64 {

	return *obj.obj.ReplayedPacket

}

// Number of LLR-eligible frames replayed.
// ReplayedPacket returns a uint64
func (obj *ultraEthernetMetric) HasReplayedPacket() bool {
	return obj.obj.ReplayedPacket != nil
}

// Number of LLR-eligible frames replayed.
// SetReplayedPacket sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetReplayedPacket(value uint64) UltraEthernetMetric {

	obj.obj.ReplayedPacket = &value
	return obj
}

// Number of bytes replayed.
// ReplayedByte returns a uint64
func (obj *ultraEthernetMetric) ReplayedByte() uint64 {

	return *obj.obj.ReplayedByte

}

// Number of bytes replayed.
// ReplayedByte returns a uint64
func (obj *ultraEthernetMetric) HasReplayedByte() bool {
	return obj.obj.ReplayedByte != nil
}

// Number of bytes replayed.
// SetReplayedByte sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetReplayedByte(value uint64) UltraEthernetMetric {

	obj.obj.ReplayedByte = &value
	return obj
}

// The current LLR transmit sequence number.
// TxSeq returns a uint64
func (obj *ultraEthernetMetric) TxSeq() uint64 {

	return *obj.obj.TxSeq

}

// The current LLR transmit sequence number.
// TxSeq returns a uint64
func (obj *ultraEthernetMetric) HasTxSeq() bool {
	return obj.obj.TxSeq != nil
}

// The current LLR transmit sequence number.
// SetTxSeq sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxSeq(value uint64) UltraEthernetMetric {

	obj.obj.TxSeq = &value
	return obj
}

// The current number of unacknowledged transmitted frames.
// TxOutstandingSeq returns a uint64
func (obj *ultraEthernetMetric) TxOutstandingSeq() uint64 {

	return *obj.obj.TxOutstandingSeq

}

// The current number of unacknowledged transmitted frames.
// TxOutstandingSeq returns a uint64
func (obj *ultraEthernetMetric) HasTxOutstandingSeq() bool {
	return obj.obj.TxOutstandingSeq != nil
}

// The current number of unacknowledged transmitted frames.
// SetTxOutstandingSeq sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetTxOutstandingSeq(value uint64) UltraEthernetMetric {

	obj.obj.TxOutstandingSeq = &value
	return obj
}

// Number of LLR_INIT control ordered sets received.
// RxInitCtlOs returns a uint64
func (obj *ultraEthernetMetric) RxInitCtlOs() uint64 {

	return *obj.obj.RxInitCtlOs

}

// Number of LLR_INIT control ordered sets received.
// RxInitCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasRxInitCtlOs() bool {
	return obj.obj.RxInitCtlOs != nil
}

// Number of LLR_INIT control ordered sets received.
// SetRxInitCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxInitCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.RxInitCtlOs = &value
	return obj
}

// Number of LLR_INIT_ECHO control ordered sets received.
// RxInitEchoCtlOs returns a uint64
func (obj *ultraEthernetMetric) RxInitEchoCtlOs() uint64 {

	return *obj.obj.RxInitEchoCtlOs

}

// Number of LLR_INIT_ECHO control ordered sets received.
// RxInitEchoCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasRxInitEchoCtlOs() bool {
	return obj.obj.RxInitEchoCtlOs != nil
}

// Number of LLR_INIT_ECHO control ordered sets received.
// SetRxInitEchoCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxInitEchoCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.RxInitEchoCtlOs = &value
	return obj
}

// Number of LLR_ACK control ordered sets received.
// RxAckCtlOs returns a uint64
func (obj *ultraEthernetMetric) RxAckCtlOs() uint64 {

	return *obj.obj.RxAckCtlOs

}

// Number of LLR_ACK control ordered sets received.
// RxAckCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasRxAckCtlOs() bool {
	return obj.obj.RxAckCtlOs != nil
}

// Number of LLR_ACK control ordered sets received.
// SetRxAckCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxAckCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.RxAckCtlOs = &value
	return obj
}

// Number of LLR_NACK control ordered sets received.
// RxNackCtlOs returns a uint64
func (obj *ultraEthernetMetric) RxNackCtlOs() uint64 {

	return *obj.obj.RxNackCtlOs

}

// Number of LLR_NACK control ordered sets received.
// RxNackCtlOs returns a uint64
func (obj *ultraEthernetMetric) HasRxNackCtlOs() bool {
	return obj.obj.RxNackCtlOs != nil
}

// Number of LLR_NACK control ordered sets received.
// SetRxNackCtlOs sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxNackCtlOs(value uint64) UltraEthernetMetric {

	obj.obj.RxNackCtlOs = &value
	return obj
}

// Number of LLR_ACK/LLR_NACK control ordered sets received with a sequence number error.
// RxAckNackSeqError returns a uint64
func (obj *ultraEthernetMetric) RxAckNackSeqError() uint64 {

	return *obj.obj.RxAckNackSeqError

}

// Number of LLR_ACK/LLR_NACK control ordered sets received with a sequence number error.
// RxAckNackSeqError returns a uint64
func (obj *ultraEthernetMetric) HasRxAckNackSeqError() bool {
	return obj.obj.RxAckNackSeqError != nil
}

// Number of LLR_ACK/LLR_NACK control ordered sets received with a sequence number error.
// SetRxAckNackSeqError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxAckNackSeqError(value uint64) UltraEthernetMetric {

	obj.obj.RxAckNackSeqError = &value
	return obj
}

// Number of LLR-eligible frames received with a good FCS.
// RxOk returns a uint64
func (obj *ultraEthernetMetric) RxOk() uint64 {

	return *obj.obj.RxOk

}

// Number of LLR-eligible frames received with a good FCS.
// RxOk returns a uint64
func (obj *ultraEthernetMetric) HasRxOk() bool {
	return obj.obj.RxOk != nil
}

// Number of LLR-eligible frames received with a good FCS.
// SetRxOk sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxOk(value uint64) UltraEthernetMetric {

	obj.obj.RxOk = &value
	return obj
}

// Number of LLR-eligible frames received with a poisoned FCS.
// RxPoisoned returns a uint64
func (obj *ultraEthernetMetric) RxPoisoned() uint64 {

	return *obj.obj.RxPoisoned

}

// Number of LLR-eligible frames received with a poisoned FCS.
// RxPoisoned returns a uint64
func (obj *ultraEthernetMetric) HasRxPoisoned() bool {
	return obj.obj.RxPoisoned != nil
}

// Number of LLR-eligible frames received with a poisoned FCS.
// SetRxPoisoned sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxPoisoned(value uint64) UltraEthernetMetric {

	obj.obj.RxPoisoned = &value
	return obj
}

// Number of LLR-eligible frames received with a bad FCS.
// RxBad returns a uint64
func (obj *ultraEthernetMetric) RxBad() uint64 {

	return *obj.obj.RxBad

}

// Number of LLR-eligible frames received with a bad FCS.
// RxBad returns a uint64
func (obj *ultraEthernetMetric) HasRxBad() bool {
	return obj.obj.RxBad != nil
}

// Number of LLR-eligible frames received with a bad FCS.
// SetRxBad sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxBad(value uint64) UltraEthernetMetric {

	obj.obj.RxBad = &value
	return obj
}

// Number of LLR-eligible frames received with a good FCS that had the expected sequence number.
// RxExpectedSeqGood returns a uint64
func (obj *ultraEthernetMetric) RxExpectedSeqGood() uint64 {

	return *obj.obj.RxExpectedSeqGood

}

// Number of LLR-eligible frames received with a good FCS that had the expected sequence number.
// RxExpectedSeqGood returns a uint64
func (obj *ultraEthernetMetric) HasRxExpectedSeqGood() bool {
	return obj.obj.RxExpectedSeqGood != nil
}

// Number of LLR-eligible frames received with a good FCS that had the expected sequence number.
// SetRxExpectedSeqGood sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxExpectedSeqGood(value uint64) UltraEthernetMetric {

	obj.obj.RxExpectedSeqGood = &value
	return obj
}

// Number of LLR-eligible frames received with a poisoned FCS that had the expected sequence number.
// RxExpectedSeqPoisoned returns a uint64
func (obj *ultraEthernetMetric) RxExpectedSeqPoisoned() uint64 {

	return *obj.obj.RxExpectedSeqPoisoned

}

// Number of LLR-eligible frames received with a poisoned FCS that had the expected sequence number.
// RxExpectedSeqPoisoned returns a uint64
func (obj *ultraEthernetMetric) HasRxExpectedSeqPoisoned() bool {
	return obj.obj.RxExpectedSeqPoisoned != nil
}

// Number of LLR-eligible frames received with a poisoned FCS that had the expected sequence number.
// SetRxExpectedSeqPoisoned sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxExpectedSeqPoisoned(value uint64) UltraEthernetMetric {

	obj.obj.RxExpectedSeqPoisoned = &value
	return obj
}

// Number of LLR-eligible frames received with a bad FCS that had the expected sequence number.
// RxExpectedSeqBad returns a uint64
func (obj *ultraEthernetMetric) RxExpectedSeqBad() uint64 {

	return *obj.obj.RxExpectedSeqBad

}

// Number of LLR-eligible frames received with a bad FCS that had the expected sequence number.
// RxExpectedSeqBad returns a uint64
func (obj *ultraEthernetMetric) HasRxExpectedSeqBad() bool {
	return obj.obj.RxExpectedSeqBad != nil
}

// Number of LLR-eligible frames received with a bad FCS that had the expected sequence number.
// SetRxExpectedSeqBad sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxExpectedSeqBad(value uint64) UltraEthernetMetric {

	obj.obj.RxExpectedSeqBad = &value
	return obj
}

// Number of LLR-eligible frames received that indicated a missing frame in the sequence.
// RxMissingSeq returns a uint64
func (obj *ultraEthernetMetric) RxMissingSeq() uint64 {

	return *obj.obj.RxMissingSeq

}

// Number of LLR-eligible frames received that indicated a missing frame in the sequence.
// RxMissingSeq returns a uint64
func (obj *ultraEthernetMetric) HasRxMissingSeq() bool {
	return obj.obj.RxMissingSeq != nil
}

// Number of LLR-eligible frames received that indicated a missing frame in the sequence.
// SetRxMissingSeq sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxMissingSeq(value uint64) UltraEthernetMetric {

	obj.obj.RxMissingSeq = &value
	return obj
}

// Number of LLR-eligible frames received that had a duplicate sequence number.
// RxDuplicateSeq returns a uint64
func (obj *ultraEthernetMetric) RxDuplicateSeq() uint64 {

	return *obj.obj.RxDuplicateSeq

}

// Number of LLR-eligible frames received that had a duplicate sequence number.
// RxDuplicateSeq returns a uint64
func (obj *ultraEthernetMetric) HasRxDuplicateSeq() bool {
	return obj.obj.RxDuplicateSeq != nil
}

// Number of LLR-eligible frames received that had a duplicate sequence number.
// SetRxDuplicateSeq sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxDuplicateSeq(value uint64) UltraEthernetMetric {

	obj.obj.RxDuplicateSeq = &value
	return obj
}

// Number of times the receiver has detected the start of a replay.
// RxReplay returns a uint64
func (obj *ultraEthernetMetric) RxReplay() uint64 {

	return *obj.obj.RxReplay

}

// Number of times the receiver has detected the start of a replay.
// RxReplay returns a uint64
func (obj *ultraEthernetMetric) HasRxReplay() bool {
	return obj.obj.RxReplay != nil
}

// Number of times the receiver has detected the start of a replay.
// SetRxReplay sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxReplay(value uint64) UltraEthernetMetric {

	obj.obj.RxReplay = &value
	return obj
}

// The current expected LLR receive sequence number.
// RxNextSeq returns a uint64
func (obj *ultraEthernetMetric) RxNextSeq() uint64 {

	return *obj.obj.RxNextSeq

}

// The current expected LLR receive sequence number.
// RxNextSeq returns a uint64
func (obj *ultraEthernetMetric) HasRxNextSeq() bool {
	return obj.obj.RxNextSeq != nil
}

// The current expected LLR receive sequence number.
// SetRxNextSeq sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxNextSeq(value uint64) UltraEthernetMetric {

	obj.obj.RxNextSeq = &value
	return obj
}

// The minimum observed spacing between received LLR_INIT control ordered sets.
// LlrInitCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) LlrInitCtlOsSpacingMin() uint64 {

	return *obj.obj.LlrInitCtlOsSpacingMin

}

// The minimum observed spacing between received LLR_INIT control ordered sets.
// LlrInitCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) HasLlrInitCtlOsSpacingMin() bool {
	return obj.obj.LlrInitCtlOsSpacingMin != nil
}

// The minimum observed spacing between received LLR_INIT control ordered sets.
// SetLlrInitCtlOsSpacingMin sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetLlrInitCtlOsSpacingMin(value uint64) UltraEthernetMetric {

	obj.obj.LlrInitCtlOsSpacingMin = &value
	return obj
}

// Number of received LLR_INIT control ordered set minimum spacing violations.
// LlrInitCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) LlrInitCtlOsSpacingError() uint64 {

	return *obj.obj.LlrInitCtlOsSpacingError

}

// Number of received LLR_INIT control ordered set minimum spacing violations.
// LlrInitCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) HasLlrInitCtlOsSpacingError() bool {
	return obj.obj.LlrInitCtlOsSpacingError != nil
}

// Number of received LLR_INIT control ordered set minimum spacing violations.
// SetLlrInitCtlOsSpacingError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetLlrInitCtlOsSpacingError(value uint64) UltraEthernetMetric {

	obj.obj.LlrInitCtlOsSpacingError = &value
	return obj
}

// The minimum observed spacing between received LLR_ACK/LLR_NACK control ordered sets.
// LlrAckNackCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) LlrAckNackCtlOsSpacingMin() uint64 {

	return *obj.obj.LlrAckNackCtlOsSpacingMin

}

// The minimum observed spacing between received LLR_ACK/LLR_NACK control ordered sets.
// LlrAckNackCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) HasLlrAckNackCtlOsSpacingMin() bool {
	return obj.obj.LlrAckNackCtlOsSpacingMin != nil
}

// The minimum observed spacing between received LLR_ACK/LLR_NACK control ordered sets.
// SetLlrAckNackCtlOsSpacingMin sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetLlrAckNackCtlOsSpacingMin(value uint64) UltraEthernetMetric {

	obj.obj.LlrAckNackCtlOsSpacingMin = &value
	return obj
}

// Number of received LLR_ACK/LLR_NACK control ordered set minimum spacing violations.
// LlrAckNackCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) LlrAckNackCtlOsSpacingError() uint64 {

	return *obj.obj.LlrAckNackCtlOsSpacingError

}

// Number of received LLR_ACK/LLR_NACK control ordered set minimum spacing violations.
// LlrAckNackCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) HasLlrAckNackCtlOsSpacingError() bool {
	return obj.obj.LlrAckNackCtlOsSpacingError != nil
}

// Number of received LLR_ACK/LLR_NACK control ordered set minimum spacing violations.
// SetLlrAckNackCtlOsSpacingError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetLlrAckNackCtlOsSpacingError(value uint64) UltraEthernetMetric {

	obj.obj.LlrAckNackCtlOsSpacingError = &value
	return obj
}

// Number of received LLR_INIT_ECHO control ordered sets whose init sequence did not match the transmit sequence.
// LlrInitEchoInitSeqMismatch returns a uint64
func (obj *ultraEthernetMetric) LlrInitEchoInitSeqMismatch() uint64 {

	return *obj.obj.LlrInitEchoInitSeqMismatch

}

// Number of received LLR_INIT_ECHO control ordered sets whose init sequence did not match the transmit sequence.
// LlrInitEchoInitSeqMismatch returns a uint64
func (obj *ultraEthernetMetric) HasLlrInitEchoInitSeqMismatch() bool {
	return obj.obj.LlrInitEchoInitSeqMismatch != nil
}

// Number of received LLR_INIT_ECHO control ordered sets whose init sequence did not match the transmit sequence.
// SetLlrInitEchoInitSeqMismatch sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetLlrInitEchoInitSeqMismatch(value uint64) UltraEthernetMetric {

	obj.obj.LlrInitEchoInitSeqMismatch = &value
	return obj
}

// Number of received LLR_ACK control ordered sets dropped.
// RxAckCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) RxAckCtlOsDropped() uint64 {

	return *obj.obj.RxAckCtlOsDropped

}

// Number of received LLR_ACK control ordered sets dropped.
// RxAckCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) HasRxAckCtlOsDropped() bool {
	return obj.obj.RxAckCtlOsDropped != nil
}

// Number of received LLR_ACK control ordered sets dropped.
// SetRxAckCtlOsDropped sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxAckCtlOsDropped(value uint64) UltraEthernetMetric {

	obj.obj.RxAckCtlOsDropped = &value
	return obj
}

// Number of received LLR_NACK control ordered sets dropped.
// RxNackCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) RxNackCtlOsDropped() uint64 {

	return *obj.obj.RxNackCtlOsDropped

}

// Number of received LLR_NACK control ordered sets dropped.
// RxNackCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) HasRxNackCtlOsDropped() bool {
	return obj.obj.RxNackCtlOsDropped != nil
}

// Number of received LLR_NACK control ordered sets dropped.
// SetRxNackCtlOsDropped sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxNackCtlOsDropped(value uint64) UltraEthernetMetric {

	obj.obj.RxNackCtlOsDropped = &value
	return obj
}

// Number of received LLR_INIT control ordered sets dropped.
// RxInitCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) RxInitCtlOsDropped() uint64 {

	return *obj.obj.RxInitCtlOsDropped

}

// Number of received LLR_INIT control ordered sets dropped.
// RxInitCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) HasRxInitCtlOsDropped() bool {
	return obj.obj.RxInitCtlOsDropped != nil
}

// Number of received LLR_INIT control ordered sets dropped.
// SetRxInitCtlOsDropped sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxInitCtlOsDropped(value uint64) UltraEthernetMetric {

	obj.obj.RxInitCtlOsDropped = &value
	return obj
}

// Number of received LLR_INIT_ECHO control ordered sets dropped.
// RxInitEchoCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) RxInitEchoCtlOsDropped() uint64 {

	return *obj.obj.RxInitEchoCtlOsDropped

}

// Number of received LLR_INIT_ECHO control ordered sets dropped.
// RxInitEchoCtlOsDropped returns a uint64
func (obj *ultraEthernetMetric) HasRxInitEchoCtlOsDropped() bool {
	return obj.obj.RxInitEchoCtlOsDropped != nil
}

// Number of received LLR_INIT_ECHO control ordered sets dropped.
// SetRxInitEchoCtlOsDropped sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetRxInitEchoCtlOsDropped(value uint64) UltraEthernetMetric {

	obj.obj.RxInitEchoCtlOsDropped = &value
	return obj
}

// Number of received UE control ordered sets with a frame header error.
// UeRxCtlOsFrameHeaderError returns a uint64
func (obj *ultraEthernetMetric) UeRxCtlOsFrameHeaderError() uint64 {

	return *obj.obj.UeRxCtlOsFrameHeaderError

}

// Number of received UE control ordered sets with a frame header error.
// UeRxCtlOsFrameHeaderError returns a uint64
func (obj *ultraEthernetMetric) HasUeRxCtlOsFrameHeaderError() bool {
	return obj.obj.UeRxCtlOsFrameHeaderError != nil
}

// Number of received UE control ordered sets with a frame header error.
// SetUeRxCtlOsFrameHeaderError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetUeRxCtlOsFrameHeaderError(value uint64) UltraEthernetMetric {

	obj.obj.UeRxCtlOsFrameHeaderError = &value
	return obj
}

// Number of received UE control ordered set intra-frame spacing violations.
// UeRxCtlOsIntraFrameSpacingError returns a uint64
func (obj *ultraEthernetMetric) UeRxCtlOsIntraFrameSpacingError() uint64 {

	return *obj.obj.UeRxCtlOsIntraFrameSpacingError

}

// Number of received UE control ordered set intra-frame spacing violations.
// UeRxCtlOsIntraFrameSpacingError returns a uint64
func (obj *ultraEthernetMetric) HasUeRxCtlOsIntraFrameSpacingError() bool {
	return obj.obj.UeRxCtlOsIntraFrameSpacingError != nil
}

// Number of received UE control ordered set intra-frame spacing violations.
// SetUeRxCtlOsIntraFrameSpacingError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetUeRxCtlOsIntraFrameSpacingError(value uint64) UltraEthernetMetric {

	obj.obj.UeRxCtlOsIntraFrameSpacingError = &value
	return obj
}

// The minimum observed spacing between received UE control ordered sets.
// UeCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) UeCtlOsSpacingMin() uint64 {

	return *obj.obj.UeCtlOsSpacingMin

}

// The minimum observed spacing between received UE control ordered sets.
// UeCtlOsSpacingMin returns a uint64
func (obj *ultraEthernetMetric) HasUeCtlOsSpacingMin() bool {
	return obj.obj.UeCtlOsSpacingMin != nil
}

// The minimum observed spacing between received UE control ordered sets.
// SetUeCtlOsSpacingMin sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetUeCtlOsSpacingMin(value uint64) UltraEthernetMetric {

	obj.obj.UeCtlOsSpacingMin = &value
	return obj
}

// Number of received UE control ordered set minimum spacing violations.
// UeCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) UeCtlOsSpacingError() uint64 {

	return *obj.obj.UeCtlOsSpacingError

}

// Number of received UE control ordered set minimum spacing violations.
// UeCtlOsSpacingError returns a uint64
func (obj *ultraEthernetMetric) HasUeCtlOsSpacingError() bool {
	return obj.obj.UeCtlOsSpacingError != nil
}

// Number of received UE control ordered set minimum spacing violations.
// SetUeCtlOsSpacingError sets the uint64 value in the UltraEthernetMetric object
func (obj *ultraEthernetMetric) SetUeCtlOsSpacingError(value uint64) UltraEthernetMetric {

	obj.obj.UeCtlOsSpacingError = &value
	return obj
}

// The CBFC sender per virtual channel credit state.
// CbfcSenderVirtualChannels returns a []UltraEthernetCbfcVcMetric
func (obj *ultraEthernetMetric) CbfcSenderVirtualChannels() UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	if len(obj.obj.CbfcSenderVirtualChannels) == 0 {
		obj.obj.CbfcSenderVirtualChannels = []*otg.UltraEthernetCbfcVcMetric{}
	}
	if obj.cbfcSenderVirtualChannelsHolder == nil {
		obj.cbfcSenderVirtualChannelsHolder = newUltraEthernetMetricUltraEthernetCbfcVcMetricIter(&obj.obj.CbfcSenderVirtualChannels).setMsg(obj)
	}
	return obj.cbfcSenderVirtualChannelsHolder
}

type ultraEthernetMetricUltraEthernetCbfcVcMetricIter struct {
	obj                            *ultraEthernetMetric
	ultraEthernetCbfcVcMetricSlice []UltraEthernetCbfcVcMetric
	fieldPtr                       *[]*otg.UltraEthernetCbfcVcMetric
}

func newUltraEthernetMetricUltraEthernetCbfcVcMetricIter(ptr *[]*otg.UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	return &ultraEthernetMetricUltraEthernetCbfcVcMetricIter{fieldPtr: ptr}
}

type UltraEthernetMetricUltraEthernetCbfcVcMetricIter interface {
	setMsg(*ultraEthernetMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	Items() []UltraEthernetCbfcVcMetric
	Add() UltraEthernetCbfcVcMetric
	Append(items ...UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	Set(index int, newObj UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	Clear() UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	clearHolderSlice() UltraEthernetMetricUltraEthernetCbfcVcMetricIter
	appendHolderSlice(item UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter
}

func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) setMsg(msg *ultraEthernetMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	obj.clearHolderSlice()
	for _, val := range *obj.fieldPtr {
		obj.appendHolderSlice(&ultraEthernetCbfcVcMetric{obj: val})
	}
	obj.obj = msg
	return obj
}

func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) Items() []UltraEthernetCbfcVcMetric {
	return obj.ultraEthernetCbfcVcMetricSlice
}

func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) Add() UltraEthernetCbfcVcMetric {
	newObj := &otg.UltraEthernetCbfcVcMetric{}
	*obj.fieldPtr = append(*obj.fieldPtr, newObj)
	newLibObj := &ultraEthernetCbfcVcMetric{obj: newObj}
	newLibObj.setDefault()
	obj.ultraEthernetCbfcVcMetricSlice = append(obj.ultraEthernetCbfcVcMetricSlice, newLibObj)
	return newLibObj
}

func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) Append(items ...UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	for _, item := range items {
		newObj := item.msg()
		*obj.fieldPtr = append(*obj.fieldPtr, newObj)
		obj.ultraEthernetCbfcVcMetricSlice = append(obj.ultraEthernetCbfcVcMetricSlice, item)
	}
	return obj
}

func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) Set(index int, newObj UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	(*obj.fieldPtr)[index] = newObj.msg()
	obj.ultraEthernetCbfcVcMetricSlice[index] = newObj
	return obj
}
func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) Clear() UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	if len(*obj.fieldPtr) > 0 {
		*obj.fieldPtr = []*otg.UltraEthernetCbfcVcMetric{}
		obj.ultraEthernetCbfcVcMetricSlice = []UltraEthernetCbfcVcMetric{}
	}
	return obj
}
func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) clearHolderSlice() UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	if len(obj.ultraEthernetCbfcVcMetricSlice) > 0 {
		obj.ultraEthernetCbfcVcMetricSlice = []UltraEthernetCbfcVcMetric{}
	}
	return obj
}
func (obj *ultraEthernetMetricUltraEthernetCbfcVcMetricIter) appendHolderSlice(item UltraEthernetCbfcVcMetric) UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	obj.ultraEthernetCbfcVcMetricSlice = append(obj.ultraEthernetCbfcVcMetricSlice, item)
	return obj
}

// The CBFC receiver per virtual channel credit state.
// CbfcReceiverVirtualChannels returns a []UltraEthernetCbfcVcMetric
func (obj *ultraEthernetMetric) CbfcReceiverVirtualChannels() UltraEthernetMetricUltraEthernetCbfcVcMetricIter {
	if len(obj.obj.CbfcReceiverVirtualChannels) == 0 {
		obj.obj.CbfcReceiverVirtualChannels = []*otg.UltraEthernetCbfcVcMetric{}
	}
	if obj.cbfcReceiverVirtualChannelsHolder == nil {
		obj.cbfcReceiverVirtualChannelsHolder = newUltraEthernetMetricUltraEthernetCbfcVcMetricIter(&obj.obj.CbfcReceiverVirtualChannels).setMsg(obj)
	}
	return obj.cbfcReceiverVirtualChannelsHolder
}

func (obj *ultraEthernetMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if len(obj.obj.CbfcSenderVirtualChannels) != 0 {

		if set_default {
			obj.CbfcSenderVirtualChannels().clearHolderSlice()
			for _, item := range obj.obj.CbfcSenderVirtualChannels {
				obj.CbfcSenderVirtualChannels().appendHolderSlice(&ultraEthernetCbfcVcMetric{obj: item})
			}
		}
		for _, item := range obj.CbfcSenderVirtualChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

	if len(obj.obj.CbfcReceiverVirtualChannels) != 0 {

		if set_default {
			obj.CbfcReceiverVirtualChannels().clearHolderSlice()
			for _, item := range obj.obj.CbfcReceiverVirtualChannels {
				obj.CbfcReceiverVirtualChannels().appendHolderSlice(&ultraEthernetCbfcVcMetric{obj: item})
			}
		}
		for _, item := range obj.CbfcReceiverVirtualChannels().Items() {
			item.validateObj(vObj, set_default)
		}

	}

}

func (obj *ultraEthernetMetric) setDefault() {

}
