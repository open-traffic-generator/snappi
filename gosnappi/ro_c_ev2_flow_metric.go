package gosnappi

import (
	"fmt"
	"strings"

	"github.com/ghodss/yaml"
	otg "github.com/open-traffic-generator/snappi/gosnappi/otg"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ***** RoCEv2FlowMetric *****
type roCEv2FlowMetric struct {
	validation
	obj          *otg.RoCEv2FlowMetric
	marshaller   marshalRoCEv2FlowMetric
	unMarshaller unMarshalRoCEv2FlowMetric
}

func NewRoCEv2FlowMetric() RoCEv2FlowMetric {
	obj := roCEv2FlowMetric{obj: &otg.RoCEv2FlowMetric{}}
	obj.setDefault()
	return &obj
}

func (obj *roCEv2FlowMetric) msg() *otg.RoCEv2FlowMetric {
	return obj.obj
}

func (obj *roCEv2FlowMetric) setMsg(msg *otg.RoCEv2FlowMetric) RoCEv2FlowMetric {

	proto.Merge(obj.obj, msg)
	return obj
}

type marshalroCEv2FlowMetric struct {
	obj *roCEv2FlowMetric
}

type marshalRoCEv2FlowMetric interface {
	// ToProto marshals RoCEv2FlowMetric to protobuf object *otg.RoCEv2FlowMetric
	ToProto() (*otg.RoCEv2FlowMetric, error)
	// ToPbText marshals RoCEv2FlowMetric to protobuf text
	ToPbText() (string, error)
	// ToYaml marshals RoCEv2FlowMetric to YAML text
	ToYaml() (string, error)
	// ToJson marshals RoCEv2FlowMetric to JSON text
	ToJson() (string, error)
}

type unMarshalroCEv2FlowMetric struct {
	obj *roCEv2FlowMetric
}

type unMarshalRoCEv2FlowMetric interface {
	// FromProto unmarshals RoCEv2FlowMetric from protobuf object *otg.RoCEv2FlowMetric
	FromProto(msg *otg.RoCEv2FlowMetric) (RoCEv2FlowMetric, error)
	// FromPbText unmarshals RoCEv2FlowMetric from protobuf text
	FromPbText(value string) error
	// FromYaml unmarshals RoCEv2FlowMetric from YAML text
	FromYaml(value string) error
	// FromJson unmarshals RoCEv2FlowMetric from JSON text
	FromJson(value string) error
}

func (obj *roCEv2FlowMetric) Marshal() marshalRoCEv2FlowMetric {
	if obj.marshaller == nil {
		obj.marshaller = &marshalroCEv2FlowMetric{obj: obj}
	}
	return obj.marshaller
}

func (obj *roCEv2FlowMetric) Unmarshal() unMarshalRoCEv2FlowMetric {
	if obj.unMarshaller == nil {
		obj.unMarshaller = &unMarshalroCEv2FlowMetric{obj: obj}
	}
	return obj.unMarshaller
}

func (m *marshalroCEv2FlowMetric) ToProto() (*otg.RoCEv2FlowMetric, error) {
	err := m.obj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return m.obj.msg(), nil
}

func (m *unMarshalroCEv2FlowMetric) FromProto(msg *otg.RoCEv2FlowMetric) (RoCEv2FlowMetric, error) {
	newObj := m.obj.setMsg(msg)
	err := newObj.validateToAndFrom()
	if err != nil {
		return nil, err
	}
	return newObj, nil
}

func (m *marshalroCEv2FlowMetric) ToPbText() (string, error) {
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

func (m *unMarshalroCEv2FlowMetric) FromPbText(value string) error {
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

func (m *marshalroCEv2FlowMetric) ToYaml() (string, error) {
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

func (m *unMarshalroCEv2FlowMetric) FromYaml(value string) error {
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

func (m *marshalroCEv2FlowMetric) ToJson() (string, error) {
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

func (m *unMarshalroCEv2FlowMetric) FromJson(value string) error {
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

func (obj *roCEv2FlowMetric) validateToAndFrom() error {
	// emptyVars()
	obj.validateObj(&obj.validation, true)
	return obj.validationResult()
}

func (obj *roCEv2FlowMetric) validate() error {
	// emptyVars()
	obj.validateObj(&obj.validation, false)
	return obj.validationResult()
}

func (obj *roCEv2FlowMetric) String() string {
	str, err := obj.Marshal().ToYaml()
	if err != nil {
		return err.Error()
	}
	return str
}

func (obj *roCEv2FlowMetric) Clone() (RoCEv2FlowMetric, error) {
	vErr := obj.validate()
	if vErr != nil {
		return nil, vErr
	}
	newObj := NewRoCEv2FlowMetric()
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

// RoCEv2FlowMetric is roCEv2 Flow statistics information.
type RoCEv2FlowMetric interface {
	Validation
	// msg marshals RoCEv2FlowMetric to protobuf object *otg.RoCEv2FlowMetric
	// and doesn't set defaults
	msg() *otg.RoCEv2FlowMetric
	// setMsg unmarshals RoCEv2FlowMetric from protobuf object *otg.RoCEv2FlowMetric
	// and doesn't set defaults
	setMsg(*otg.RoCEv2FlowMetric) RoCEv2FlowMetric
	// provides marshal interface
	Marshal() marshalRoCEv2FlowMetric
	// provides unmarshal interface
	Unmarshal() unMarshalRoCEv2FlowMetric
	// validate validates RoCEv2FlowMetric
	validate() error
	// A stringer function
	String() string
	// Clones the object
	Clone() (RoCEv2FlowMetric, error)
	validateToAndFrom() error
	validateObj(vObj *validation, set_default bool)
	setDefault()
	// FlowName returns string, set in RoCEv2FlowMetric.
	FlowName() string
	// SetFlowName assigns string provided by user to RoCEv2FlowMetric
	SetFlowName(value string) RoCEv2FlowMetric
	// HasFlowName checks if FlowName has been set in RoCEv2FlowMetric
	HasFlowName() bool
	// TrafficItem returns string, set in RoCEv2FlowMetric.
	TrafficItem() string
	// SetTrafficItem assigns string provided by user to RoCEv2FlowMetric
	SetTrafficItem(value string) RoCEv2FlowMetric
	// HasTrafficItem checks if TrafficItem has been set in RoCEv2FlowMetric
	HasTrafficItem() bool
	// PortTx returns string, set in RoCEv2FlowMetric.
	PortTx() string
	// SetPortTx assigns string provided by user to RoCEv2FlowMetric
	SetPortTx(value string) RoCEv2FlowMetric
	// HasPortTx checks if PortTx has been set in RoCEv2FlowMetric
	HasPortTx() bool
	// PortRx returns string, set in RoCEv2FlowMetric.
	PortRx() string
	// SetPortRx assigns string provided by user to RoCEv2FlowMetric
	SetPortRx(value string) RoCEv2FlowMetric
	// HasPortRx checks if PortRx has been set in RoCEv2FlowMetric
	HasPortRx() bool
	// SrcQp returns uint64, set in RoCEv2FlowMetric.
	SrcQp() uint64
	// SetSrcQp assigns uint64 provided by user to RoCEv2FlowMetric
	SetSrcQp(value uint64) RoCEv2FlowMetric
	// HasSrcQp checks if SrcQp has been set in RoCEv2FlowMetric
	HasSrcQp() bool
	// DestQp returns uint64, set in RoCEv2FlowMetric.
	DestQp() uint64
	// SetDestQp assigns uint64 provided by user to RoCEv2FlowMetric
	SetDestQp(value uint64) RoCEv2FlowMetric
	// HasDestQp checks if DestQp has been set in RoCEv2FlowMetric
	HasDestQp() bool
	// SrcIpv4 returns string, set in RoCEv2FlowMetric.
	SrcIpv4() string
	// SetSrcIpv4 assigns string provided by user to RoCEv2FlowMetric
	SetSrcIpv4(value string) RoCEv2FlowMetric
	// HasSrcIpv4 checks if SrcIpv4 has been set in RoCEv2FlowMetric
	HasSrcIpv4() bool
	// DestIpv4 returns string, set in RoCEv2FlowMetric.
	DestIpv4() string
	// SetDestIpv4 assigns string provided by user to RoCEv2FlowMetric
	SetDestIpv4(value string) RoCEv2FlowMetric
	// HasDestIpv4 checks if DestIpv4 has been set in RoCEv2FlowMetric
	HasDestIpv4() bool
	// DataFrameTx returns uint64, set in RoCEv2FlowMetric.
	DataFrameTx() uint64
	// SetDataFrameTx assigns uint64 provided by user to RoCEv2FlowMetric
	SetDataFrameTx(value uint64) RoCEv2FlowMetric
	// HasDataFrameTx checks if DataFrameTx has been set in RoCEv2FlowMetric
	HasDataFrameTx() bool
	// DataFramesRx returns uint64, set in RoCEv2FlowMetric.
	DataFramesRx() uint64
	// SetDataFramesRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetDataFramesRx(value uint64) RoCEv2FlowMetric
	// HasDataFramesRx checks if DataFramesRx has been set in RoCEv2FlowMetric
	HasDataFramesRx() bool
	// FrameDelta returns uint64, set in RoCEv2FlowMetric.
	FrameDelta() uint64
	// SetFrameDelta assigns uint64 provided by user to RoCEv2FlowMetric
	SetFrameDelta(value uint64) RoCEv2FlowMetric
	// HasFrameDelta checks if FrameDelta has been set in RoCEv2FlowMetric
	HasFrameDelta() bool
	// DataFramesRetransmitted returns uint64, set in RoCEv2FlowMetric.
	DataFramesRetransmitted() uint64
	// SetDataFramesRetransmitted assigns uint64 provided by user to RoCEv2FlowMetric
	SetDataFramesRetransmitted(value uint64) RoCEv2FlowMetric
	// HasDataFramesRetransmitted checks if DataFramesRetransmitted has been set in RoCEv2FlowMetric
	HasDataFramesRetransmitted() bool
	// FrameSequenceError returns uint64, set in RoCEv2FlowMetric.
	FrameSequenceError() uint64
	// SetFrameSequenceError assigns uint64 provided by user to RoCEv2FlowMetric
	SetFrameSequenceError(value uint64) RoCEv2FlowMetric
	// HasFrameSequenceError checks if FrameSequenceError has been set in RoCEv2FlowMetric
	HasFrameSequenceError() bool
	// TxBytes returns uint64, set in RoCEv2FlowMetric.
	TxBytes() uint64
	// SetTxBytes assigns uint64 provided by user to RoCEv2FlowMetric
	SetTxBytes(value uint64) RoCEv2FlowMetric
	// HasTxBytes checks if TxBytes has been set in RoCEv2FlowMetric
	HasTxBytes() bool
	// RxBytes returns uint64, set in RoCEv2FlowMetric.
	RxBytes() uint64
	// SetRxBytes assigns uint64 provided by user to RoCEv2FlowMetric
	SetRxBytes(value uint64) RoCEv2FlowMetric
	// HasRxBytes checks if RxBytes has been set in RoCEv2FlowMetric
	HasRxBytes() bool
	// DataTxRate returns uint64, set in RoCEv2FlowMetric.
	DataTxRate() uint64
	// SetDataTxRate assigns uint64 provided by user to RoCEv2FlowMetric
	SetDataTxRate(value uint64) RoCEv2FlowMetric
	// HasDataTxRate checks if DataTxRate has been set in RoCEv2FlowMetric
	HasDataTxRate() bool
	// DataRxRate returns uint64, set in RoCEv2FlowMetric.
	DataRxRate() uint64
	// SetDataRxRate assigns uint64 provided by user to RoCEv2FlowMetric
	SetDataRxRate(value uint64) RoCEv2FlowMetric
	// HasDataRxRate checks if DataRxRate has been set in RoCEv2FlowMetric
	HasDataRxRate() bool
	// MessageTx returns uint64, set in RoCEv2FlowMetric.
	MessageTx() uint64
	// SetMessageTx assigns uint64 provided by user to RoCEv2FlowMetric
	SetMessageTx(value uint64) RoCEv2FlowMetric
	// HasMessageTx checks if MessageTx has been set in RoCEv2FlowMetric
	HasMessageTx() bool
	// MessageCompleteRx returns uint64, set in RoCEv2FlowMetric.
	MessageCompleteRx() uint64
	// SetMessageCompleteRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetMessageCompleteRx(value uint64) RoCEv2FlowMetric
	// HasMessageCompleteRx checks if MessageCompleteRx has been set in RoCEv2FlowMetric
	HasMessageCompleteRx() bool
	// MessageFail returns uint64, set in RoCEv2FlowMetric.
	MessageFail() uint64
	// SetMessageFail assigns uint64 provided by user to RoCEv2FlowMetric
	SetMessageFail(value uint64) RoCEv2FlowMetric
	// HasMessageFail checks if MessageFail has been set in RoCEv2FlowMetric
	HasMessageFail() bool
	// FlowCompletionTime returns uint64, set in RoCEv2FlowMetric.
	FlowCompletionTime() uint64
	// SetFlowCompletionTime assigns uint64 provided by user to RoCEv2FlowMetric
	SetFlowCompletionTime(value uint64) RoCEv2FlowMetric
	// HasFlowCompletionTime checks if FlowCompletionTime has been set in RoCEv2FlowMetric
	HasFlowCompletionTime() bool
	// AvgLatency returns uint64, set in RoCEv2FlowMetric.
	AvgLatency() uint64
	// SetAvgLatency assigns uint64 provided by user to RoCEv2FlowMetric
	SetAvgLatency(value uint64) RoCEv2FlowMetric
	// HasAvgLatency checks if AvgLatency has been set in RoCEv2FlowMetric
	HasAvgLatency() bool
	// MinLatency returns uint64, set in RoCEv2FlowMetric.
	MinLatency() uint64
	// SetMinLatency assigns uint64 provided by user to RoCEv2FlowMetric
	SetMinLatency(value uint64) RoCEv2FlowMetric
	// HasMinLatency checks if MinLatency has been set in RoCEv2FlowMetric
	HasMinLatency() bool
	// MaxLatency returns uint64, set in RoCEv2FlowMetric.
	MaxLatency() uint64
	// SetMaxLatency assigns uint64 provided by user to RoCEv2FlowMetric
	SetMaxLatency(value uint64) RoCEv2FlowMetric
	// HasMaxLatency checks if MaxLatency has been set in RoCEv2FlowMetric
	HasMaxLatency() bool
	// EcnCeRx returns uint64, set in RoCEv2FlowMetric.
	EcnCeRx() uint64
	// SetEcnCeRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetEcnCeRx(value uint64) RoCEv2FlowMetric
	// HasEcnCeRx checks if EcnCeRx has been set in RoCEv2FlowMetric
	HasEcnCeRx() bool
	// CnpTx returns uint64, set in RoCEv2FlowMetric.
	CnpTx() uint64
	// SetCnpTx assigns uint64 provided by user to RoCEv2FlowMetric
	SetCnpTx(value uint64) RoCEv2FlowMetric
	// HasCnpTx checks if CnpTx has been set in RoCEv2FlowMetric
	HasCnpTx() bool
	// CnpRx returns uint64, set in RoCEv2FlowMetric.
	CnpRx() uint64
	// SetCnpRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetCnpRx(value uint64) RoCEv2FlowMetric
	// HasCnpRx checks if CnpRx has been set in RoCEv2FlowMetric
	HasCnpRx() bool
	// AckTx returns uint64, set in RoCEv2FlowMetric.
	AckTx() uint64
	// SetAckTx assigns uint64 provided by user to RoCEv2FlowMetric
	SetAckTx(value uint64) RoCEv2FlowMetric
	// HasAckTx checks if AckTx has been set in RoCEv2FlowMetric
	HasAckTx() bool
	// AckRx returns uint64, set in RoCEv2FlowMetric.
	AckRx() uint64
	// SetAckRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetAckRx(value uint64) RoCEv2FlowMetric
	// HasAckRx checks if AckRx has been set in RoCEv2FlowMetric
	HasAckRx() bool
	// NakTx returns uint64, set in RoCEv2FlowMetric.
	NakTx() uint64
	// SetNakTx assigns uint64 provided by user to RoCEv2FlowMetric
	SetNakTx(value uint64) RoCEv2FlowMetric
	// HasNakTx checks if NakTx has been set in RoCEv2FlowMetric
	HasNakTx() bool
	// NakRx returns uint64, set in RoCEv2FlowMetric.
	NakRx() uint64
	// SetNakRx assigns uint64 provided by user to RoCEv2FlowMetric
	SetNakRx(value uint64) RoCEv2FlowMetric
	// HasNakRx checks if NakRx has been set in RoCEv2FlowMetric
	HasNakRx() bool
	// FirstTimestamp returns string, set in RoCEv2FlowMetric.
	FirstTimestamp() string
	// SetFirstTimestamp assigns string provided by user to RoCEv2FlowMetric
	SetFirstTimestamp(value string) RoCEv2FlowMetric
	// HasFirstTimestamp checks if FirstTimestamp has been set in RoCEv2FlowMetric
	HasFirstTimestamp() bool
	// LastTimestamp returns string, set in RoCEv2FlowMetric.
	LastTimestamp() string
	// SetLastTimestamp assigns string provided by user to RoCEv2FlowMetric
	SetLastTimestamp(value string) RoCEv2FlowMetric
	// HasLastTimestamp checks if LastTimestamp has been set in RoCEv2FlowMetric
	HasLastTimestamp() bool
}

// Flow Name.
// FlowName returns a string
func (obj *roCEv2FlowMetric) FlowName() string {

	return *obj.obj.FlowName

}

// Flow Name.
// FlowName returns a string
func (obj *roCEv2FlowMetric) HasFlowName() bool {
	return obj.obj.FlowName != nil
}

// Flow Name.
// SetFlowName sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetFlowName(value string) RoCEv2FlowMetric {

	obj.obj.FlowName = &value
	return obj
}

// Traffic Item Name.
// TrafficItem returns a string
func (obj *roCEv2FlowMetric) TrafficItem() string {

	return *obj.obj.TrafficItem

}

// Traffic Item Name.
// TrafficItem returns a string
func (obj *roCEv2FlowMetric) HasTrafficItem() bool {
	return obj.obj.TrafficItem != nil
}

// Traffic Item Name.
// SetTrafficItem sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetTrafficItem(value string) RoCEv2FlowMetric {

	obj.obj.TrafficItem = &value
	return obj
}

// The name of the transmit port
// PortTx returns a string
func (obj *roCEv2FlowMetric) PortTx() string {

	return *obj.obj.PortTx

}

// The name of the transmit port
// PortTx returns a string
func (obj *roCEv2FlowMetric) HasPortTx() bool {
	return obj.obj.PortTx != nil
}

// The name of the transmit port
// SetPortTx sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetPortTx(value string) RoCEv2FlowMetric {

	obj.obj.PortTx = &value
	return obj
}

// The name of the receive port
// PortRx returns a string
func (obj *roCEv2FlowMetric) PortRx() string {

	return *obj.obj.PortRx

}

// The name of the receive port
// PortRx returns a string
func (obj *roCEv2FlowMetric) HasPortRx() bool {
	return obj.obj.PortRx != nil
}

// The name of the receive port
// SetPortRx sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetPortRx(value string) RoCEv2FlowMetric {

	obj.obj.PortRx = &value
	return obj
}

// Current Source QP number.
// SrcQp returns a uint64
func (obj *roCEv2FlowMetric) SrcQp() uint64 {

	return *obj.obj.SrcQp

}

// Current Source QP number.
// SrcQp returns a uint64
func (obj *roCEv2FlowMetric) HasSrcQp() bool {
	return obj.obj.SrcQp != nil
}

// Current Source QP number.
// SetSrcQp sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetSrcQp(value uint64) RoCEv2FlowMetric {

	obj.obj.SrcQp = &value
	return obj
}

// Current destination QP number.
// DestQp returns a uint64
func (obj *roCEv2FlowMetric) DestQp() uint64 {

	return *obj.obj.DestQp

}

// Current destination QP number.
// DestQp returns a uint64
func (obj *roCEv2FlowMetric) HasDestQp() bool {
	return obj.obj.DestQp != nil
}

// Current destination QP number.
// SetDestQp sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDestQp(value uint64) RoCEv2FlowMetric {

	obj.obj.DestQp = &value
	return obj
}

// Current source address.
// SrcIpv4 returns a string
func (obj *roCEv2FlowMetric) SrcIpv4() string {

	return *obj.obj.SrcIpv4

}

// Current source address.
// SrcIpv4 returns a string
func (obj *roCEv2FlowMetric) HasSrcIpv4() bool {
	return obj.obj.SrcIpv4 != nil
}

// Current source address.
// SetSrcIpv4 sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetSrcIpv4(value string) RoCEv2FlowMetric {

	obj.obj.SrcIpv4 = &value
	return obj
}

// Current destination address.
// DestIpv4 returns a string
func (obj *roCEv2FlowMetric) DestIpv4() string {

	return *obj.obj.DestIpv4

}

// Current destination address.
// DestIpv4 returns a string
func (obj *roCEv2FlowMetric) HasDestIpv4() bool {
	return obj.obj.DestIpv4 != nil
}

// Current destination address.
// SetDestIpv4 sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDestIpv4(value string) RoCEv2FlowMetric {

	obj.obj.DestIpv4 = &value
	return obj
}

// Current number of data frames transmitted.
// DataFrameTx returns a uint64
func (obj *roCEv2FlowMetric) DataFrameTx() uint64 {

	return *obj.obj.DataFrameTx

}

// Current number of data frames transmitted.
// DataFrameTx returns a uint64
func (obj *roCEv2FlowMetric) HasDataFrameTx() bool {
	return obj.obj.DataFrameTx != nil
}

// Current number of data frames transmitted.
// SetDataFrameTx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDataFrameTx(value uint64) RoCEv2FlowMetric {

	obj.obj.DataFrameTx = &value
	return obj
}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *roCEv2FlowMetric) DataFramesRx() uint64 {

	return *obj.obj.DataFramesRx

}

// Current number of data frames received.
// DataFramesRx returns a uint64
func (obj *roCEv2FlowMetric) HasDataFramesRx() bool {
	return obj.obj.DataFramesRx != nil
}

// Current number of data frames received.
// SetDataFramesRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDataFramesRx(value uint64) RoCEv2FlowMetric {

	obj.obj.DataFramesRx = &value
	return obj
}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *roCEv2FlowMetric) FrameDelta() uint64 {

	return *obj.obj.FrameDelta

}

// Current differnece between tx and rx data frames
// FrameDelta returns a uint64
func (obj *roCEv2FlowMetric) HasFrameDelta() bool {
	return obj.obj.FrameDelta != nil
}

// Current differnece between tx and rx data frames
// SetFrameDelta sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetFrameDelta(value uint64) RoCEv2FlowMetric {

	obj.obj.FrameDelta = &value
	return obj
}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *roCEv2FlowMetric) DataFramesRetransmitted() uint64 {

	return *obj.obj.DataFramesRetransmitted

}

// Current number of data frames re-transmitted.
// DataFramesRetransmitted returns a uint64
func (obj *roCEv2FlowMetric) HasDataFramesRetransmitted() bool {
	return obj.obj.DataFramesRetransmitted != nil
}

// Current number of data frames re-transmitted.
// SetDataFramesRetransmitted sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDataFramesRetransmitted(value uint64) RoCEv2FlowMetric {

	obj.obj.DataFramesRetransmitted = &value
	return obj
}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *roCEv2FlowMetric) FrameSequenceError() uint64 {

	return *obj.obj.FrameSequenceError

}

// Current number of frame sequence errors.
// FrameSequenceError returns a uint64
func (obj *roCEv2FlowMetric) HasFrameSequenceError() bool {
	return obj.obj.FrameSequenceError != nil
}

// Current number of frame sequence errors.
// SetFrameSequenceError sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetFrameSequenceError(value uint64) RoCEv2FlowMetric {

	obj.obj.FrameSequenceError = &value
	return obj
}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *roCEv2FlowMetric) TxBytes() uint64 {

	return *obj.obj.TxBytes

}

// Current number of bytes transmitted.
// TxBytes returns a uint64
func (obj *roCEv2FlowMetric) HasTxBytes() bool {
	return obj.obj.TxBytes != nil
}

// Current number of bytes transmitted.
// SetTxBytes sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetTxBytes(value uint64) RoCEv2FlowMetric {

	obj.obj.TxBytes = &value
	return obj
}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *roCEv2FlowMetric) RxBytes() uint64 {

	return *obj.obj.RxBytes

}

// Current number of bytes received.
// RxBytes returns a uint64
func (obj *roCEv2FlowMetric) HasRxBytes() bool {
	return obj.obj.RxBytes != nil
}

// Current number of bytes received.
// SetRxBytes sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetRxBytes(value uint64) RoCEv2FlowMetric {

	obj.obj.RxBytes = &value
	return obj
}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *roCEv2FlowMetric) DataTxRate() uint64 {

	return *obj.obj.DataTxRate

}

// Current rate at which data is transmitted in Gbps.
// DataTxRate returns a uint64
func (obj *roCEv2FlowMetric) HasDataTxRate() bool {
	return obj.obj.DataTxRate != nil
}

// Current rate at which data is transmitted in Gbps.
// SetDataTxRate sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDataTxRate(value uint64) RoCEv2FlowMetric {

	obj.obj.DataTxRate = &value
	return obj
}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *roCEv2FlowMetric) DataRxRate() uint64 {

	return *obj.obj.DataRxRate

}

// Current rate at which data is received in Gbps.
// DataRxRate returns a uint64
func (obj *roCEv2FlowMetric) HasDataRxRate() bool {
	return obj.obj.DataRxRate != nil
}

// Current rate at which data is received in Gbps.
// SetDataRxRate sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetDataRxRate(value uint64) RoCEv2FlowMetric {

	obj.obj.DataRxRate = &value
	return obj
}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *roCEv2FlowMetric) MessageTx() uint64 {

	return *obj.obj.MessageTx

}

// Current number of Message transmitted.
// MessageTx returns a uint64
func (obj *roCEv2FlowMetric) HasMessageTx() bool {
	return obj.obj.MessageTx != nil
}

// Current number of Message transmitted.
// SetMessageTx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetMessageTx(value uint64) RoCEv2FlowMetric {

	obj.obj.MessageTx = &value
	return obj
}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *roCEv2FlowMetric) MessageCompleteRx() uint64 {

	return *obj.obj.MessageCompleteRx

}

// Current number of Message Complete received.
// MessageCompleteRx returns a uint64
func (obj *roCEv2FlowMetric) HasMessageCompleteRx() bool {
	return obj.obj.MessageCompleteRx != nil
}

// Current number of Message Complete received.
// SetMessageCompleteRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetMessageCompleteRx(value uint64) RoCEv2FlowMetric {

	obj.obj.MessageCompleteRx = &value
	return obj
}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *roCEv2FlowMetric) MessageFail() uint64 {

	return *obj.obj.MessageFail

}

// Current number of Message Fail count.
// MessageFail returns a uint64
func (obj *roCEv2FlowMetric) HasMessageFail() bool {
	return obj.obj.MessageFail != nil
}

// Current number of Message Fail count.
// SetMessageFail sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetMessageFail(value uint64) RoCEv2FlowMetric {

	obj.obj.MessageFail = &value
	return obj
}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *roCEv2FlowMetric) FlowCompletionTime() uint64 {

	return *obj.obj.FlowCompletionTime

}

// Current flow comletion time in ms.
// FlowCompletionTime returns a uint64
func (obj *roCEv2FlowMetric) HasFlowCompletionTime() bool {
	return obj.obj.FlowCompletionTime != nil
}

// Current flow comletion time in ms.
// SetFlowCompletionTime sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetFlowCompletionTime(value uint64) RoCEv2FlowMetric {

	obj.obj.FlowCompletionTime = &value
	return obj
}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *roCEv2FlowMetric) AvgLatency() uint64 {

	return *obj.obj.AvgLatency

}

// Current average latency measured in ns.
// AvgLatency returns a uint64
func (obj *roCEv2FlowMetric) HasAvgLatency() bool {
	return obj.obj.AvgLatency != nil
}

// Current average latency measured in ns.
// SetAvgLatency sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetAvgLatency(value uint64) RoCEv2FlowMetric {

	obj.obj.AvgLatency = &value
	return obj
}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *roCEv2FlowMetric) MinLatency() uint64 {

	return *obj.obj.MinLatency

}

// Current minimum latency measured in ns.
// MinLatency returns a uint64
func (obj *roCEv2FlowMetric) HasMinLatency() bool {
	return obj.obj.MinLatency != nil
}

// Current minimum latency measured in ns.
// SetMinLatency sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetMinLatency(value uint64) RoCEv2FlowMetric {

	obj.obj.MinLatency = &value
	return obj
}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *roCEv2FlowMetric) MaxLatency() uint64 {

	return *obj.obj.MaxLatency

}

// Current maximum latency measured in ns.
// MaxLatency returns a uint64
func (obj *roCEv2FlowMetric) HasMaxLatency() bool {
	return obj.obj.MaxLatency != nil
}

// Current maximum latency measured in ns.
// SetMaxLatency sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetMaxLatency(value uint64) RoCEv2FlowMetric {

	obj.obj.MaxLatency = &value
	return obj
}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *roCEv2FlowMetric) EcnCeRx() uint64 {

	return *obj.obj.EcnCeRx

}

// Current number of ECN-CE Recevied.
// EcnCeRx returns a uint64
func (obj *roCEv2FlowMetric) HasEcnCeRx() bool {
	return obj.obj.EcnCeRx != nil
}

// Current number of ECN-CE Recevied.
// SetEcnCeRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetEcnCeRx(value uint64) RoCEv2FlowMetric {

	obj.obj.EcnCeRx = &value
	return obj
}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *roCEv2FlowMetric) CnpTx() uint64 {

	return *obj.obj.CnpTx

}

// Current number of CNP transmitted.
// CnpTx returns a uint64
func (obj *roCEv2FlowMetric) HasCnpTx() bool {
	return obj.obj.CnpTx != nil
}

// Current number of CNP transmitted.
// SetCnpTx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetCnpTx(value uint64) RoCEv2FlowMetric {

	obj.obj.CnpTx = &value
	return obj
}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *roCEv2FlowMetric) CnpRx() uint64 {

	return *obj.obj.CnpRx

}

// Current number of CNP received.
// CnpRx returns a uint64
func (obj *roCEv2FlowMetric) HasCnpRx() bool {
	return obj.obj.CnpRx != nil
}

// Current number of CNP received.
// SetCnpRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetCnpRx(value uint64) RoCEv2FlowMetric {

	obj.obj.CnpRx = &value
	return obj
}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *roCEv2FlowMetric) AckTx() uint64 {

	return *obj.obj.AckTx

}

// Current number of ACK transmitted.
// AckTx returns a uint64
func (obj *roCEv2FlowMetric) HasAckTx() bool {
	return obj.obj.AckTx != nil
}

// Current number of ACK transmitted.
// SetAckTx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetAckTx(value uint64) RoCEv2FlowMetric {

	obj.obj.AckTx = &value
	return obj
}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *roCEv2FlowMetric) AckRx() uint64 {

	return *obj.obj.AckRx

}

// Current number of ACK received.
// AckRx returns a uint64
func (obj *roCEv2FlowMetric) HasAckRx() bool {
	return obj.obj.AckRx != nil
}

// Current number of ACK received.
// SetAckRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetAckRx(value uint64) RoCEv2FlowMetric {

	obj.obj.AckRx = &value
	return obj
}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *roCEv2FlowMetric) NakTx() uint64 {

	return *obj.obj.NakTx

}

// Current number of NAK transmitted.
// NakTx returns a uint64
func (obj *roCEv2FlowMetric) HasNakTx() bool {
	return obj.obj.NakTx != nil
}

// Current number of NAK transmitted.
// SetNakTx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetNakTx(value uint64) RoCEv2FlowMetric {

	obj.obj.NakTx = &value
	return obj
}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *roCEv2FlowMetric) NakRx() uint64 {

	return *obj.obj.NakRx

}

// Current number of NAK received.
// NakRx returns a uint64
func (obj *roCEv2FlowMetric) HasNakRx() bool {
	return obj.obj.NakRx != nil
}

// Current number of NAK received.
// SetNakRx sets the uint64 value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetNakRx(value uint64) RoCEv2FlowMetric {

	obj.obj.NakRx = &value
	return obj
}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *roCEv2FlowMetric) FirstTimestamp() string {

	return *obj.obj.FirstTimestamp

}

// First Timestamp.
// FirstTimestamp returns a string
func (obj *roCEv2FlowMetric) HasFirstTimestamp() bool {
	return obj.obj.FirstTimestamp != nil
}

// First Timestamp.
// SetFirstTimestamp sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetFirstTimestamp(value string) RoCEv2FlowMetric {

	obj.obj.FirstTimestamp = &value
	return obj
}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *roCEv2FlowMetric) LastTimestamp() string {

	return *obj.obj.LastTimestamp

}

// Last Timestamp.
// LastTimestamp returns a string
func (obj *roCEv2FlowMetric) HasLastTimestamp() bool {
	return obj.obj.LastTimestamp != nil
}

// Last Timestamp.
// SetLastTimestamp sets the string value in the RoCEv2FlowMetric object
func (obj *roCEv2FlowMetric) SetLastTimestamp(value string) RoCEv2FlowMetric {

	obj.obj.LastTimestamp = &value
	return obj
}

func (obj *roCEv2FlowMetric) validateObj(vObj *validation, set_default bool) {
	if set_default {
		obj.setDefault()
	}

	if obj.obj.SrcIpv4 != nil {

		err := obj.validateIpv4(obj.SrcIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowMetric.SrcIpv4"))
		}

	}

	if obj.obj.DestIpv4 != nil {

		err := obj.validateIpv4(obj.DestIpv4())
		if err != nil {
			vObj.validationErrors = append(vObj.validationErrors, fmt.Sprintf("%s %s", err.Error(), "on RoCEv2FlowMetric.DestIpv4"))
		}

	}

}

func (obj *roCEv2FlowMetric) setDefault() {

}
